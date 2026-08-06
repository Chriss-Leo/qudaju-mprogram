<script setup lang="ts">
import AppTabBar from "../../components/AppTabBar.vue";
import { getPlatformAdapter } from "../../platforms";

function back() {
  uni.navigateBack({ fail: () => uni.redirectTo({ url: "/pages/index/index" }) });
}
function demo(message: string) {
  uni.showToast({ title: message, icon: "none" });
}
</script>

<template>
  <view class="order-page">
    <view class="top-nav">
      <button class="back" aria-label="返回" @tap="back"><view /></button>
      <text>订单详情</text>
      <button class="chat" aria-label="联系服务者" @tap="demo('聊天能力将在联调阶段接入')"><view /><view /></button>
    </view>

    <view class="status-panel">
      <view class="status-panel__icon"><view>✓</view></view>
      <view><text class="status-panel__title">服务者已接单</text><text class="status-panel__copy">请按预约时间到达集合地点</text></view>
    </view>

    <view class="progress-card">
      <view v-for="(step, index) in ['已下单', '已接单', '待服务', '已完成']" :key="step" class="progress-step" :class="{ done: index <= 1 }">
        <view class="progress-step__dot">{{ index + 1 }}</view><text>{{ step }}</text>
      </view>
      <view class="progress-line progress-line--1" /><view class="progress-line progress-line--2" /><view class="progress-line progress-line--3" />
    </view>

    <view class="order-card provider-row">
      <view class="avatar"><text>林</text></view>
      <view class="provider-copy"><text>林野 · 城市向导</text><text>身份已核验 · 30分钟内回复</text></view>
      <button @tap="demo('已为你准备联系入口')">联系</button>
    </view>

    <view class="order-card detail-card">
      <text class="card-title">梧桐区漫步与城市摄影</text>
      <view class="detail-row"><text>预约时间</text><text>今天 14:00—16:00</text></view>
      <view class="detail-row"><text>服务方式</text><text>线下见面</text></view>
      <view class="detail-row"><text>集合地点</text><text>上海图书馆东馆入口</text></view>
      <view class="detail-row"><text>订单编号</text><text>QD202608050128</text></view>
      <view class="detail-row total"><text>实付金额</text><text>¥128.00</text></view>
    </view>

    <view class="safety-note">
      <view class="safety-note__icon">!</view>
      <view><text>线下服务安全提示</text><text>请在公共区域见面，不要脱离平台私下转账。</text></view>
    </view>

    <view class="order-card menu-card">
      <view @tap="demo('取消退款规则将在订单接口接入后展示')"><text>取消与退款</text><text>›</text></view>
      <view @tap="getPlatformAdapter().openCustomerService()"><text>平台客服</text><text>›</text></view>
      <view @tap="demo('订单投诉入口已预留')"><text>投诉与申诉</text><text>›</text></view>
    </view>

    <AppTabBar active="orders" />
  </view>
</template>

