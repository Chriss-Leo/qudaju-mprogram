<script setup lang="ts">
type TabName = "home" | "discover" | "orders" | "profile";
const props = defineProps<{ active: TabName }>();
const emit = defineEmits<{ (event: "discover"): void }>();
const safeBottom = uni.getWindowInfo().safeAreaInsets?.bottom || 0;
function home() {
  if (props.active !== "home") uni.redirectTo({ url: "/pages/index/index" });
  else uni.pageScrollTo({ scrollTop: 0, duration: 300 });
}
function discover() {
  if (props.active === "home") emit("discover");
  else uni.navigateTo({ url: "/pages/search/index" });
}
function profile() {
  if (props.active !== "profile") uni.redirectTo({ url: "/pages/profile/index" });
}
</script>

<template>
  <!-- #ifdef H5 -->
  <Teleport to="body">
  <!-- #endif -->
  <view class="tabbar" :style="{ '--nav-safe-bottom': safeBottom + 'px' }">
    <view class="tabbar__crest" />
    <button class="tabbar__item" :class="{ active: active === 'home' }" aria-label="首页" @tap="home">
      <image :src="active === 'home' ? '/static/entry/icons/home.svg' : '/static/entry/icons/home-muted.svg'" mode="aspectFit" />
      <text>首页</text>
    </button>
    <button class="tabbar__item tabbar__discover" aria-label="发现同城活动" @tap="discover">
      <view class="tabbar__plus"><image src="/static/entry/icons/plus.svg" mode="aspectFit" /></view>
      <text>发布</text>
    </button>
    <button class="tabbar__item" :class="{ active: active === 'profile' }" aria-label="我的" @tap="profile">
      <image :src="active === 'profile' ? '/static/entry/icons/profile.svg' : '/static/entry/icons/profile-muted.svg'" mode="aspectFit" />
      <text>我的</text>
    </button>
  </view>
  <!-- #ifdef H5 -->
  </Teleport>
  <!-- #endif -->
</template>

<style scoped lang="scss">
.tabbar { position: fixed; z-index: 30; left: 0; right: 0; bottom: 0; height: calc(136rpx + var(--nav-safe-bottom, 0px)); padding: 10rpx 45rpx var(--nav-safe-bottom, 0px); display: flex; border-top: 2rpx solid #bdc7bc; border-radius: 36rpx 36rpx 0 0; background: #fafaf6; }
.tabbar__crest { position: absolute; top: -39rpx; left: 50%; width: 186rpx; height: 76rpx; transform: translateX(-50%); border: 2rpx solid #bdc7bc; border-bottom: 0; border-radius: 50% 50% 0 0; background: #fafaf6; pointer-events: none; }
.tabbar__crest::after { content: ""; position: absolute; left: 0; right: 0; bottom: -2rpx; height: 38rpx; background: #fafaf6; }
.tabbar__item { position: relative; flex: 1; min-width: 0; height: 114rpx; padding: 8rpx 0; margin: 0; display: flex; flex-direction: column; align-items: center; justify-content: center; gap: 7rpx; color: #73826f; background: transparent; font-size: 25rpx; line-height: 1.3; font-weight: 500; }
.tabbar__item image { width: 50rpx; height: 50rpx; }
.tabbar__item.active, .tabbar__discover { color: #153f35; font-weight: 600; }
.tabbar__discover { justify-content: flex-end; padding-bottom: 10rpx; overflow: visible; }
.tabbar__plus { position: absolute; top: -36rpx; width: 108rpx; height: 108rpx; display: flex; align-items: center; justify-content: center; border: 3rpx solid #fafaf6; border-radius: 50%; background: #fa704b; }
.tabbar__plus image { width: 59rpx; height: 59rpx; }
.tabbar__item:active { opacity: .7; }
.tabbar__item::after { border: 0; }
.tabbar__item:focus-visible { outline: 2rpx solid #153f35; outline-offset: -4rpx; border-radius: 12rpx; }
/* #ifdef H5 */
@media (min-width: 600px) {
  .tabbar { left: 50%; right: auto; width: 960px; transform: translateX(-50%) scale(var(--qudaju-preview-scale, 1)); transform-origin: bottom center; }
}
/* #endif */
</style>
