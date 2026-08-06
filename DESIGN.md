---
name: "趣搭局"
description: "把城市兴趣服务组织成一面可信、可筛选、可预约的活动海报墙。"
colors:
  forest-masthead: "#0F4838"
  forest-ink: "#153F35"
  coral-ticket: "#EF644A"
  warm-paper: "#FFFDF6"
  cool-canvas: "#F3F5F0"
  search-paper: "#F7F5EE"
  mint-stub: "#B8CFBF"
  mint-band: "#CBDACF"
  masthead-cream: "#F7F3E8"
  muted-ink: "#66736C"
  perforation: "#C6D0C8"
  stamp-brown: "#6C3527"
typography:
  display:
    fontFamily: "-apple-system, BlinkMacSystemFont, PingFang SC, Microsoft YaHei, sans-serif"
    fontSize: "70rpx"
    fontWeight: 900
    lineHeight: 1.12
    letterSpacing: "-3rpx"
  wordmark:
    fontFamily: "-apple-system, BlinkMacSystemFont, PingFang SC, Microsoft YaHei, sans-serif"
    fontSize: "49rpx"
    fontWeight: 900
    lineHeight: 1
    letterSpacing: "-5rpx"
  headline:
    fontFamily: "-apple-system, BlinkMacSystemFont, PingFang SC, Microsoft YaHei, sans-serif"
    fontSize: "34rpx"
    fontWeight: 900
    lineHeight: 1.1
    letterSpacing: "-1rpx"
  title:
    fontFamily: "-apple-system, BlinkMacSystemFont, PingFang SC, Microsoft YaHei, sans-serif"
    fontSize: "28rpx"
    fontWeight: 900
    lineHeight: 1.25
  body:
    fontFamily: "-apple-system, BlinkMacSystemFont, PingFang SC, Microsoft YaHei, sans-serif"
    fontSize: "22rpx"
    fontWeight: 400
    lineHeight: 1.5
  label:
    fontFamily: "-apple-system, BlinkMacSystemFont, PingFang SC, Microsoft YaHei, sans-serif"
    fontSize: "18rpx"
    fontWeight: 700
    lineHeight: 1.2
rounded:
  detail: "8rpx"
  paper: "12rpx"
  ticket: "14rpx"
  control: "18rpx"
  pill: "999rpx"
  circle: "50%"
spacing:
  xs: "8rpx"
  sm: "13rpx"
  md: "17rpx"
  lg: "24rpx"
  xl: "34rpx"
  section: "42rpx"
components:
  button-primary:
    backgroundColor: "{colors.forest-ink}"
    textColor: "{colors.warm-paper}"
    typography: "{typography.body}"
    rounded: "{rounded.ticket}"
    padding: "0 28rpx"
    height: "72rpx"
  button-coral:
    backgroundColor: "{colors.coral-ticket}"
    textColor: "{colors.warm-paper}"
    typography: "{typography.body}"
    rounded: "{rounded.ticket}"
    padding: "0 28rpx"
    height: "72rpx"
  chip-active:
    backgroundColor: "{colors.forest-ink}"
    textColor: "{colors.warm-paper}"
    typography: "{typography.label}"
    rounded: "{rounded.pill}"
    padding: "0 18rpx"
    height: "52rpx"
  search-rail:
    backgroundColor: "{colors.search-paper}"
    textColor: "{colors.muted-ink}"
    typography: "{typography.body}"
    rounded: "{rounded.pill}"
    padding: "7rpx 7rpx 7rpx 24rpx"
    height: "82rpx"
  activity-ticket:
    backgroundColor: "{colors.warm-paper}"
    textColor: "{colors.forest-ink}"
    rounded: "{rounded.ticket}"
    padding: "19rpx"
  bottom-navigation:
    backgroundColor: "{colors.warm-paper}"
    textColor: "{colors.forest-ink}"
    height: "104rpx"
---

# Design System: 趣搭局

## Overview

