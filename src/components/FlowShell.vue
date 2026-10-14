<script setup lang="ts">
import { ref } from "vue";
import { onResize } from "@dcloudio/uni-app";
import SurfaceHeader from "./SurfaceHeader.vue";
import { flowBack } from "../utils/flow-navigation";
defineProps<{ title: string; fallback?: string; scrollTo?: string }>();
const height = ref(uni.getWindowInfo().windowHeight);
onResize(({ size }) => { if (size) height.value = size.windowHeight; });
</script>
<template>
  <view class="flow-shell" :style="{ '--flow-height': height + 'px' }" data-design-contract="Approved chat and registration 2026-10-14: white task surfaces, blue controls, orange primary actions, safe fixed header and footer, independent scrolling.">
    <SurfaceHeader :title="title" back @back="flowBack(fallback)" />
    <view v-if="$slots.top" class="flow-shell__top"><slot name="top" /></view>
    <scroll-view class="flow-shell__scroll" scroll-y :show-scrollbar="false" :scroll-into-view="scrollTo" :scroll-with-animation="true"><slot /></scroll-view>
    <view v-if="$slots.footer" class="flow-shell__footer"><slot name="footer" /></view>
  </view>
</template>
<style scoped lang="scss">
.flow-shell { height: var(--flow-height); display: flex; flex-direction: column; overflow: hidden; background: $color-white; color: $color-text; }
.flow-shell :deep(.surface-header), .flow-shell__top, .flow-shell__footer { flex-shrink: 0; }
.flow-shell :deep(.surface-header) { padding-bottom: 14rpx; }
.flow-shell__scroll { flex: 1; height: 0; min-height: 0; }
.flow-shell__footer { padding: 18rpx 28rpx calc(18rpx + env(safe-area-inset-bottom)); background: $color-white; }
/* #ifdef H5 */
@media (min-width: 600px) { .flow-shell { height: var(--qudaju-preview-layout-height, var(--flow-height)); } }
/* #endif */
</style>
