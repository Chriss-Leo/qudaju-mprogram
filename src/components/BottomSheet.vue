<script setup lang="ts">
defineProps<{ visible: boolean; title: string }>();
const emit = defineEmits<{ close: [] }>();
</script>

<template>
  <view v-if="visible" class="sheet-layer" @touchmove.stop.prevent>
    <view class="sheet-layer__mask" @tap="emit('close')" />
    <view class="sheet">
      <view class="sheet__handle" />
      <view class="sheet__header">
        <text class="sheet__title">{{ title }}</text>
        <button class="sheet__close" aria-label="关闭" @tap="emit('close')">×</button>
      </view>
      <slot />
    </view>
  </view>
</template>

<style scoped lang="scss">
.sheet-layer { position: fixed; z-index: 80; inset: 0; }
.sheet-layer__mask { position: absolute; inset: 0; background: $color-compatOverlay7; }
.sheet {
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  max-height: 84vh;
  overflow-y: auto;
  padding: 16rpx 40rpx calc(34rpx + env(safe-area-inset-bottom));
  border-top: 8rpx solid $color-primary;
  border-radius: $radius-v18rpx_18rpx_0_0;
  background: $color-inverseWarm;
}
.sheet__handle { width: 80rpx; height: 6rpx; margin: 0 auto 20rpx; border-radius: $radius-v8rpx; background: $color-borderStrong; }
.sheet__header { height: 72rpx; display: flex; align-items: center; justify-content: space-between; }
.sheet__title { color: $color-text; font-size: $font-size-v34rpx; font-weight: $font-weight-v900; }
.sheet__close { width: 58rpx; height: 58rpx; padding: 0; border: 2rpx solid $color-border; border-radius: $radius-v12rpx; background: $color-compatF3F5F0; color: $color-text; font-size: $font-size-v38rpx; line-height: $line-height-v52rpx; }
/* #ifdef H5 */
@media (min-width: 600px) {
  .sheet { left: 50%; right: auto; width: 960px; max-height: calc(84vh / var(--qudaju-preview-scale, 1)); transform: translateX(-50%) scale(var(--qudaju-preview-scale, 1)); transform-origin: bottom center; }
}
/* #endif */
</style>
