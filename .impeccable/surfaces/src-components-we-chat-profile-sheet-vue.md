---
version: 1
slug: "src-components-we-chat-profile-sheet-vue"
primary_target: "src/components/WeChatProfileSheet.vue"
related_targets: ["src/pages/login/index.vue", "src/pages/profile/index.vue"]
---

# 微信头像昵称确认弹层

2026-10-15，Operate 模式。用户明确要求微信原生 chooseAvatar 和 nickname 能力，并在底部弹出确认，不跳独立资料页。

复用 BottomSheet：白色底、共享圆角、蓝色头像操作、橙色白字「确认并继续」，全部主题值来自 src/theme/tokens.json。不改主页和导航设计。

微信按钮使用 open-type="chooseAvatar"，回调读取 e.detail.avatarUrl 临时路径并预览；微信输入框使用 type="nickname"、v-model、blur/change 及 form submit 取值，支持用户选用微信昵称。非微信只保留普通图片选择和昵称输入，不伪造原生微信授权。

登录流程已改为原生手机号一键注册，不再打开本弹层。「我的 → 编辑资料」自愿打开，回填已有头像昵称。确认时 uni.uploadFile 上传文件与昵称，已有头像且仅改昵称时不重复上传。H5的BottomSheet与导航同样挂载到body，避免导航覆盖保存操作；微信不使用Teleport。

空值、头像加载失败、提交中和错误重试均可见；保存失败保留输入，保存中阻止重复提交和关闭。退出登录关闭弹层，迟到的保存结果不能恢复已退出会话。

仍需微信开发者工具与真机验证原生选择、昵称建议、键盘和安全区；H5替身接口测试不代表原生能力验收。不会自动无操作读取微信用户资料，也不调用旧 getUserProfile。
