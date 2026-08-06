<script setup lang="ts">
import { computed, ref } from "vue";
import { onLoad } from "@dcloudio/uni-app";
import ServiceVisual from "../../components/ServiceVisual.vue";
import { services } from "../../mocks/services";

const selectedTime = ref("今天 14:00");
const liked = ref(false);
const serviceId = ref(services[0].id);
const service = computed(() => services.find((item) => item.id === serviceId.value) || services[0]);

onLoad((query) => {
  if (query?.id) serviceId.value = query.id;
});

function back() {
  uni.navigateBack({ fail: () => uni.redirectTo({ url: "/pages/index/index" }) });
}
function reserve() {
  uni.navigateTo({ url: `/pages/order-submit/index?id=${service.value.id}&time=${encodeURIComponent(selectedTime.value)}` });
}
function consult() {
  uni.showToast({ title: "咨询功能将在联调阶段接入", icon: "none" });
}
function demo(message: string) { uni.showToast({ title: message, icon: "none" }); }
</script>

<template>
  <view class="detail-page">
    <view class="hero">
      <ServiceVisual :type="service.visual" />
      <view class="hero__nav">
        <button class="circle-action back" aria-label="返回" @tap="back"><view /></button>
        <view class="hero__right">
          <button class="circle-action share" aria-label="分享" @tap="demo('分享能力将在平台联调阶段接入')"><view /><view /><view /></button>
          <button class="circle-action heart" :class="{ liked }" aria-label="收藏" @tap="liked = !liked"><view>♡</view></button>
        </view>
      </view>
    </view>

    <view class="provider-card">
      <view class="avatar"><text>林</text></view>
      <view class="provider-card__info">
        <view class="provider-card__name"><text>{{ service.provider }}</text><text class="verified">身份已核验</text></view>
        <view class="provider-card__meta"><text>3年经验</text><text>·</text><text>30分钟内回复</text><text>·</text><text>服务{{ service.orders }}次</text></view>
      </view>
    </view>

    <view class="content-card service-copy">
      <text class="service-title">{{ service.title }}</text>
      <view class="tag-row"><text v-for="tag in service.tags" :key="tag">{{ tag }}</text><text>{{ service.mode }}</text></view>
      <text class="description">服务范围、集合地点、具体流程和交付内容均在订单内确认，可先咨询服务者了解细节。</text>
    </view>

    <view class="trust-strip">
      <view><text class="trust-icon">◇</text><text>平台担保</text></view>
      <view><text class="trust-icon">↶</text><text>未服务可退</text></view>
      <view><text class="trust-icon">◉</text><text>客服介入</text></view>
    </view>

    <view class="content-card booking-card">
      <view class="content-heading"><text>选择可约时间</text><text class="more" @tap="demo('日期选择器已预留')">更多日期</text></view>
      <scroll-view scroll-x :show-scrollbar="false">
        <view class="time-list">
          <button v-for="time in ['今天 14:00', '明天 10:00', '周六 13:00', '周日 16:00']" :key="time" :class="{ selected: selectedTime === time }" @tap="selectedTime = time">{{ time }}</button>
        </view>
      </scroll-view>
    </view>

    <view class="content-card review-card">
      <view class="content-heading"><text>用户评价</text><text class="review-count" @tap="demo('评价列表已预留')">全部 {{ service.orders }}</text></view>
      <view class="review-box">
        <text>路线安排舒服，讲解自然，拍照也很有耐心。集合地点沟通非常顺畅。</text>
        <view><text>南枝</text><text>2天前</text></view>
      </view>
    </view>

    <view class="detail-actions">
      <button class="consult" @tap="consult">先咨询</button>
      <button class="reserve" @tap="reserve"><text>¥{{ service.price }}起</text> 立即预约</button>
    </view>
  </view>
</template>

