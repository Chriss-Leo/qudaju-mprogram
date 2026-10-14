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
- `src/pages/messages` / `chat-invitation` / `chat`：本地消息分栏、邀请回应、文字聊天与关闭后的可读历史。
- `src/pages/group-apply` / `my-registrations` / `my-hosted` / `group-members`：报名留言、报名记录、局主活动与成员审核管理。
- `src/components`：导航、卡片、抽屉和兼容旧页的组件。
- `src/styles/flows.scss`：报名与聊天页面共用的标签、卡片、按钮、提示和空状态样式；`FlowShell` 统一固定头尾与独立滚动区。
- `src/stores/meetups.ts`：本地示例身份、报名、邀请、文字消息、结束状态与人数计算，不是服务端业务实现。
- `src/utils/use-flow-id.ts`：同步页面 ID，兼容 H5 缓存页仅查询参数变更；只读取所属路由，避免其他缓存页被误更新。
- `src/theme`：统一主题配置及生成产物，详见THEME.md；不得在组件重复定义。
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

## 本地报名与聊天调试

当前身份固定为 `me`，局主示例活动为 `my-game`；数据保存在 uni 存储键 `qudaju-local-meetups-v1`。需要恢复初始示例时，在目标平台开发工具中删除这一键并重新加载，不要清除其他用户存储。无效或旧版本数据会退回初始示例。

你通过 `getGroup` 读取随报名状态变化的活动人数和头像，不直接修改 `groupCards`。待审核、待付款不占正式名额；模拟付款完成和免费加入会检查名额上限。客户端检查局主所有权、邀请收件人、报名关系、邀请状态与聊天开启条件；接入后端时必须在服务端重新校验并处理并发，不能信任这些客户端状态。

报名提交与聊天发送仅在本地生效。「去付款」必须保留明确的不收款模拟确认框；收费正式报名不能通过本地退出伪装退款。退出或结束活动后，以邀请的 `closed` 标记冻结关闭状态；恢复存储时也冻结已经失效的旧邀请。重新报名不会重开旧会话，必须创建并接受新邀请，已接受旧会话的历史仍可查看。真实登录、IM、推送、支付、退款及基金结算均不在此模型中。

新增路由回归应覆盖未授权成员页、未知 ID、邀请未接受/已处理/已失效、退出与结束后的关闭历史、接受→退出→重新报名后旧会话仍关闭、空记录、小屏换行和 H5 同路由不同 ID。2026-10-14 的页面截图在 `local-tests/flows-20261014/`，批准的参考图在 `docs/design/review-20261014/`；它们不进入运行时静态目录。

## 回归检查

```bash
QUDADJU_PLAYWRIGHT_PATH=/path/to/playwright node scripts/check-entry.cjs
```

需可用Playwright及Chromium；默认读取 `http://127.0.0.1:5173/`，可用 `QUDADJU_BASE_URL` 覆盖。此依赖不是运行时业务包。

图标重建：`node scripts/prepare-entry-assets.mjs`。标题字体重建：`python3 scripts/prepare-rounded-font.py /path/to/ResourceHanRoundedCN-Bold.ttf`；需要fontTools和Brotli。字体源文件从作者发布页获取，见素材说明。不要将完整中文字体加入静态目录。
