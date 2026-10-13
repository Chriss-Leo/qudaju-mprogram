<script setup lang="ts">
import { computed, ref, onMounted, onUnmounted } from "vue";
import { onLoad, onShow, onResize } from "@dcloudio/uni-app";
import { theme } from "../../theme";
import { groupCards } from "../../mocks/group-cards";
import SurfaceHeader from "../../components/SurfaceHeader.vue";
import UiIcon from "../../components/UiIcon.vue";

const groupId = ref("game-01");
const viewportHeight = ref(uni.getWindowInfo().windowHeight);
onResize(({ size }) => { if (size) viewportHeight.value = size.windowHeight; });
const imageFailed = ref(false);
const group = computed(() => Object.prototype.hasOwnProperty.call(groupCards, groupId.value) ? groupCards[groupId.value] : undefined);
const photos: Record<string, string> = {
  "game-01": "/static/profile-a/boardgame-hero.jpg",
  "photo-01": "/static/posterwall/city-photo-walk.jpg",
  "city-guide-01": "/static/posterwall/city-photo-walk.jpg",
  "outdoor-01": "/static/posterwall/outdoor-walk.jpg",
  "guide-02": "/static/posterwall/outdoor-walk.jpg",
};
const tags: Record<string, string[]> = { 桌游: ["轻策略", "新手友好"], 陪拍: ["城市漫步", "互拍"], 向导: ["街区漫步", "城市故事"], 户外: ["轻徒步", "慢节奏"] };
const facts = computed(() => group.value ? [
  { icon: "clock", value: group.value.time, label: "活动时间" },
  { icon: "location", value: group.value.area, label: "公开区域" },
  { icon: "people", value: group.value.joined + " / " + group.value.capacity + "人", label: "已报名" },
  { icon: "wallet", value: group.value.fund ? "¥" + group.value.fund + "/人" : "免费", label: group.value.fund ? "活动基金" : "费用" },
] : []);
const rule = computed(() => {
  if (!group.value) return "";
  const paid = group.value.fund > 0;
  if (group.value.admission === "局主审核") return paid
    ? "局主审核通过后付款，付款完成才占正式名额。通过后可查看集合细节。"
    : "提交申请后，局主会进行审核。通过后占正式名额，并可查看集合细节。";
  return paid ? "付款完成后占正式名额，报名成功后可查看集合细节。" : "报名成功后占正式名额，并可查看集合细节。";
});
onLoad(query => { if (query?.id) groupId.value = query.id; imageFailed.value = false; });
// H5 reuses the route component when only the query changes (including history).
// #ifdef H5
function syncQuery() {
  const hash = window.location.hash;
  if (!hash.startsWith("#/pages/service-detail/index")) return;
  groupId.value = new URLSearchParams(hash.split("?")[1] || "").get("id") || "game-01";
  imageFailed.value = false;
}
onMounted(() => { syncQuery(); window.addEventListener("hashchange", syncQuery); });
onUnmounted(() => window.removeEventListener("hashchange", syncQuery));
// #endif
onShow(() => { uni.setNavigationBarColor({ frontColor: "#000000", backgroundColor: theme.colors.white }); });
function back() { uni.navigateBack({ fail: () => uni.redirectTo({ url: "/pages/groups/index" }) }); }
function preview() { const url = photos[groupId.value]; if (url) uni.previewImage({ urls: [url], current: url }); }
function apply() {
  uni.showModal({ title: "报名功能尚未接入", content: "当前为组局详情示例。报名、局主审核及付款将在后端联调后开放，本次不会提交申请或收款。", showCancel: false, confirmText: "知道了", confirmColor: theme.colors.primary });
}
function hostInfo() { uni.showToast({ title: "局主信息卡片尚未接入", icon: "none" }); }
</script>

