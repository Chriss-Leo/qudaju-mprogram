<script setup lang="ts">
import { computed, ref, watch } from "vue";
import { getGroup } from "../stores/meetups";
import { openFlow } from "../utils/flow-navigation";
const props = defineProps<{ groupId: string; label?: string }>();
const group = computed(() => getGroup(props.groupId));
const failed = ref(false); watch(() => props.groupId, () => { failed.value = false; });
</script>
<template><button v-if="group" class="activity-context" @tap="openFlow('service-detail', groupId)"><image v-if="!failed" :src="group.photos[0]" mode="aspectFill" @error="failed = true" /><view v-else class="activity-context__placeholder">{{ group.category }}</view><view class="activity-context__copy"><text>{{ group.title }}</text><text>{{ group.time }} · {{ group.area }}</text></view><text v-if="label" class="activity-context__label">{{ label }}</text></button></template>
<style scoped lang="scss">
.activity-context { display: flex; align-items: center; gap: 18rpx; width: 100%; margin: 0; padding: 20rpx; border-radius: $radius-v24rpx; background: $color-card; color: $color-text; text-align: left; }
.activity-context image, .activity-context__placeholder { flex: 0 0 96rpx; width: 96rpx; height: 96rpx; border-radius: $radius-v16rpx; }
.activity-context__placeholder { display: flex; align-items: center; justify-content: center; background: $color-tag; color: $color-primary; font-size: $font-size-v24rpx; }
.activity-context__copy { flex: 1; min-width: 0; }
.activity-context__copy text:first-child { display: block; overflow: hidden; white-space: nowrap; text-overflow: ellipsis; font-size: $font-size-v28rpx; font-weight: $font-weight-v600; line-height: $line-height-v1p4; }
.activity-context__copy text:last-child { display: block; margin-top: 12rpx; color: $color-textMuted; font-size: $font-size-v23rpx; line-height: $line-height-v1p4; }
.activity-context__label { flex-shrink: 0; color: $color-primary; font-size: $font-size-v23rpx; }
.activity-context:active { opacity: .75; }
.activity-context::after { border: 0; }
</style>
