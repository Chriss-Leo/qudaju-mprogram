<script setup lang="ts">
const workbench = [
  { label: "发布服务", icon: "publish" },
  { label: "档期管理", icon: "calendar" },
  { label: "订单履约", icon: "order" },
  { label: "收益提现", icon: "income" },
];
function back() { uni.navigateBack({ fail: () => uni.redirectTo({ url: "/pages/profile/index" }) }); }
function demo(title: string) { uni.showToast({ title: `${title}能力已预留`, icon: "none" }); }
</script>

<template>
  <view class="provider-page">
    <view class="provider-hero">
      <view class="top-nav">
        <button class="back" aria-label="返回" @tap="back"><view /></button><text>工作台</text><button class="gear" aria-label="设置" @tap="demo('服务者设置')"><view /><view /></button>
      </view>
      <view class="provider-profile">
        <view class="avatar"><text>林</text></view>
        <view><view class="name-line"><text>林野</text><text>已认证</text></view><text class="role">城市向导 · 摄影陪拍</text></view>
      </view>
      <view class="metrics">
        <view><text>¥2,680</text><text>本月收益</text></view>
        <view><text>3</text><text>待处理订单</text></view>
        <view><text>4.9</text><text>服务评分</text></view>
      </view>
    </view>

    <view class="workbench-card">
      <view v-for="item in workbench" :key="item.label" @tap="demo(item.label)"><view class="workbench-icon" :class="`workbench-icon--${item.icon}`"><i /><i /></view><text>{{ item.label }}</text></view>
    </view>

    <view class="section-heading"><view><text>待接单</text><text>及时处理能提高推荐权重</text></view><text @tap="demo('全部订单')">全部订单 ›</text></view>
    <view class="task-card">
      <view class="task-card__top"><view><text>梧桐区漫步与城市摄影</text><text>今天 14:00—16:00</text></view><text class="countdown">剩余 18 分钟</text></view>
      <view class="customer-row"><view class="customer-avatar"><text>周</text></view><view><text>周末探险家</text><text>线下见面 · 徐汇区</text></view><text class="amount">¥128</text></view>
      <view class="task-actions"><button @tap="demo('婉拒订单')">婉拒</button><button @tap="demo('接单')">立即接单</button></view>
    </view>

    <view class="tips-card"><view>i</view><text>完成身份认证、维护档期并及时回复，可以获得更多附近推荐。</text></view>

    <view class="schedule-section">
      <view class="section-heading schedule-heading"><view><text>今日档期</text><text>保持可约时间准确</text></view><text @tap="demo('档期管理')">管理档期 ›</text></view>
      <view class="schedule-ticket">
        <view><text>09:00</text><text>可预约</text></view>
        <view class="busy"><text>14:00</text><text>已锁定</text></view>
        <view><text>19:00</text><text>可预约</text></view>
      </view>
    </view>
  </view>
</template>

