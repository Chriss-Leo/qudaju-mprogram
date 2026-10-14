<script setup lang="ts">
import { computed, ref, onMounted, onUnmounted } from "vue";
import { onLoad, onShow, onResize } from "@dcloudio/uni-app";
import { theme } from "../../theme";
import { getGroup, myRegistration, isMyHost, statusLabel, meetupState } from "../../stores/meetups";
import { openFlow } from "../../utils/flow-navigation";
import { ensureLogin } from "../../utils/auth-navigation";
import SurfaceHeader from "../../components/SurfaceHeader.vue";
import UiIcon from "../../components/UiIcon.vue";
import GroupPhotoGrid from "../../components/GroupPhotoGrid.vue";
import ParticipantAvatarStack from "../../components/ParticipantAvatarStack.vue";

const groupId = ref("game-01");
const viewportHeight = ref(uni.getWindowInfo().windowHeight);
onResize(({ size }) => { if (size) viewportHeight.value = size.windowHeight; });
const imageFailed = ref(false);
const hostFailed = ref(false);
const group = computed(() => getGroup(groupId.value));
const registration = computed(() => myRegistration(groupId.value));
const closed = computed(() => meetupState.ended.includes(groupId.value));
const own = computed(() => isMyHost(groupId.value));
const applicationLabel = computed(() => own.value ? '管理报名成员' : closed.value ? '活动已结束' : registration.value && ['review','payment'].includes(registration.value.status) ? statusLabel[registration.value.status] : group.value?.isJoinedByMe ? '已报名' : group.value?.admission === '局主审核' ? '申请加入' : '报名加入');
const tags: Record<string, string[]> = { 桌游: ["轻策略", "新手友好"], 陪拍: ["城市漫步", "互拍"], 向导: ["街区漫步", "城市故事"], 户外: ["轻徒步", "慢节奏"] };
const facts = computed(() => group.value ? [
  { key: "time", value: group.value.time, label: "活动时间" },
  { key: "area", value: group.value.area, label: "公开区域" },
] : []);
const remaining = computed(() => group.value ? Math.max(0, group.value.capacity - group.value.joined) : 0);
const steps = computed(() => {
  if (!group.value) return [];
  const reviewed = group.value.admission === "局主审核";
  const paid = group.value.fund > 0;
  return [
    { icon: "info", title: reviewed ? "提交申请" : "提交报名", description: reviewed ? "局主审核报名信息" : "无需局主审核" },
    ...(reviewed ? [{ icon: "people", title: "审核通过", description: paid ? "等待付款，尚未占正式名额" : "正式占名额，可查看集合细节" }] : []),
    ...(paid ? [{ icon: "wallet", title: "完成付款", description: "正式占名额，可查看集合细节" }] : !reviewed ? [{ icon: "people", title: "报名成功", description: "正式占名额，可查看集合细节" }] : []),
  ];
});
const rule = computed(() => {
  if (!group.value) return "";
  const paid = group.value.fund > 0;
  if (group.value.admission === "局主审核") return paid
    ? "局主审核通过后付款，付款完成才占正式名额。通过后可查看集合细节。"
    : "提交申请后，局主会进行审核。通过后占正式名额，并可查看集合细节。";
  return paid ? "付款完成后占正式名额，报名成功后可查看集合细节。" : "报名成功后占正式名额，并可查看集合细节。";
});
onLoad(query => { if (query?.id) groupId.value = query.id; imageFailed.value = false; hostFailed.value = false; });
// H5 reuses the route component when only the query changes (including history).
// #ifdef H5
function syncQuery() {
  const hash = window.location.hash;
  if (!hash.startsWith("#/pages/service-detail/index")) return;
  groupId.value = new URLSearchParams(hash.split("?")[1] || "").get("id") || "game-01";
  imageFailed.value = false;
  hostFailed.value = false;
}
onMounted(() => { syncQuery(); window.addEventListener("hashchange", syncQuery); });
onUnmounted(() => window.removeEventListener("hashchange", syncQuery));
// #endif
onShow(() => { uni.setNavigationBarColor({ frontColor: "#000000", backgroundColor: theme.colors.white }); });
function back() { uni.navigateBack({ fail: () => uni.redirectTo({ url: "/pages/groups/index" }) }); }
function preview() { const urls = group.value?.photos.slice(0, 3) || []; if (urls.length) uni.previewImage({ urls, current: urls[0] }); }
function apply() {
  if (!ensureLogin(`/pages/service-detail/index?id=${encodeURIComponent(groupId.value)}`)) return;
  if (own.value) openFlow('group-members', groupId.value);
  else if (registration.value && ['review','payment','joined','ongoing'].includes(registration.value.status)) openFlow('my-registrations');
  else if (!closed.value) openFlow('group-apply', groupId.value);
}
function hostInfo() { uni.showToast({ title: "局主信息卡片尚未接入", icon: "none" }); }
</script>

