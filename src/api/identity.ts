import { identityBaseURL, identityPaths, requireIdentityURL } from "../config/identity";
export interface AuthUser { id: string; nickname: string; avatarUrl: string }
export interface LoginResult { accessToken: string; tokenType: "Bearer"; expiresIn: number; user: AuthUser }
interface ResponseEntity<T> { code: string; msg: string; data: T; time: string }
export class UnauthorizedError extends Error {}
const messages: Record<string, string> = { wechat_login_failed: "微信登录失败，请重新尝试", wechat_phone_failed: "手机号验证失败，请重新授权；若持续失败，请检查小程序手机号能力配置", account_disabled: "账号不可用", "3100001": "请求参数异常", "3010107": "服务繁忙，请稍后重试", rate_limited: "请求过于频繁，请稍后重试" };
async function request<T>(path: string, method: "GET" | "POST", token?: string, data?: object): Promise<T> {
  const base = requireIdentityURL();
  return new Promise((resolve, reject) => uni.request({
    url: base + path, method, data, timeout: path === identityPaths.login ? 20000 : 12000,
    header: { "Content-Type": "application/json", "Accept-Language": "zh-CN", ...(token ? { Authorization: `Bearer ${token}` } : {}) },
    success(response) {
      if (response.statusCode === 401) { reject(new UnauthorizedError("登录已失效，请重新登录")); return; }
      if (response.statusCode < 200 || response.statusCode >= 300) { reject(new Error(response.statusCode === 429 ? "请求过于频繁，请稍后重试" : "登录服务暂不可用")); return; }
      const body = response.data as ResponseEntity<T>;
      if (!body || body.code !== "1000000") { reject(new Error(messages[body?.code] || "操作失败，请稍后重试")); return; }
      resolve(body.data);
    },
    fail: () => reject(new Error("无法连接登录服务，请检查网络后重试")),
  }));
}
export function validUser(value: unknown): value is AuthUser {
  const user = value as AuthUser;
  return !!user && typeof user.id === "string" && /^[1-9]\d*$/.test(user.id) && typeof user.nickname === "string" && typeof user.avatarUrl === "string" && (!user.avatarUrl || /^https:\/\//.test(user.avatarUrl) || (import.meta.env.DEV && !!identityBaseURL && user.avatarUrl.startsWith(identityBaseURL + "/v1/media/avatars/")));
}
function normalizeUser(user: AuthUser): AuthUser {
  if (user && /^\/v1\/media\/avatars\/[a-f0-9]{48}\.jpg$/.test(user.avatarUrl)) return { ...user, avatarUrl: requireIdentityURL() + user.avatarUrl };
  return user;
}
export function hasProfile(user?: AuthUser | null) { return !!user && !!user.avatarUrl && user.nickname.trim() !== "微信用户" && Array.from(user.nickname.trim()).length >= 2; }
function saveProfile(token: string, nickname: string, avatarPath?: string): Promise<AuthUser> {
  const url = requireIdentityURL() + identityPaths.profile;
  return new Promise((resolve, reject) => {
    const success = (response: { statusCode: number; data: unknown }) => {
      if (response.statusCode === 401) { reject(new UnauthorizedError("登录已失效，请重新登录")); return; }
      try {
        const body = (typeof response.data === "string" ? JSON.parse(response.data) : response.data) as ResponseEntity<AuthUser>;
        if (response.statusCode !== 200 || body?.code !== "1000000") throw new Error(body?.code === "3100001" ? "请填写2–20字昵称，并选择2MB以内的JPG或PNG头像" : messages[body?.code] || "资料保存失败，请重试");
        const user = normalizeUser(body.data); if (!validUser(user) || !hasProfile(user)) throw new Error("资料保存结果异常，请重试"); resolve(user);
      } catch (error) { reject(error instanceof Error ? error : new Error("资料保存失败")); }
    };
    const fail = () => reject(new Error("无法保存资料，请检查网络后重试"));
    if (avatarPath) uni.uploadFile({ url, filePath: avatarPath, name: "avatar", formData: { nickname }, header: { Authorization: `Bearer ${token}`, "Accept-Language": "zh-CN" }, timeout: 30000, success, fail });
    else uni.request({ url, method: "POST", data: { nickname }, header: { Authorization: `Bearer ${token}`, "Content-Type": "application/x-www-form-urlencoded", "Accept-Language": "zh-CN" }, timeout: 12000, success, fail });
  });
}
export const identityAPI = {
  login: async (code: string, phoneCode: string) => { const result = await request<LoginResult>(identityPaths.login, "POST", undefined, { code, phoneCode }); return { ...result, user: normalizeUser(result.user) }; },
  me: async (token: string) => normalizeUser(await request<AuthUser>(identityPaths.me, "GET", token)),
  logout: (token: string) => request<null>(identityPaths.logout, "POST", token),
  saveProfile,
};
