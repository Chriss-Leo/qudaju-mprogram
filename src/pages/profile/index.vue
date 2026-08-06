<script setup lang="ts">
import AppTabBar from "../../components/AppTabBar.vue";

const orderShortcuts = [
  { label: "待付款", icon: "wallet" },
  { label: "待接单", icon: "ticket" },
  { label: "进行中", icon: "route" },
  { label: "待评价", icon: "review" },
];
const menus = ["我的收藏", "评价记录", "优惠与活动", "投诉与客服", "关于趣搭局"];
function demo(title: string) { uni.showToast({ title: `${title}功能已预留`, icon: "none" }); }
function goOrders() { uni.redirectTo({ url: "/pages/order-detail/index" }); }
function goProvider() { uni.navigateTo({ url: "/pages/provider/index" }); }
function openRecent() { uni.navigateTo({ url: "/pages/service-detail/index?id=city-guide-01" }); }
</script>

<template>
  <view class="profile-page">
    <view class="page-heading"><text>我的</text><button class="settings" aria-label="设置" @tap="demo('设置')"><view /><view /></button></view>
    <view class="profile-hero">
      <view class="avatar"><text>周</text></view>
      <view class="profile-copy"><text>周末探险家</text><text>已加入趣搭局 128 天</text></view>
      <text class="chevron">›</text>
    </view>

    <view class="order-card">
      <view class="card-heading"><text>我的订单</text><text @tap="goOrders">查看全部 ›</text></view>
      <view class="shortcut-grid">
        <view v-for="item in orderShortcuts" :key="item.label" @tap="goOrders">
          <view class="shortcut-icon" :class="`shortcut-icon--${item.icon}`"><i /><i /></view><text>{{ item.label }}</text>
        </view>
      </view>
    </view>

    <view class="provider-banner" @tap="goProvider">
      <view class="provider-emblem"><view /><view /></view>
      <view><text>服务者中心</text><text>发布你的技能，管理订单与收益</text></view>
      <text class="provider-banner__arrow">›</text>
    </view>

    <view class="menu-card">
      <view v-for="(item, index) in menus" :key="item" @tap="demo(item)">
        <view class="menu-icon" :class="`menu-icon--${index}`"><view /></view><text>{{ item }}</text><text>›</text>
      </view>
    </view>

    <view class="recent-section">
      <view class="recent-heading"><text>最近浏览</text><text @tap="openRecent">继续看看 ›</text></view>
      <view class="recent-ticket" @tap="openRecent">
        <image src="/static/posterwall/city-photo-walk.jpg" mode="aspectFill" />
        <view><text>梧桐区漫步与城市摄影</text><text>林野 · 城市向导</text><view><text>今天可约</text><text>1.8km</text></view></view>
      </view>
    </view>
    <AppTabBar active="profile" />
  </view>
</template>

