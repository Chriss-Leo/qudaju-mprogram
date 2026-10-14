# 素材与授权

## 运行时素材

| 目录 | 用途 |
| --- | --- |
| src/static/brand | 当前白橙透明字标 |
| src/static/entry/city-walk.jpg | 首页摄影底图 |
| src/static/entry/fonts | 资源圆体Bold标题子集、OFL许可证 |
| src/static/entry/icons | 五个固定导航SVG及授权说明 |
| src/static/groups | 三组三图条示例摄影 |
| src/static/group-detail | 独立活动照片与4位虚构报名头像，用于1–3图网格与头像堆叠；来源见manifest.json |
| src/static/posterwall | 旧交易页及部分新组局详情复用的示例摄影 |
| src/static/profile-a | A方案独立透明票券、个人头像、局主头像及桌游详情主图；生成记录见manifest.json |
| src/static/ui-icons | 主题生成的Phosphor功能图标及身份卡曲线接缝SVG |

首页、组局、旧服务及A方案人物与活动图均为AI生成的虚构示例素材；不代表真实用户、真实活动、实拍或真实参与记录。正式运营素材需单独确认授权。

A方案运行时素材为独立生成的资产，不是从批准的整页UI图裁切。`src/static/profile-a/manifest.json`记录各资产的提示词、尺寸、格式、透明性与生成来源；票券PNG保留透明背景，不含等级、徽章或获得权益的含义。个人与局主头像为虚构人物，桌游主图为虚构活动场景。

品牌字标当前仍有边缘瑕疵，按用户确认直接使用；后续可替换为清洁字标，不应静默重绘。

2026-10-14的新资产来自独立照片条和独立头像联系表的素材提取与压缩，不从整页设计图裁切UI。照片网格、标题、费用、头像布局均为真实组件。已有桌游独立主图继续用作详情封面。

## 字体

[Resource Han Rounded](https://github.com/CyanoHao/Resource-Han-Rounded/releases/tag/v0.990)，CN Bold子集，SIL OFL1.1。许可证保留在字体目录。脚本同时生成WOFF2和base64 TypeScript供uni.loadFontFace使用。

## 图标

导航及A方案功能图标来自[Phosphor](https://github.com/phosphor-icons/core)，MIT；保留其授权说明，旧页面仍有uni-ui图标。`scripts/generate-theme.mjs`从主题配置生成五个导航SVG的颜色，以及功能图标的蓝/灰蓝/近黑版本。身份卡曲线接缝为同脚本生成的项目自有SVG，白色填充同样读取主题；这些生成文件不可手改。

## 非运行时参考

`docs/design/card-layout-comparison.png`中B是已选卡片结构；三张旧详情概念图为历史参考，不是当前实现依据。`.impeccable/mocks/A-profile-detail.png`是用户批准的我的与组局详情参考，实际页面使用独立素材和真实组件，并保留原首页与导航；合同见同目录的 `a-contract.md`。`docs/brand/miniapp-icon.png`为保留的小程序对外图候选最终版，发布时仍需用户确认。

旧候选图、提示词和大原图已可恢复归档，见清理记录。不把设计稿当作页面静态资源。