**Creative North Star: "城市活动海报墙 / B 全幅城市活动海报"**

趣搭局不是一组普通的圆角服务卡，而是一面正在更新的城市活动海报墙：森林绿报头像街区公告栏，暖白内容像被收集和叠放的纸张，珊瑚票根、虚线撕口、票孔与印章把每项服务变成一张可以赴约的节目票。界面必须同时保留生活气和平台可信度，装饰语法用于解释层级、状态与交易，而不是遮蔽真实信息。

首页完整表达“B 全幅城市活动海报”：白色品牌栏只承载字标、右上位置胶囊和独立搜索按钮；其下的竖屏纪实摄影满宽展开，左下森林绿撕纸大标题与右下珊瑚票根 CTA 形成唯一首屏焦点。四张横向分类票、推荐工具栏与全幅摄影海报流顺序承接主视觉。搜索、详情、下单、订单、个人中心和服务者中心不复制首页构图，但必须继承纸张、票券、撕线、印章、森林绿与珊瑚主操作的同一视觉语法。

所有活动摄影均为 **synthetic demo imagery（AI 生成的演示摄影）**，只用于 MVP 视觉与裁切验证，不代表真实服务者、真实顾客、真实活动或平台实拍。文字、价格、认证、时间、位置和按钮必须保持为独立语义界面元素，禁止烘焙进图片。

**Key Characteristics:**

- 白色品牌栏、右上位置胶囊与独立搜索按钮共同建立清晰的城市入口。
- 全幅竖屏纪实摄影以左下深绿撕纸标题、右下珊瑚票根 CTA 形成首屏海报焦点。
- 暖白纸张、珊瑚票根、撕线、票孔和印章构成可复用的票券材料语言。
- 纪实感活动摄影承载人物与场景，界面层承载全部事实与操作。
- 真实功能筛选同时提供推荐、本周、3km 内快捷条件和完整底部筛选抽屉。
- CSS 几何图标统一导航、分类、筛选和工作台，不依赖 emoji 或外部图标字体。
- 全宽底部导航跨页面固定，安全区和内容留白同步适配。

## Colors

配色像城市公告栏上的油墨与纸张：深绿稳定背景和信息骨架，珊瑚只在票根、价格、紧迫状态和关键动作上发声，薄荷绿连接认证、保障与次级信息。

### Primary

- **报头森林绿**：用于搜索、个人与服务者中心报头，以及首页主摄影上的撕纸标题块和活动海报信息条。
- **油墨森林绿**：用于正文标题、主按钮、激活筛选、CSS 图标和关键边线。

### Secondary

- **票根珊瑚**：用于主票根、价格、倒计时、主交易按钮、选中态和手绘下划线；它不是通用背景色。

### Tertiary

- **存根薄荷绿**：用于认证、横向分类票、标签与弱强调。
- **信息薄荷带**：用于票券时间地点条、保障条和状态分区。
- **印章棕**：仅用于圆形品牌印章等模拟油墨痕迹，不用于正文。

### Neutral

- **活动暖白纸**：票券、抽屉、底部导航和主要内容卡的标准表面。
- **冷调画布**：纸张之间的页面底色，也是票孔负形的颜色来源。
- **搜索纸**：输入轨道、评论框等需要更柔和区分的内嵌表面。
- **报头奶油白**：深绿背景上的标题、图标和主文案。
- **褪色油墨**：服务者、距离、说明和元数据。
- **撕线灰绿**：虚线、票券分隔和弱边界。

**The Coral Stub Rule.** 珊瑚色只服务于“可行动、可交易或像票根一样可撕下”的信息；单屏不把多个普通容器同时染成珊瑚色。

**The Paper Contrast Rule.** 暖白纸张必须落在冷调画布或深绿报头上，禁止把纸张与页面背景都退回无层级的纯白。

## Typography

**Display Font:** 系统无衬线中文字体栈（优先 `PingFang SC`）  
**Body Font:** 同一系统字体栈  
**Label Font:** 同一系统字体栈，数字启用等宽数字特征时使用 `tabular-nums`

