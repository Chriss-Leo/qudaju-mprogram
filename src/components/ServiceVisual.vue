<script setup lang="ts">
defineProps<{ type: "guide" | "photo" | "game" | "outdoor"; compact?: boolean }>();

const images = {
  guide: "/static/posterwall/city-photo-walk.jpg",
  photo: "/static/posterwall/outdoor-walk.jpg",
  game: "/static/posterwall/boardgame-night.jpg",
  outdoor: "/static/posterwall/hero-frisbee.jpg",
};
</script>

<template>
  <view class="visual" :class="[`visual--${type}`, { 'visual--compact': compact }]">
    <image class="visual__photo" :src="images[type]" mode="aspectFill" />
    <view class="visual__wash" />
    <view class="visual__badge"><text>{{ type === 'guide' ? '城市漫步' : type === 'photo' ? '记录此刻' : type === 'game' ? '轻松组局' : '一起出发' }}</text></view>
  </view>
</template>

<style scoped lang="scss">
.visual { position: relative; width: 100%; height: 286rpx; overflow: hidden; background: #9aaf9f; }
.visual--compact { height: 220rpx; border-radius: 12rpx; }
.visual__photo { width: 100%; height: 100%; display: block; }
.visual__wash { position: absolute; inset: 0; background: linear-gradient(180deg, transparent 45%, rgba(11, 45, 35, .62)); }
.visual__badge { position: absolute; left: 17rpx; bottom: 15rpx; padding: 8rpx 13rpx; color: #153f35; background: #b8cfbf; font-size: 19rpx; font-weight: 800; transform: rotate(-1.5deg); }
</style>
