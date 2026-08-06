export type MiniappPlatform = "wechat" | "alipay" | "xhs" | "h5" | "unknown";

export interface PlatformAdapter {
  name: MiniappPlatform;
  bootstrap: () => void;
  requestPayment: (orderId: string) => Promise<void>;
  openCustomerService: () => Promise<void>;
}
