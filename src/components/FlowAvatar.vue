<script setup lang="ts">
import { ref, watch } from "vue";
import type { GroupParticipant } from "../mocks/group-cards";
const props = defineProps<{ person?: GroupParticipant; large?: boolean }>();
const failed = ref(false); watch(() => props.person?.avatar, () => { failed.value = false; });
</script>
<template><image v-if="person && !failed" class="flow-avatar" :class="{ 'flow-avatar--large': large }" :src="person.avatar" mode="aspectFill" :aria-label="person.nickname" @error="failed = true" /><view v-else class="flow-avatar flow-avatar--fallback" :class="{ 'flow-avatar--large': large }">{{ person?.nickname.slice(0,1) || '友' }}</view></template>
<style scoped lang="scss">
.flow-avatar { display: block; flex: 0 0 88rpx; width: 88rpx; height: 88rpx; border-radius: $radius-v50percent; background: $color-tag; }
.flow-avatar--large { flex-basis: 200rpx; width: 200rpx; height: 200rpx; }
.flow-avatar--fallback { display: flex; align-items: center; justify-content: center; color: $color-primary; font-size: $font-size-v32rpx; font-weight: $font-weight-v600; }
</style>
