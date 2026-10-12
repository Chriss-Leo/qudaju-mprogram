<script setup lang="ts">
import type { GroupCardData } from "../mocks/group-cards";
defineProps<{ group: GroupCardData; distance: string }>();
const emit = defineEmits<{ (event: "open"): void }>();
function preview(gallery: string) {
  uni.previewImage({ urls: [gallery], current: gallery });
}
</script>

<template>
  <view class="group-card" role="button" tabindex="0" :aria-label="`查看${group.title}`" @tap="emit('open')" @keydown.enter="emit('open')">
    <view class="group-card__host">
      <view class="group-card__avatar" aria-hidden="true"><text>{{ group.host.slice(0, 1) }}</text></view>
      <view class="group-card__identity">
        <text class="group-card__name">{{ group.host }}</text>
        <text class="group-card__context">{{ group.category }} · {{ group.admission }}</text>
      </view>
    </view>
    <text class="group-card__title">{{ group.title }}</text>
    <text class="group-card__description">{{ group.description }}</text>
    <image class="group-card__gallery" :src="group.gallery" mode="aspectFill" :aria-label="`${group.category}活动示例三图`" @tap.stop="preview(group.gallery)" />
    <view class="group-card__meta">
      <view><view class="group-card__clock" aria-hidden="true" /><text>{{ group.time }}</text></view>
      <view><uni-icons type="location" size="16" color="#66786f" /><text>{{ group.area }} · {{ distance }}</text></view>
    </view>
    <view class="group-card__footer">
      <text v-if="group.fund === 0" class="group-card__free">免费</text>
      <view v-else class="group-card__fund"><text>活动基金</text><text>¥{{ group.fund }}</text><text>/人</text></view>
      <view class="group-card__capacity"><text>已报名 </text><text class="group-card__joined">{{ group.joined }}</text><text>/{{ group.capacity }}人</text></view>
    </view>
  </view>
</template>

<style scoped lang="scss">
.group-card { padding: 30rpx; border-radius: 16rpx; color: #153f35; background: #fff; }
.group-card:focus-visible { outline: 3rpx solid #153f35; outline-offset: 3rpx; }
.group-card__host { display: flex; align-items: center; gap: 16rpx; margin-bottom: 22rpx; }
.group-card__avatar { flex: 0 0 72rpx; height: 72rpx; display: flex; align-items: center; justify-content: center; border-radius: 50%; background: #e6eee6; color: #153f35; font-size: 30rpx; font-weight: 600; }
.group-card__identity { display: flex; flex-direction: column; gap: 4rpx; }
.group-card__name { font-size: 29rpx; line-height: 1.35; font-weight: 600; }
.group-card__context { color: #66786f; font-size: 23rpx; line-height: 1.4; }
.group-card__title { display: block; font-size: 34rpx; line-height: 1.45; font-weight: 700; overflow-wrap: anywhere; }
.group-card__description { display: block; margin-top: 8rpx; color: #596d62; font-size: 26rpx; line-height: 1.6; }
.group-card__gallery { display: block; width: 100%; height: 207rpx; margin-top: 22rpx; border-radius: 8rpx; background: #e7ece7; }
.group-card__meta { display: flex; flex-wrap: wrap; gap: 10rpx 24rpx; margin: 22rpx 0; color: #66786f; font-size: 24rpx; line-height: 1.5; }
.group-card__meta > view { display: flex; align-items: center; gap: 7rpx; }
.group-card__clock { position: relative; width: 25rpx; height: 25rpx; flex: 0 0 25rpx; border: 2rpx solid currentColor; border-radius: 50%; }
.group-card__clock::after { content: ""; position: absolute; left: 11rpx; top: 4rpx; width: 6rpx; height: 9rpx; border-left: 2rpx solid currentColor; border-bottom: 2rpx solid currentColor; }
.group-card__footer { display: flex; flex-wrap: wrap; align-items: baseline; justify-content: space-between; gap: 12rpx; padding-top: 20rpx; border-top: 1rpx solid #e6ebe5; }
.group-card__free { font-size: 31rpx; font-weight: 700; }
.group-card__fund { display: flex; align-items: baseline; gap: 5rpx; color: #bf482b; font-size: 23rpx; }
.group-card__fund text:nth-child(2) { font-size: 32rpx; font-weight: 700; }
.group-card__capacity { color: #66786f; font-size: 24rpx; white-space: nowrap; }
.group-card__joined { color: #bf482b; font-size: 31rpx; font-weight: 600; }
</style>
