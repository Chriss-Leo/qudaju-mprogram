import type { PlatformAdapter } from "./types";

export const createDemoAdapter = (name: PlatformAdapter["name"]): PlatformAdapter => ({
  name,
  bootstrap() {
    uni.setStorageSync("qudaju-platform", name);
  },
  async requestPayment() {
    uni.showToast({ title: "支付能力将在联调阶段接入", icon: "none" });
  },
  async openCustomerService() {
    uni.showToast({ title: "客服能力将在联调阶段接入", icon: "none" });
  },
});
