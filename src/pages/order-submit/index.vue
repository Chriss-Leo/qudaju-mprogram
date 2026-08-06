<script setup lang="ts">
import { computed, ref } from "vue";
import { onLoad } from "@dcloudio/uni-app";
import { services } from "../../mocks/services";
import { getPlatformAdapter } from "../../platforms";

const serviceId = ref(services[0].id);
const bookingTime = ref("今天 14:00");
const note = ref("");
const submitting = ref(false);
const service = computed(() => services.find((item) => item.id === serviceId.value) || services[0]);
const serviceImage = computed(() => ({
  guide: "/static/posterwall/city-photo-walk.jpg",
  photo: "/static/posterwall/outdoor-walk.jpg",
  game: "/static/posterwall/boardgame-night.jpg",
  outdoor: "/static/posterwall/hero-frisbee.jpg",
})[service.value.visual]);

onLoad((query) => {
  if (query?.id) serviceId.value = query.id;
  if (query?.time) bookingTime.value = decodeURIComponent(query.time);
});

function back() { uni.navigateBack(); }
async function submitOrder() {
  if (submitting.value) return;
  submitting.value = true;
  await getPlatformAdapter().requestPayment(`demo-${service.value.id}`);
  uni.showToast({ title: "演示订单已创建", icon: "success" });
  setTimeout(() => uni.redirectTo({ url: "/pages/order-detail/index" }), 450);
}
</script>

<template>
  <view class="submit-page">
    <view class="top-nav"><button aria-label="返回" @tap="back"><view /></button><text>确认订单</text><view /></view>

    <view class="summary-card">
      <view class="summary-visual"><image :src="serviceImage" mode="aspectFill" /></view>
      <view class="summary-copy"><text>{{ service.title }}</text><text>{{ service.provider }}</text><text>¥{{ service.price }} / 次</text></view>
    </view>

    <view class="form-card">
      <view class="form-row"><text>预约时间</text><text>{{ bookingTime }} ›</text></view>
      <view class="form-row"><text>服务方式</text><text>{{ service.mode }}</text></view>
      <view class="form-row"><text>服务区域</text><text>上海市徐汇区 ›</text></view>
      <view class="form-row contact"><text>联系手机</text><text>138 **** 8026</text></view>
    </view>

    <view class="note-card"><text>给服务者留言</text><textarea v-model="note" maxlength="100" placeholder="选填，例如集合地点或需要提前准备的事项" /><text class="counter">{{ note.length }}/100</text></view>

    <view class="assurance"><view>✓</view><view><text>资金由平台托管</text><text>服务确认完成后再结算给服务者</text></view></view>

    <view class="notice-ticket">
      <view class="notice-ticket__head"><text>下单前确认</text><text>ORDER NOTE</text></view>
      <view><text>变更时间</text><text>请先与服务者确认新的可约档期</text></view>
      <view><text>取消退款</text><text>按服务开始时间与平台规则处理</text></view>
      <view><text>线下安全</text><text>建议在公共区域集合并全程平台沟通</text></view>
    </view>

    <view class="submit-bar">
      <view><text>合计</text><text>¥{{ service.price }}</text></view>
      <button :loading="submitting" @tap="submitOrder">微信/平台支付</button>
    </view>
  </view>
</template>

