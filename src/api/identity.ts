import { identityPaths, requireIdentityURL } from "../config/identity";
export interface AuthUser { id: string; nickname: string; avatarUrl: string }
export interface LoginResult { accessToken: string; tokenType: "Bearer"; expiresIn: number; user: AuthUser }
interface ResponseEntity<T> { code: string; msg: string; data: T; time: string }
export class UnauthorizedError extends Error {}
const messages: Record<string, string> = { wechat_login_failed: "微信登录失败，请重新尝试", account_disabled: "账号不可用", "3100001": "请求参数异常", "3010107": "服务繁忙，请稍后重试", rate_limited: "请求过于频繁，请稍后重试" };
async function request<T>(path: string, method: "GET" | "POST", token?: string, data?: object): Promise<T> {
  const base = requireIdentityURL();
  return new Promise((resolve, reject) => uni.request({
    url: base + path, method, data, timeout: 12000,
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
  return !!user && typeof user.id === "string" && /^[1-9]\d*$/.test(user.id) && typeof user.nickname === "string" && typeof user.avatarUrl === "string" && (!user.avatarUrl || /^https:\/\//.test(user.avatarUrl));
}
export const identityAPI = {
  login: (code: string) => request<LoginResult>(identityPaths.login, "POST", undefined, { code }),
  me: (token: string) => request<AuthUser>(identityPaths.me, "GET", token),
  logout: (token: string) => request<null>(identityPaths.logout, "POST", token),
};
