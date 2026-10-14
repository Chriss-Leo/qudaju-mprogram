<script setup lang="ts">
import { theme } from "../../theme";
import { computed, onUnmounted, ref, watch } from "vue";
import { onResize, onShow } from "@dcloudio/uni-app";
import AppTabBar from "../../components/AppTabBar.vue";
import GroupFeedCard from "../../components/GroupFeedCard.vue";
import { getGroup } from "../../stores/meetups";
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
  uni.setNavigationBarColor({ frontColor: "#000000", backgroundColor: theme.colors.white });
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
    const group = getGroup(item.id)!;
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
        <view class="region-trigger" role="button" aria-label="选择省市区"><uni-icons type="location" size="25" :color="theme.colors.text" /><text>{{ currentCity }}<text class="location-separator"> · </text>{{ currentArea === '上海' ? '全城' : currentArea }}</text></view>
      </picker>
      <!-- #endif -->
      <!-- #ifndef MP-WEIXIN -->
      <button aria-label="切换所在区域" @tap="locationVisible = true"><uni-icons type="location" size="25" :color="theme.colors.text" /><text>{{ currentCity }}<text class="location-separator"> · </text>{{ currentArea === '上海' ? '全城' : currentArea }}</text></button>
      <!-- #endif -->
    </view>
    <view class="filter-rail">
      <button v-for="item in ['全部', '推荐', '距离', '本周']" :key="item" :class="{ active: quickFilter === item }" @tap="selectQuick(item)">{{ item }}</button>
      <button class="rail-search" :class="{ 'rail-search--filtered': hasSearchFilters }" aria-label="搜索与筛选" @tap="openSearch"><uni-icons type="search" size="25" :color="theme.colors.text" /></button>
    </view>
    </view>
    <scroll-view scroll-y class="groups-scroll">
    <view v-if="displayedServices.length" class="group-feed">
      <GroupFeedCard
        v-for="item in displayedServices"
        :key="item.id"
        :group="getGroup(item.id)!"
        :distance="item.distance"
        @open="openService(item.id)"
      />
    </view>

    <view v-else class="empty-poster">
      <text>暂时没有匹配的活动</text><button @tap="resetFilters">回到推荐</button>
    </view>
    </scroll-view>
    </view>

    <AppTabBar active="discover" owner-route="pages/groups/index" @discover="explore" />

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
        <view class="search-dialog__header"><text>找到想去的局</text><button aria-label="关闭搜索" @tap="filterVisible = false"><uni-icons type="closeempty" size="25" :color="theme.colors.text" /></button></view>
        <view class="search-dialog__input"><uni-icons type="search" size="22" :color="theme.colors.text" /><input v-model="draftFilters.keyword" placeholder="搜索标题、玩法或标签" confirm-type="search" @confirm="applyFilters" /></view>
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
.home-page { height: var(--groups-page-height, 100vh); overflow: hidden; color: $color-text; background: $color-white; }
.activity-discovery { height: 100%; display: flex; flex-direction: column; }
.groups-header { flex: 0 0 auto; background: $color-white; }
.groups-scroll { flex: 1; height: 0; min-height: 0; }
.discovery-location button::after, .search-dialog button::after, .filter-rail button::after { border: 0; }
.discovery-location { padding: calc(var(--menu-button-top, 0px) + 12rpx) var(--menu-right-reserve, 17px) 8rpx 30rpx; }
.discovery-location button, .region-trigger { display: flex; align-items: center; gap: 10rpx; width: max-content; min-height: 80rpx; margin: 0; padding: 0; color: $color-text; background: transparent; font-size: $font-size-v28rpx; line-height: $line-height-v1p5; font-weight: $font-weight-v500; }
.location-separator { color: $color-textPlaceholder; }
.filter-rail .rail-search { display: flex; align-items: center; justify-content: center; min-width: 60rpx; }
.rail-search--filtered::before { content: ""; position: absolute; top: 10rpx; right: 4rpx; width: 10rpx; height: 10rpx; border-radius: $radius-v50percent; background: $color-accent; }
.search-overlay { position: fixed; inset: 0; z-index: 100; }
.search-overlay__mask { position: absolute; inset: 0; background: $color-compatOverlay28; }
.search-dialog { position: absolute; left: 50%; top: 50%; transform: translate(-50%, -50%); width: 92%; max-height: 86vh; padding: 30rpx; border-radius: $radius-v32rpx; background: $color-inverseWarm; color: $color-text; display: flex; flex-direction: column; }
.search-dialog__header { display: flex; align-items: center; justify-content: space-between; gap: 20rpx; margin-bottom: 24rpx; font-size: $font-size-v35rpx; font-weight: $font-weight-v700; }
.search-dialog__header button { display: flex; align-items: center; justify-content: center; width: 60rpx; height: 60rpx; margin: 0; padding: 0; background: transparent; }
.search-dialog__input { display: flex; align-items: center; gap: 14rpx; min-height: 82rpx; padding: 0 22rpx; border-radius: $radius-v22rpx; background: $color-input; }
.search-dialog__input input { flex: 1; min-width: 0; font-size: $font-size-v27rpx; color: $color-text; }
.search-dialog__filters { min-height: 0; height: 560rpx; margin: 22rpx 0; }
.search-choice { display: flex; align-items: flex-start; gap: 18rpx; padding: 12rpx 0; }
.search-choice > text { width: 116rpx; flex-shrink: 0; padding-top: 17rpx; font-size: $font-size-v25rpx; line-height: $line-height-v1p35; font-weight: $font-weight-v600; }
.search-choice > view { flex: 1; min-width: 0; display: flex; flex-wrap: wrap; gap: 12rpx; }
.search-choice button { min-width: 112rpx; height: 66rpx; padding: 0 18rpx; margin: 0; border: 1rpx solid $color-border; border-radius: $radius-v14rpx; background: $color-white; color: $color-textSecondary; font-size: $font-size-v24rpx; line-height: $line-height-v64rpx; }
.search-choice button.selected { border-color: $color-primary; color: $color-white; background: $color-primary; }
.search-dialog__actions { display: flex; gap: 18rpx; }
.search-dialog__actions button { flex: 1; height: 80rpx; margin: 0; padding: 0; border: 1rpx solid $color-primary; border-radius: $radius-v20rpx; background: $color-primary; color: $color-white; font-size: $font-size-v28rpx; line-height: $line-height-v78rpx; }
.search-dialog__actions button:first-child { border-color: $color-primary; color: $color-primary; background: $color-white; }
.search-dialog__actions button:last-child { flex: 1.5; border-color: $color-primary; background: $color-primary; color: $color-white; font-weight: $font-weight-v600; }
/* #ifdef H5 */
@media (min-width: 600px) { .home-page { height: var(--qudaju-preview-layout-height, 100vh); } .search-dialog { width: 880px; max-height: calc(86vh / var(--qudaju-preview-scale, 1)); transform: translate(-50%, -50%) scale(var(--qudaju-preview-scale, 1)); } }
/* #endif */
@media (max-height: 700px) { .search-dialog__filters { height: 380rpx; } }
.activity-discovery { position: relative; background: $color-white; }
.group-feed { display: flex; flex-direction: column; gap: 24rpx; padding: 24rpx 24rpx calc(184rpx + env(safe-area-inset-bottom)); }
/* #ifdef H5 */
/* #endif */
button { transition: transform .16s ease, opacity .16s ease; }
button:active { transform: scale(.98); opacity: .86; }




