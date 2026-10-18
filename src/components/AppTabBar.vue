<script setup lang="ts">
import { onMounted, onUnmounted, ref } from "vue";
import { unreadCount } from "../stores/meetups";
import { ensureLogin } from "../utils/auth-navigation";
import { isAuthenticated } from "../stores/auth";
import { withNativeLoading } from "../utils/loading";
type TabName = "home" | "discover" | "orders" | "profile";
const props = defineProps<{ active: TabName; ownerRoute: string }>();
const emit = defineEmits<{ (event: "discover"): void }>();
const safeBottom = uni.getWindowInfo().safeAreaInsets?.bottom || 0;
const isVisible = ref(true);
const publishing = ref(false);
// H5 teleport escapes the hidden page's DOM; keep only the current page's bar.
// #ifdef H5
function syncVisibility() {
  const route = window.location.hash.replace(/^#\/?/, "").split("?")[0] || "pages/index/index";
  isVisible.value = route === props.ownerRoute;
}
// Ownership comes from the rendering page, never the current stack: HMR can
// remount a cached page while a different page is on top of that stack.
syncVisibility();
onMounted(() => {
  syncVisibility();
  window.addEventListener("hashchange", syncVisibility);
});
onUnmounted(() => window.removeEventListener("hashchange", syncVisibility));
// #endif
function home() {
  if (props.active !== "home") uni.redirectTo({ url: "/pages/index/index" });
  else uni.pageScrollTo({ scrollTop: 0, duration: 300 });
}
async function discover() {
  if (publishing.value) return;
  publishing.value = true;
  try {
    // Temporary preview of the platform's native loading indicator.
    await withNativeLoading("加载中", () => new Promise<void>(resolve => setTimeout(resolve, 3000)));
    if (!ensureLogin(`/${props.ownerRoute}`)) return;
    emit("discover");
    if (props.active !== "home" && props.active !== "profile") uni.showToast({ title: "发布组局功能开发中", icon: "none" });
  } finally {
    publishing.value = false;
  }
}
function profile() {
  if (props.active !== "profile") uni.redirectTo({ url: "/pages/profile/index" });
}
</script>

<template>
  <!-- #ifdef H5 -->
  <Teleport to="body">
  <!-- #endif -->
  <view v-if="isVisible" class="tabbar" :style="{ '--nav-safe-bottom': safeBottom + 'px' }">
    <view class="tabbar__crest" />
    <button class="tabbar__item" :class="{ active: active === 'home' }" aria-label="首页" @tap="home">
      <image :src="active === 'home' ? '/static/entry/icons/home.svg' : '/static/entry/icons/home-muted.svg'" mode="aspectFit" />
      <text>首页</text>
    </button>
    <button class="tabbar__item tabbar__discover" :disabled="publishing" aria-label="发现同城活动" @tap="discover">
      <view class="tabbar__plus"><image src="/static/entry/icons/plus.svg" mode="aspectFit" /></view>
      <text>发布</text>
    </button>
    <button class="tabbar__item" :class="{ active: active === 'profile' }" aria-label="我的" @tap="profile">
      <image :src="active === 'profile' ? '/static/entry/icons/profile.svg' : '/static/entry/icons/profile-muted.svg'" mode="aspectFit" />
      <view v-if="isAuthenticated && unreadCount" class="tabbar__unread" :aria-label="`${unreadCount}条未读消息`" />
      <text>我的</text>
    </button>
  </view>
  <!-- #ifdef H5 -->
  </Teleport>
  <!-- #endif -->
</template>

<style scoped lang="scss">
.tabbar { position: fixed; z-index: 30; left: 0; right: 0; bottom: 0; box-sizing: border-box; height: calc(136rpx + var(--nav-safe-bottom, 0px)); padding: 10rpx 45rpx var(--nav-safe-bottom, 0px); display: grid; grid-template-columns: repeat(3, minmax(0, 1fr)); align-items: start; border-top: 2rpx solid $color-border; border-radius: $radius-v36rpx_36rpx_0_0; background: $color-white; font-family: $font-body; }
.tabbar__crest { position: absolute; top: -39rpx; left: 50%; width: 186rpx; height: 76rpx; transform: translateX(-50%); border: 2rpx solid $color-border; border-bottom: 0; border-radius: $radius-v50percent_50percent_0_0; background: $color-white; pointer-events: none; }
.tabbar__crest::after { content: ""; position: absolute; left: 0; right: 0; bottom: -2rpx; height: 38rpx; background: $color-white; }
.tabbar__item { position: relative; box-sizing: border-box; width: 100%; min-width: 0; height: 114rpx; padding: 8rpx 0; margin: 0; display: flex; flex-direction: column; align-items: center; justify-content: center; gap: 7rpx; color: $color-textInactive; background: transparent; font-family: inherit; font-size: $font-size-v25rpx; line-height: $line-height-v1p3; font-weight: $font-weight-v500; transform: none; }
.tabbar__item image { width: 50rpx; height: 50rpx; }
.tabbar__unread { position: absolute; top: 12rpx; left: calc(50% + 17rpx); width: 14rpx; height: 14rpx; border: 2rpx solid $color-white; border-radius: $radius-v50percent; background: $color-accent; }
.tabbar__item.active, .tabbar__discover { color: $color-text; font-weight: $font-weight-v600; }
.tabbar__discover { justify-content: flex-end; padding-bottom: 10rpx; overflow: visible; }
.tabbar__plus { position: absolute; left: 50%; top: -36rpx; transform: translateX(-50%); box-sizing: border-box; width: 108rpx; height: 108rpx; display: flex; align-items: center; justify-content: center; border: 3rpx solid $color-white; border-radius: $radius-v50percent; background: $color-accent; }
.tabbar__plus image { width: 59rpx; height: 59rpx; }
.tabbar__item:active { opacity: .7; }
.tabbar__item::after { border: 0; }
.tabbar__item:focus-visible { outline: 2rpx solid $color-primary; outline-offset: -4rpx; border-radius: $radius-v12rpx; }
/* #ifdef H5 */
@media (min-width: 600px) {
  .tabbar { left: 50%; right: auto; width: 960px; transform: translateX(-50%) scale(var(--qudaju-preview-scale, 1)); transform-origin: bottom center; }
}
/* #endif */
</style>
