<script setup lang="ts">
import { theme } from "../../theme";
import { computed, ref } from "vue";
import { onLoad, onResize, onShow } from "@dcloudio/uni-app";
import { roundedFontSource } from "../../assets/rounded-font";
import AppTabBar from "../../components/AppTabBar.vue";
const windowInfo = uni.getWindowInfo();
const entryHeight = ref(windowInfo.windowHeight);
const entryStyle = computed(() => ({
  "--entry-height": entryHeight.value + "px",
  "--entry-safe-bottom": (windowInfo.safeAreaInsets?.bottom || 0) + "px",
}));
const posterReady = ref(false);
const posterFailed = ref(false);
onResize(({ size }) => { if (size) entryHeight.value = size.windowHeight; });
onShow(() => { uni.setNavigationBarColor({ frontColor: "#ffffff", backgroundColor: theme.colors.primary }); });
onLoad(() => { uni.loadFontFace({ global: true, family: theme.fontFaces.display, source: roundedFontSource }); });
function explore() { uni.navigateTo({ url: "/pages/groups/index" }); }
function publish() { uni.showToast({ title: "发布组局功能开发中", icon: "none" }); }
</script>

<template>
  <view class="home-page" :style="entryStyle" @touchmove.stop.prevent>
    <view class="entry-cover">
      <image
        v-if="!posterFailed"
        class="entry-cover__image"
        :class="{ 'entry-cover__image--ready': posterReady }"
        src="/static/entry/city-walk.jpg"
        mode="aspectFill"
        aria-label="朋友在梧桐街区漫步"
        @load="posterReady = true"
        @error="posterFailed = true"
      />
      <view class="entry-cover__shade" />
      <view class="entry-brand">
        <image class="entry-brand__logo" src="/static/brand/qudaju-wordmark-orange-ivory-v1.png" mode="aspectFit" aria-label="趣搭局" />
        <text class="entry-brand__tagline">好搭子 · 好时光</text>
      </view>
      <view class="entry-headline">
        <text>今天，</text>
        <view><text>找些好搭子</text><view class="entry-headline__dot" /></view>
      </view>
      <view class="entry-discover">
        <button class="entry-discover__link" @tap="explore"><text>发现同城好局</text><view class="entry-discover__arrow"><uni-icons type="arrow-right" size="23" :color="theme.colors.ivory" /></view></button>
      </view>
      <button v-if="posterFailed" class="entry-retry" @tap="posterFailed = false">重新加载背景</button>
    </view>
    <AppTabBar active="home" owner-route="pages/index/index" @discover="publish" />
  </view>
</template>

<style scoped lang="scss">
.home-page { height: var(--entry-height, 100vh); overflow: hidden; background: $color-primary; }
.entry-cover { position: relative; height: var(--entry-height, 100vh); background: $color-primary; overflow: hidden; }
.entry-cover__image { position: absolute; inset: 0; width: 100%; height: 100%; opacity: 0; }
.entry-cover__image--ready { opacity: 1; }
.entry-cover__shade { position: absolute; inset: 0; pointer-events: none; background: linear-gradient(180deg, $color-compatOverlay12, transparent 30%, transparent 52%, $color-compatOverlay13 72%, $color-compatOverlay14); }
.entry-brand { position: absolute; left: 0; right: 0; top: 13%; display: flex; flex-direction: column; align-items: center; gap: 13rpx; color: $color-inverseWarm; }
.entry-brand__logo { display: block; width: 380rpx; height: 100rpx; }
.entry-brand__tagline { padding-left: 7rpx; font-size: $font-size-v36rpx; line-height: $line-height-v1p4; font-weight: $font-weight-v700; letter-spacing: 7rpx; color: $color-accentForeground; }
.entry-headline { position: absolute; top: 55%; left: 6.2%; right: 4%; color: $color-inverseWarm; font-family: $font-display; font-size: $font-size-v85rpx; font-weight: $font-weight-v350; line-height: $line-height-v1p23; letter-spacing: 0; }
.entry-headline > text { display: block; }
.entry-headline > view { display: flex; align-items: baseline; white-space: nowrap; }
.entry-headline__dot { flex: 0 0 auto; width: 22rpx; height: 22rpx; margin-left: 7rpx; border-radius: $radius-v50percent; background: $color-accent; transform: translateY(.08em); }
.entry-discover { position: absolute; left: 6.2%; right: 6.2%; bottom: calc(192rpx + var(--entry-safe-bottom, 0px)); color: $color-ivory; }
.entry-discover button { margin: 0; padding: 0; color: inherit; background: transparent; border-radius: $radius-v0; }
.entry-discover__link { display: flex; align-items: center; gap: 24rpx; width: 100%; font-family: $font-editorial; font-size: $font-size-v48rpx; line-height: $line-height-v1p3; text-align: left; }
.entry-discover__arrow { display: flex; align-items: center; justify-content: center; width: 52rpx; height: 52rpx; border: 1rpx solid $color-ivory; border-radius: $radius-v50percent; }
.entry-retry { position: absolute; top: 42%; left: 50%; transform: translateX(-50%); padding: 0 24rpx; color: $color-primary; background: $color-inverseWarm; font-size: $font-size-v25rpx; }
.entry-discover button::after { border: 0; }
/* #ifdef H5 */
@media (min-width: 600px) { .home-page, .entry-cover { height: var(--qudaju-preview-layout-height, 100vh); } }
/* #endif */
@media (max-height: 700px) { .entry-headline { top: 49%; font-size: $font-size-v78rpx; } .entry-discover__link { font-size: $font-size-v42rpx; } }
</style>
