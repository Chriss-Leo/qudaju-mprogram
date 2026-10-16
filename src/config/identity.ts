// Public API origin only. Never put AppSecret or session_key in VITE_* variables.
export const identityBaseURL = (import.meta.env.VITE_IDENTITY_BASE_URL || "").replace(/\/+$/, "");
export const identityPaths = { login: "/v1/auth/wechat/login", me: "/v1/auth/me", logout: "/v1/auth/logout", profile: "/v1/auth/me/profile" } as const;
export function requireIdentityURL() {
  if (!identityBaseURL) throw new Error("微信登录服务尚未配置，请配置 VITE_IDENTITY_BASE_URL");
  if (!/^https:\/\//.test(identityBaseURL) && !(import.meta.env.DEV && /^http:\/\//.test(identityBaseURL))) throw new Error("登录服务必须使用HTTPS");
  return identityBaseURL;
}
