<script setup lang="ts">
import { ref } from "vue";
import { onLoad } from "@dcloudio/uni-app";
import SurfaceHeader from "../../components/SurfaceHeader.vue";
import { theme } from "../../theme";
import { isAuthenticated, wechatLogin } from "../../stores/auth";
import { safeReturnURL } from "../../utils/auth-navigation";
const next = ref("/pages/profile/index"); const busy = ref(false); const error = ref("");
onLoad(query => { next.value = safeReturnURL(String(query?.next || "/pages/profile/index")); });
function back() { uni.reLaunch({ url: "/pages/index/index" }); }
async function onGetPhoneNumber(event: { detail?: { code?: string; errMsg?: string; errno?: number } }) {
  if (busy.value) return;
  const detail = event.detail;
  if (!detail?.code || detail.errMsg !== "getPhoneNumber:ok") {
    const message = detail?.errMsg || "";
    if (/deny|cancel/i.test(message)) {
      error.value = "你已取消手机号授权，可以继续浏览，或再次点击登录。";
    } else if (/no permission|no privilege|permission denied/i.test(message) || detail?.errno === 102) {
      error.value = "当前小程序没有获取手机号的权限，请检查当前 AppID 的手机号能力。";
    } else if (message === "getPhoneNumber:ok" && !detail?.code) {
      error.value = "微信未返回手机号授权凭证，请使用新版微信真机重试。";
    } else {
      error.value = "未能获取手机号授权，请重试，并确认小程序已开通手机号能力。";
    }
    // Only platform diagnostics: never log the event, phone code, or encrypted data.
    if (import.meta.env.DEV) {
      console.warn("[wechat-phone] native authorization failed", { errMsg: message, errno: detail?.errno, hasCode: !!detail?.code });
      error.value += `（${message || "无错误详情"}${detail?.errno !== undefined ? `；错误码 ${detail.errno}` : ""}）`;
    }
    return;
  }
  busy.value = true; error.value = "";
  try { await wechatLogin(detail.code); if (isAuthenticated.value) uni.redirectTo({ url: next.value }); }
  catch (value) { error.value = value instanceof Error ? value.message : "登录失败，请重试"; }
  finally { busy.value = false; }
}
function unsupportedLogin() { error.value = "请在微信小程序中使用手机号一键登录。"; }
</script>
<template>
  <view class="login-page">
    <SurfaceHeader title="微信登录" back @back="back" />
    <view class="login-body">
      <view class="login-brand">
        <image class="login-brand__logo" src="/static/brand/qudaju-wordmark-orange-ivory-v1.png" mode="aspectFit" aria-label="趣搭局" />
        <text class="login-brand__tagline">好搭子 · 好时光</text>
      </view>
      <text class="login-description">登录后发布组局、申请报名，<br />管理你的活动与聊天邀请。</text>
      <!-- #ifdef MP-WEIXIN -->
      <button class="login-action action-accent" open-type="getPhoneNumber" :disabled="busy" @getphonenumber="onGetPhoneNumber">
        <view class="login-action__content"><uni-icons class="login-action__wechat" type="weixin" size="46rpx" :color="theme.colors.accentForeground" aria-hidden="true" /><text>微信手机号一键登录</text></view>
      </button>
      <!-- #endif -->
      <!-- #ifndef MP-WEIXIN -->
      <button class="login-action action-accent" @tap="unsupportedLogin"><view class="login-action__content"><uni-icons class="login-action__wechat" type="weixin" size="46rpx" :color="theme.colors.accentForeground" aria-hidden="true" /><text>微信手机号一键登录</text></view></button>
      <!-- #endif -->
      <text v-if="error" class="login-error" role="alert">{{ error }}</text>
      <text class="login-hint">授权手机号用于注册和登录，首次登录自动创建账号。头像与昵称可稍后在「我的」中设置。</text>
      <button class="login-guest" @tap="back">先逛逛同城好局</button>
    </view>
  </view>
</template>
<style scoped lang="scss">
.login-page { min-height: 100vh; color: $color-text; background: $color-white; }
.login-body { padding: 100rpx 56rpx calc(40rpx + env(safe-area-inset-bottom)); }
.login-brand { display: flex; flex-direction: column; align-items: center; gap: 13rpx; }
.login-brand__logo { display: block; width: 380rpx; max-width: 100%; height: 100rpx; }
.login-brand__tagline { padding-left: 7rpx; font-size: $font-size-v36rpx; line-height: $line-height-v1p4; font-weight: $font-weight-v700; letter-spacing: 7rpx; color: $color-accent; }
.login-description { display: block; margin-top: 48rpx; color: $color-textSecondary; font-size: $font-size-v28rpx; line-height: $line-height-v1p75; }
.login-action { width: 100%; min-height: 44px; padding: 26rpx 20rpx; margin-top: 70rpx; border-radius: $radius-v24rpx; font-size: $font-size-v30rpx; font-weight: $font-weight-v600; line-height: $line-height-v1p4; }
.login-error { display: block; margin-top: 24rpx; color: $color-textSecondary; font-size: $font-size-v26rpx; line-height: $line-height-v1p6; }
.login-action__content { display: flex; align-items: center; justify-content: center; gap: 16rpx; }
.login-action__wechat { display: flex; align-items: center; justify-content: center; width: 46rpx; height: 46rpx; flex: 0 0 auto; }
.login-hint { display: block; margin-top: 28rpx; color: $color-textMuted; font-size: $font-size-v23rpx; line-height: $line-height-v1p6; }
.login-guest { min-height: 44px; padding: 20rpx 0; margin: 30rpx 0 0; color: $color-primary; background: transparent; border-radius: $radius-v0; font-size: $font-size-v27rpx; line-height: $line-height-v1p4; }
</style>
