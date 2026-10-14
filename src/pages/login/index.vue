<script setup lang="ts">
import { ref } from "vue";
import { onLoad } from "@dcloudio/uni-app";
import SurfaceHeader from "../../components/SurfaceHeader.vue";
import UiIcon from "../../components/UiIcon.vue";
import { isAuthenticated, restoreAuth, wechatLogin } from "../../stores/auth";
import { safeReturnURL } from "../../utils/auth-navigation";
const next = ref("/pages/profile/index"); const busy = ref(false); const error = ref("");
onLoad(query => { next.value = safeReturnURL(String(query?.next || "/pages/profile/index")); });
function back() { uni.reLaunch({ url: "/pages/index/index" }); }
async function login() {
  if (busy.value) return; busy.value = true; error.value = "";
  try { if (!await restoreAuth()) await wechatLogin(); if (isAuthenticated.value) uni.redirectTo({ url: next.value }); }
  catch (value) { error.value = value instanceof Error ? value.message : "登录失败，请重试"; }
  finally { busy.value = false; }
}
</script>
<template><view class="login-page"><SurfaceHeader title="微信登录" back @back="back" /><view class="login-body"><view class="login-brand"><UiIcon name="people" /></view><text class="login-title">好搭子，好时光</text><text class="login-description">登录后发布组局、申请报名，<br />管理你的活动与聊天邀请。</text><button class="login-action action-accent" :disabled="busy" :loading="busy" @tap="login">{{ busy ? '正在登录' : '微信登录' }}</button><text v-if="error" class="login-error" role="alert">{{ error }}</text><text class="login-hint">仅用于建立账号身份，不自动获取微信头像、昵称或手机号。</text><button class="login-guest" @tap="back">先逛逛同城好局</button></view></view></template>
<style scoped lang="scss">
.login-page { min-height: 100vh; color: $color-text; background: $color-white; }
.login-body { padding: 100rpx 56rpx calc(40rpx + env(safe-area-inset-bottom)); }
.login-brand { display: flex; align-items: center; justify-content: center; width: 120rpx; height: 120rpx; border-radius: $radius-v32rpx; background: $color-tag; }
.login-brand :deep(.ui-icon) { width: 64rpx; height: 64rpx; }
.login-title { display: block; margin-top: 34rpx; font-size: $font-size-v42rpx; font-weight: $font-weight-v600; line-height: $line-height-v1p4; }
.login-description { display: block; margin-top: 20rpx; color: $color-textSecondary; font-size: $font-size-v28rpx; line-height: $line-height-v1p75; }
.login-action { width: 100%; min-height: 44px; padding: 26rpx 20rpx; margin-top: 70rpx; border-radius: $radius-v24rpx; font-size: $font-size-v30rpx; font-weight: $font-weight-v600; line-height: $line-height-v1p4; }
.login-error { display: block; margin-top: 24rpx; color: $color-textSecondary; font-size: $font-size-v26rpx; line-height: $line-height-v1p6; }
.login-hint { display: block; margin-top: 28rpx; color: $color-textMuted; font-size: $font-size-v23rpx; line-height: $line-height-v1p6; }
.login-guest { min-height: 44px; padding: 20rpx 0; margin: 30rpx 0 0; color: $color-primary; background: transparent; border-radius: $radius-v0; font-size: $font-size-v27rpx; line-height: $line-height-v1p4; }
</style>