<style scoped lang="scss">
.profile-page { min-height: 100vh; padding: calc(52rpx + env(safe-area-inset-top)) 34rpx 170rpx; }
.page-heading { height: 74rpx; display: flex; align-items: center; justify-content: space-between; font-size: 42rpx; font-weight: 900; }
.settings { position: relative; width: 66rpx; height: 66rpx; margin: 0; padding: 0; border-radius: 50%; background: #e5eae4; }
.settings view:first-child { position: absolute; width: 29rpx; height: 29rpx; left: 15rpx; top: 15rpx; border: 5rpx dotted #38453d; border-radius: 50%; }.settings view:last-child { position: absolute; width: 8rpx; height: 8rpx; left: 29rpx; top: 29rpx; border-radius: 50%; background: #38453d; }
.profile-hero { min-height: 174rpx; padding: 27rpx 5rpx; display: flex; align-items: center; }
.avatar { position: relative; width: 108rpx; height: 108rpx; flex: 0 0 auto; overflow: hidden; border: 7rpx solid #fff; border-radius: 50%; background: #d7e1da; box-shadow: 0 8rpx 24rpx rgba(26,50,40,.1); }
.avatar { display: grid; place-items: center; color: #fffdf6; background: #153f35; }.avatar text { font-size: 42rpx; font-weight: 900; transform: rotate(-5deg); }
.profile-copy { margin-left: 22rpx; display: flex; flex: 1; flex-direction: column; gap: 12rpx; }.profile-copy text:first-child { font-size: 33rpx; font-weight: 800; }.profile-copy text:last-child { color: #8a948d; font-size: 21rpx; }.chevron { color: #a2aaa4; font-size: 45rpx; }
.order-card, .menu-card { border: 1rpx solid #e2e7e1; border-radius: 32rpx; background: #fffefa; }
.order-card { padding: 30rpx 27rpx; }
.card-heading { display: flex; align-items: center; justify-content: space-between; }.card-heading text:first-child { font-size: 29rpx; font-weight: 800; }.card-heading text:last-child { color: #89938c; font-size: 21rpx; }
.shortcut-grid { margin-top: 30rpx; display: grid; grid-template-columns: repeat(4, 1fr); }
.shortcut-grid > view { display: flex; flex-direction: column; align-items: center; gap: 12rpx; color: #4b5650; font-size: 21rpx; }
.shortcut-icon { width: 62rpx; height: 62rpx; display: grid; place-items: center; border-radius: 22rpx; color: #173f34; background: #e6eee8; font-size: 26rpx; font-weight: 800; }
.shortcut-icon { position: relative; }.shortcut-icon i { position: absolute; display: block; }
.shortcut-icon--wallet i:first-child { width: 31rpx; height: 23rpx; border: 4rpx solid currentColor; border-radius: 5rpx; }.shortcut-icon--wallet i:last-child { width: 12rpx; height: 8rpx; right: 10rpx; border: 3rpx solid currentColor; border-radius: 3rpx; background: #b8cfbf; }
.shortcut-icon--ticket i:first-child { width: 32rpx; height: 24rpx; border: 4rpx solid currentColor; }.shortcut-icon--ticket i:last-child { width: 3rpx; height: 18rpx; border-left: 3rpx dashed currentColor; }
.shortcut-icon--route i:first-child { width: 29rpx; height: 18rpx; border: 4rpx solid currentColor; border-left: 0; border-bottom: 0; transform: skewX(-28deg); }.shortcut-icon--route i:last-child { width: 7rpx; height: 7rpx; right: 12rpx; top: 12rpx; border: 3rpx solid currentColor; border-radius: 50%; }
.shortcut-icon--review i:first-child { width: 29rpx; height: 23rpx; border: 4rpx solid currentColor; border-radius: 8rpx; }.shortcut-icon--review i:last-child { width: 10rpx; height: 10rpx; left: 14rpx; bottom: 12rpx; border: 4rpx solid currentColor; border-top: 0; border-right: 0; transform: rotate(-25deg); }
.provider-banner { min-height: 150rpx; margin-top: 24rpx; padding: 28rpx; display: flex; align-items: center; border-radius: 32rpx; color: #fff; background: #173f34; }
.provider-emblem { position: relative; width: 68rpx; height: 68rpx; flex: 0 0 auto; border-radius: 22rpx; background: rgba(255,255,255,.14); }.provider-emblem view:first-child { position: absolute; width: 31rpx; height: 34rpx; left: 18rpx; top: 15rpx; border: 4rpx solid #fff; border-radius: 8rpx 8rpx 15rpx 15rpx; }.provider-emblem view:last-child { position: absolute; width: 13rpx; height: 4rpx; left: 28rpx; top: 33rpx; background: #fff; transform: rotate(-45deg); }
.provider-banner > view:nth-child(2) { margin-left: 20rpx; display: flex; flex: 1; flex-direction: column; gap: 10rpx; }.provider-banner > view:nth-child(2) text:first-child { font-size: 29rpx; font-weight: 800; }.provider-banner > view:nth-child(2) text:last-child { color: rgba(255,255,255,.66); font-size: 20rpx; }.provider-banner__arrow { font-size: 43rpx; }
.menu-card { margin-top: 24rpx; overflow: hidden; }
.menu-card > view { min-height: 90rpx; margin: 0 27rpx; display: grid; grid-template-columns: 52rpx 1fr auto; align-items: center; border-bottom: 1rpx solid #e5e9e4; font-size: 24rpx; }.menu-card > view:last-child { border-bottom: 0; }.menu-card > view > text:last-child { color: #a2aaa4; font-size: 36rpx; }
.menu-icon { position: relative; width: 42rpx; height: 42rpx; border-radius: 14rpx; background: #edf1ec; }.menu-icon view { position: absolute; inset: 11rpx; border: 3rpx solid #526158; border-radius: 50%; }.menu-icon--1 view { border-radius: 4rpx; transform: rotate(45deg); }.menu-icon--2 view { border-radius: 3rpx; }.menu-icon--3 view { inset: 10rpx 8rpx; border-radius: 10rpx; }.menu-icon--4 view { border-style: dotted; }
.profile-page { padding-top: calc(42rpx + env(safe-area-inset-top)); color: #153f35; background: linear-gradient(#0f4838 0 270rpx, #f3f5f0 270rpx); }
.page-heading { color: #fffdf6; }.settings { border: 2rpx solid rgba(255,253,246,.4); border-radius: 14rpx; background: transparent; }.settings view:first-child { border-color: #fffdf6; }.settings view:last-child { background: #fffdf6; }
.profile-hero { margin-top: 20rpx; padding: 24rpx; color: #153f35; background: #fffdf6; box-shadow: 0 13rpx 28rpx rgba(10,45,35,.2); transform: rotate(-.5deg); }
.order-card, .menu-card { border: 0; border-radius: 12rpx; background: #fffdf6; box-shadow: 0 8rpx 20rpx rgba(21,63,53,.06); }
.shortcut-icon { border-radius: 8rpx; color: #153f35; background: #b8cfbf; }
.provider-banner { position: relative; overflow: hidden; border-radius: 12rpx; background: #153f35; }.provider-banner::after { content: "PROVIDER PASS"; position: absolute; right: 20rpx; bottom: 10rpx; color: rgba(255,255,255,.14); font-size: 16rpx; letter-spacing: 2rpx; }
.menu-card > view { border-bottom-style: dashed; }.menu-icon { border-radius: 6rpx; background: #e2eae3; }
.recent-section { margin-top: 34rpx; }.recent-heading { display: flex; align-items: center; justify-content: space-between; }.recent-heading text:first-child { font-size: 29rpx; font-weight: 900; }.recent-heading text:last-child { color: #718078; font-size: 21rpx; }
.recent-ticket { position: relative; min-height: 150rpx; margin-top: 17rpx; display: grid; grid-template-columns: 184rpx 1fr; overflow: hidden; border-radius: 12rpx; background: #fffdf6; box-shadow: 0 8rpx 20rpx rgba(21,63,53,.07); }.recent-ticket::before, .recent-ticket::after { content: ""; position: absolute; top: 50%; width: 20rpx; height: 20rpx; border-radius: 50%; background: #f3f5f0; transform: translateY(-50%); }.recent-ticket::before { left: -10rpx; }.recent-ticket::after { right: -10rpx; }.recent-ticket image { width: 184rpx; height: 150rpx; }.recent-ticket > view { min-width: 0; padding: 20rpx; display: flex; flex-direction: column; }.recent-ticket > view > text:first-child { font-size: 25rpx; line-height: 1.35; font-weight: 900; }.recent-ticket > view > text:nth-child(2) { margin-top: 9rpx; color: #718078; font-size: 19rpx; }.recent-ticket > view > view { margin-top: auto; display: flex; justify-content: space-between; color: #3f5b4e; font-size: 18rpx; }
</style>