<style scoped lang="scss">
.order-page { min-height: 100vh; padding: calc(46rpx + env(safe-area-inset-top)) 34rpx 170rpx; }
.top-nav { height: 72rpx; display: grid; grid-template-columns: 64rpx 1fr 64rpx; align-items: center; text-align: center; font-size: 32rpx; font-weight: 800; }
.top-nav button { width: 64rpx; height: 64rpx; margin: 0; padding: 0; border-radius: 50%; background: #e8ece7; }
.back view { width: 19rpx; height: 19rpx; margin: 20rpx 0 0 25rpx; border: 4rpx solid #263129; border-top: 0; border-right: 0; transform: rotate(45deg); }
.chat { position: relative; }
.chat view:first-child { position: absolute; width: 30rpx; height: 24rpx; left: 15rpx; top: 16rpx; border: 4rpx solid #263129; border-radius: 10rpx; }
.chat view:last-child { position: absolute; width: 10rpx; height: 10rpx; left: 20rpx; top: 39rpx; border: 4rpx solid #263129; border-top: 0; border-right: 0; transform: rotate(-20deg); }
.status-panel { min-height: 180rpx; margin-top: 28rpx; padding: 34rpx; display: flex; align-items: center; gap: 24rpx; border-radius: 34rpx; color: #fff; background: #173f34; }
.status-panel__icon { width: 72rpx; height: 72rpx; display: grid; place-items: center; border-radius: 24rpx; background: rgba(255,255,255,.13); }
.status-panel__icon view { width: 38rpx; height: 38rpx; display: grid; place-items: center; border: 4rpx solid #fff; border-radius: 50%; font-size: 22rpx; }
.status-panel > view:last-child { display: flex; flex-direction: column; }
.status-panel__title { font-size: 34rpx; font-weight: 800; }
.status-panel__copy { margin-top: 12rpx; color: rgba(255,255,255,.7); font-size: 22rpx; }
.progress-card { position: relative; margin-top: 22rpx; padding: 28rpx 16rpx; display: grid; grid-template-columns: repeat(4, 1fr); border-radius: 28rpx; background: #fff; }
.progress-step { position: relative; z-index: 2; display: flex; flex-direction: column; align-items: center; gap: 12rpx; color: #9aa29c; font-size: 20rpx; }
.progress-step__dot { width: 44rpx; height: 44rpx; display: grid; place-items: center; border: 4rpx solid #d7ddd7; border-radius: 50%; background: #fff; }
.progress-step.done { color: #173f34; font-weight: 700; }
.progress-step.done .progress-step__dot { border-color: #173f34; color: #fff; background: #173f34; }
.progress-line { position: absolute; top: 49rpx; width: 17%; height: 4rpx; background: #dfe4de; }
.progress-line--1 { left: 17%; background: #173f34; }.progress-line--2 { left: 42%; }.progress-line--3 { left: 67%; }
.order-card { margin-top: 22rpx; border-radius: 28rpx; background: #fff; }
.provider-row { padding: 28rpx; display: flex; align-items: center; }
.avatar { width: 76rpx; height: 76rpx; flex: 0 0 auto; overflow: hidden; border-radius: 50%; background: #d6aa86; }
.avatar view { width: 55rpx; height: 67rpx; margin: 17rpx auto 0; border-radius: 50% 50% 20% 20%; background: #2d312c; }
.provider-copy { min-width: 0; margin-left: 20rpx; display: flex; flex: 1; flex-direction: column; gap: 9rpx; }
.provider-copy text:first-child { font-size: 27rpx; font-weight: 700; }.provider-copy text:last-child { color: #87918a; font-size: 20rpx; }
.provider-row button { width: auto; height: 60rpx; margin: 0; padding: 0 21rpx; border-radius: 20rpx; color: #173f34; background: #e8f0eb; font-size: 22rpx; }
.detail-card { padding: 30rpx; }
.card-title { display: block; padding-bottom: 25rpx; border-bottom: 1rpx solid #e4e8e2; font-size: 29rpx; font-weight: 800; }
.detail-row { min-height: 68rpx; display: flex; align-items: center; justify-content: space-between; color: #78837b; font-size: 23rpx; }
.detail-row text:last-child { max-width: 66%; color: #2c3730; text-align: right; }
.detail-row.total { margin-top: 10rpx; padding-top: 20rpx; border-top: 1rpx solid #e4e8e2; }
.detail-row.total text:last-child { color: #e75a40; font-size: 30rpx; font-weight: 800; }
.safety-note { margin-top: 22rpx; padding: 24rpx 28rpx; display: flex; gap: 18rpx; border-radius: 26rpx; color: #805a31; background: #fff0d8; }
.safety-note__icon { width: 42rpx; height: 42rpx; flex: 0 0 auto; display: grid; place-items: center; border: 3rpx solid currentColor; border-radius: 50%; font-weight: 800; }
.safety-note > view:last-child { display: flex; flex-direction: column; gap: 8rpx; }.safety-note text:first-child { font-size: 24rpx; font-weight: 700; }.safety-note text:last-child { color: #98744e; font-size: 20rpx; }
.menu-card { overflow: hidden; }
.menu-card > view { height: 84rpx; margin: 0 28rpx; display: flex; align-items: center; justify-content: space-between; border-bottom: 1rpx solid #e5e9e4; font-size: 25rpx; }
.menu-card > view:last-child { border-bottom: 0; }.menu-card > view text:last-child { color: #a5ada7; font-size: 36rpx; }
.avatar { display: grid; place-items: center; color: #fffdf6; background: #153f35; }.avatar text { font-size: 28rpx; font-weight: 900; transform: rotate(-5deg); }.avatar view { display: none; }
.order-page { color: #153f35; background: #f3f5f0; }.top-nav button { border-radius: 14rpx; }
.status-panel { position: relative; overflow: hidden; border-radius: 12rpx; background: #0f4838; }.status-panel::after { content: "QD / ACCEPTED"; position: absolute; right: 22rpx; bottom: 15rpx; color: rgba(255,253,246,.16); font-size: 18rpx; letter-spacing: 3rpx; }
.progress-card, .order-card { border-radius: 12rpx; background: #fffdf6; box-shadow: 0 8rpx 22rpx rgba(21,63,53,.07); }.progress-card { border-top: 3rpx solid #ef644a; }
.progress-step__dot { border-radius: 8rpx; }.progress-step.done .progress-step__dot { background: #153f35; }
.provider-row button { border-radius: 999rpx; background: #b8cfbf; }.card-title, .menu-card > view { border-bottom-style: dashed; }
.safety-note { border: 2rpx dashed #d5aa74; border-radius: 8rpx; }
</style>
