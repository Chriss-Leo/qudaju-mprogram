<script setup lang="ts">
import { computed, watch } from "vue";
import { onShow } from "@dcloudio/uni-app";
import { authState, isAuthenticated, restoreAuth } from "../stores/auth";
import { openLogin } from "../utils/auth-navigation";
const props = defineProps<{ next?: string }>();
const next = computed(() => props.next || (() => {
  const page = getCurrentPages().slice(-1)[0] as { route?: string; options?: Record<string,string> } | undefined;
  const path = page?.route || "pages/profile/index";
  const id = page?.options?.id;
  // H5 page cache may not expose options; the current hash contains its route.
  if (typeof window !== "undefined") { const value = window.location.hash.replace(/^#/, ""); if (value.startsWith("/pages/")) return value; }
  return `/${path}${id ? `?id=${encodeURIComponent(id)}` : ""}`;
})());
function login() { openLogin(next.value); }
onShow(async () => { if (!await restoreAuth()) login(); });
watch(isAuthenticated, (value, previous) => { if (!value && previous) login(); });
</script>
<template><slot v-if="isAuthenticated" /><view v-else-if="authState.status !== 'checking'" class="auth-block"><text>登录后继续</text><text>使用微信登录，管理你的报名、发布与消息。</text><button class="auth-block__button action-accent" @tap="login">微信登录</button></view></template>
<style scoped lang="scss">
.auth-block { min-height: 100vh; padding: calc(120rpx + env(safe-area-inset-top)) 40rpx 40rpx; background: $color-white; color: $color-text; text-align: center; }
.auth-block > text:first-child { display: block; font-size: $font-size-v38rpx; font-weight: $font-weight-v600; line-height: $line-height-v1p4; }
.auth-block > text:nth-child(2) { display: block; margin-top: 22rpx; color: $color-textMuted; font-size: $font-size-v27rpx; line-height: $line-height-v1p6; }
.auth-block__button { margin-top: 40rpx; min-height: 44px; padding: 20rpx; border-radius: $radius-v24rpx; font-size: $font-size-v28rpx; line-height: $line-height-v1p4; }
</style>
