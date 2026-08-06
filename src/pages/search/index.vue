<script setup lang="ts">
import { computed, ref } from "vue";
import { onLoad } from "@dcloudio/uni-app";
import AppTabBar from "../../components/AppTabBar.vue";
import BottomSheet from "../../components/BottomSheet.vue";
import ServiceCard from "../../components/ServiceCard.vue";
import { services } from "../../mocks/services";

const keyword = ref("城市向导");
const sort = ref("推荐");
const filterVisible = ref(false);
const certifiedOnly = ref(true);
const todayOnly = ref(false);

onLoad((query) => {
  if (query?.keyword) keyword.value = decodeURIComponent(query.keyword);
});

const filtered = computed(() => {
  let result = !keyword.value || ["城市向导", "更多分类"].includes(keyword.value)
    ? [...services]
    : services.filter((item) => item.title.includes(keyword.value) || item.tags.some((tag) => tag.includes(keyword.value)));
  if (todayOnly.value) result = result.filter((item) => item.available.includes("今天"));
  if (sort.value === "距离") result.sort((a, b) => Number.parseFloat(a.distance) - Number.parseFloat(b.distance));
  return result;
});

function back() {
  uni.navigateBack({ fail: () => uni.redirectTo({ url: "/pages/index/index" }) });
}
function openService(id: string) {
  uni.navigateTo({ url: `/pages/service-detail/index?id=${id}` });
}
</script>

<template>
  <view class="search-page">
    <view class="search-nav">
      <button class="round-button back-button" aria-label="返回" @tap="back"><view /></button>
      <view class="search-field">
        <view class="search-glyph" />
        <input v-model="keyword" confirm-type="search" placeholder="搜索服务或服务者" />
        <button v-if="keyword" class="clear-button" aria-label="清空" @tap="keyword = ''">×</button>
      </view>
      <button class="cancel-button" @tap="back">取消</button>
    </view>

    <view class="result-toolbar">
      <text>共 {{ filtered.length || 0 }} 个结果</text>
      <view class="result-toolbar__actions">
        <text v-for="item in ['推荐', '距离']" :key="item" :class="{ active: sort === item }" @tap="sort = item">{{ item }}</text>
        <button class="filter-pill" @tap="filterVisible = true"><view class="sliders"><view /><view /></view>筛选</button>
      </view>
    </view>

    <view v-if="filtered.length" class="result-list">
      <ServiceCard v-for="item in filtered" :key="item.id" :item="item" horizontal @select="openService" />
    </view>
    <view v-else class="empty-state">
      <view class="empty-state__illustration"><view /><view /></view>
      <text class="empty-state__title">暂时没有匹配的服务</text>
      <text class="empty-state__copy">换个关键词，或者放宽区域与时间条件试试</text>
      <button @tap="keyword = ''">查看全部服务</button>
    </view>

    <AppTabBar active="home" />

    <BottomSheet :visible="filterVisible" title="快速筛选" @close="filterVisible = false">
      <view class="quick-filter">
        <text>只看已认证服务者</text><switch :checked="certifiedOnly" color="#173f34" @change="certifiedOnly = !certifiedOnly" />
      </view>
      <view class="quick-filter">
        <text>今天可约</text><switch :checked="todayOnly" color="#173f34" @change="todayOnly = !todayOnly" />
      </view>
      <button class="apply-button" @tap="filterVisible = false">应用筛选</button>
    </BottomSheet>
  </view>
</template>

