import { createDemoAdapter } from "../shared";

export const wechatAdapter = {
  ...createDemoAdapter("wechat"),
  loginCode: () => new Promise<string>((resolve, reject) => uni.login({
    provider: "weixin", timeout: 10000,
    success: result => result.code ? resolve(result.code) : reject(new Error("微信未返回登录凭证，请重试")),
    fail: () => reject(new Error("微信登录失败，请检查开发者工具登录状态后重试")),
  })),
};
