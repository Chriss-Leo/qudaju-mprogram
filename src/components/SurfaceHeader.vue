<script setup lang="ts">
import UiIcon from "./UiIcon.vue";
defineProps<{ title: string; back?: boolean; messages?: boolean }>();
const emit = defineEmits<{ (event: "back"): void; (event: "messages"): void }>();
const info = uni.getWindowInfo();
let capsule: ReturnType<typeof uni.getMenuButtonBoundingClientRect> | null = null;
try {
  const bounds = typeof uni.getMenuButtonBoundingClientRect === "function" ? uni.getMenuButtonBoundingClientRect() : null;
  if (bounds && bounds.top >= 0 && bounds.height > 0 && bounds.width > 0 && bounds.left > 0) capsule = bounds;
} catch { /* Unsupported platforms use the status-bar fallback. */ }
const layout = {
  "--header-top": `${capsule ? capsule.top : (info.statusBarHeight || 0) + 12}px`,
  "--header-height": `${capsule?.height || Math.max(44, Math.min(info.windowWidth, 960) * 84 / 750)}px`,
  "--header-reserve": `${capsule ? Math.max(16, info.windowWidth - capsule.left + 12) : 0}px`,
};
</script>

<template>
  <view class="surface-header" :style="layout">
    <view class="surface-header__row">
      <button v-if="back" class="header-button header-button--back" aria-label="返回" @tap="emit('back')"><UiIcon name="back" tone="black" /></button>
      <text class="surface-header__title" :class="{ 'surface-header__title--detail': back }">{{ title }}</text>
      <button v-if="messages" class="header-button" aria-label="消息" @tap="emit('messages')"><UiIcon name="message" tone="black" /></button>
    </view>
  </view>
</template>

<style scoped lang="scss">
.surface-header { padding: var(--header-top) 28rpx 28rpx; background: $color-white; }
.surface-header__row { position: relative; height: var(--header-height); min-height: var(--header-height); padding-right: var(--header-reserve); display: flex; align-items: center; gap: 12rpx; }
.surface-header__title { flex: 1; font-size: $font-size-v42rpx; font-weight: $font-weight-v700; line-height: $line-height-v1p3; color: $color-text; }
.surface-header__title--detail { position: absolute; left: 84rpx; right: 84rpx; font-size: $font-size-v32rpx; text-align: center; overflow: hidden; text-overflow: ellipsis; white-space: nowrap; }
.header-button { display: flex; justify-content: center; align-items: center; margin: 0; padding: 0; width: 72rpx; height: 80rpx; background: transparent; border-radius: $radius-v16rpx; }
.header-button:active { background: $color-card; }
.header-button { border: 0; }
.header-button--back { width: 84rpx; height: 84rpx; justify-content: flex-start; }
/* Compensate the caret glyph's built-in inset, aligning its visible tip to the page gutter. */
.header-button--back :deep(.ui-icon) { transform: translateX(-12rpx); }
.header-button::after { border: 0; }
.header-button:focus-visible { outline: 3rpx solid $color-primary; }
</style>
