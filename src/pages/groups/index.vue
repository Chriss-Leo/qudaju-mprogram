<script setup lang="ts">
import { computed, onUnmounted, ref, watch } from "vue";
import { onResize, onShow } from "@dcloudio/uni-app";
import AppTabBar from "../../components/AppTabBar.vue";
import GroupFeedCard from "../../components/GroupFeedCard.vue";
import { groupCards } from "../../mocks/group-cards";
import BottomSheet from "../../components/BottomSheet.vue";
import { services } from "../../mocks/services";

const locationVisible = ref(false);
const filterVisible = ref(false);
const currentArea = ref("静安");
const currentCity = ref("上海");
const selectedRegion = ref<string[]>(["上海市", "上海市", "静安区"]);
const quickFilter = ref("全部");
const blankFilters = () => ({ keyword: "", type: "不限", fee: "不限", distance: "不限", time: "不限", admission: "不限" });
const appliedFilters = ref(blankFilters());
const draftFilters = ref(blankFilters());
const filterGroups = [
  { key: "type" as const, label: "类型", options: ["不限", "桌游", "陪拍", "向导", "户外", "交友"] },
  { key: "fee" as const, label: "费用", options: ["不限", "免费", "收费"] },
  { key: "distance" as const, label: "距离", options: ["不限", "1km内", "3km内", "5km内"] },
  { key: "time" as const, label: "时间", options: ["不限", "今天", "明天", "本周"] },
  { key: "admission" as const, label: "报名方式", options: ["不限", "局主审核", "直接加入"] },
];
const areaApplied = ref(false);
const hasSearchFilters = computed(() => Object.entries(appliedFilters.value).some(([key, value]) => key === "keyword" ? !!value.trim() : value !== "不限"));
// #ifdef H5
let previousBodyOverflow = "";
watch(filterVisible, (visible) => {
  if (visible) { previousBodyOverflow = document.body.style.overflow; document.body.style.overflow = "hidden"; }
  else document.body.style.overflow = previousBodyOverflow;
});
onUnmounted(() => { if (filterVisible.value) document.body.style.overflow = previousBodyOverflow; });
// #endif
const windowInfo = uni.getWindowInfo();
const pageHeight = ref(windowInfo.windowHeight);
onResize(({ size }) => { if (size) pageHeight.value = size.windowHeight; });
onShow(() => {
  uni.setNavigationBarColor({ frontColor: "#000000", backgroundColor: "#fbfcf8" });
});

function explore() {
  uni.showToast({ title: "发布组局功能开发中", icon: "none" });
}
const statusBarHeight = windowInfo.statusBarHeight || 0;
const menuButtonRect = typeof uni.getMenuButtonBoundingClientRect === "function"
  ? uni.getMenuButtonBoundingClientRect()
  : null;
const menuTop = menuButtonRect?.top || statusBarHeight;
const menuRightReserve = menuButtonRect
  ? Math.max(12, windowInfo.windowWidth - menuButtonRect.left + 8)
  : 17;
const brandBarStyle = {
  "--menu-button-top": `${menuTop}px`,
  "--menu-right-reserve": `${menuRightReserve}px`,
};

const areas = ["全部区域", "徐汇区", "浦东新区", "黄浦区", "静安区", "长宁区", "闵行区", "普陀区"];

const displayedServices = computed(() => {
  const feedOrder = ["game-01", "photo-01", "outdoor-01", "city-guide-01", "guide-02"];
  let result = [...services].sort((a, b) => feedOrder.indexOf(a.id) - feedOrder.indexOf(b.id));
  const filters = appliedFilters.value;
  result = result.filter((item) => {
    const group = groupCards[item.id];
    if (filters.type !== "不限" && group.category !== filters.type) return false;
    if (filters.fee === "免费" && group.fund !== 0) return false;
    if (filters.fee === "收费" && group.fund === 0) return false;
    if (filters.admission !== "不限" && group.admission !== filters.admission) return false;
    if (filters.distance !== "不限" && parseFloat(item.distance) > parseFloat(filters.distance)) return false;
    const requestedTime = quickFilter.value === "本周" ? "本周" : filters.time;
    if (requestedTime !== "不限" && !item.available.includes(requestedTime === "本周" ? "本周末" : requestedTime as "今天" | "明天")) return false;
    const keyword = filters.keyword.trim().toLowerCase();
    if (keyword && ![group.title, group.description, group.host, group.category, ...item.tags].join(" ").toLowerCase().includes(keyword)) return false;
    const regions: Record<string, string> = { "game-01": "静安", "photo-01": "徐汇", "outdoor-01": "静安", "city-guide-01": "徐汇", "guide-02": "静安" };
    if (areaApplied.value && currentCity.value !== "上海") return false;
    if (areaApplied.value && currentArea.value !== "上海" && regions[item.id] !== currentArea.value) return false;
    return true;
  });
  if (quickFilter.value === "距离") result.sort((a, b) => parseFloat(a.distance) - parseFloat(b.distance));
  return result;
});