<template>
  <view class="meetup-detail" :style="{ height: viewportHeight + 'px' }" data-design-contract="THESIS: photograph-led meetup, no service booking cards. OWN-WORLD: white, blue utilities, orange application. STORY: activity then facts and rules. FIRST VIEWPORT: photo, title, four facts, host. FORM: user-approved A 2026-10-13. FINISH: unreviewed and undocumented is unfinished; this build ends with the finish review, the verdict, and DESIGN.md">
    <!-- THESIS: The approved A meetup photograph and open reading hierarchy, not service booking cards.
         OWN-WORLD: white, blue utility icons, orange application; independent photo and semantic controls.
         STORY: activity, facts, host, review rules, application. FIRST VIEWPORT: photo, title, tags, four facts, host.
         FORM: User-approved A, 2026-10-13. FINISH: unreviewed and undocumented is unfinished; this build ends with the finish review, the verdict, and DESIGN.md -->
    <SurfaceHeader title="组局详情" back @back="back" />
    <scroll-view class="detail-scroll" scroll-y :show-scrollbar="false">
    <view v-if="group" class="detail-content">
      <button class="hero-button" aria-label="查看活动示例大图" @tap="preview">
        <image v-if="!imageFailed" class="detail-photo" :src="photos[groupId]" mode="aspectFill" @error="imageFailed = true" />
        <view v-else class="photo-error">图片暂时未能加载</view>
      </button>
      <text class="detail-title">{{ group.title }}</text>
      <view class="detail-tags"><text>{{ group.category }}</text><text v-for="tag in tags[group.category] || []" :key="tag">{{ tag }}</text></view>
      <view class="detail-facts">
        <view v-for="fact in facts" :key="fact.icon" class="detail-fact"><UiIcon :name="fact.icon" /><text class="detail-fact__value">{{ fact.value }}</text><text class="detail-fact__label">{{ fact.label }}</text></view>
      </view>
      <button class="detail-host" @tap="hostInfo">
        <image v-if="groupId === 'game-01'" class="detail-host__avatar" src="/static/profile-a/host-avatar.jpg" mode="aspectFill" aria-label="示例局主头像" />
        <view v-else class="detail-host__avatar detail-host__initial">{{ group.host.slice(0, 1) }}</view>
        <view class="detail-host__copy"><text>{{ group.host }}</text><text>本次活动发起人</text></view>
        <UiIcon class="detail-host__chevron" name="chevron" tone="muted" />
      </button>
      <view class="detail-about"><text class="detail-heading">活动介绍</text><text class="detail-description">{{ group.description }}</text></view>
      <view class="detail-rule"><UiIcon name="info" /><view><text>{{ group.admission }}</text><text>{{ rule }}</text></view></view>
      <text class="detail-demo">活动、人物及图片均为示例，暂不开放真实报名。</text>
    </view>
    <view v-else class="detail-empty"><text class="detail-heading">暂未找到这个局</text><text>返回列表，看看其他同城活动。</text><button @tap="back">返回列表</button></view>
    </scroll-view>
    <view v-if="group" class="detail-actions">
      <view class="detail-price">
        <text class="detail-price__label">{{ group.fund ? '活动基金' : '活动费用' }}</text>
        <view class="detail-price__amount"><text v-if="group.fund" class="detail-price__unit">¥</text><text class="detail-price__value">{{ group.fund || '免费' }}</text><text v-if="group.fund" class="detail-price__unit">/ 人</text></view>
      </view>
      <button class="detail-apply action-accent" @tap="apply"><text>{{ group.admission === '局主审核' ? '申请加入' : '报名加入' }}</text></button>
    </view>
  </view>
</template>