<template>
  <view class="meetup-detail" :style="{ height: viewportHeight + 'px' }" data-design-contract="Approved 2026-10-14: edge-to-edge photograph, continuous white panel, single-line title, host portrait, enrolled avatar stack, fixed application footer.">
    <SurfaceHeader title="组局详情" back @back="back" />
    <scroll-view class="detail-scroll" scroll-y :show-scrollbar="false">
    <view v-if="group" class="detail-content">
      <button class="hero-button" aria-label="查看活动示例大图" @tap="preview">
        <image v-if="group.photos[0] && !imageFailed" class="detail-photo" :src="group.photos[0]" mode="aspectFill" @error="imageFailed = true" />
        <view v-else class="photo-error">图片暂时未能加载</view>
      </button>
      <view class="detail-panel">
      <text class="detail-title" :aria-label="group.title">{{ group.title }}</text>
      <view class="detail-tags"><text>{{ group.category }}</text><text v-for="tag in tags[group.category] || []" :key="tag">{{ tag }}</text></view>
      <view class="detail-facts">
      <button class="detail-host" @tap="hostInfo">
        <image v-if="!hostFailed" class="detail-host__avatar" :src="group.hostAvatar" mode="aspectFill" :aria-label="`${group.host}示例头像`" @error="hostFailed = true" />
        <view v-else class="detail-host__avatar detail-host__initial">{{ group.host.slice(0, 1) }}</view>
        <view class="detail-host__copy"><text>{{ group.host }}</text><text>发起人</text></view>
      </button>
        <view v-for="fact in facts" :key="fact.key" class="detail-fact"><text class="detail-fact__label">{{ fact.label }}</text><text class="detail-fact__value">{{ fact.value }}</text></view>
      </view>
      <view class="detail-enrolled">
        <text class="detail-heading">已报名</text>
        <view class="detail-enrolled__row"><ParticipantAvatarStack :participants="group.participants" :total="group.joined" /><view class="detail-enrolled__count"><view><text>{{ group.joined }}</text><text> / {{ group.capacity }} 人</text></view><text>{{ remaining ? `还剩 ${remaining} 个名额` : '名额已满' }}</text></view></view>
      </view>
      <view class="detail-about"><text class="detail-heading">活动介绍</text><text class="detail-description">{{ group.description }}</text><GroupPhotoGrid v-if="group.photos.length > 1" class="detail-about__photos" :photos="group.photos.slice(1)" :label="group.category" /></view>
      <view class="detail-place"><text class="detail-heading">集合地点</text><view><UiIcon name="location" /><view><text>{{ group.area }}</text><text>正式报名成功后可查看集合细节</text></view></view></view>
      <view class="detail-joining"><text class="detail-heading">加入方式</text><text class="detail-description">{{ rule }}</text><view class="detail-rule"><view v-for="step in steps" :key="step.title" class="detail-step"><view class="detail-step__icon"><UiIcon :name="step.icon" /></view><view><text>{{ step.title }}</text><text>{{ step.description }}</text></view></view></view></view>
      <text class="detail-demo">活动、人物及图片均为示例，暂不开放真实报名。</text>
      </view>
    </view>
    <view v-else class="detail-empty"><text class="detail-heading">暂未找到这个局</text><text>返回列表，看看其他同城活动。</text><button @tap="back">返回列表</button></view>
    </scroll-view>
    <view v-if="group" class="detail-actions">
      <view class="detail-price">
        <text class="detail-price__label">{{ group.fund ? '活动基金' : '活动费用' }}</text>
        <view class="detail-price__amount"><text v-if="group.fund" class="detail-price__unit">¥</text><text class="detail-price__value">{{ group.fund || '免费' }}</text><text v-if="group.fund" class="detail-price__unit">/ 人</text></view>
      </view>
      <button class="detail-apply action-accent" :class="{ 'detail-apply--joined': group.isJoinedByMe || closed }" :disabled="closed && !own" @tap="apply"><text>{{ applicationLabel }}</text></button>
    </view>
  </view>
