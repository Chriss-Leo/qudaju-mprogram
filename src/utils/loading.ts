// uni-app maps these APIs to wx.showLoading / wx.hideLoading on WeChat.
// Share ownership so concurrent operations cannot hide another operation's loading.
let pending = 0;
export async function withNativeLoading<T>(title: string, operation: () => Promise<T>): Promise<T> {
  if (pending++ === 0) uni.showLoading({ title, mask: true });
  try {
    return await operation();
  } finally {
    if (--pending === 0) uni.hideLoading();
  }
}
