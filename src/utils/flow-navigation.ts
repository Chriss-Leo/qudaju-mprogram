export function flowBack(fallback = "/pages/profile/index") { uni.navigateBack({ fail: () => uni.redirectTo({ url: fallback }) }); }
export function openFlow(path: string, id?: string) { uni.navigateTo({ url: `/pages/${path}/index${id ? `?id=${encodeURIComponent(id)}` : ""}` }); }
export function flowError(error: unknown) { uni.showToast({ title: error instanceof Error ? error.message : "操作失败，请重试", icon: "none" }); }