</template>

<style scoped lang="scss">
.meetup-detail { display: flex; flex-direction: column; overflow: hidden; background: $color-white; color: $color-text; }
.meetup-detail :deep(.surface-header) { flex-shrink: 0; padding-bottom: 14rpx; }
.detail-scroll { flex: 1; height: 0; min-height: 0; }
.detail-content { padding: 0; }
.hero-button { display: block; margin: 0; padding: 0; width: 100%; height: 540rpx; overflow: hidden; border-radius: $radius-v0; background: $color-card; }
.detail-panel { position: relative; margin-top: -32rpx; padding: 36rpx 28rpx 40rpx; border-radius: $radius-v44rpx_44rpx_0_0; background: $color-white; }
.detail-photo { display: block; width: 100%; height: 100%; }
.photo-error { height: 100%; display: flex; align-items: center; justify-content: center; color: $color-textMuted; font-size: $font-size-v26rpx; }
.detail-title { display: block; overflow: hidden; text-overflow: ellipsis; white-space: nowrap; font-size: $font-size-v38rpx; font-weight: $font-weight-v600; line-height: $line-height-v1p4; }
.detail-tags { display: flex; flex-wrap: wrap; gap: 14rpx; margin-top: 20rpx; }
.detail-tags text { padding: 7rpx 18rpx; color: $color-primary; background: $color-tag; border-radius: $radius-v999rpx; font-size: $font-size-v23rpx; line-height: $line-height-v1p4; }
.detail-facts { margin-top: 24rpx; padding: 32rpx 0; display: grid; grid-template-columns: repeat(3, minmax(0, 1fr)); align-items: center; border-bottom: 1rpx solid $color-border; }
.detail-fact { display: flex; flex-direction: column; align-items: flex-start; gap: 8rpx; min-width: 0; text-align: left; }
.detail-fact__value { display: block; white-space: nowrap; font-size: $font-size-v28rpx; font-weight: $font-weight-v600; line-height: $line-height-v1p4; }
.detail-fact__label { display: block; font-size: $font-size-v23rpx; color: $color-textMuted; line-height: $line-height-v1p4; }
.detail-host { display: flex; align-items: center; justify-content: flex-start; gap: 16rpx; width: 100%; min-width: 0; min-height: 44px; margin: 0; padding: 0; border-radius: $radius-v0; background: transparent; text-align: left; }
.detail-host__avatar { display: block; flex: 0 0 80rpx; width: 80rpx; height: 80rpx; border-radius: $radius-v50percent; background: $color-tag; }
.detail-host__initial { display: flex; align-items: center; justify-content: center; color: $color-primary; font-size: $font-size-v32rpx; }
.detail-host__copy { display: flex; flex-direction: column; gap: 8rpx; min-width: 0; }
.detail-host__copy text:first-child { overflow: hidden; text-overflow: ellipsis; white-space: nowrap; font-size: $font-size-v28rpx; font-weight: $font-weight-v600; line-height: $line-height-v1p4; color: $color-text; }
.detail-host__copy text:last-child { font-size: $font-size-v23rpx; line-height: $line-height-v1p4; color: $color-textMuted; }
.detail-enrolled { padding: 30rpx 0; border-bottom: 1rpx solid $color-border; }
.detail-enrolled__row { display: flex; align-items: center; justify-content: space-between; gap: 16rpx; margin-top: 20rpx; }
.detail-enrolled__count { flex-shrink: 0; }
.detail-enrolled__count > view { font-size: $font-size-v28rpx; line-height: $line-height-v1p4; }
.detail-enrolled__count > view text:first-child { color: $color-primary; font-size: $font-size-v42rpx; font-weight: $font-weight-v700; }
.detail-enrolled__count > text { display: block; margin-top: 5rpx; color: $color-textMuted; font-size: $font-size-v23rpx; line-height: $line-height-v1p4; }
.detail-about { padding: 32rpx 0 24rpx; }
.detail-heading { display: block; font-size: $font-size-v32rpx; font-weight: $font-weight-v700; line-height: $line-height-v1p4; }
.detail-description { display: block; margin-top: 18rpx; color: $color-textSecondary; font-size: $font-size-v28rpx; line-height: $line-height-v1p75; }
.detail-about__photos { margin-top: 22rpx; }
.detail-place { padding: 30rpx 0; border-top: 1rpx solid $color-border; border-bottom: 1rpx solid $color-border; }
.detail-place > view { display: flex; gap: 20rpx; margin-top: 20rpx; }
.detail-place text:not(.detail-heading) { display: block; font-size: $font-size-v28rpx; line-height: $line-height-v1p5; }
.detail-place > view text:last-child { margin-top: 8rpx; color: $color-textMuted; font-size: $font-size-v24rpx; }
.detail-joining { padding-top: 30rpx; }
.detail-rule { display: flex; flex-direction: column; gap: 24rpx; margin-top: 20rpx; padding: 26rpx; border-radius: $radius-v24rpx; background: $color-input; }
.detail-step { display: flex; align-items: flex-start; gap: 20rpx; }
.detail-step__icon { width: 64rpx; height: 64rpx; display: flex; align-items: center; justify-content: center; flex-shrink: 0; border-radius: $radius-v50percent; background: $color-tag; }
.detail-step__icon :deep(.ui-icon) { width: 36rpx; height: 36rpx; }
.detail-step > view:last-child { flex: 1; min-width: 0; }
.detail-step text:first-child { display: block; color: $color-text; font-size: $font-size-v28rpx; font-weight: $font-weight-v600; line-height: $line-height-v1p4; }
.detail-step text:last-child { display: block; margin-top: 8rpx; color: $color-textSecondary; font-size: $font-size-v24rpx; line-height: $line-height-v1p6; }
.detail-demo { display: block; margin-top: 24rpx; color: $color-textMuted; font-size: $font-size-v21rpx; line-height: $line-height-v1p5; }
.detail-actions { flex-shrink: 0; padding: 22rpx 28rpx calc(22rpx + env(safe-area-inset-bottom)); display: grid; grid-template-columns: minmax(164rpx, 1fr) minmax(0, 2.1fr); align-items: center; gap: 30rpx; border-top: 1rpx solid $color-border; background: $color-white; font-family: $font-body; }
.detail-price { display: flex; flex-direction: column; gap: 6rpx; min-width: 0; }
.detail-price__label { color: $color-textMuted; font-size: $font-size-v21rpx; line-height: $line-height-v1p4; }
.detail-price__amount { display: flex; align-items: baseline; gap: 5rpx; }
.detail-price__value { color: $color-accent; font-size: $font-size-v42rpx; font-weight: $font-weight-v700; line-height: $line-height-v1p3; }
.detail-price__unit { color: $color-accent; font-size: $font-size-v22rpx; line-height: $line-height-v1p4; }
.detail-actions button { height: 96rpx; margin: 0; padding: 0 28rpx; display: flex; justify-content: center; align-items: center; border-radius: $radius-v28rpx; color: $color-accentForeground; background: $color-accent; font-size: $font-size-v29rpx; font-weight: $font-weight-v600; line-height: $line-height-v1p3; }
.detail-actions button:active { opacity: .8; }
.detail-actions .detail-apply--joined { color: $color-textMuted; background: $color-card; opacity: 1; }
.hero-button, .detail-actions button { border: 0; }
.detail-host::after, .hero-button::after, .detail-actions button::after { border: 0; }
button:focus-visible { outline: 3rpx solid $color-primary; outline-offset: 3rpx; }
.detail-empty { padding: 100rpx 32rpx; display: flex; flex-direction: column; gap: 24rpx; color: $color-textSecondary; font-size: $font-size-v27rpx; }
.detail-empty button { margin: 18rpx 0; color: $color-white; background: $color-primary; border-radius: $radius-v24rpx; font-size: $font-size-v28rpx; }
@media (max-width: 350px) { .detail-enrolled__count > text { font-size: $font-size-v21rpx; } }
/* #ifdef H5 */
.meetup-detail { height: var(--qudaju-preview-layout-height, 100dvh) !important; }
/* #endif */
</style>