**Character:** 字体依靠高字重、紧凑字距、轻微旋转和票据编号形成街头海报感，而不是依赖外部展示字体。中文必须保持快速可读，英文编号、`NO.`、`CURRENT CITY` 等只作为短标签和票据索引。

### Hierarchy

- **Display**（900，超大，紧凑）：首页问句与少量首屏主张；最多两行。
- **Wordmark**（900，极紧字距）：趣搭局品牌字标，单字允许轻微错位旋转。
- **Headline**（900，紧凑）：模块标题、服务详情标题和工作台区块。
- **Title**（900）：活动票券标题、订单标题和列表主信息。
- **Body**（400–700）：筛选、说明、服务信息和按钮文案。
- **Label**（700）：票号、认证、距离、时间、短标签与英文票据标记。

**The Poster Weight Rule.** 报头与票券标题可以重，但正文不能整页加粗；每个纸面容器只保留一个最强文字层级。

**The Semantic Type Rule.** 价格、日期、编号和位置始终由真实文本渲染，摄影素材中不得出现承担产品语义的文字。

## Layout

系统以 750rpx 小程序设计宽度为基准。常规页面水平边距为 24–34rpx；票券内部使用 16–28rpx；模块节奏通常为 24–42rpx。所有页面顶部叠加 `env(safe-area-inset-top)`，固定底栏和主操作栏叠加 `env(safe-area-inset-bottom)`。

首页采用单列全幅海报构图。顶部白色品牌栏高约 126rpx（另加顶部安全区），左侧是不规则字标，右侧是位置胶囊与独立搜索图标。主视觉为 592rpx 高的满宽竖屏纪实摄影；左下森林绿撕纸标题块覆盖约 440rpx 宽，右下珊瑚票根从画面右侧伸出，两个叠层共同构成首屏唯一焦点。主视觉之后依次是四列横向分类票、推荐/本周/3km 内/筛选工具栏，以及带文字贴纸、印章和价格章的全幅活动摄影海报流。

其他页面使用“继承材料，不照搬构图”的规则：搜索页以深绿搜索报头和粘性结果工具栏承接票券列表；详情页以纪实摄影开场，服务者票券叠压图片；下单和订单页把摘要、进度、注意事项组织成票据组；个人与服务者中心用深绿身份报头、暖白通行证和节目票组织功能。

位置控件固定在白色品牌栏右上，与字标同层，使用白底描边胶囊、城市/区域文本和定位几何图标；独立搜索按钮紧邻其右，不把搜索塞入位置胶囊或主视觉。首页筛选由“推荐 / 本周 / 3km 内”快捷条件与完整底部抽屉组成，区域、服务方式、可约时间、最低价和最高价均影响真实展示结果或状态，不制作无效筛选装饰。

全宽底部导航固定贴合屏幕左右和底部，高度包含安全区。所有带导航页面必须预留至少 150–170rpx 的尾部空间；详情和支付页的固定操作栏同样不得遮挡最后一块内容。

**The Cross-Page Adaptation Rule.** 每个页面至少继承两种材料信号（报头、纸张、撕线、票孔、印章、摄影或票根），但只在首页使用白色品牌栏接全幅竖屏主海报与单列全幅摄影流。

## Elevation & Depth

深度来自摄影上的撕纸/票根覆盖、纸张叠层、轻微旋转、票孔负形和克制的绿色环境阴影。首页主摄影本身贴边且无卡片投影，深度由左下撕纸标题与右下票根制造；活动摄影海报、身份通行证、工作台票和固定导航才获得阴影。遮罩抽屉使用深色 scrim 与纸面顶部绿边建立模态层级。

### Shadow Vocabulary