function openSearch() {
  draftFilters.value = { ...appliedFilters.value };
  filterVisible.value = true;
}
function openService(id: string) { uni.navigateTo({ url: `/pages/service-detail/index?id=${id}` }); }
function chooseArea(area: string) {
  currentCity.value = "上海";
  currentArea.value = area === "全部区域" ? "上海" : area.replace("区", "");
  selectedRegion.value = ["上海市", "上海市", area === "全部区域" ? "静安区" : area];
  areaApplied.value = true;
  locationVisible.value = false;
}
function changeRegion(event: { detail: { value: string[] } }) {
  const value = event.detail.value;
  if (value.length !== 3) return;
  selectedRegion.value = [...value];
  currentCity.value = (value[1] || value[0]).replace(/市$/, "");
  currentArea.value = value[2].replace(/区$/, "");
  areaApplied.value = true;
}
function selectQuick(value: string) {
  quickFilter.value = value;
  if (value === "全部") appliedFilters.value = blankFilters();
}
function resetFilters() {
  appliedFilters.value = blankFilters();
  quickFilter.value = "全部";
  areaApplied.value = false;
}
function applyFilters() {
  appliedFilters.value = { ...draftFilters.value };
  quickFilter.value = "全部";
  filterVisible.value = false;
}
function demo(message: string) { uni.showToast({ title: message, icon: "none" }); }
</script>

<template>
  <view class="home-page" :style="{ '--groups-page-height': pageHeight + 'px' }" data-design-contract="group-social-feed">
    <view id="activity-discovery" class="activity-discovery">
    <view class="groups-header">
    <view class="discovery-location" :style="brandBarStyle">
      <!-- #ifdef MP-WEIXIN -->
      <picker mode="region" :value="selectedRegion" @change="changeRegion">
        <view class="region-trigger" role="button" aria-label="选择省市区"><uni-icons type="location" size="25" color="#153f35" /><text>{{ currentCity }}<text class="location-separator"> · </text>{{ currentArea === '上海' ? '全城' : currentArea }}</text></view>
      </picker>
      <!-- #endif -->
      <!-- #ifndef MP-WEIXIN -->
      <button aria-label="切换所在区域" @tap="locationVisible = true"><uni-icons type="location" size="25" color="#153f35" /><text>{{ currentCity }}<text class="location-separator"> · </text>{{ currentArea === '上海' ? '全城' : currentArea }}</text></button>
      <!-- #endif -->
    </view>
    <view class="filter-rail">
      <button v-for="item in ['全部', '推荐', '距离', '本周']" :key="item" :class="{ active: quickFilter === item }" @tap="selectQuick(item)">{{ item }}</button>
      <button class="rail-search" :class="{ 'rail-search--filtered': hasSearchFilters }" aria-label="搜索与筛选" @tap="openSearch"><uni-icons type="search" size="25" color="#153f35" /></button>
    </view>
    </view>
    <scroll-view scroll-y class="groups-scroll">
    <view v-if="displayedServices.length" class="group-feed">
      <GroupFeedCard
        v-for="item in displayedServices"
        :key="item.id"
        :group="groupCards[item.id]"
        :distance="item.distance"
        @open="openService(item.id)"
      />
    </view>

    <view v-else class="empty-poster">
      <text>暂时没有匹配的活动</text><button @tap="resetFilters">回到推荐</button>
    </view>
    </scroll-view>
    </view>

    <AppTabBar active="discover" @discover="explore" />

    <!-- #ifdef H5 -->
    <Teleport to="body">
    <!-- #endif -->
    <BottomSheet :visible="locationVisible" title="选择区域" @close="locationVisible = false">
      <view class="city-row">
        <view><text class="city-label">当前城市</text><text class="city-name">上海市</text></view>
        <button class="city-switch" @tap="demo('城市切换能力已预留')">切换城市</button>
      </view>
      <view class="area-grid">
        <button v-for="area in areas" :key="area" :class="{ selected: currentArea === area.replace('区', '') || (currentArea === '上海' && area === '全部区域') }" @tap="chooseArea(area)">{{ area }}</button>
      </view>
    </BottomSheet>

    <view v-if="filterVisible" class="search-overlay" @keydown.esc="filterVisible = false">
      <view class="search-overlay__mask" @tap="filterVisible = false" @touchmove.stop.prevent />
      <view class="search-dialog" role="dialog" aria-modal="true" aria-label="找到想去的局">
        <view class="search-dialog__header"><text>找到想去的局</text><button aria-label="关闭搜索" @tap="filterVisible = false"><uni-icons type="closeempty" size="25" color="#153f35" /></button></view>
        <view class="search-dialog__input"><uni-icons type="search" size="22" color="#153f35" /><input v-model="draftFilters.keyword" placeholder="搜索标题、玩法或标签" confirm-type="search" @confirm="applyFilters" /></view>
        <scroll-view scroll-y class="search-dialog__filters">
          <view v-for="group in filterGroups" :key="group.key" class="search-choice">
            <text>{{ group.label }}</text>
            <view><button v-for="option in group.options" :key="option" :class="{ selected: draftFilters[group.key] === option }" @tap="draftFilters[group.key] = option">{{ option }}</button></view>
          </view>
        </scroll-view>
        <view class="search-dialog__actions"><button @tap="draftFilters = blankFilters()">重置</button><button @tap="applyFilters">搜索好局</button></view>
      </view>
    </view>
    <!-- #ifdef H5 -->
    </Teleport>
    <!-- #endif -->
  </view>
