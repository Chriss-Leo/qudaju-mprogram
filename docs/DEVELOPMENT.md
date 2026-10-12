# 开发指南

## 启动与检查

```bash
npm ci
npm run dev:h5
npm run type-check
npm run build:mp-weixin
npm run build:mp-alipay
npm run build:mp-xhs
npm run build:h5
```

构建产物为 `dist/build/<平台>`。微信导入 `dist/build/mp-weixin`，配置本人的小程序AppID与开发者权限。开发者工具登录过期需重新登录，不能靠改业务代码修复。

## 目录职责

- `src/pages/index`：固定首页。
- `src/pages/groups`：组局列表、地址、统一搜索筛选。
- `src/components`：导航、卡片、抽屉和兼容旧页的组件。
- `src/mocks`：示例数据。groups仍借用services的id/距离/可用时间；待独立活动模型接入。
- `src/platforms`：微信、支付宝、小红书能力适配，目前为演示接口。
- `src/static`：运行时素材；不得混入候选稿或大原图。
- `scripts`：导航SVG导出、标题字体子集、H5回归检查。
- `docs/design`：有限的确认/待实现参考稿，不打包进小程序。
- `docs/brand`：小程序对外图标，不是页面运行依赖。

旧详情、订单、搜索、个人中心和服务者页面当前仍可访问，参见进展清单，不要误删它们依赖的数据/素材。

## 平台差异

业务平台API收敛在适配层；原生UI兼容可在页面条件编译。微信省市区组件仅微信启用，H5/支付宝/小红书暂用上海区域备用，不支持全国滚轮。

原生picker及支付等必须在目标开发者工具或真机验证；H5测试不代表原生运行完全通过。

## 回归检查

```bash
QUDADJU_PLAYWRIGHT_PATH=/path/to/playwright node scripts/check-entry.cjs
```

需可用Playwright及Chromium；默认读取 `http://127.0.0.1:5173/`，可用 `QUDADJU_BASE_URL` 覆盖。此依赖不是运行时业务包。

图标重建：`node scripts/prepare-entry-assets.mjs`。标题字体重建：`python3 scripts/prepare-rounded-font.py /path/to/ResourceHanRoundedCN-Bold.ttf`；需要fontTools和Brotli。字体源文件从作者发布页获取，见素材说明。不要将完整中文字体加入静态目录。
