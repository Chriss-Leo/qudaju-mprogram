import type { PlatformAdapter } from "./types";
import { createDemoAdapter } from "./shared";
import { wechatAdapter } from "./wechat";
import { alipayAdapter } from "./alipay";
import { xhsAdapter } from "./xhs";

export function getPlatformAdapter(): PlatformAdapter {
  let adapter = createDemoAdapter("unknown");
  // #ifdef H5
  adapter = createDemoAdapter("h5");
  // #endif
  // #ifdef MP-WEIXIN
  adapter = wechatAdapter;
  // #endif
  // #ifdef MP-ALIPAY
  adapter = alipayAdapter;
  // #endif
  // #ifdef MP-XHS
  adapter = xhsAdapter;
  // #endif
  return adapter;
}
