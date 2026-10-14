<script setup lang="ts">
import { computed } from "vue";
import FlowShell from "../../components/FlowShell.vue";
import ActivityContext from "../../components/ActivityContext.vue";
import { endGroup, getGroup, meetupState } from "../../stores/meetups";
import { flowError, openFlow } from "../../utils/flow-navigation";
const group = computed(() => getGroup("my-game"));
const ended = computed(() => meetupState.ended.includes("my-game"));
const waiting = computed(() => meetupState.registrations.filter(r => r.groupId === "my-game" && r.status === "review").length);
function finish() { uni.showModal({ title: "结束这个局？", content: "结束后停止报名及聊天发送，已接受的聊天记录保留可读。此操作不能撤销。", success: result => { if (result.confirm) try { endGroup("my-game"); } catch (error) { flowError(error); } } }); }
</script>
<template><FlowShell title="我发起的"><view class="flow-content"><view v-if="group" class="flow-card"><ActivityContext group-id="my-game" /><view class="hosted-summary"><text class="flow-badge" :class="{ 'flow-badge--neutral': ended }">{{ ended ? '已结束' : '招募中' }}</text><text>{{ group.joined }} / {{ group.capacity }} 人报名 · {{ waiting }} 个待审核</text></view><view class="flow-actions"><button v-if="!ended" class="flow-button flow-button--muted" @tap="finish">结束活动</button><button class="flow-button" @tap="openFlow('group-members', 'my-game')">管理报名成员</button></view></view><view v-else class="flow-empty"><text>还没有发起活动</text><text>你发布的活动会出现在这里。</text></view><view class="flow-notice">活动与成员为本地记录，用于体验报名管理流程。</view></view></FlowShell></template>
<style scoped lang="scss">
.hosted-summary { display: flex; flex-wrap: wrap; gap: 16rpx; align-items: center; margin-top: 26rpx; color: $color-textMuted; font-size: $font-size-v25rpx; line-height: $line-height-v1p5; }
</style>