<style scoped lang="scss">
.meetup-detail { display: flex; flex-direction: column; overflow: hidden; background: $color-white; color: $color-text; }
.meetup-detail :deep(.surface-header) { flex-shrink: 0; padding-bottom: 14rpx; }
.detail-scroll { flex: 1; height: 0; min-height: 0; }
.detail-content { padding: 20rpx 28rpx 40rpx; }
.hero-button { display: block; margin: 0; padding: 0; width: 100%; height: 452rpx; overflow: hidden; border-radius: $radius-v24rpx; background: $color-card; }
.detail-photo { display: block; width: 100%; height: 100%; }
.photo-error { height: 100%; display: flex; align-items: center; justify-content: center; color: $color-textMuted; font-size: $font-size-v26rpx; }
.detail-title { display: block; margin-top: 34rpx; font-size: $font-size-v38rpx; font-weight: $font-weight-v700; line-height: $line-height-v1p4; overflow-wrap: anywhere; }
.detail-tags { display: flex; flex-wrap: wrap; gap: 14rpx; margin-top: 20rpx; }
.detail-tags text { padding: 7rpx 18rpx; color: $color-primary; background: $color-tag; border-radius: $radius-v999rpx; font-size: $font-size-v23rpx; line-height: $line-height-v1p4; }
.detail-facts { margin: 34rpx 0 32rpx; padding: 30rpx 8rpx; display: grid; grid-template-columns: repeat(4, minmax(0, 1fr)); border-radius: $radius-v24rpx; background: $color-card; }
.detail-fact { position: relative; display: flex; align-items: center; flex-direction: column; gap: 12rpx; padding: 0 4rpx; }
.detail-fact + .detail-fact::before { content: ""; position: absolute; top: 16rpx; bottom: 10rpx; left: 0; width: 1rpx; background: $color-border; }
.detail-fact__value { font-size: $font-size-v24rpx; font-weight: $font-weight-v500; line-height: $line-height-v1p4; text-align: center; }
.detail-fact__label { font-size: $font-size-v21rpx; color: $color-textMuted; line-height: $line-height-v1p4; }
.detail-host { display: flex; align-items: center; gap: 20rpx; width: 100%; margin: 0; padding: 24rpx 0; border-top: 1rpx solid $color-border; border-bottom: 1rpx solid $color-border; border-radius: $radius-v0; background: transparent; text-align: left; }
.detail-host__avatar { display: block; flex: 0 0 96rpx; width: 96rpx; height: 96rpx; border-radius: $radius-v50percent; background: $color-tag; }
.detail-host__initial { display: flex; align-items: center; justify-content: center; color: $color-primary; font-size: $font-size-v32rpx; }
.detail-host__copy { display: flex; flex-direction: column; gap: 8rpx; flex: 1; }
.detail-host__copy text:first-child { font-size: $font-size-v29rpx; font-weight: $font-weight-v600; line-height: $line-height-v1p4; color: $color-text; }
.detail-host__copy text:last-child { font-size: $font-size-v23rpx; line-height: $line-height-v1p4; color: $color-textMuted; }
.detail-host__chevron { width: 28rpx; height: 28rpx; }
.detail-about { padding: 32rpx 0 24rpx; }
.detail-heading { display: block; font-size: $font-size-v32rpx; font-weight: $font-weight-v700; line-height: $line-height-v1p4; }
.detail-description { display: block; margin-top: 18rpx; color: $color-textSecondary; font-size: $font-size-v27rpx; line-height: $line-height-v1p75; }
.detail-rule { display: flex; align-items: flex-start; gap: 16rpx; padding: 26rpx; border-radius: $radius-v24rpx; background: $color-input; }
.detail-rule > view { flex: 1; min-width: 0; }
.detail-rule text:first-child { display: block; color: $color-primary; font-size: $font-size-v28rpx; font-weight: $font-weight-v600; line-height: $line-height-v1p4; }
.detail-rule text:last-child { display: block; margin-top: 10rpx; color: $color-textSecondary; font-size: $font-size-v24rpx; line-height: $line-height-v1p6; }
.detail-demo { display: block; margin-top: 24rpx; color: $color-textMuted; font-size: $font-size-v21rpx; line-height: $line-height-v1p5; }
.detail-actions { flex-shrink: 0; padding: 22rpx 28rpx calc(22rpx + env(safe-area-inset-bottom)); display: grid; grid-template-columns: minmax(164rpx, 1fr) minmax(0, 2.1fr); align-items: center; gap: 30rpx; border-top: 1rpx solid $color-border; background: $color-white; font-family: $font-body; }
.detail-price { display: flex; flex-direction: column; gap: 6rpx; min-width: 0; }
.detail-price__label { color: $color-textMuted; font-size: $font-size-v21rpx; line-height: $line-height-v1p4; }
.detail-price__amount { display: flex; align-items: baseline; gap: 5rpx; }
.detail-price__value { color: $color-accent; font-size: $font-size-v42rpx; font-weight: $font-weight-v700; line-height: $line-height-v1p3; }
.detail-price__unit { color: $color-accent; font-size: $font-size-v22rpx; line-height: $line-height-v1p4; }
.detail-actions button { height: 96rpx; margin: 0; padding: 0 28rpx; display: flex; justify-content: center; align-items: center; border-radius: $radius-v28rpx; color: $color-accentForeground; background: $color-accent; font-size: $font-size-v29rpx; font-weight: $font-weight-v600; line-height: $line-height-v1p3; }
.detail-actions button:active { opacity: .8; }
.hero-button, .detail-actions button { border: 0; }
.detail-host::after, .hero-button::after, .detail-actions button::after { border: 0; }
button:focus-visible { outline: 3rpx solid $color-primary; outline-offset: 3rpx; }
.detail-empty { padding: 100rpx 32rpx; display: flex; flex-direction: column; gap: 24rpx; color: $color-textSecondary; font-size: $font-size-v27rpx; }
.detail-empty button { margin: 18rpx 0; color: $color-white; background: $color-primary; border-radius: $radius-v24rpx; font-size: $font-size-v28rpx; }
@media (max-width: 350px) { .detail-facts { grid-template-columns: repeat(2, minmax(0, 1fr)); gap: 24rpx 0; } .detail-fact:nth-child(3)::before { display: none; } }
/* #ifdef H5 */
.meetup-detail { height: var(--qudaju-preview-layout-height, 100dvh) !important; }
/* #endif */
</style>
