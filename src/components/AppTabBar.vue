<script setup lang="ts">
type TabName = "home" | "orders" | "profile";

const props = defineProps<{ active: TabName }>();

const tabs: Array<{ key: TabName; label: string; path: string }> = [
  { key: "home", label: "首页", path: "/pages/index/index" },
  { key: "orders", label: "订单", path: "/pages/order-detail/index" },
  { key: "profile", label: "我的", path: "/pages/profile/index" },
];

function go(tab: (typeof tabs)[number]) {
  if (tab.key === props.active) return;
  uni.redirectTo({ url: tab.path });
}
</script>

<template>
  <view class="tabbar">
    <view
      v-for="tab in tabs"
      :key="tab.key"
      class="tabbar__item"
      :class="{ 'tabbar__item--active': active === tab.key }"
      @tap="go(tab)"
    >
      <view class="nav-icon" :class="`nav-icon--${tab.key}`">
        <view class="nav-icon__shape" />
      </view>
      <text>{{ tab.label }}</text>
    </view>
  </view>
</template>

<style scoped lang="scss">
.tabbar {
  position: fixed;
  z-index: 30;
  left: 0;
  right: 0;
  bottom: 0;
  height: calc(126rpx + env(safe-area-inset-bottom));
  padding: 13rpx 54rpx env(safe-area-inset-bottom);
  border-top: 1rpx solid rgba(21, 63, 53, 0.12);
  background: rgba(253, 253, 249, 0.98);
  box-shadow: 0 -8rpx 24rpx rgba(21, 63, 53, 0.05);
  display: flex;
}
.tabbar__item {
  flex: 1;
  min-width: 0;
  color: #7e8981;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 7rpx;
  font-size: 25rpx;
  font-weight: 800;
}
.tabbar__item--active { color: #153f35; }
.tabbar__item--active::before { display: none; }
.nav-icon { position: relative; width: 58rpx; height: 50rpx; }
.nav-icon__shape { position: absolute; inset: 5rpx; border: 4rpx solid currentColor; }
.nav-icon--home .nav-icon__shape { inset: 6rpx 0 3rpx; overflow: hidden; border: 0; border-radius: 9rpx; background: currentColor; clip-path: polygon(50% 0, 100% 43%, 88% 43%, 88% 100%, 12% 100%, 12% 43%, 0 43%); }
.nav-icon--home::after { content: ""; position: absolute; z-index: 1; width: 13rpx; height: 20rpx; left: 23rpx; bottom: 3rpx; border-radius: 5rpx 5rpx 0 0; background: #fffdf6; }
.tabbar__item--active .nav-icon--home::after { background: #ef644a; }
.nav-icon--orders .nav-icon__shape { inset: 5rpx 8rpx; border-radius: 7rpx; }
.nav-icon--orders::before { content: ""; position: absolute; z-index: 1; left: 20rpx; top: 2rpx; width: 18rpx; height: 9rpx; border-radius: 3rpx; background: #fffdf6; border: 3rpx solid currentColor; }
.nav-icon--orders::after { content: ""; position: absolute; left: 21rpx; top: 19rpx; width: 17rpx; height: 3rpx; background: currentColor; box-shadow: 0 10rpx 0 currentColor; }
.nav-icon--profile .nav-icon__shape { inset: 1rpx 14rpx 25rpx; border-radius: 50%; }
.nav-icon--profile::after { content: ""; position: absolute; left: 4rpx; right: 4rpx; bottom: 1rpx; height: 20rpx; border: 4rpx solid currentColor; border-radius: 22rpx 22rpx 8rpx 8rpx; }
</style>
