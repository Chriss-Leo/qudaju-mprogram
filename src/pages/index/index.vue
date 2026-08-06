<script setup lang="ts">
import { computed, ref } from "vue";
import AppTabBar from "../../components/AppTabBar.vue";
import BottomSheet from "../../components/BottomSheet.vue";
import { services } from "../../mocks/services";
import type { ServiceItem } from "../../types/service";

const locationVisible = ref(false);
const filterVisible = ref(false);
const currentArea = ref("静安");
const currentSort = ref("推荐");
const serviceMode = ref("线下见面");
const bookingTime = ref("今天");
const minPrice = ref("");
const maxPrice = ref("");
const filtersApplied = ref(false);
const quickFilter = ref("推荐");
const windowInfo = uni.getWindowInfo();
const statusBarHeight = windowInfo.statusBarHeight || 0;
const menuButtonRect = typeof uni.getMenuButtonBoundingClientRect === "function"
  ? uni.getMenuButtonBoundingClientRect()
  : null;
const menuTop = menuButtonRect?.top || statusBarHeight;
const menuHeight = menuButtonRect?.height || 32;
const menuRightReserve = menuButtonRect
  ? Math.max(12, windowInfo.windowWidth - menuButtonRect.left + 8)
  : 17;
const brandBarStyle = {
  "--menu-button-top": `${menuTop}px`,
  "--menu-button-height": `${menuHeight}px`,
  "--menu-right-reserve": `${menuRightReserve}px`,
};

const programs = [
  { label: "桌游组局", icon: "dice", keyword: "桌游", tone: "paper" },
  { label: "摄影陪拍", icon: "camera", keyword: "摄影", tone: "mint" },
  { label: "城市向导", icon: "sign", keyword: "城市向导", tone: "paper" },
  { label: "户外活动", icon: "mountain", keyword: "户外", tone: "mint" },
];
const areas = ["全部区域", "徐汇区", "浦东新区", "黄浦区", "静安区", "长宁区", "闵行区", "普陀区"];

const displayedServices = computed(() => {
  let result = [...services];
  if (filtersApplied.value) {
    result = result.filter((item) => serviceMode.value === "均可" || item.mode === serviceMode.value);
    if (bookingTime.value !== "自定义") {
      result = result.filter((item) => item.available.includes(bookingTime.value as "今天" | "明天" | "本周末"));
    }
    if (minPrice.value) result = result.filter((item) => item.price >= Number(minPrice.value));
    if (maxPrice.value) result = result.filter((item) => item.price <= Number(maxPrice.value));
  }
  if (quickFilter.value === "3km内") result = result.filter((item) => Number.parseFloat(item.distance) <= 3);
  if (currentSort.value === "距离") result.sort((a, b) => Number.parseFloat(a.distance) - Number.parseFloat(b.distance));
  if (currentSort.value === "价格") result.sort((a, b) => a.price - b.price);
  return result;
});