- **珊瑚票根投影**（`0 10rpx 22rpx rgba(47, 30, 22, .16)`）：首页主海报右下 CTA，使票根从摄影上抬起。
- **海报抬升**（`0 11rpx 28rpx rgba(16, 63, 51, .13)`）：首页单列活动摄影海报和主要服务票券。
- **标准纸张**（`0 8rpx 20rpx rgba(21, 63, 53, .06–.07)`）：菜单、订单信息和最近浏览等普通纸张。
- **身份通行证**（`0 13rpx 28rpx rgba(10, 45, 35, .20)`）：深绿报头上叠放的个人身份纸张。
- **底部导航**（`0 -12rpx 32rpx rgba(21, 63, 53, .08)`）：全宽固定导航与内容的弱分离。

**The Paper Stack Rule.** 阴影必须解释真实叠放关系；禁止给每一行、每个 chip 或每个按钮添加阴影。

## Shapes

票券主体使用 12–14rpx 的小圆角，表单和底部抽屉使用 16–18rpx，快捷筛选和搜索轨道使用 999rpx 胶囊。圆形只保留给头像、票孔、印章和少量状态符号；页面不再以 32–44rpx 的大圆角白卡作为默认容器。

票券分隔线使用 2–3rpx 灰绿虚线；左右半圆缺口使用与页面画布同色的圆形负形；珊瑚票根通过虚线连接主票面。印章允许约 -5° 至 -7° 的旋转，叠纸允许约 -2° 至 1.2° 的偏转，功能文字本身不随意倾斜。

CSS 几何图标由边框、矩形、圆形、伪元素和 `clip-path` 构成，线宽通常为 3–4rpx。首页横向分类票、独立搜索按钮、筛选工具栏、底部导航、个人快捷入口和服务者工作台必须使用同一几何密度，不引入 emoji、图标字体或风格突变的插画符号。

**The Functional Distress Rule.** 撕线、票孔和印章必须出现在票据语义或状态凭证上；不要把每个普通分隔线都做成仿旧效果。

## Components

### Buttons

- **Primary:** 森林绿实心、暖白文字，小票券圆角；用于搜索确认、应用筛选和非交易主操作。
- **Coral Transaction:** 珊瑚实心、暖白文字；用于立即预约、支付、接单和强交易动作。
- **Secondary:** 暖白或透明纸面、森林绿文字、2rpx 低对比边线；用于重置、婉拒和咨询。
- **State:** 小程序以 tap/active 反馈为主，不依赖 hover；操作中必须阻止重复提交。所有按钮保留清晰文字，几何图标只辅助含义。

### Chips

- 快捷筛选使用胶囊；默认透明纸面加灰绿边线，激活为森林绿底暖白字。
- 完整筛选项使用 16rpx 控件圆角；默认浅纸面，选中为珊瑚底暖白字。
- 认证和内容标签使用薄荷底、森林绿字和直角或极小圆角，形似贴在票面上的标签。

### Cards / Containers

- **首页主海报:** 满宽竖屏纪实摄影直接贴边；左下覆盖森林绿撕纸大标题，右下覆盖带虚线撕口的珊瑚 CTA 票根。主海报不是圆角卡，也不叠放背纸。
- **横向分类票:** 四张等宽小票紧密排成一行，暖白与薄荷交替；相邻票之间以连续半圆票孔分隔，并使用统一 CSS 几何分类图标。
- **活动海报流:** 每项服务使用一张全幅横向摄影海报，标题、类别与标签以深绿/暖白/薄荷贴纸叠在照片左侧，价格章固定右下；可选圆形印章只用于确有状态含义的活动。
- **跨页服务票券:** 搜索结果等非首页场景继续使用纪实摄影 + 信息主体 + 虚线元数据带，左右票孔从画布中“咬出”。
- **服务/订单票:** 暖白纸张、小圆角、虚线分区、单层环境阴影；必要时使用珊瑚顶边指示主信息。
- **状态凭证:** 深绿或薄荷信息带承载认证、平台保障、档期与履约状态。

### Inputs / Fields

