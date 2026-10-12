# 素材与授权

## 运行时素材

| 目录 | 用途 |
| --- | --- |
| src/static/brand | 当前白橙透明字标 |
| src/static/entry/city-walk.jpg | 首页摄影底图 |
| src/static/entry/fonts | 资源圆体Bold标题子集、OFL许可证 |
| src/static/entry/icons | 五个固定导航SVG及授权说明 |
| src/static/groups | 三组三图条示例摄影 |
| src/static/posterwall | 四张旧详情/订单仍引用的摄影 |

组局和旧服务照片均为AI生成的示例素材；不代表真实用户、真实活动或授权肖像。正式运营素材需单独确认授权。

品牌字标当前仍有边缘瑕疵，按用户确认直接使用；后续可替换为清洁字标，不应静默重绘。

## 字体

[Resource Han Rounded](https://github.com/CyanoHao/Resource-Han-Rounded/releases/tag/v0.990)，CN Bold子集，SIL OFL1.1。许可证保留在字体目录。脚本同时生成WOFF2和base64 TypeScript供uni.loadFontFace使用。

## 图标

导航来自[Phosphor](https://github.com/phosphor-icons/core)，MIT；局部功能图标使用uni-ui。导航导出脚本仅生成当前使用的五个SVG。

## 非运行时参考

`docs/design/card-layout-comparison.png`中B是已选卡片结构；三张详情概念图待实现。`docs/brand/miniapp-icon.png`为保留的小程序对外图候选最终版，发布时仍需用户确认。

旧候选图、提示词和大原图已可恢复归档，见清理记录。不把设计稿当作页面静态资源。