function openSearch(keyword = "") {
  uni.navigateTo({ url: `/pages/search/index${keyword ? `?keyword=${encodeURIComponent(keyword)}` : ""}` });
}
function openService(id: string) { uni.navigateTo({ url: `/pages/service-detail/index?id=${id}` }); }
function chooseArea(area: string) {
  currentArea.value = area === "全部区域" ? "上海" : area.replace("区", "");
  locationVisible.value = false;
}
function selectQuick(value: string) {
  quickFilter.value = value;
  if (value === "推荐") {
    filtersApplied.value = false;
    currentSort.value = "推荐";
  } else if (value === "本周") {
    bookingTime.value = "本周末";
    filtersApplied.value = true;
    currentSort.value = "推荐";
  } else {
    filtersApplied.value = false;
    currentSort.value = "距离";
  }
}
function resetFilters() {
  serviceMode.value = "线下见面";
  bookingTime.value = "今天";
  minPrice.value = "";
  maxPrice.value = "";
  filtersApplied.value = false;
  quickFilter.value = "推荐";
}
function applyFilters() {
  filtersApplied.value = true;
  quickFilter.value = bookingTime.value === "本周末" ? "本周" : "筛选";
  filterVisible.value = false;
}
function serviceImage(item: ServiceItem) {
  const imageMap: Record<ServiceItem["visual"], string> = {
    guide: "/static/posterwall/city-photo-walk.jpg",
    photo: "/static/posterwall/outdoor-walk.jpg",
    game: "/static/posterwall/boardgame-night.jpg",
    outdoor: "/static/posterwall/hero-frisbee.jpg",
  };
  return imageMap[item.visual];
}
function posterLabel(item: ServiceItem) {
  return ({ guide: "城市漫步", photo: "街头纪实", game: "桌游组局", outdoor: "户外搭子" } as const)[item.visual];
}
function posterTitleLines(item: ServiceItem): [string, string] {
  const breaks: Record<string, [string, string]> = {
    "梧桐区漫步与城市摄影": ["梧桐区漫步", "与城市摄影"],
    "街头氛围感摄影陪拍": ["街头氛围感", "摄影陪拍"],
    "周末桌游轻松组局": ["周末桌游", "轻松组局"],
    "滨江日落飞盘新手局": ["滨江日落飞盘", "新手局"],
    "夜游苏河湾建筑故事线": ["夜游苏河湾", "建筑故事线"],
  };
  return breaks[item.title] || [item.title, "同城兴趣服务"];
}
function demo(message: string) { uni.showToast({ title: message, icon: "none" }); }
</script>

