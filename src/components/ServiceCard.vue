<script setup lang="ts">
import type { ServiceItem } from "../types/service";
import ServiceVisual from "./ServiceVisual.vue";

defineProps<{ item: ServiceItem; horizontal?: boolean }>();
const emit = defineEmits<{ select: [id: string] }>();
</script>

<template>
  <view class="service-card" :class="{ 'service-card--horizontal': horizontal }" @tap="emit('select', item.id)">
    <ServiceVisual :type="item.visual" :compact="horizontal" />
    <view class="service-card__body">
      <text class="service-card__title">{{ item.title }}</text>
      <view class="service-card__meta">
        <text class="verified">已认证</text>
        <text>{{ item.provider }}</text>
        <text>· {{ item.distance }}</text>
      </view>
      <view class="service-card__tags">
        <text v-for="tag in item.tags" :key="tag">{{ tag }}</text>
      </view>
      <view class="service-card__footer">
        <text class="rating">★ {{ item.rating }} <text class="muted">{{ item.orders }}次服务</text></text>
        <text class="price">¥{{ item.price }}<text>起</text></text>
      </view>
    </view>
  </view>
</template>

<style scoped lang="scss">
.service-card { position: relative; min-width: 0; overflow: hidden; border-radius: $radius-v14rpx; background: $color-card; box-shadow: 0 12rpx 28rpx $color-compatOverlay5; }
.service-card::before, .service-card::after { content: ""; position: absolute; z-index: 3; top: 286rpx; width: 22rpx; height: 22rpx; border-radius: $radius-v50percent; background: $color-compatF3F5F0; transform: translateY(-50%); }
.service-card::before { left: -11rpx; }.service-card::after { right: -11rpx; }
.service-card__body { padding: 22rpx 24rpx 24rpx; border-top: 2rpx dashed $color-compatC6D0C8; }
.service-card__title { display: block; color: $color-primary; font-size: $font-size-v29rpx; line-height: $line-height-v1p35; font-weight: $font-weight-v900; }
.service-card__meta { display: flex; align-items: center; gap: 8rpx; margin-top: 14rpx; color: $color-compat627168; font-size: $font-size-v21rpx; white-space: nowrap; }
.verified { padding: 5rpx 9rpx; color: $color-primary; background: $color-compatB8CFBF; font-weight: $font-weight-v700; }
.service-card__tags { margin-top: 14rpx; display: flex; flex-wrap: wrap; gap: 9rpx; }
.service-card__tags text { padding: 7rpx 12rpx; color: $color-compat506158; background: $color-compatE7EEE8; font-size: $font-size-v20rpx; }
.service-card__footer { margin-top: 20rpx; display: flex; align-items: flex-end; justify-content: space-between; }
.rating { color: $color-compatD9503B; font-size: $font-size-v22rpx; }
.muted { margin-left: 6rpx; color: $color-compat909992; }
.price { color: $color-compatE85F46; font-size: $font-size-v32rpx; font-weight: $font-weight-v900; }
.price text { margin-left: 3rpx; font-size: $font-size-v20rpx; font-weight: $font-weight-v500; }
.service-card--horizontal { display: grid; grid-template-columns: 238rpx minmax(0, 1fr); margin-bottom: 23rpx; border-radius: $radius-v13rpx; box-shadow: 0 10rpx 24rpx $color-compatOverlay6; background: $color-inverseWarm; }
.service-card--horizontal::before, .service-card--horizontal::after { top: auto; bottom: -11rpx; }
.service-card--horizontal .service-card__body { padding: 18rpx 20rpx; border-top: 0; border-left: 2rpx dashed $color-compatC6D0C8; }
.service-card--horizontal .service-card__title { font-size: $font-size-v30rpx; }
.service-card--horizontal .service-card__tags { display: none; }
.service-card--horizontal .service-card__meta { flex-wrap: wrap; white-space: normal; line-height: $line-height-v1p6; }
.service-card--horizontal .service-card__footer { margin-top: 24rpx; }
</style>