<style scoped lang="scss">
.provider-page { min-height: 100vh; padding-bottom: 80rpx; background: $color-compatF2F4EF; }
.provider-hero { min-height: 510rpx; padding: calc(48rpx + env(safe-area-inset-top)) 34rpx 100rpx; color: $color-white; background: $color-primary; }
.top-nav { height: 70rpx; display: grid; grid-template-columns: 64rpx 1fr 64rpx; align-items: center; text-align: center; font-size: $font-size-v33rpx; font-weight: $font-weight-v800; }
.top-nav button { width: 64rpx; height: 64rpx; margin: 0; padding: 0; border-radius: $radius-v50percent; background: $color-compatOverlay31; }
.back view { width: 19rpx; height: 19rpx; margin: 20rpx 0 0 25rpx; border: 4rpx solid $color-white; border-top: 0; border-right: 0; transform: rotate(45deg); }
.gear { position: relative; }.gear view:first-child { position: absolute; width: 25rpx; height: 25rpx; left: 16rpx; top: 16rpx; border: 5rpx dotted $color-white; border-radius: $radius-v50percent; }.gear view:last-child { position: absolute; width: 7rpx; height: 7rpx; left: 29rpx; top: 29rpx; border-radius: $radius-v50percent; background: $color-white; }
.provider-profile { margin-top: 42rpx; display: flex; align-items: center; }
.avatar { width: 94rpx; height: 94rpx; flex: 0 0 auto; display: grid; place-items: center; border: 5rpx solid $color-compatOverlay32; border-radius: $radius-v50percent; color: $color-primary; background: $color-compatB8CFBF; }.avatar text { font-size: $font-size-v39rpx; font-weight: $font-weight-v900; transform: rotate(-5deg); }
.provider-profile > view:last-child { margin-left: 23rpx; }.name-line { display: flex; align-items: center; gap: 13rpx; }.name-line text:first-child { font-size: $font-size-v34rpx; font-weight: $font-weight-v800; }.name-line text:last-child { padding: 6rpx 12rpx; border-radius: $radius-v10rpx; color: $color-primary; background: $color-compatDFECE4; font-size: $font-size-v19rpx; font-weight: $font-weight-v700; }.role { display: block; margin-top: 12rpx; color: $color-compatOverlay33; font-size: $font-size-v22rpx; }
.metrics { margin-top: 42rpx; display: grid; grid-template-columns: repeat(3, 1fr); }.metrics > view { display: flex; flex-direction: column; align-items: center; gap: 11rpx; border-right: 1rpx solid $color-compatOverlay21; }.metrics > view:last-child { border-right: 0; }.metrics text:first-child { font-size: $font-size-v34rpx; font-weight: $font-weight-v800; }.metrics text:last-child { color: $color-compatOverlay34; font-size: $font-size-v20rpx; }
.workbench-card { position: relative; z-index: 2; margin: -65rpx 34rpx 0; padding: 30rpx 22rpx; display: grid; grid-template-columns: repeat(4, 1fr); border-radius: $radius-v34rpx; background: $color-white; box-shadow: 0 16rpx 42rpx $color-compatOverlay35; }
.workbench-card > view { display: flex; flex-direction: column; align-items: center; gap: 12rpx; color: $color-compat4A5650; font-size: $font-size-v21rpx; }
.workbench-icon { width: 68rpx; height: 68rpx; display: grid; place-items: center; border-radius: $radius-v22rpx; color: $color-primary; background: $color-compatE5EEE8; font-size: $font-size-v27rpx; font-weight: $font-weight-v800; }
.workbench-icon { position: relative; }.workbench-icon i { position: absolute; display: block; }
.workbench-icon--publish i:first-child { width: 30rpx; height: 4rpx; background: currentColor; }.workbench-icon--publish i:last-child { width: 4rpx; height: 30rpx; background: currentColor; }
.workbench-icon--calendar i:first-child { width: 32rpx; height: 28rpx; border: 4rpx solid currentColor; border-radius: $radius-v5rpx; }.workbench-icon--calendar i:last-child { width: 22rpx; height: 3rpx; top: 23rpx; background: currentColor; box-shadow: 0 8rpx 0 currentColor; }
.workbench-icon--order i:first-child { width: 31rpx; height: 35rpx; border: 4rpx solid currentColor; border-radius: $radius-v5rpx; }.workbench-icon--order i:last-child { width: 14rpx; height: 7rpx; border: 4rpx solid currentColor; border-top: 0; border-right: 0; transform: rotate(-45deg); }
.workbench-icon--income i:first-child { width: 35rpx; height: 24rpx; border: 4rpx solid currentColor; border-radius: $radius-v5rpx; }.workbench-icon--income i:last-child { width: 9rpx; height: 9rpx; border: 3rpx solid currentColor; border-radius: $radius-v50percent; }
.section-heading { margin: 42rpx 34rpx 20rpx; display: flex; align-items: flex-end; justify-content: space-between; }.section-heading > view { display: flex; flex-direction: column; gap: 8rpx; }.section-heading > view text:first-child { font-size: $font-size-v34rpx; font-weight: $font-weight-v900; }.section-heading > view text:last-child { color: $color-compat89938C; font-size: $font-size-v20rpx; }.section-heading > text { color: $color-compat758078; font-size: $font-size-v21rpx; }
.task-card { margin: 0 34rpx; padding: 28rpx; border-radius: $radius-v30rpx; background: $color-white; }
.task-card__top { padding-bottom: 24rpx; display: flex; justify-content: space-between; border-bottom: 1rpx solid $color-compatE4E8E2; }.task-card__top > view { display: flex; flex-direction: column; gap: 10rpx; }.task-card__top > view text:first-child { font-size: $font-size-v27rpx; font-weight: $font-weight-v800; }.task-card__top > view text:last-child { color: $color-compat7E8881; font-size: $font-size-v21rpx; }.countdown { align-self: flex-start; padding: 7rpx 11rpx; border-radius: $radius-v10rpx; color: $color-compatD05038; background: $color-compatFAE8E1; font-size: $font-size-v18rpx; }
.customer-row { padding: 24rpx 0; display: flex; align-items: center; }.customer-avatar { width: 64rpx; height: 64rpx; display: grid; place-items: center; border-radius: $radius-v50percent; color: $color-inverseWarm; background: $color-compat315A4C; }.customer-avatar text { font-size: $font-size-v25rpx; font-weight: $font-weight-v900; }.customer-row > view:nth-child(2) { margin-left: 17rpx; display: flex; flex: 1; flex-direction: column; gap: 7rpx; }.customer-row > view:nth-child(2) text:first-child { font-size: $font-size-v23rpx; font-weight: $font-weight-v700; }.customer-row > view:nth-child(2) text:last-child { color: $color-compat8B948E; font-size: $font-size-v19rpx; }.amount { color: $color-compatE7583F; font-size: $font-size-v29rpx; font-weight: $font-weight-v800; }
.task-actions { display: grid; grid-template-columns: 1fr 1.45fr; gap: 15rpx; }.task-actions button { height: 72rpx; border-radius: $radius-v22rpx; font-size: $font-size-v24rpx; }.task-actions button:first-child { border: 2rpx solid $color-compatDFE4DE; background: $color-compatF5F7F3; color: $color-compat4B5750; }.task-actions button:last-child { background: $color-accent; color: $color-white; font-weight: $font-weight-v700; }
.tips-card { margin: 24rpx 34rpx 0; padding: 24rpx; display: flex; align-items: center; gap: 15rpx; border-radius: $radius-v24rpx; color: $color-compat68766D; background: $color-compatE5EBE5; font-size: $font-size-v21rpx; line-height: $line-height-v1p5; }.tips-card view { width: 36rpx; height: 36rpx; flex: 0 0 auto; display: grid; place-items: center; border: 3rpx solid $color-compat66766C; border-radius: $radius-v50percent; font-weight: $font-weight-v800; }
.provider-page { background: $color-compatF3F5F0; }.provider-hero { background: $color-primary; }.top-nav button { border-radius: $radius-v14rpx; }
.name-line text:last-child { border-radius: $radius-v0; background: $color-compatB8CFBF; }
.workbench-card { border-radius: $radius-v12rpx; background: $color-inverseWarm; transform: rotate(-.4deg); }.workbench-card::before, .workbench-card::after { content: ""; position: absolute; top: 50%; width: 22rpx; height: 22rpx; border-radius: $radius-v50percent; background: $color-compatF3F5F0; }.workbench-card::before { left: -11rpx; }.workbench-card::after { right: -11rpx; }
.workbench-icon { border-radius: $radius-v8rpx; background: $color-compatB8CFBF; }.task-card { position: relative; border-radius: $radius-v12rpx; background: $color-inverseWarm; box-shadow: 0 10rpx 24rpx $color-compatOverlay36; }.task-card__top { border-bottom-style: dashed; }
.countdown { border-radius: $radius-v0; color: $color-compat7A2F24; background: $color-compatEF9A87; }.task-actions button { border-radius: $radius-v12rpx; }.task-actions button:last-child { background: $color-accent; }
.tips-card { border: 2rpx dashed $color-compatAAB9AE; border-radius: $radius-v8rpx; background: $color-compatE5ECE6; }
.schedule-section { margin-top: 38rpx; }.schedule-heading { margin-top: 0; }.schedule-ticket { margin: 0 34rpx; display: grid; grid-template-columns: repeat(3, 1fr); overflow: hidden; border: 2rpx dashed $color-compatAEBCB2; border-radius: $radius-v8rpx; background: $color-inverseWarm; }.schedule-ticket > view { min-height: 104rpx; display: flex; flex-direction: column; align-items: center; justify-content: center; gap: 8rpx; border-right: 1rpx solid $color-compatDFE5DF; }.schedule-ticket > view:last-child { border-right: 0; }.schedule-ticket text:first-child { font-size: $font-size-v26rpx; font-weight: $font-weight-v900; font-variant-numeric: tabular-nums; }.schedule-ticket text:last-child { color: $color-compat668073; font-size: $font-size-v18rpx; }.schedule-ticket .busy { color: $color-inverseWarm; background: $color-primary; }.schedule-ticket .busy text:last-child { color: $color-compatB8CFBF; }
</style>