<template>
  <view class="home-page" data-design-contract="posterwall-b-fullbleed">
    <!--
      THESIS: 首页是一张可直接进入活动的城市海报，不是搜索框和卡片组件的集合。
      OWN-WORLD: 森林绿油墨、珊瑚票根、冷白纸面、全幅纪实摄影、撕纸文字块与横向票券分类。
      STORY: 用户确认区域，看见今日主活动，按类别或距离筛选，再从摄影海报进入服务详情。
      FIRST VIEWPORT: 白色品牌栏在顶，摄影主视觉满宽展开，左下深绿撕纸标题与右下珊瑚票根形成唯一焦点，分类票紧接其后。
      FORM: B 风格全幅城市活动海报；seed key posterwall-b-fullbleed.
      FINISH: unreviewed and undocumented is unfinished; this build ends with the finish review, the verdict, and DESIGN.md
    -->

    <view class="brand-bar" :style="brandBarStyle">
      <image class="wordmark-image" src="/static/posterwall/wordmark-b-transparent.png" mode="aspectFit" aria-label="趣搭局" />
      <view class="brand-actions">
        <button class="location-control" @tap="locationVisible = true">
          <text>上海 · {{ currentArea }}</text>
          <uni-icons type="location" size="18" color="#153f35" />
        </button>
        <button class="search-button" aria-label="搜索" @tap="openSearch()"><view /></button>
      </view>
    </view>

    <image class="hero-poster-link" src="/static/posterwall/home-hero-b-exact.png" mode="widthFix" @tap="openSearch('城市活动')" />

    <view class="program-strip">
      <button
        v-for="program in programs"
        :key="program.label"
        class="program-ticket"
        :class="`program-ticket--${program.tone}`"
        @tap="openSearch(program.keyword)"
      >
        <view class="program-icon" :class="`program-icon--${program.icon}`">
          <view class="program-icon__part" />
          <view class="program-icon__part" />
          <view class="program-icon__part" />
        </view>
        <text>{{ program.label }}</text>
      </button>
    </view>

    <view class="filter-rail">
      <button v-for="item in ['推荐', '本周', '3km内']" :key="item" :class="{ active: quickFilter === item }" @tap="selectQuick(item)">{{ item }}</button>
      <button class="filter-button" @tap="filterVisible = true"><view class="funnel" /><text>筛选</text></button>
    </view>

    <view v-if="displayedServices.length" class="poster-feed">
      <view
        v-for="(item, index) in displayedServices"
        :key="item.id"
        class="activity-poster"
        :class="index % 2 === 0 ? 'activity-poster--even' : 'activity-poster--odd'"
        @tap="openService(item.id)"
      >
        <image :src="serviceImage(item)" mode="aspectFill" />
        <view class="activity-poster__shade" />
        <view class="activity-poster__copy">
          <text class="poster-kicker">{{ posterLabel(item) }}</text>
          <text class="poster-title poster-title--first">{{ posterTitleLines(item)[0] }}</text>
          <text class="poster-title poster-title--second">{{ posterTitleLines(item)[1] }}</text>
          <text class="poster-note">{{ item.tags.join(' / ') }}</text>
        </view>
        <view v-if="index === 1" class="friendly-stamp"><text>新手友好</text><text>氛围轻松</text></view>
        <view class="poster-price"><text>¥{{ item.price }}</text><text>起</text></view>
      </view>
    </view>

    <view v-else class="empty-poster">
      <text>暂时没有匹配的活动</text><button @tap="resetFilters">回到推荐</button>
    </view>

    <AppTabBar active="home" />

    <BottomSheet :visible="locationVisible" title="选择服务区域" @close="locationVisible = false">
      <view class="city-row">
        <view><text class="city-label">当前城市</text><text class="city-name">上海市</text></view>
        <button class="city-switch" @tap="demo('城市切换能力已预留')">切换城市</button>
      </view>
      <view class="area-grid">
        <button v-for="area in areas" :key="area" :class="{ selected: currentArea === area.replace('区', '') || (currentArea === '上海' && area === '全部区域') }" @tap="chooseArea(area)">{{ area }}</button>
      </view>
    </BottomSheet>

    <BottomSheet :visible="filterVisible" title="筛选活动" @close="filterVisible = false">
      <view class="filter-group">
        <text class="filter-group__label">服务方式</text>
        <view class="choice-grid"><button v-for="mode in ['线下见面', '线上远程', '均可']" :key="mode" :class="{ selected: serviceMode === mode }" @tap="serviceMode = mode">{{ mode }}</button></view>
      </view>
      <view class="filter-group">
        <text class="filter-group__label">可约时间</text>
        <view class="choice-grid choice-grid--four"><button v-for="time in ['今天', '明天', '本周末', '自定义']" :key="time" :class="{ selected: bookingTime === time }" @tap="bookingTime = time">{{ time }}</button></view>
      </view>
      <view class="filter-group">
        <text class="filter-group__label">价格范围</text>
        <view class="price-fields"><view><input v-model="minPrice" type="number" placeholder="最低价" /></view><text>-</text><view><input v-model="maxPrice" type="number" placeholder="最高价" /></view></view>
      </view>
      <view class="sheet-actions"><button class="button-secondary" @tap="resetFilters">重置</button><button class="button-primary" @tap="applyFilters">查看 {{ displayedServices.length }} 项</button></view>
    </BottomSheet>
  </view>
</template>