.filter-rail { height: 76rpx; padding: 0 30rpx; display: flex; align-items: stretch; justify-content: space-between; border-bottom: 1rpx solid $color-border; background: $color-white; }
.filter-rail button { position: relative; width: auto; height: 100%; flex: 0 0 auto; margin: 0; padding: 0; color: $color-textMuted; background: transparent; font-size: $font-size-v27rpx; line-height: $line-height-v76rpx; font-weight: $font-weight-v800; }.filter-rail button.active { color: $color-text; font-weight: $font-weight-v900; }.filter-rail button.active::after { content: ""; position: absolute; top: auto; right: auto; left: 50%; bottom: 2rpx; width: 40rpx; height: 3rpx; margin: 0; border: 0; border-radius: $radius-v0; background: $color-primary; transform: translateX(-50%); }

.empty-poster { margin: 28rpx 20rpx; padding: 56rpx 28rpx; display: flex; flex-direction: column; align-items: center; border: 2rpx dashed $color-compatBDADB4; background: $color-inverseWarm; }.empty-poster text { font-size: $font-size-v26rpx; font-weight: $font-weight-v800; }.empty-poster button { margin-top: 24rpx; padding: 0 28rpx; color: $color-white; background: $color-primary; }

.city-row { padding: 20rpx 0 28rpx; display: flex; align-items: center; justify-content: space-between; }.city-row > view { display: flex; flex-direction: column; gap: 8rpx; }.city-label { color: $color-textMuted; font-size: $font-size-v20rpx; }.city-name { font-size: $font-size-v39rpx; font-weight: $font-weight-v900; }.city-switch { width: auto; margin: 0; padding: 0 23rpx; border: 2rpx solid $color-compatCEC0C6; color: $color-text; background: $color-inverseWarm; }
.area-grid { display: grid; grid-template-columns: repeat(4, 1fr); gap: 13rpx; }.area-grid button { height: 68rpx; margin: 0; padding: 0; border: 2rpx solid $color-border; border-radius: $radius-v8rpx; color: $color-textSecondary; background: $color-white; font-size: $font-size-v21rpx; }.area-grid button.selected { border-color: $color-primary; color: $color-white; background: $color-primary; }

@media (prefers-reduced-motion: reduce) { button { transition: none; } }

</style>
