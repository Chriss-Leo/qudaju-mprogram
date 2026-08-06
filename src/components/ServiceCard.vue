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
.service-card { position: relative; min-width: 0; overflow: hidden; border-radius: 14rpx; background: #fffdf6; box-shadow: 0 12rpx 28rpx rgba(21, 63, 53, 0.1); }
.service-card::before, .service-card::after { content: ""; position: absolute; z-index: 3; top: 286rpx; width: 22rpx; height: 22rpx; border-radius: 50%; background: #f3f5f0; transform: translateY(-50%); }
.service-card::before { left: -11rpx; }.service-card::after { right: -11rpx; }
.service-card__body { padding: 22rpx 24rpx 24rpx; border-top: 2rpx dashed #c6d0c8; }
.service-card__title { display: block; color: #153f35; font-size: 29rpx; line-height: 1.35; font-weight: 900; }
.service-card__meta { display: flex; align-items: center; gap: 8rpx; margin-top: 14rpx; color: #627168; font-size: 21rpx; white-space: nowrap; }
.verified { padding: 5rpx 9rpx; color: #153f35; background: #b8cfbf; font-weight: 700; }
.service-card__tags { margin-top: 14rpx; display: flex; flex-wrap: wrap; gap: 9rpx; }
.service-card__tags text { padding: 7rpx 12rpx; color: #506158; background: #e7eee8; font-size: 20rpx; }
.service-card__footer { margin-top: 20rpx; display: flex; align-items: flex-end; justify-content: space-between; }
.rating { color: #d9503b; font-size: 22rpx; }
.muted { margin-left: 6rpx; color: #909992; }
.price { color: #e85f46; font-size: 32rpx; font-weight: 900; }
.price text { margin-left: 3rpx; font-size: 20rpx; font-weight: 500; }
.service-card--horizontal { display: grid; grid-template-columns: 238rpx minmax(0, 1fr); margin-bottom: 23rpx; border-radius: 13rpx; box-shadow: 0 10rpx 24rpx rgba(21, 63, 53, .09); background: #fffdf6; }
.service-card--horizontal::before, .service-card--horizontal::after { top: auto; bottom: -11rpx; }
.service-card--horizontal .service-card__body { padding: 18rpx 20rpx; border-top: 0; border-left: 2rpx dashed #c6d0c8; }
.service-card--horizontal .service-card__title { font-size: 30rpx; }
.service-card--horizontal .service-card__tags { display: none; }
.service-card--horizontal .service-card__meta { flex-wrap: wrap; white-space: normal; line-height: 1.6; }
.service-card--horizontal .service-card__footer { margin-top: 24rpx; }
</style>