<style scoped lang="scss">
.home-page { min-height: 100vh; padding-bottom: calc(160rpx + env(safe-area-inset-bottom)); overflow-x: hidden; color: #103f33; background: white; }
button { transition: transform .16s ease, opacity .16s ease; }
button:active { transform: scale(.98); opacity: .86; }

.brand-bar { box-sizing: border-box; height: calc(var(--menu-button-top, 0px) + var(--menu-button-height, 32px) + 12rpx); padding: var(--menu-button-top, 0px) var(--menu-right-reserve, 17px) 12rpx 26rpx; display: flex; align-items: center; gap: 20rpx; background: #fff; }
.wordmark-image { width: 210rpx; height: 112rpx; display: block; }
.brand-actions { min-width: 0; margin-left: auto; display: flex; align-items: center; gap: 18rpx; }
.location-control { width: auto; height: 62rpx; margin: 0; padding: 0 22rpx; display: flex; align-items: center; gap: 12rpx; border: 2rpx solid #d9ddda; border-radius: 999rpx; color: #17211c; background: #fff; font-size: 24rpx; font-weight: 700; box-shadow: 0 4rpx 14rpx rgba(18,57,46,.04); }
.location-icon { flex: 0 0 auto; }
.search-button { position: relative; width: 58rpx; height: 58rpx; margin: 0; padding: 0; background: transparent; }.search-button view { position: absolute; width: 34rpx; height: 34rpx; left: 0rpx; top: 5rpx; border: 4rpx solid #153f35; border-radius: 50%; }.search-button view::after { content: ""; position: absolute; width: 13rpx; height: 4rpx; right: -9rpx; bottom: -5rpx; border-radius: 4rpx; background: #153f35; transform: rotate(47deg); transform-origin: left center; }

.hero-poster-link { width: 100%; height: auto; margin-top: 18rpx; display: block; background: transparent; }

.program-strip { padding: 24rpx 18rpx 26rpx; display: grid; grid-template-columns: repeat(4, minmax(0, 1fr)); gap: 11rpx; border-bottom: 1rpx solid #dfe4df; background: #fbfcf8; }
.program-ticket { position: relative; height: 152rpx; margin: 0; padding: 15rpx 3rpx 12rpx; display: flex; flex-direction: column; align-items: center; justify-content: center; gap: 13rpx; overflow: hidden; border: 2rpx solid #dcded5; border-radius: 0; color: #121713; background: #fffdf6; font-size: 25rpx; line-height: 1.1; font-weight: 900; clip-path: polygon(4% 0,96% 0,96% 6%,100% 12%,96% 18%,100% 24%,96% 30%,100% 36%,96% 42%,100% 48%,96% 54%,100% 60%,96% 66%,100% 72%,96% 78%,100% 84%,96% 90%,100% 96%,96% 100%,4% 100%,0 94%,4% 88%,0 82%,4% 76%,0 70%,4% 64%,0 58%,4% 52%,0 46%,4% 40%,0 34%,4% 28%,0 22%,4% 16%,0 10%,4% 4%); }
.program-ticket--paper { background: rgba(255,248,224,.96); }
.program-ticket--mint { background: #c8dacb; }
.program-ticket > text, .program-icon { position: relative; z-index: 2; }
.program-icon { width: 57rpx; height: 50rpx; color: #111612; }.program-icon__part { position: absolute; display: block; box-sizing: border-box; }
.program-icon--dice .program-icon__part:first-child { left: 4rpx; top: 7rpx; width: 29rpx; height: 37rpx; border: 3rpx solid currentColor; border-radius: 4rpx; transform: rotate(-7deg); }.program-icon--dice .program-icon__part:nth-child(2) { right: 0; top: 18rpx; width: 26rpx; height: 26rpx; border: 3rpx solid currentColor; transform: rotate(18deg); }.program-icon--dice .program-icon__part:last-child { width: 5rpx; height: 5rpx; left: 14rpx; top: 18rpx; border-radius: 50%; background: currentColor; box-shadow: 10rpx 11rpx 0 currentColor, 27rpx 8rpx 0 currentColor; }
.program-icon--camera .program-icon__part:first-child { inset: 12rpx 1rpx 2rpx; border: 3rpx solid currentColor; border-radius: 4rpx; }.program-icon--camera .program-icon__part:nth-child(2) { left: 20rpx; top: 18rpx; width: 15rpx; height: 15rpx; border: 3rpx solid currentColor; border-radius: 50%; }.program-icon--camera .program-icon__part:last-child { left: 10rpx; top: 5rpx; width: 22rpx; height: 10rpx; border: 3rpx solid currentColor; border-bottom: 0; }
.program-icon--sign .program-icon__part:first-child { left: 26rpx; top: 1rpx; width: 4rpx; height: 48rpx; background: currentColor; }.program-icon--sign .program-icon__part:nth-child(2), .program-icon--sign .program-icon__part:last-child { left: 6rpx; width: 44rpx; height: 15rpx; border: 3rpx solid currentColor; background: #fffdf6; transform: skewX(-15deg); }.program-icon--sign .program-icon__part:nth-child(2) { top: 6rpx; }.program-icon--sign .program-icon__part:last-child { top: 27rpx; }.program-ticket--mint .program-icon--sign .program-icon__part:nth-child(2), .program-ticket--mint .program-icon--sign .program-icon__part:last-child { background: #c8dacb; }
.program-icon--mountain .program-icon__part:first-child { left: 1rpx; bottom: 3rpx; width: 33rpx; height: 33rpx; border: 3rpx solid currentColor; border-right: 0; border-bottom: 0; transform: rotate(45deg); }.program-icon--mountain .program-icon__part:nth-child(2) { right: 0; bottom: 5rpx; width: 27rpx; height: 27rpx; border: 3rpx solid currentColor; border-right: 0; border-bottom: 0; transform: rotate(45deg); }.program-icon--mountain .program-icon__part:last-child { left: 11rpx; top: 26rpx; width: 33rpx; border-top: 3rpx solid currentColor; transform: rotate(-8deg); }

.filter-rail { height: 76rpx; padding: 0 30rpx; display: flex; align-items: stretch; justify-content: space-between; border-bottom: 1rpx solid #d7ded8; background: #fbfcf8; }
.filter-rail button { position: relative; width: auto; height: 100%; flex: 0 0 auto; margin: 0; padding: 0; color: #70756f; background: transparent; font-size: 27rpx; line-height: 76rpx; font-weight: 800; }.filter-rail button.active { color: #0f4838; font-weight: 900; }.filter-rail button.active::after { content: ""; position: absolute; left: 50%; bottom: 2rpx; width: 40rpx; height: 3rpx; border-radius: 0; background: #0f4838; transform: translateX(-50%); }
.filter-button { display: flex; align-items: center; justify-content: flex-end; gap: 11rpx; }.funnel { width: 25rpx; height: 25rpx; background: currentColor; clip-path: polygon(0 3%,100% 3%,62% 48%,62% 100%,42% 82%,42% 48%); }

.poster-feed { padding: 22rpx 20rpx 48rpx; display: flex; flex-direction: column; gap: 18rpx; }
.activity-poster { position: relative; height: 352rpx; overflow: hidden; border-radius: 12rpx; background: #719080; box-shadow: 0 11rpx 28rpx rgba(16,63,51,.13); }
.activity-poster image { width: 100%; height: 100%; display: block; }.activity-poster__shade { position: absolute; inset: 0; background: linear-gradient(90deg, rgba(5,31,24,.7) 0, rgba(5,31,24,.18) 56%, transparent 76%); }
.activity-poster__copy { position: absolute; left: 0; top: 72rpx; max-width: 72%; display: flex; flex-direction: column; align-items: flex-start; }
.poster-kicker, .poster-title, .poster-note { clip-path: polygon(.8% 0,4% 1.5%,8% 0,12% 1.5%,16% 0,20% 1.5%,24% 0,28% 1.5%,32% 0,36% 1.5%,40% 0,44% 1.5%,48% 0,52% 1.5%,56% 0,60% 1.5%,64% 0,68% 1.5%,72% 0,76% 1.5%,80% 0,84% 1.5%,88% 0,92% 1.5%,96% 0,99.2% 0,100% 10%,99.2% 20%,100% 30%,99.2% 40%,100% 50%,99.2% 60%,100% 70%,99.2% 80%,100% 90%,99.2% 100%,96% 100%,92% 98.5%,88% 100%,84% 98.5%,80% 100%,76% 98.5%,72% 100%,68% 98.5%,64% 100%,60% 98.5%,56% 100%,52% 98.5%,48% 100%,44% 98.5%,40% 100%,36% 98.5%,32% 100%,28% 98.5%,24% 100%,20% 98.5%,16% 100%,12% 98.5%,8% 100%,4% 98.5%,.8% 100%,0 90%,.8% 80%,0 70%,.8% 60%,0 50%,.8% 40%,0 30%,.8% 20%,0 10%); }
.poster-kicker { padding: 9rpx 27rpx; font-size: 20rpx; font-weight: 800; }
.poster-title { padding: 8rpx 30rpx 7rpx; color: #172019; background: rgba(255,248,224,.96); font-size: 31rpx; line-height: 1.08; font-weight: 900; }
.poster-title--first { margin-top: 6rpx; }
.poster-title--second { margin-top: 3rpx; }
.poster-note { margin-top: 6rpx; padding: 9rpx 29rpx; font-size: 20rpx; font-weight: 700; transform: rotate(1.8deg); transform-origin: left center; }
.activity-poster--even .poster-kicker { color: #f7faf3; background: rgba(13,78,61,.96); }
.activity-poster--even .poster-note { color: #21483d; background: rgba(197,220,200,.96); }
.activity-poster--odd .poster-kicker { color: #21483d; background: rgba(197,220,200,.96); }
.activity-poster--odd .poster-note { color: #f7faf3; background: rgba(13,78,61,.96); }
.friendly-stamp { position: absolute; right: 18rpx; top: 18rpx; width: 104rpx; height: 104rpx; display: flex; flex-direction: column; align-items: center; justify-content: center; border: 2rpx dashed #783326; border-radius: 50%; color: #2c241e; background: #ef735c; font-size: 19rpx; font-weight: 900; transform: rotate(11deg); }.poster-price { position: absolute; right: 18rpx; bottom: 17rpx; padding: 7rpx 12rpx; color: #fffdf6; background: rgba(12,65,50,.88); }.poster-price text:first-child { font-size: 27rpx; font-weight: 900; }.poster-price text:last-child { margin-left: 3rpx; font-size: 16rpx; }
.empty-poster { margin: 28rpx 20rpx; padding: 56rpx 28rpx; display: flex; flex-direction: column; align-items: center; border: 2rpx dashed #aab8ad; background: #fffdf6; }.empty-poster text { font-size: 26rpx; font-weight: 800; }.empty-poster button { margin-top: 24rpx; padding: 0 28rpx; color: #fff; background: #0f4838; }

.city-row { padding: 20rpx 0 28rpx; display: flex; align-items: center; justify-content: space-between; }.city-row > view { display: flex; flex-direction: column; gap: 8rpx; }.city-label { color: #728078; font-size: 20rpx; }.city-name { font-size: 39rpx; font-weight: 900; }.city-switch { width: auto; margin: 0; padding: 0 23rpx; border: 2rpx solid #b7c2b9; color: #153f35; background: #fffdf6; }
.area-grid { display: grid; grid-template-columns: repeat(4, 1fr); gap: 13rpx; }.area-grid button, .choice-grid button { height: 68rpx; margin: 0; padding: 0; border: 2rpx solid #d5ddd6; border-radius: 8rpx; color: #49554e; background: #f2f4ef; font-size: 21rpx; }.area-grid button.selected, .choice-grid button.selected { border-color: #ef644a; color: #fff; background: #ef644a; }
.filter-group { margin-top: 30rpx; }.filter-group__label { display: block; margin-bottom: 16rpx; font-size: 24rpx; font-weight: 900; }.choice-grid { display: grid; grid-template-columns: repeat(3, 1fr); gap: 13rpx; }.choice-grid--four { grid-template-columns: repeat(4, 1fr); }
.price-fields { display: grid; grid-template-columns: 1fr auto 1fr; align-items: center; gap: 15rpx; }.price-fields > view { height: 72rpx; padding: 0 20rpx; display: flex; align-items: center; border: 2rpx solid #ccd6ce; border-radius: 8rpx; background: #f7f8f4; }.price-fields input { width: 100%; font-size: 22rpx; }
.sheet-actions { margin-top: 36rpx; display: grid; grid-template-columns: .8fr 1.35fr; gap: 14rpx; }.sheet-actions button { height: 78rpx; margin: 0; border-radius: 10rpx; font-size: 24rpx; font-weight: 900; }.button-secondary { border: 2rpx solid #bcc8bf; color: #153f35; background: #fffdf6; }.button-primary { color: #fffdf6; background: #0f4838; }

@media (prefers-reduced-motion: reduce) { button { transition: none; } }

</style>
