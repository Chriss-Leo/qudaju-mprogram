<script setup lang="ts">
import UiIcon from "./UiIcon.vue";
defineProps<{ title: string; back?: boolean; messages?: boolean }>();
const emit = defineEmits<{ (event: "back"): void; (event: "messages"): void }>();
const info = uni.getWindowInfo();
const capsule = typeof uni.getMenuButtonBoundingClientRect === "function" ? uni.getMenuButtonBoundingClientRect() : null;
const layout = {
  "--header-top": `${capsule?.top || info.statusBarHeight || 12}px`,
  "--header-height": `${capsule?.height || 36}px`,
  "--header-reserve": `${capsule ? Math.max(16, info.windowWidth - capsule.left + 12) : 0}px`,
};
</script>

<template>
  <view class="surface-header" :style="layout">
    <view class="surface-header__row">
      <button v-if="back" class="header-button" aria-label="返回" @tap="emit('back')"><UiIcon name="back" tone="black" /></button>
      <text class="surface-header__title" :class="{ 'surface-header__title--detail': back }">{{ title }}</text>
      <button v-if="messages" class="header-button" aria-label="消息" @tap="emit('messages')"><UiIcon name="message" tone="black" /></button>
    </view>
  </view>
</template>

<style scoped lang="scss">
.surface-header { padding: max(var(--header-top), env(safe-area-inset-top)) 28rpx 28rpx; background: $color-white; }
.surface-header__row { min-height: max(var(--header-height), 84rpx); padding-right: var(--header-reserve); display: flex; align-items: center; gap: 12rpx; }
.surface-header__title { flex: 1; font-size: $font-size-v42rpx; font-weight: $font-weight-v700; line-height: $line-height-v1p3; color: $color-text; }
.surface-header__title--detail { font-size: $font-size-v32rpx; text-align: center; padding-right: 72rpx; }
.header-button { display: flex; justify-content: center; align-items: center; margin: 0; padding: 0; width: 72rpx; height: 80rpx; background: transparent; border-radius: $radius-v16rpx; }
.header-button:active { background: $color-card; }
.header-button { border: 0; }
.header-button::after { border: 0; }
.header-button:focus-visible { outline: 3rpx solid $color-primary; }
</style>
