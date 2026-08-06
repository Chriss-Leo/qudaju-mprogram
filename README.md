# 趣搭局小程序

基于 uni-app、Vue 3、TypeScript 和 Vite 搭建的趣搭局 MVP 小程序端。当前实现以 `qudaju-ui` 的页面设计稿为依据，包含可交互的首页筛选与核心交易页面。

## 已实现页面

- 首页：分类、附近优选、位置选择、排序和筛选抽屉
- 搜索结果：关键词搜索、排序、快速筛选和无结果状态
- 服务详情：服务者信息、服务说明、保障、档期、评价和预约入口
- 确认订单：预约信息、联系方式、留言、平台支付适配入口
- 订单详情：接单状态、履约进度、订单信息、安全提示和售后入口
- 个人中心：订单快捷入口、收藏评价、客服和服务者中心入口
- 服务者中心：经营数据、常用功能、待接单处理

## 平台目录

平台差异能力统一放在 `src/platforms`：

- `src/platforms/wechat`：微信小程序
- `src/platforms/alipay`：支付宝小程序
- `src/platforms/xhs`：小红书小程序

支付、客服、登录、分享等平台 API 后续只在适配层实现，页面通过统一接口调用。

## 本地运行

```bash
npm install
npm run dev:h5
```

## 构建

```bash
npm run build:h5
npm run build:mp-weixin
npm run build:mp-alipay
npm run build:mp-xhs
```

对应产物位于：

- `dist/build/h5`
- `dist/build/mp-weixin`
- `dist/build/mp-alipay`
- `dist/build/mp-xhs`

导入相应平台开发者工具前，需要先在 `src/manifest.json` 配置各平台 App ID。

## 检查

```bash
npm run type-check
```

当前页面使用本地示例数据，支付、客服、认证和提现入口为演示状态，已保留平台适配接口。
