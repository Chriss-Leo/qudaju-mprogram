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
.sheet-layer__mask { position: absolute; inset: 0; background: rgba(10, 23, 18, 0.48); }
.sheet {
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  max-height: 84vh;
  overflow-y: auto;
  padding: 16rpx 40rpx calc(34rpx + env(safe-area-inset-bottom));
  border-top: 8rpx solid #153f35;
  border-radius: 18rpx 18rpx 0 0;
  background: #fffdf6;
}
.sheet__handle { width: 80rpx; height: 6rpx; margin: 0 auto 20rpx; border-radius: 8rpx; background: #b8c7bd; }
.sheet__header { height: 72rpx; display: flex; align-items: center; justify-content: space-between; }
.sheet__title { color: #153f35; font-size: 34rpx; font-weight: 900; }
.sheet__close { width: 58rpx; height: 58rpx; padding: 0; border: 2rpx solid #c8d2ca; border-radius: 12rpx; background: #f3f5f0; color: #153f35; font-size: 38rpx; line-height: 52rpx; }
</style>
