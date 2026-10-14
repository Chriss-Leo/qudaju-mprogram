<script setup lang="ts">
import { ref, watch } from "vue";
import type { GroupCardData } from "../mocks/group-cards";
import GroupPhotoGrid from "./GroupPhotoGrid.vue";
import ParticipantAvatarStack from "./ParticipantAvatarStack.vue";
import UiIcon from "./UiIcon.vue";
const props = defineProps<{ group: GroupCardData; distance: string }>();
const emit = defineEmits<{ (event: "open"): void }>();
const hostFailed = ref(false);
watch(() => props.group.hostAvatar, () => { hostFailed.value = false; });
</script>

<template>
  <view class="group-card" role="button" tabindex="0" :aria-label="`查看${group.title}`" @tap="emit('open')" @keydown.enter.self="emit('open')">
    <view class="group-card__header">
    <view class="group-card__host">
      <image v-if="!hostFailed" class="group-card__avatar" :src="group.hostAvatar" mode="aspectFill" :aria-label="`${group.host}示例头像`" @error="hostFailed = true" />
      <view v-else class="group-card__avatar group-card__initial">{{ group.host.slice(0, 1) }}</view>
      <view class="group-card__identity"><text class="group-card__name">{{ group.host }}</text><text class="group-card__admission">{{ group.admission }}</text></view>
      <text class="group-card__tag">{{ group.category }}</text>
    </view>
      <view class="group-card__fee">
        <text v-if="group.fund === 0" class="group-card__price">免费</text>
        <view v-else class="group-card__fund"><text class="group-card__price">¥{{ group.fund }}</text><text>/人</text></view>
      </view>
    </view>
    <text class="group-card__title">{{ group.title }}</text>
    <text class="group-card__description">{{ group.description }}</text>
    <GroupPhotoGrid class="group-card__gallery" :photos="group.photos" :label="group.category" />
    <view class="group-card__meta">
      <view><UiIcon name="clock" /><text>{{ group.time }}</text></view>
      <view><UiIcon name="location" /><text>{{ group.area }}<text class="group-card__distance"> · {{ distance }}</text></text></view>
    </view>
    <view class="group-card__footer">
      <view class="group-card__participants"><ParticipantAvatarStack :participants="group.participants" :total="group.joined" compact /><view class="group-card__capacity"><text class="group-card__joined">{{ group.joined }}</text><text> / {{ group.capacity }} 已报名</text></view></view>
      <button class="group-card__apply" :class="{ 'group-card__apply--joined': group.isJoinedByMe }" :disabled="group.isJoinedByMe" :aria-label="`${group.isJoinedByMe ? '已报名' : '去报名'}：${group.title}`" @tap.stop="!group.isJoinedByMe && emit('open')" @keydown.enter.stop="!group.isJoinedByMe && emit('open')">{{ group.isJoinedByMe ? '已报名' : '去报名' }}</button>
    </view>
  </view>
</template>

<style scoped lang="scss">
.group-card { padding: 28rpx; min-width: 0; border-radius: $radius-v24rpx; color: $color-text; background: $color-card; }
.group-card:focus-visible { outline: 3rpx solid $color-primary; outline-offset: 3rpx; }
.group-card__header { display: grid; grid-template-columns: minmax(0, 1fr) auto; align-items: center; gap: 16rpx; margin-bottom: 22rpx; }
.group-card__host { display: flex; align-items: center; gap: 16rpx; min-width: 0; }
.group-card__avatar { display: block; flex: 0 0 72rpx; width: 72rpx; height: 72rpx; border-radius: $radius-v50percent; background: $color-tag; }
.group-card__initial { display: flex; align-items: center; justify-content: center; color: $color-primary; font-size: $font-size-v30rpx; font-weight: $font-weight-v600; }
.group-card__identity { display: flex; flex-direction: column; gap: 5rpx; min-width: 0; }
.group-card__name { overflow: hidden; text-overflow: ellipsis; white-space: nowrap; font-size: $font-size-v29rpx; line-height: $line-height-v1p35; font-weight: $font-weight-v600; }
.group-card__admission { overflow: hidden; text-overflow: ellipsis; white-space: nowrap; color: $color-textMuted; font-size: $font-size-v23rpx; line-height: $line-height-v1p4; }
.group-card__tag { flex-shrink: 0; margin-left: 10rpx; padding: 6rpx 18rpx; border-radius: $radius-v999rpx; background: $color-tag; color: $color-primary; font-size: $font-size-v23rpx; line-height: $line-height-v1p4; }
.group-card__title { display: block; overflow: hidden; text-overflow: ellipsis; white-space: nowrap; font-size: $font-size-v34rpx; line-height: $line-height-v1p45; font-weight: $font-weight-v600; }
.group-card__description { display: block; margin-top: 8rpx; color: $color-textSecondary; font-size: $font-size-v26rpx; line-height: $line-height-v1p6; }
.group-card__gallery { margin-top: 22rpx; }
.group-card__meta { display: flex; flex-wrap: wrap; gap: 12rpx 28rpx; margin: 22rpx 0; font-size: $font-size-v24rpx; line-height: $line-height-v1p5; }
.group-card__meta > view { display: flex; align-items: center; gap: 10rpx; }
.group-card__meta :deep(.ui-icon) { width: 32rpx; height: 32rpx; }
.group-card__distance { color: $color-textMuted; }
.group-card__footer { display: flex; align-items: center; justify-content: space-between; gap: 14rpx; }
.group-card__participants { display: flex; align-items: center; gap: 12rpx; min-width: 0; }
.group-card__capacity { color: $color-textMuted; font-size: $font-size-v22rpx; line-height: $line-height-v1p4; white-space: nowrap; }
.group-card__joined { color: $color-primary; font-size: $font-size-v28rpx; font-weight: $font-weight-v600; }
.group-card__price { color: $color-accent; font-size: $font-size-v32rpx; font-weight: $font-weight-v700; white-space: nowrap; }
.group-card__fund { display: flex; align-items: baseline; gap: 4rpx; color: $color-textSecondary; font-size: $font-size-v22rpx; white-space: nowrap; }
.group-card__fee { justify-self: end; }
.group-card__apply { display: flex; flex-shrink: 0; align-items: center; justify-content: center; box-sizing: border-box; width: 160rpx; min-height: 44px; margin: 0; padding: 12rpx 18rpx; border: 1px solid $color-primary; border-radius: $radius-v16rpx; background: transparent; color: $color-primary; font-size: $font-size-v24rpx; font-weight: $font-weight-v600; line-height: $line-height-v1p4; white-space: nowrap; }
.group-card__apply::after { border: 0; }
.group-card .group-card__apply--joined { border-color: $color-borderStrong; background: transparent; color: $color-textMuted; opacity: 1; }
.group-card__apply:not([disabled]):active { opacity: .7; }
.group-card__apply:focus-visible { outline: 3rpx solid $color-primary; outline-offset: 2rpx; }
@media (max-width: 350px) { .group-card { padding: 22rpx; } .group-card__participants { gap: 8rpx; } .group-card__capacity { font-size: $font-size-v20rpx; } }
</style>
