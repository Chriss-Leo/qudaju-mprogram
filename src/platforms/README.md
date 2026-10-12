# 平台适配层

wechat、alipay、xhs分别保留微信、支付宝、小红书能力入口。统一接口定义在types.ts；目前使用演示实现，未完成真实登录、支付或客服。

页面通过getPlatformAdapter调用业务API。原生UI如微信region picker允许页面内条件编译；其他端不能假设同样支持。

不要删除支付宝和小红书目录，即使当前优先开发微信。
