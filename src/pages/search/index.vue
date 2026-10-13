<script setup lang="ts">
import { theme } from "../../theme";
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

    <AppTabBar active="home" owner-route="pages/search/index" />

    <BottomSheet :visible="filterVisible" title="快速筛选" @close="filterVisible = false">
      <view class="quick-filter">
        <text>只看已认证服务者</text><switch :checked="certifiedOnly" :color="theme.colors.primary" @change="certifiedOnly = !certifiedOnly" />
      </view>
      <view class="quick-filter">
        <text>今天可约</text><switch :checked="todayOnly" :color="theme.colors.primary" @change="todayOnly = !todayOnly" />
      </view>
      <button class="apply-button" @tap="filterVisible = false">应用筛选</button>
    </BottomSheet>
  </view>
</template>

<style scoped lang="scss">
.search-page { min-height: 100vh; padding: calc(54rpx + env(safe-area-inset-top)) 34rpx 170rpx; background: $color-compatFFFEFA; }
.search-nav { display: grid; grid-template-columns: 66rpx minmax(0, 1fr) 66rpx; align-items: center; gap: 15rpx; }
.round-button { width: 66rpx; height: 66rpx; margin: 0; padding: 0; border-radius: $radius-v50percent; background: $color-compatF0F2EE; }
.back-button view { width: 20rpx; height: 20rpx; margin: 21rpx 0 0 25rpx; border: 4rpx solid $color-compat263129; border-top: 0; border-right: 0; transform: rotate(45deg); }
.search-field { height: 74rpx; padding: 0 14rpx 0 22rpx; display: flex; align-items: center; gap: 15rpx; border: 2rpx solid $color-compat183F34; border-radius: $radius-v26rpx; background: $color-white; }
.search-field input { min-width: 0; flex: 1; color: $color-compat263129; font-size: $font-size-v26rpx; }
.search-glyph { position: relative; width: 22rpx; height: 22rpx; flex: 0 0 auto; border: 4rpx solid $color-primary; border-radius: $radius-v50percent; }
.search-glyph::after { content: ""; position: absolute; width: 13rpx; height: 4rpx; right: -10rpx; bottom: -5rpx; background: $color-primary; transform: rotate(45deg); }
.clear-button { width: 46rpx; height: 46rpx; padding: 0; border-radius: $radius-v50percent; background: $color-compatE9EDE8; color: $color-compat6B776F; font-size: $font-size-v30rpx; line-height: $line-height-v42rpx; }
.cancel-button { width: auto; margin: 0; padding: 0; background: transparent; color: $color-compat354139; font-size: $font-size-v24rpx; }
.result-toolbar { margin-top: 40rpx; padding-bottom: 18rpx; display: flex; align-items: center; justify-content: space-between; border-bottom: 1rpx solid $color-compatE4E8E2; color: $color-compat858F88; font-size: $font-size-v22rpx; }
.result-toolbar__actions { display: flex; align-items: center; gap: 30rpx; }
.result-toolbar__actions > text.active { color: $color-primary; font-weight: $font-weight-v800; }
.filter-pill { height: 58rpx; margin: 0; padding: 0 17rpx; display: flex; align-items: center; gap: 8rpx; border-radius: $radius-v20rpx; background: $color-compatEEF2ED; color: $color-compat24332B; font-size: $font-size-v22rpx; }
.sliders { width: 24rpx; display: flex; flex-direction: column; gap: 6rpx; }
.sliders view { height: 3rpx; background: currentColor; }
.sliders view:last-child { width: 15rpx; align-self: flex-end; }
.result-list { margin-top: 4rpx; }
.empty-state { min-height: 760rpx; padding: 150rpx 40rpx 0; display: flex; flex-direction: column; align-items: center; text-align: center; }
.empty-state__illustration { position: relative; width: 190rpx; height: 150rpx; border-radius: $radius-v80rpx_80rpx_28rpx_28rpx; background: $color-compatE5EEE7; }
.empty-state__illustration > view:first-child { position: absolute; width: 70rpx; height: 70rpx; left: 60rpx; top: 26rpx; border: 8rpx solid $color-compat507565; border-radius: $radius-v50percent; }
.empty-state__illustration > view:last-child { position: absolute; width: 44rpx; height: 8rpx; right: 36rpx; bottom: 42rpx; border-radius: $radius-v8rpx; background: $color-compat507565; transform: rotate(45deg); }
.empty-state__title { margin-top: 34rpx; font-size: $font-size-v30rpx; font-weight: $font-weight-v800; }
.empty-state__copy { margin-top: 14rpx; color: $color-compat89938C; font-size: $font-size-v23rpx; line-height: $line-height-v1p6; }
.empty-state button { margin-top: 34rpx; height: 72rpx; padding: 0 28rpx; border-radius: $radius-v24rpx; color: $color-white; background: $color-primary; font-size: $font-size-v24rpx; }
.quick-filter { min-height: 94rpx; display: flex; align-items: center; justify-content: space-between; border-bottom: 1rpx solid $color-compatE5E9E4; font-size: $font-size-v26rpx; }
.apply-button { height: 86rpx; margin-top: 34rpx; border-radius: $radius-v28rpx; color: $color-white; background: $color-primary; font-size: $font-size-v27rpx; font-weight: $font-weight-v700; }
.search-page { padding: 0 24rpx 170rpx; background: $color-compatF3F5F0; }
.search-nav { margin: 0 -24rpx; padding: calc(42rpx + env(safe-area-inset-top)) 24rpx 27rpx; grid-template-columns: 60rpx minmax(0,1fr) 58rpx; color: $color-inverseWarm; background: $color-primary; }
.round-button { width: 58rpx; height: 58rpx; border: 2rpx solid $color-compatOverlay26; border-radius: $radius-v14rpx; background: transparent; }
.back-button view { margin: 17rpx 0 0 22rpx; border-color: $color-inverseWarm; }
.search-field { height: 70rpx; border: 3rpx solid $color-inverseWarm; border-radius: $radius-v999rpx; background: $color-inverseWarm; }
.cancel-button { color: $color-inverseWarm; }
.result-toolbar { position: sticky; z-index: 5; top: 0; margin: 0 -24rpx; padding: 22rpx 24rpx 18rpx; border-bottom: 1rpx solid $color-compatCBD4CD; background: $color-compatOverlay27; }
.filter-pill { border: 2rpx solid $color-primary; border-radius: $radius-v999rpx; color: $color-white; background: $color-primary; }
.result-list { margin-top: 23rpx; }
.empty-state { min-height: 660rpx; padding-top: 100rpx; border: 2rpx dashed $color-compatAEBDB3; background: $color-inverseWarm; }
.empty-state__illustration { border-radius: $radius-v8rpx; background: $color-compatB8CFBF; transform: rotate(-2deg); }
.empty-state button, .apply-button { border-radius: $radius-v14rpx; background: $color-primary; }
</style>
