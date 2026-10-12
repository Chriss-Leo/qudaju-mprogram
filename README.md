# 趣搭局小程序

同城兴趣组局平台。技术栈：uni-app + Vue 3 + TypeScript + Vite。

## 文档入口

- [产品方案](PRODUCT.md)：功能范围、报名与聊天规则。
- [设计规范](DESIGN.md)：当前确认的视觉与布局。
- [开发指南](docs/DEVELOPMENT.md)：运行、构建、目录和平台差异。
- [实现进展](docs/STATUS.md)：已实现、待重构和待开发。
- [素材说明](docs/ASSETS.md)：图片、字体、图标与授权。
- [清理记录](docs/CLEANUP.md)：移除范围及恢复位置。

## 快速运行

```bash
npm ci
npm run dev:h5
```

打开终端输出的本地地址。微信使用 `npm run build:mp-weixin`，在开发者工具导入 `dist/build/mp-weixin`。AppID 在 `src/manifest.json` 配置；本地工具配置保留在项目根目录。

当前可用：首页固定摄影入口、组局列表、筛选、微信原生省市区选择。数据是本地示例；发布、真实报名、支付与聊天未接入。详情等旧服务页面仍被引用，尚未按新产品重构。