- 首页不展示搜索输入轨道；搜索是白色品牌栏右侧的独立 CSS 放大镜按钮，点击后进入搜索结果页。
- 搜索结果页保留同一胶囊搜索语言，并置于深绿报头内。
- 价格输入、留言和筛选字段使用浅纸面、低对比边线或虚线；不得伪装为装饰票号。
- 空值、清除和输入结果必须可见；最低价、最高价、时间和服务方式参与真实筛选或订单状态。

### Navigation

- 底部导航为全宽暖白纸面，三项均分，图标与文字垂直排列。
- 激活项使用森林绿文字和顶部 5rpx 珊瑚短线，不使用大块悬浮胶囊。
- 图标由 CSS 几何绘制：首页为实心屋形、订单为票据/清单、我的为人物轮廓。
- 深绿页面顶部的返回、设置和动作按钮使用小方形轮廓，保持报头层级清晰。

### Location Control

- 位置切换始终位于首页白色品牌栏右上角，与品牌同层，独立搜索按钮位于其右侧。
- 使用白底、浅灰描边的胶囊，显示“城市 · 区域”和 CSS 定位图标；点击打开区域选择底部抽屉。
- 搜索按钮保持透明、无外框，以 34rpx 圆环和斜柄绘制放大镜，不与位置控件合并。

### Bottom Sheet

- 暖白纸面、18rpx 顶部圆角、8rpx 森林绿顶边，最大高度 84vh。
- 使用深色遮罩、拖拽条和方形关闭按钮；显示时阻止背景滚动。
- 区域、服务方式、日期和价格控件必须形成完整可操作流程，主按钮文案反馈当前结果数量。

### Documentary Imagery

- 图像采用自然光、真实皮肤与材质、轻微胶片颗粒的城市纪实摄影语气，避免商业棚拍、网红摆拍、过度 HDR 和霓虹夜店感。
- 首页主活动使用满宽 592rpx 高的竖屏安全裁切；首页活动流使用 352rpx 高的全幅横向裁切；跨页服务票券继续使用接近 5:4 的中心裁切，均通过 `aspectFill` 适配容器。
- 当前 `src/static/posterwall` 中全部摄影均为 synthetic demo imagery；更换为真实授权素材前必须保留该说明。

## Do's and Don'ts

### Do:

- **Do** 把位置胶囊和独立搜索按钮放在白色品牌栏右上，与品牌同层，并提供真实区域选择状态。
- **Do** 让首页主摄影满宽贴边，以左下深绿撕纸标题和右下珊瑚票根构成唯一首屏焦点。
- **Do** 用暖白纸张、虚线撕口、票孔、薄荷信息带和珊瑚票根组织服务与订单信息。
- **Do** 让快捷筛选与完整筛选真正改变结果、排序或订单状态，并提供可恢复的空状态。
- **Do** 用 CSS 几何图标维持分类、导航、筛选、订单和工作台的一致线宽与形态。
- **Do** 在跨页面适配时保留材料语言，同时为搜索、详情、下单、订单和中心页选择各自合适的构图。
- **Do** 明确标注当前全部活动摄影为 synthetic demo imagery，并把所有事实与操作留在语义界面层。

### Don't:

- **Don't** 回退到大圆角白卡、悬浮胶囊底栏和通用模板式分区作为默认视觉语言。
- **Don't** 把位置切换放到主标题下方、把搜索塞进主视觉，或添加没有业务作用的右上角图标。
- **Don't** 把首页重新拆成深绿报头、搜索胶囊和左侧节目单/右侧内容流的分栏布局。
- **Don't** 使用 emoji、混杂风格的图标库或图片内嵌按钮与文案代替 CSS 几何图标和真实组件。
- **Don't** 把珊瑚色铺满普通内容区；它应集中在票根、价格、选中状态和主交易动作。
- **Don't** 为仿旧而滥用旋转、阴影、撕线和印章；每种效果必须对应纸张层级、票据边界或状态凭证。
- **Don't** 将 synthetic demo imagery 描述为实拍、真实用户活动或真实服务者资料。
