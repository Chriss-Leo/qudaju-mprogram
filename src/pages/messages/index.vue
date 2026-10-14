<script setup lang="ts">
import { computed, ref } from "vue";
import FlowShell from "../../components/FlowShell.vue";
import FlowAvatar from "../../components/FlowAvatar.vue";
import UiIcon from "../../components/UiIcon.vue";
import { getGroup, meetupState, me, people, relationshipActive } from "../../stores/meetups";
import { openFlow } from "../../utils/flow-navigation";
const tab = ref("邀请");
const tabs = ["聊天", "通知", "邀请"];
const invites = computed(() => meetupState.invitations.filter(i => i.to === me.id).slice().reverse());
const chats = computed(() => meetupState.invitations.filter(i => (i.to === me.id || i.from === me.id) && i.status === "accepted"));
function select(value: string) { tab.value = value; if (value === "通知") meetupState.notificationsRead = true; }
function peer(id: string) { const i = meetupState.invitations.find(i => i.id === id)!; return people[i.from === me.id ? i.to : i.from]; }
function lastMessage(id: string) { return meetupState.messages.filter(m => m.invitationId === id).slice(-1)[0]?.text || "聊天邀请已接受，可以开始交流。"; }
</script>
<template>
  <FlowShell title="消息">
    <template #top><view class="flow-tabs"><button v-for="item in tabs" :key="item" class="flow-tab" :class="{ 'flow-tab--active': tab === item }" @tap="select(item)">{{ item }}<text v-if="item === '邀请' && invites.some(i => !i.read && i.status === 'pending' && relationshipActive(i))" class="message-dot" /></button></view></template>
    <view class="flow-content">
      <template v-if="tab === '邀请'">
        <view v-for="i in invites" :key="i.id" class="flow-card invitation-row">
          <view class="message-identity"><FlowAvatar :person="people[i.from]" /><view class="message-titles"><text class="flow-name message-name">{{ people[i.from].nickname }}</text><text class="message-heading">邀请你聊一聊</text></view><text class="flow-badge message-status" :class="{ 'flow-badge--neutral': i.status !== 'pending' || !relationshipActive(i) }">{{ !relationshipActive(i) ? i.status === 'accepted' ? '已关闭' : '已失效' : i.status === 'accepted' ? '已接受' : i.status === 'rejected' ? '已拒绝' : '待回应' }}</text></view>
          <view class="message-details"><view class="message-details__copy"><text class="message-activity">{{ getGroup(i.groupId)?.title }}</text><text class="flow-muted message-note">{{ i.note }}</text></view><button class="flow-button message-action" @tap="openFlow(i.status === 'accepted' ? 'chat' : 'chat-invitation', i.id)">{{ i.status === 'accepted' ? '查看聊天' : '查看邀请' }}</button></view>
        </view>
        <view v-if="!invites.length" class="flow-empty"><text>暂时没有聊天邀请</text><text>报名活动后，局主可以邀请你交流。</text></view>
        <view class="flow-notice message-explanation"><UiIcon name="info" tone="muted" /><text>接受邀请后才可聊天；活动结束或退出后停止发送。</text></view>
      </template>
      <template v-else-if="tab === '聊天'">
        <button v-for="i in chats" :key="i.id" class="conversation-row" @tap="openFlow('chat', i.id)"><FlowAvatar :person="peer(i.id)" /><view class="flow-copy"><text class="flow-name">{{ peer(i.id).nickname }}</text><text class="flow-muted conversation-preview">{{ lastMessage(i.id) }}</text><text class="flow-muted">{{ getGroup(i.groupId)?.title }}</text></view></button>
        <view v-if="!chats.length" class="flow-empty"><text>还没有开启的聊天</text><text>接受局主的邀请后，会话会出现在这里。</text></view>
      </template>
      <template v-else><view class="flow-card"><text class="flow-name">报名进度，随时查看</text><text class="flow-muted">审核、付款与参加状态集中在「我的报名」。</text><view class="flow-actions"><button class="flow-button" @tap="openFlow('my-registrations')">查看报名</button></view></view><view class="flow-notice">当前通知为本地流程提示，不包含推送服务。</view></template>
    </view>
  </FlowShell>
</template>
<style scoped lang="scss">
.message-dot { width: 12rpx; height: 12rpx; margin-left: 12rpx; border-radius: $radius-v50percent; background: $color-accent; }
.invitation-row { padding: 24rpx; }
.message-identity { display: flex; align-items: center; gap: 20rpx; }
.message-titles { display: flex; flex: 1; min-width: 0; flex-direction: column; gap: 8rpx; }
.message-name { overflow: hidden; white-space: nowrap; text-overflow: ellipsis; }
.message-status { flex-shrink: 0; align-self: flex-start; }
.message-heading { font-size: $font-size-v27rpx; font-weight: $font-weight-v600; line-height: $line-height-v1p5; }
.message-details { display: grid; grid-template-columns: minmax(0, 1fr) auto; align-items: center; gap: 18rpx; margin-top: 22rpx; }
.message-details__copy { min-width: 0; }
.message-activity { display: block; min-width: 0; overflow: hidden; text-overflow: ellipsis; white-space: nowrap; color: $color-textSecondary; font-size: $font-size-v24rpx; line-height: $line-height-v1p5; }
.message-note { margin-top: 10rpx; overflow-wrap: anywhere; }
.message-action { width: auto; padding: 10rpx 16rpx; }
.message-explanation :deep(.ui-icon) { width: 32rpx; height: 32rpx; }
.message-explanation text { min-width: 0; }
.conversation-row { display: flex; align-items: flex-start; gap: 20rpx; width: 100%; padding: 28rpx 0; margin: 0; text-align: left; background: transparent; border-radius: $radius-v0; border-bottom: 1rpx solid $color-border; }
.conversation-preview { overflow: hidden; white-space: nowrap; text-overflow: ellipsis; }
.conversation-row::after { border: 0; }
</style>