</template>

<style scoped lang="scss">
.home-page { height: var(--groups-page-height, 100vh); overflow: hidden; color: #103f33; background: #fbfcf8; }
.activity-discovery { height: 100%; display: flex; flex-direction: column; }
.groups-header { flex: 0 0 auto; background: #fbfcf8; }
.groups-scroll { flex: 1; height: 0; min-height: 0; }
.discovery-location button::after, .search-dialog button::after, .filter-rail button::after { border: 0; }
.discovery-location { padding: calc(var(--menu-button-top, 0px) + 12rpx) var(--menu-right-reserve, 17px) 8rpx 30rpx; }
.discovery-location button, .region-trigger { display: flex; align-items: center; gap: 10rpx; width: max-content; min-height: 80rpx; margin: 0; padding: 0; color: #153f35; background: transparent; font-size: 28rpx; line-height: 1.5; font-weight: 500; }
.location-separator { color: #8a988f; }
.filter-rail .rail-search { display: flex; align-items: center; justify-content: center; min-width: 60rpx; }
.rail-search--filtered::before { content: ""; position: absolute; top: 10rpx; right: 4rpx; width: 10rpx; height: 10rpx; border-radius: 50%; background: #fa704b; }
.search-overlay { position: fixed; inset: 0; z-index: 100; }
.search-overlay__mask { position: absolute; inset: 0; background: rgba(10,23,18,.5); }
.search-dialog { position: absolute; left: 50%; top: 50%; transform: translate(-50%, -50%); width: 92%; max-height: 86vh; padding: 30rpx; border-radius: 32rpx; background: #fffdf6; color: #153f35; display: flex; flex-direction: column; }
.search-dialog__header { display: flex; align-items: center; justify-content: space-between; gap: 20rpx; margin-bottom: 24rpx; font-size: 35rpx; font-weight: 700; }
.search-dialog__header button { display: flex; align-items: center; justify-content: center; width: 60rpx; height: 60rpx; margin: 0; padding: 0; background: transparent; }
.search-dialog__input { display: flex; align-items: center; gap: 14rpx; min-height: 82rpx; padding: 0 22rpx; border-radius: 22rpx; background: #edf3ec; }
.search-dialog__input input { flex: 1; min-width: 0; font-size: 27rpx; color: #153f35; }
.search-dialog__filters { min-height: 0; height: 560rpx; margin: 22rpx 0; }
.search-choice { display: flex; align-items: flex-start; gap: 18rpx; padding: 12rpx 0; }
.search-choice > text { width: 116rpx; flex-shrink: 0; padding-top: 17rpx; font-size: 25rpx; line-height: 1.35; font-weight: 600; }
.search-choice > view { flex: 1; min-width: 0; display: flex; flex-wrap: wrap; gap: 12rpx; }
.search-choice button { min-width: 112rpx; height: 66rpx; padding: 0 18rpx; margin: 0; border: 1rpx solid #d4ddd6; border-radius: 14rpx; background: #fff; color: #52685b; font-size: 24rpx; line-height: 64rpx; }
.search-choice button.selected { border-color: #153f35; color: #fff; background: #153f35; }
.search-dialog__actions { display: flex; gap: 18rpx; }
.search-dialog__actions button { flex: 1; height: 80rpx; margin: 0; padding: 0; border: 1rpx solid #153f35; border-radius: 20rpx; background: transparent; color: #153f35; font-size: 28rpx; line-height: 78rpx; }
.search-dialog__actions button:last-child { flex: 1.5; border-color: #fa704b; background: #fa704b; color: #fff; font-weight: 600; }
/* #ifdef H5 */
@media (min-width: 600px) { .home-page { height: var(--qudaju-preview-layout-height, 100vh); } .search-dialog { width: 880px; max-height: calc(86vh / var(--qudaju-preview-scale, 1)); transform: translate(-50%, -50%) scale(var(--qudaju-preview-scale, 1)); } }
/* #endif */
@media (max-height: 700px) { .search-dialog__filters { height: 380rpx; } }
.activity-discovery { position: relative; background: #fbfcf8; }
.group-feed { display: flex; flex-direction: column; gap: 24rpx; padding: 24rpx 24rpx calc(184rpx + env(safe-area-inset-bottom)); }
/* #ifdef H5 */
/* #endif */
button { transition: transform .16s ease, opacity .16s ease; }
button:active { transform: scale(.98); opacity: .86; }




.filter-rail { height: 76rpx; padding: 0 30rpx; display: flex; align-items: stretch; justify-content: space-between; border-bottom: 1rpx solid #d7ded8; background: #fbfcf8; }
.filter-rail button { position: relative; width: auto; height: 100%; flex: 0 0 auto; margin: 0; padding: 0; color: #70756f; background: transparent; font-size: 27rpx; line-height: 76rpx; font-weight: 800; }.filter-rail button.active { color: #0f4838; font-weight: 900; }.filter-rail button.active::after { content: ""; position: absolute; top: auto; right: auto; left: 50%; bottom: 2rpx; width: 40rpx; height: 3rpx; margin: 0; border: 0; border-radius: 0; background: #0f4838; transform: translateX(-50%); }

.empty-poster { margin: 28rpx 20rpx; padding: 56rpx 28rpx; display: flex; flex-direction: column; align-items: center; border: 2rpx dashed #aab8ad; background: #fffdf6; }.empty-poster text { font-size: 26rpx; font-weight: 800; }.empty-poster button { margin-top: 24rpx; padding: 0 28rpx; color: #fff; background: #0f4838; }

.city-row { padding: 20rpx 0 28rpx; display: flex; align-items: center; justify-content: space-between; }.city-row > view { display: flex; flex-direction: column; gap: 8rpx; }.city-label { color: #728078; font-size: 20rpx; }.city-name { font-size: 39rpx; font-weight: 900; }.city-switch { width: auto; margin: 0; padding: 0 23rpx; border: 2rpx solid #b7c2b9; color: #153f35; background: #fffdf6; }
.area-grid { display: grid; grid-template-columns: repeat(4, 1fr); gap: 13rpx; }.area-grid button { height: 68rpx; margin: 0; padding: 0; border: 2rpx solid #d5ddd6; border-radius: 8rpx; color: #49554e; background: #f2f4ef; font-size: 21rpx; }.area-grid button.selected { border-color: #ef644a; color: #fff; background: #ef644a; }

@media (prefers-reduced-motion: reduce) { button { transition: none; } }

</style>