<style scoped lang="scss">
.detail-page { min-height: 100vh; padding-bottom: calc(140rpx + env(safe-area-inset-bottom)); background: #fff; }
.hero { position: relative; height: 470rpx; overflow: hidden; }
.hero :deep(.visual) { height: 470rpx; }
.hero :deep(.visual__figure) { width: 72rpx; height: 150rpx; left: 64%; bottom: 62rpx; }
.hero :deep(.visual__head) { width: 49rpx; height: 49rpx; top: -36rpx; left: 11rpx; }
.hero__nav { position: absolute; left: 34rpx; right: 34rpx; top: calc(48rpx + env(safe-area-inset-top)); display: flex; align-items: center; justify-content: space-between; }
.hero__right { display: flex; gap: 16rpx; }
.circle-action { width: 72rpx; height: 72rpx; margin: 0; padding: 0; border-radius: 50%; color: #fff; background: rgba(18, 31, 26, .48); backdrop-filter: blur(10px); }
.back view { width: 22rpx; height: 22rpx; margin: 23rpx 0 0 27rpx; border: 5rpx solid #fff; border-top: 0; border-right: 0; transform: rotate(45deg); }
.heart view { font-size: 49rpx; line-height: 66rpx; }
.heart.liked { color: #fff; background: #eb5c42; }
.share { position: relative; }
.share view { position: absolute; width: 10rpx; height: 10rpx; border: 4rpx solid #fff; border-radius: 50%; }
.share view:nth-child(1) { left: 18rpx; top: 30rpx; }
.share view:nth-child(2) { right: 17rpx; top: 17rpx; }
.share view:nth-child(3) { right: 17rpx; bottom: 15rpx; }
.share::before, .share::after { content: ""; position: absolute; width: 28rpx; height: 4rpx; left: 27rpx; background: #fff; transform-origin: left center; }
.share::before { top: 32rpx; transform: rotate(-28deg); }
.share::after { top: 37rpx; transform: rotate(28deg); }
.provider-card { position: relative; z-index: 2; margin-top: -42rpx; padding: 34rpx 38rpx; min-height: 154rpx; display: flex; align-items: center; border-radius: 44rpx 44rpx 0 0; background: #fff; }
.avatar { position: relative; width: 94rpx; height: 94rpx; flex: 0 0 auto; overflow: hidden; border-radius: 50%; background: #d8b391; }
.avatar__face { position: absolute; width: 49rpx; height: 62rpx; left: 23rpx; top: 21rpx; border-radius: 50% 50% 45% 45%; background: #d6a27b; }
.avatar__hair { position: absolute; z-index: 2; width: 73rpx; height: 66rpx; left: 10rpx; top: 4rpx; border-radius: 50% 50% 42% 55%; background: #2e2d27; }
.provider-card__info { min-width: 0; margin-left: 24rpx; }
.provider-card__name { display: flex; align-items: center; gap: 14rpx; font-size: 30rpx; font-weight: 700; }
.verified { padding: 7rpx 13rpx; border-radius: 10rpx; color: #336c58; background: #e2eee7; font-size: 20rpx; font-weight: 600; }
.provider-card__meta { margin-top: 15rpx; display: flex; gap: 16rpx; color: #7e8881; font-size: 21rpx; }
.content-card { padding: 34rpx 40rpx; border-top: 1rpx solid #e5e9e4; background: #fff; }
.service-title { display: block; font-size: 38rpx; font-weight: 900; }
.tag-row { margin-top: 24rpx; display: flex; gap: 12rpx; }
.tag-row text { padding: 10rpx 16rpx; border-radius: 10rpx; color: #6d786f; background: #f2f4f0; font-size: 22rpx; }
.description { display: block; margin-top: 24rpx; color: #6d776f; font-size: 25rpx; line-height: 1.75; }
.trust-strip { min-height: 84rpx; padding: 0 40rpx; display: flex; align-items: center; justify-content: space-between; color: #336653; background: #e8f0eb; font-size: 23rpx; }
.trust-strip > view { display: flex; align-items: center; gap: 9rpx; }
.trust-icon { font-size: 31rpx; font-weight: 800; }
.content-heading { display: flex; align-items: center; justify-content: space-between; font-size: 30rpx; font-weight: 700; }
.more { color: #eb5c42; font-size: 24rpx; font-weight: 500; }
.time-list { width: max-content; margin-top: 24rpx; padding-right: 40rpx; display: flex; gap: 16rpx; }
.time-list button { height: 72rpx; padding: 0 27rpx; border: 2rpx solid #e0e5df; border-radius: 22rpx; background: #f5f7f3; color: #27322b; font-size: 24rpx; }
.time-list button.selected { border-color: #eb5c42; color: #fff; background: #eb5c42; }
.review-count { color: #7c867f; font-size: 23rpx; font-weight: 500; }
.review-box { margin-top: 24rpx; padding: 27rpx; border-radius: 24rpx; background: #f5f7f3; color: #28342d; font-size: 24rpx; line-height: 1.65; }
.review-box view { margin-top: 22rpx; display: flex; justify-content: space-between; color: #9ba39d; font-size: 21rpx; }
.detail-actions { position: fixed; z-index: 20; left: 0; right: 0; bottom: 0; padding: 20rpx 40rpx calc(20rpx + env(safe-area-inset-bottom)); display: grid; grid-template-columns: 1fr 1.35fr; gap: 18rpx; border-top: 1rpx solid #e1e6e0; background: rgba(255,255,255,.97); }
.detail-actions button { height: 88rpx; border-radius: 28rpx; font-size: 28rpx; }
.consult { border: 2rpx solid #dfe4de; color: #253128; background: #f5f7f3; }
.reserve { color: #fff; background: #eb5c42; font-weight: 700; }
.reserve text { margin-right: 10rpx; font-size: 24rpx; font-weight: 500; }
.avatar { display: grid; place-items: center; color: #fffdf6; background: #153f35; }.avatar text { font-size: 36rpx; font-weight: 900; transform: rotate(-5deg); }.avatar__face, .avatar__hair { display: none; }
.detail-page { color: #153f35; background: #f3f5f0; }
.provider-card { margin: -34rpx 24rpx 0; padding: 26rpx; border-radius: 14rpx; background: #fffdf6; box-shadow: 0 14rpx 30rpx rgba(21,63,53,.14); }
.provider-card::before, .provider-card::after { content: ""; position: absolute; top: 50%; width: 22rpx; height: 22rpx; border-radius: 50%; background: #f3f5f0; }.provider-card::before { left: -11rpx; }.provider-card::after { right: -11rpx; }
.verified { border-radius: 0; color: #153f35; background: #b8cfbf; }
.content-card { margin: 18rpx 24rpx 0; padding: 28rpx; border: 0; border-radius: 12rpx; background: #fffdf6; box-shadow: 0 9rpx 22rpx rgba(21,63,53,.07); }
.service-copy { border-top: 4rpx solid #ef644a; }.tag-row text { border-radius: 0; color: #153f35; background: #e1ebe4; }
.trust-strip { margin: 18rpx 24rpx 0; border-top: 2rpx dashed #93aa9b; border-bottom: 2rpx dashed #93aa9b; background: #cbdacf; }
.time-list button { border-radius: 999rpx; background: transparent; }.time-list button.selected { border-color: #153f35; background: #153f35; }
.review-box { border: 2rpx dashed #bdc9c0; border-radius: 8rpx; background: #f7f5ee; }
.detail-actions { border-top: 1rpx solid rgba(21,63,53,.16); background: #fffdf6; }.detail-actions button { border-radius: 14rpx; }.reserve { background: #ef644a; }
</style>
