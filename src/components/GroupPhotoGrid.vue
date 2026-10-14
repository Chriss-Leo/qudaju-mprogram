<script setup lang="ts">
import { computed, ref } from "vue";
const props = defineProps<{ photos: string[]; label: string }>();
const visible = computed(() => props.photos.slice(0, 3));
const failed = ref<Record<string, boolean>>({});
function preview(photo: string) { uni.previewImage({ urls: props.photos.slice(0, 3), current: photo }); }
</script>

<template>
  <view v-if="visible.length" class="photo-grid" :class="`photo-grid--${visible.length}`">
    <button v-for="(photo, index) in visible" :key="photo" class="photo-grid__item" :aria-label="`查看${label}第${index + 1}张照片`" @tap.stop="preview(photo)">
      <image v-if="!failed[photo]" :src="photo" mode="aspectFill" @error="failed[photo] = true" />
      <text v-else>图片暂未加载</text>
    </button>
  </view>
</template>

<style scoped lang="scss">
.photo-grid { display: grid; grid-template-columns: repeat(3, minmax(0, 1fr)); gap: 8rpx; }
.photo-grid--1 { grid-template-columns: minmax(0, 1fr); }
.photo-grid--2 { grid-template-columns: repeat(2, minmax(0, 1fr)); }
.photo-grid__item { position: relative; display: flex; align-items: center; justify-content: center; width: 100%; aspect-ratio: 1; margin: 0; padding: 0; overflow: hidden; border: 0; border-radius: $radius-v16rpx; background: $color-card; color: $color-textMuted; font-size: $font-size-v22rpx; line-height: $line-height-v1p5; }
.photo-grid--1 .photo-grid__item { aspect-ratio: 3 / 2; }
.photo-grid__item image { position: absolute; inset: 0; width: 100%; height: 100%; }
.photo-grid__item::after { border: 0; }
.photo-grid__item:active { opacity: .85; }
.photo-grid__item:focus-visible { outline: 3rpx solid $color-primary; outline-offset: 2rpx; }
</style>
