# 趣搭局工程约定

## 主题配置（强制）

- 配色、字体、字号、字重、行高、通用圆角等统一维护在 `src/theme/tokens.json`。
- 不得在页面、组件或图标导出脚本手写主题色号或通用字体值。
- Vue样式使用uni.scss提供的Sass变量；运行时属性从 `src/theme/index.ts` 导入theme。
- `src/theme/generated.scss`和导航SVG由脚本生成，不可手改。修改配置后执行 `npm run theme:generate`。
- 新页面优先使用语义色彩角色，禁止复制旧页面的compat颜色作为新方案。
- 平台API固定的frontColor黑/白字面量是例外；组件独有几何位置和照片裁切不必提升为主题参数。
- 修改后执行 `npm run type-check`（自动包含theme:check）与相应平台构建。完整使用说明见 `docs/THEME.md`。

## 产品边界

保留微信、支付宝、小红书适配目录。PRODUCT.md定义目标规则，docs/STATUS.md定义真实完成情况。不要将演示数据、旧交易页面或编译成功描述成真实业务已经接通。
