<script setup lang="ts">
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
onShow(() => { uni.setNavigationBarColor({ frontColor: "#ffffff", backgroundColor: "#153f35" }); });
onLoad(() => { uni.loadFontFace({ global: true, family: "QudajuRoundedBold", source: roundedFontSource }); });
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
        <image class="entry-brand__logo" src="/static/brand/qudaju-wordmark-white-orange-draft.png" mode="aspectFit" aria-label="趣搭局" />
        <text class="entry-brand__tagline">好搭子 · 好时光</text>
      </view>
      <view class="entry-headline">
        <text>今天，</text>
        <view><text>找些好搭子</text><view class="entry-headline__dot" /></view>
      </view>
      <view class="entry-discover">
        <button class="entry-discover__link" @tap="explore"><text>发现同城好局</text><view class="entry-discover__arrow"><uni-icons type="arrow-right" size="23" color="#fffdf6" /></view></button>
      </view>
      <button v-if="posterFailed" class="entry-retry" @tap="posterFailed = false">重新加载背景</button>
    </view>
    <AppTabBar active="home" @discover="publish" />
  </view>
</template>

<style scoped lang="scss">
.home-page { height: var(--entry-height, 100vh); overflow: hidden; background: #153f35; }
.entry-cover { position: relative; height: var(--entry-height, 100vh); background: #153f35; overflow: hidden; }
.entry-cover__image { position: absolute; inset: 0; width: 100%; height: 100%; opacity: 0; }
.entry-cover__image--ready { opacity: 1; }
.entry-cover__shade { position: absolute; inset: 0; pointer-events: none; background: linear-gradient(180deg, rgba(13,31,19,.06), transparent 30%, transparent 52%, rgba(9,27,17,.3) 72%, rgba(9,27,17,.88)); }
.entry-brand { position: absolute; left: 0; right: 0; top: 9%; display: flex; flex-direction: column; align-items: center; gap: 13rpx; color: #fffdf6; }
.entry-brand__logo { display: block; width: 380rpx; height: 100rpx; }
.entry-brand__tagline { padding-left: 7rpx; font-size: 27rpx; line-height: 1.4; font-weight: 700; letter-spacing: 7rpx; color: #fa704b; }
.entry-headline { position: absolute; top: 55%; left: 6.2%; right: 4%; color: #fffdf6; font-family: QudajuRoundedBold, "PingFang SC", sans-serif; font-size: 85rpx; font-weight: 350; line-height: 1.23; letter-spacing: 0; }
.entry-headline > text { display: block; }
.entry-headline > view { display: flex; align-items: baseline; white-space: nowrap; }
.entry-headline__dot { flex: 0 0 auto; width: 22rpx; height: 22rpx; margin-left: 7rpx; border-radius: 50%; background: #fa704b; transform: translateY(.08em); }
.entry-discover { position: absolute; left: 6.2%; right: 6.2%; bottom: calc(192rpx + var(--entry-safe-bottom, 0px)); color: #fffdf6; }
.entry-discover button { margin: 0; padding: 0; color: inherit; background: transparent; border-radius: 0; }
.entry-discover__link { display: flex; align-items: center; gap: 24rpx; width: 100%; font-family: "Songti SC", "STSong", serif; font-size: 48rpx; line-height: 1.3; text-align: left; }
.entry-discover__arrow { display: flex; align-items: center; justify-content: center; width: 52rpx; height: 52rpx; border: 1rpx solid #fffdf6; border-radius: 50%; }
.entry-retry { position: absolute; top: 42%; left: 50%; transform: translateX(-50%); padding: 0 24rpx; color: #153f35; background: #fffdf6; font-size: 25rpx; }
.entry-discover button::after { border: 0; }
/* #ifdef H5 */
@media (min-width: 600px) { .home-page, .entry-cover { height: var(--qudaju-preview-layout-height, 100vh); } }
/* #endif */
@media (max-height: 700px) { .entry-headline { top: 49%; font-size: 78rpx; } .entry-discover__link { font-size: 42rpx; } }
</style>
