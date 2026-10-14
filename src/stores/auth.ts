import { computed, reactive } from "vue";
import { identityAPI, UnauthorizedError, validUser, type AuthUser } from "../api/identity";
import { getPlatformAdapter } from "../platforms";
import { requireIdentityURL } from "../config/identity";

interface Session { accessToken: string; expiresAt: number; user: AuthUser }
const key = "qudaju-auth-session-v1";
export const authState = reactive<{ status: "guest" | "checking" | "authenticated"; session: Session | null; error: string }>({ status: "guest", session: null, error: "" });
export const isAuthenticated = computed(() => authState.status === "authenticated" && !!authState.session && authState.session.expiresAt > Date.now());
let validation: Promise<boolean> | undefined;
let loginPending: Promise<void> | undefined;
let generation = 0;
let expiry: ReturnType<typeof setTimeout> | undefined;
function validSession(value: unknown): value is Session {
  const s = value as Session;
  return !!s && typeof s.accessToken === "string" && /^[A-Za-z0-9_-]{43}$/.test(s.accessToken) && Number.isFinite(s.expiresAt) && s.expiresAt > Date.now() && s.expiresAt <= Date.now() + 86400000 && validUser(s.user);
}
export function clearAuth() {
  generation++; if (expiry) clearTimeout(expiry);
  authState.status = "guest"; authState.session = null;
  try { uni.removeStorageSync(key); } catch { /* Still clear in memory. */ }
}
function install(session: Session) {
  if (!validSession(session)) throw new Error("登录服务返回了无效会话");
  uni.setStorageSync(key, session);
  authState.session = session; authState.status = "authenticated"; authState.error = "";
  if (expiry) clearTimeout(expiry);
  expiry = setTimeout(() => clearAuth(), Math.max(0, session.expiresAt - Date.now()));
}
export async function restoreAuth(): Promise<boolean> {
  if (isAuthenticated.value) return true;
  if (validation) return validation;
  const start = generation;
  validation = (async () => {
    let cached: unknown;
    try { cached = uni.getStorageSync(key); } catch { return false; }
    if (!validSession(cached)) { clearAuth(); return false; }
    authState.status = "checking";
    try {
      const user = await identityAPI.me(cached.accessToken);
      if (generation !== start) return false;
      if (!validUser(user) || user.id !== cached.user.id) { clearAuth(); return false; }
      install({ ...cached, user }); return true;
    } catch (error) {
      if (generation !== start) return false;
      if (error instanceof UnauthorizedError) clearAuth(); else authState.status = "guest";
      authState.error = error instanceof Error ? error.message : "登录校验失败"; return false;
    }
  })().finally(() => { validation = undefined; });
  return validation;
}
export async function wechatLogin(): Promise<void> {
  if (loginPending) return loginPending;
  const start = ++generation;
  loginPending = (async () => {
    requireIdentityURL();
    const code = await getPlatformAdapter().loginCode();
    const result = await identityAPI.login(code);
    if (generation !== start) return;
    if (result.tokenType !== "Bearer" || !Number.isInteger(result.expiresIn) || result.expiresIn < 60 || result.expiresIn > 86400) throw new Error("登录服务返回了无效会话");
    install({ accessToken: result.accessToken, expiresAt: Date.now() + result.expiresIn * 1000, user: result.user });
  })().finally(() => { loginPending = undefined; });
  return loginPending;
}
export function assertAuthenticated() {
  if (!isAuthenticated.value) { clearAuth(); throw new UnauthorizedError("请先使用微信登录"); }
}
export async function logout() {
  const token = authState.session?.accessToken;
  try { if (token) await identityAPI.logout(token); }
  catch (error) { if (!(error instanceof UnauthorizedError)) throw new Error("已退出本机，服务端会话暂未能撤销；请检查网络"); }
  finally { clearAuth(); }
}
