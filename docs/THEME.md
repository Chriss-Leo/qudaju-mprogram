# 主题配置规范

## 唯一入口

修改 `src/theme/tokens.json`，不得在页面、组件或图标生成脚本另定义配色。

- `colors`：配色及遮罩；主要角色为 primary、accent、text、textMuted、white、card、tag、input、border。
- `fontFamilies` / `fontFaces`：CSS字体栈和运行时加载的字体名称。
- `fontSizes` / `fontWeights` / `lineHeights`：保留现有字阶，不在迁移时改视觉。
- `radii`：圆角；`spacing`、`iconSizes`、`opacity`：通用uni-ui配置。

`compat*` 是仍被旧页面使用的历史特殊值，为避免清理时改变视觉暂时集中保留。新功能必须优先使用语义角色，不增加随意编号的兼容值；旧页面重构后才删除不再引用的值。

## 使用方式

Sass通过uni.scss自动获得变量，例如 `$color-primary`、`$color-card`、`$font-body`、`$font-size-v26rpx`、`$radius-v16rpx`。

Vue/TypeScript从 `src/theme/index.ts` 导入 `theme`，使用 `theme.colors.primary`；uni-icons用 `:color="theme.colors.primary"`，不得再写静态色号。

字体栈可以更改，但新字体文件仍需单独授权、子集化与加载，不会因为改一个名称就自动下载。

## 生成与检查

`npm run theme:generate` 从JSON生成 `src/theme/generated.scss`，同步五个导航SVG的颜色。生成文件和SVG不是第二配置源，不得手改。

Vite启动/构建时自动生成；H5开发服务器监听tokens.json变更并更新样式。手动重建导航形状后也运行主题生成。

`npm run theme:check` 检查应用Vue文件中硬编码色号、字体、字号、字重、行高、圆角。`npm run type-check` 会先执行此规范检查。

允许的例外仅为平台API规定的 `frontColor` 黑白字面量。照片像素颜色不属于UI主题。组件专有几何位置、照片尺寸、响应式画布缩放不是通用主题参数，留在对应组件中。

## 当前主题

蓝色确认/筛选 `primary`，橙色重点操作 `accent`，黑色正文 `text`，浅灰白卡片 `card`。布局和行为不因主题修改而改变。

强制配对：所有橙色按钮背景使用accent，文字必须使用白色accentForeground，不允许黑字。新按钮使用action-accent通用样式。详情底部免费或具体费用使用accent；辅助费用说明仍使用次级文字色。该白字/现有橙色配对为用户明确指定规范，不宣称满足普通文字的WCAG AA对比度。
