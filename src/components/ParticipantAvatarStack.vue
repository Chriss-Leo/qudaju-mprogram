<script setup lang="ts">
import { computed, ref } from "vue";
import type { GroupParticipant } from "../mocks/group-cards";
const props = withDefaults(defineProps<{ participants: GroupParticipant[]; total: number; compact?: boolean }>(), { compact: false });
const visible = computed(() => props.participants.slice(0, 4));
const remaining = computed(() => Math.max(0, props.total - visible.value.length));
const failed = ref<Record<string, boolean>>({});
</script>

<template>
  <view class="avatar-stack" :class="{ 'avatar-stack--compact': compact }" role="img" :aria-label="`${total}位已正式报名的示例用户`">
    <view v-for="person in visible" :key="person.id" class="avatar-stack__portrait">
      <image v-if="person.avatar && !failed[person.avatar]" :src="person.avatar" mode="aspectFill" aria-hidden="true" @error="failed[person.avatar] = true" />
      <text v-else>{{ person.nickname.slice(0, 1) }}</text>
    </view>
    <view v-if="remaining" class="avatar-stack__portrait avatar-stack__more"><text>+{{ remaining }}</text></view>
  </view>
</template>

<style scoped lang="scss">
.avatar-stack { display: flex; flex-shrink: 0; padding-left: 2rpx; }
.avatar-stack__portrait { position: relative; width: 88rpx; height: 88rpx; flex: 0 0 88rpx; display: flex; align-items: center; justify-content: center; overflow: hidden; border: 4rpx solid $color-white; border-radius: $radius-v50percent; background: $color-tag; color: $color-primary; font-size: $font-size-v26rpx; font-weight: $font-weight-v600; }
.avatar-stack__portrait + .avatar-stack__portrait { margin-left: -24rpx; }
.avatar-stack__portrait image { display: block; width: 100%; height: 100%; }
.avatar-stack__more { background: $color-input; }
.avatar-stack--compact .avatar-stack__portrait { width: 52rpx; height: 52rpx; flex-basis: 52rpx; border-width: 3rpx; font-size: $font-size-v20rpx; }
.avatar-stack--compact .avatar-stack__portrait + .avatar-stack__portrait { margin-left: -15rpx; }
@media (max-width: 350px) { .avatar-stack:not(.avatar-stack--compact) .avatar-stack__portrait { width: 72rpx; height: 72rpx; flex-basis: 72rpx; } }
</style>