<style scoped lang="scss">
.submit-page { min-height: 100vh; padding: calc(48rpx + env(safe-area-inset-top)) 34rpx calc(150rpx + env(safe-area-inset-bottom)); }
.top-nav { height: 70rpx; display: grid; grid-template-columns: 64rpx 1fr 64rpx; align-items: center; text-align: center; font-size: 33rpx; font-weight: 800; }
.top-nav button { width: 64rpx; height: 64rpx; margin: 0; padding: 0; border-radius: 50%; background: #e6ebe5; }.top-nav button view { width: 19rpx; height: 19rpx; margin: 20rpx 0 0 25rpx; border: 4rpx solid #263129; border-top: 0; border-right: 0; transform: rotate(45deg); }
.summary-card { margin-top: 30rpx; padding: 24rpx; display: grid; grid-template-columns: 160rpx 1fr; gap: 22rpx; border-radius: 30rpx; background: #fff; }
.summary-visual { position: relative; height: 150rpx; overflow: hidden; border-radius: 22rpx; background: linear-gradient(145deg,#d8e3da,#4b7465); }.summary-visual--photo { background: linear-gradient(145deg,#edd6ca,#8c5b4e); }.summary-visual--game { background: linear-gradient(145deg,#ded7eb,#5b4d73); }.summary-visual > view:first-child { position: absolute; width: 55rpx; height: 55rpx; right: 18rpx; top: 18rpx; border-radius: 50%; background: rgba(255,240,190,.62); }.summary-visual > view:last-child { position: absolute; width: 25rpx; height: 72rpx; left: 66rpx; bottom: 12rpx; border-radius: 15rpx 15rpx 5rpx 5rpx; background: #173f34; }
.summary-copy { min-width: 0; display: flex; flex-direction: column; }.summary-copy text:first-child { font-size: 27rpx; line-height: 1.45; font-weight: 800; }.summary-copy text:nth-child(2) { margin-top: 10rpx; color: #7d8880; font-size: 21rpx; }.summary-copy text:last-child { margin-top: auto; color: #e75b41; font-size: 27rpx; font-weight: 800; }
.form-card, .note-card { margin-top: 24rpx; padding: 0 28rpx; border-radius: 30rpx; background: #fff; }.form-row { min-height: 88rpx; display: flex; align-items: center; justify-content: space-between; border-bottom: 1rpx solid #e5e9e4; font-size: 24rpx; }.form-row:last-child { border-bottom: 0; }.form-row text:last-child { color: #69756d; }.form-row.contact text:last-child { color: #2a362f; }
.note-card { position: relative; padding-top: 26rpx; padding-bottom: 24rpx; }.note-card > text:first-child { font-size: 25rpx; font-weight: 700; }.note-card textarea { width: 100%; height: 130rpx; margin-top: 18rpx; padding: 20rpx; border-radius: 20rpx; background: #f3f5f1; font-size: 23rpx; }.counter { position: absolute; right: 44rpx; bottom: 36rpx; color: #a0a8a2; font-size: 18rpx; }
.assurance { margin-top: 24rpx; padding: 23rpx 25rpx; display: flex; align-items: center; gap: 16rpx; border-radius: 24rpx; color: #315e4f; background: #e5eee8; }.assurance > view:first-child { width: 40rpx; height: 40rpx; display: grid; place-items: center; border: 3rpx solid currentColor; border-radius: 50%; font-weight: 800; }.assurance > view:last-child { display: flex; flex-direction: column; gap: 6rpx; }.assurance > view:last-child text:first-child { font-size: 23rpx; font-weight: 700; }.assurance > view:last-child text:last-child { color: #6c7d74; font-size: 19rpx; }
.submit-bar { position: fixed; z-index: 20; left: 0; right: 0; bottom: 0; padding: 18rpx 34rpx calc(18rpx + env(safe-area-inset-bottom)); display: grid; grid-template-columns: 1fr 1.3fr; gap: 20rpx; align-items: center; border-top: 1rpx solid #e1e6e0; background: rgba(255,255,255,.97); }.submit-bar > view { display: flex; align-items: baseline; gap: 10rpx; }.submit-bar > view text:first-child { color: #78837c; font-size: 21rpx; }.submit-bar > view text:last-child { color: #e75a40; font-size: 37rpx; font-weight: 900; }.submit-bar button { height: 84rpx; border-radius: 27rpx; color: #fff; background: #eb5c42; font-size: 26rpx; font-weight: 800; }
.submit-page { color: #153f35; background: #f3f5f0; }
.top-nav button { border-radius: 14rpx; }
.summary-card { position: relative; border-radius: 12rpx; background: #fffdf6; box-shadow: 0 12rpx 26rpx rgba(21,63,53,.1); }
.summary-card::before, .summary-card::after { content: ""; position: absolute; top: 50%; width: 22rpx; height: 22rpx; border-radius: 50%; background: #f3f5f0; transform: translateY(-50%); }.summary-card::before { left: -11rpx; }.summary-card::after { right: -11rpx; }
.summary-visual { border-radius: 8rpx; background: #b8cfbf; }.summary-visual image { width: 100%; height: 100%; display: block; }
.form-card, .note-card { border-radius: 12rpx; background: #fffdf6; box-shadow: 0 8rpx 20rpx rgba(21,63,53,.06); }
.form-row { border-bottom-style: dashed; }.note-card textarea { border: 2rpx dashed #bdc9c0; border-radius: 8rpx; background: #f7f5ee; }
.assurance { border-radius: 8rpx; background: #cbdacf; }
.submit-bar { background: #fffdf6; }.submit-bar button { border-radius: 14rpx; background: #ef644a; }
.notice-ticket { margin-top: 24rpx; padding: 0 25rpx; border: 2rpx dashed #aebcb2; border-radius: 8rpx; background: #fffdf6; }.notice-ticket__head { min-height: 70rpx; display: flex; align-items: center; justify-content: space-between; border-bottom: 2rpx dashed #cbd3cc; }.notice-ticket__head text:first-child { font-size: 24rpx; font-weight: 900; }.notice-ticket__head text:last-child { color: #ef644a; font-size: 16rpx; letter-spacing: 2rpx; }.notice-ticket > view:not(.notice-ticket__head) { min-height: 67rpx; display: grid; grid-template-columns: 120rpx 1fr; align-items: center; border-bottom: 1rpx solid #e3e7e2; font-size: 20rpx; }.notice-ticket > view:last-child { border-bottom: 0; }.notice-ticket > view > text:first-child { color: #153f35; font-weight: 800; }.notice-ticket > view > text:last-child { color: #6f7b73; line-height: 1.45; }
</style>
