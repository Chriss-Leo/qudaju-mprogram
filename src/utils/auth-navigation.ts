import { isAuthenticated, restoreAuth } from "../stores/auth";

const protectedPaths = new Set(["profile", "my-registrations", "my-hosted", "group-members", "messages", "chat-invitation", "chat", "group-apply", "order-submit", "order-detail", "provider"]);
export function requiresLogin(url: string) { const path = url.split("?")[0].replace(/^\//, ""); return protectedPaths.has(path.split("/")[1]); }
export function safeReturnURL(value: string) {
  if (!/^\/pages\/[a-z-]+\/index(?:\?[^#]*)?$/.test(value) || value.includes("\\") || value.length > 1000 || value.split("?")[0] === "/pages/login/index") return "/pages/profile/index";
  const allowed = new Set([...protectedPaths, "index", "groups", "service-detail"]);
  return allowed.has(value.split("/")[2]) ? value : "/pages/profile/index";
}
let opening = false;
export function openLogin(next = "/pages/profile/index") {
  const current = getCurrentPages().slice(-1)[0];
  if (opening || current?.route === "pages/login/index") return;
  opening = true;
  uni.navigateTo({ url: `/pages/login/index?next=${encodeURIComponent(safeReturnURL(next))}`, complete: () => { opening = false; } });
}
export function ensureLogin(next = "/pages/profile/index") { if (isAuthenticated.value) return true; openLogin(next); return false; }
export function installAuthNavigation() {
  for (const name of ["navigateTo", "redirectTo", "reLaunch", "switchTab"] as const) {
    uni.addInterceptor(name, { invoke(options: { url: string }) { if (requiresLogin(options.url) && !isAuthenticated.value) { openLogin(options.url); return false; } return options; } });
  }
  void restoreAuth();
}