<style scoped lang="scss">
.search-page { min-height: 100vh; padding: calc(54rpx + env(safe-area-inset-top)) 34rpx 170rpx; background: #fffefa; }
.search-nav { display: grid; grid-template-columns: 66rpx minmax(0, 1fr) 66rpx; align-items: center; gap: 15rpx; }
.round-button { width: 66rpx; height: 66rpx; margin: 0; padding: 0; border-radius: 50%; background: #f0f2ee; }
.back-button view { width: 20rpx; height: 20rpx; margin: 21rpx 0 0 25rpx; border: 4rpx solid #263129; border-top: 0; border-right: 0; transform: rotate(45deg); }
.search-field { height: 74rpx; padding: 0 14rpx 0 22rpx; display: flex; align-items: center; gap: 15rpx; border: 2rpx solid #183f34; border-radius: 26rpx; background: #fff; }
.search-field input { min-width: 0; flex: 1; color: #263129; font-size: 26rpx; }
.search-glyph { position: relative; width: 22rpx; height: 22rpx; flex: 0 0 auto; border: 4rpx solid #173f34; border-radius: 50%; }
.search-glyph::after { content: ""; position: absolute; width: 13rpx; height: 4rpx; right: -10rpx; bottom: -5rpx; background: #173f34; transform: rotate(45deg); }
.clear-button { width: 46rpx; height: 46rpx; padding: 0; border-radius: 50%; background: #e9ede8; color: #6b776f; font-size: 30rpx; line-height: 42rpx; }
.cancel-button { width: auto; margin: 0; padding: 0; background: transparent; color: #354139; font-size: 24rpx; }
.result-toolbar { margin-top: 40rpx; padding-bottom: 18rpx; display: flex; align-items: center; justify-content: space-between; border-bottom: 1rpx solid #e4e8e2; color: #858f88; font-size: 22rpx; }
.result-toolbar__actions { display: flex; align-items: center; gap: 30rpx; }
.result-toolbar__actions > text.active { color: #173f34; font-weight: 800; }
.filter-pill { height: 58rpx; margin: 0; padding: 0 17rpx; display: flex; align-items: center; gap: 8rpx; border-radius: 20rpx; background: #eef2ed; color: #24332b; font-size: 22rpx; }
.sliders { width: 24rpx; display: flex; flex-direction: column; gap: 6rpx; }
.sliders view { height: 3rpx; background: currentColor; }
.sliders view:last-child { width: 15rpx; align-self: flex-end; }
.result-list { margin-top: 4rpx; }
.empty-state { min-height: 760rpx; padding: 150rpx 40rpx 0; display: flex; flex-direction: column; align-items: center; text-align: center; }
.empty-state__illustration { position: relative; width: 190rpx; height: 150rpx; border-radius: 80rpx 80rpx 28rpx 28rpx; background: #e5eee7; }
.empty-state__illustration > view:first-child { position: absolute; width: 70rpx; height: 70rpx; left: 60rpx; top: 26rpx; border: 8rpx solid #507565; border-radius: 50%; }
.empty-state__illustration > view:last-child { position: absolute; width: 44rpx; height: 8rpx; right: 36rpx; bottom: 42rpx; border-radius: 8rpx; background: #507565; transform: rotate(45deg); }
.empty-state__title { margin-top: 34rpx; font-size: 30rpx; font-weight: 800; }
.empty-state__copy { margin-top: 14rpx; color: #89938c; font-size: 23rpx; line-height: 1.6; }
.empty-state button { margin-top: 34rpx; height: 72rpx; padding: 0 28rpx; border-radius: 24rpx; color: #fff; background: #173f34; font-size: 24rpx; }
.quick-filter { min-height: 94rpx; display: flex; align-items: center; justify-content: space-between; border-bottom: 1rpx solid #e5e9e4; font-size: 26rpx; }
.apply-button { height: 86rpx; margin-top: 34rpx; border-radius: 28rpx; color: #fff; background: #173f34; font-size: 27rpx; font-weight: 700; }
.search-page { padding: 0 24rpx 170rpx; background: #f3f5f0; }
.search-nav { margin: 0 -24rpx; padding: calc(42rpx + env(safe-area-inset-top)) 24rpx 27rpx; grid-template-columns: 60rpx minmax(0,1fr) 58rpx; color: #fffdf6; background: #0f4838; }
.round-button { width: 58rpx; height: 58rpx; border: 2rpx solid rgba(255,253,246,.45); border-radius: 14rpx; background: transparent; }
.back-button view { margin: 17rpx 0 0 22rpx; border-color: #fffdf6; }
.search-field { height: 70rpx; border: 3rpx solid #fffdf6; border-radius: 999rpx; background: #fffdf6; }
.cancel-button { color: #fffdf6; }
.result-toolbar { position: sticky; z-index: 5; top: 0; margin: 0 -24rpx; padding: 22rpx 24rpx 18rpx; border-bottom: 1rpx solid #cbd4cd; background: rgba(243,245,240,.97); }
.filter-pill { border: 2rpx solid #153f35; border-radius: 999rpx; color: #fff; background: #153f35; }
.result-list { margin-top: 23rpx; }
.empty-state { min-height: 660rpx; padding-top: 100rpx; border: 2rpx dashed #aebdb3; background: #fffdf6; }
.empty-state__illustration { border-radius: 8rpx; background: #b8cfbf; transform: rotate(-2deg); }
.empty-state button, .apply-button { border-radius: 14rpx; background: #153f35; }
</style>
