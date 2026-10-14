<script setup lang="ts">
import { computed, nextTick, ref, watch } from "vue";
import { useFlowId } from "../../utils/use-flow-id";
import FlowShell from "../../components/FlowShell.vue";
import FlowAvatar from "../../components/FlowAvatar.vue";
import ActivityContext from "../../components/ActivityContext.vue";
import UiIcon from "../../components/UiIcon.vue";
import { canChat, me, meetupState, people, sendMessage } from "../../stores/meetups";
import { flowError, openFlow } from "../../utils/flow-navigation";
const id = useFlowId("chat"); const draft = ref(""); const scrollTo = ref("");
const invitation = computed(() => meetupState.invitations.find(i => i.id === id.value && (i.from === me.id || i.to === me.id)));
const peer = computed(() => invitation.value ? people[invitation.value.from === me.id ? invitation.value.to : invitation.value.from] : undefined);
const accepted = computed(() => invitation.value?.status === "accepted");
const enabled = computed(() => canChat(invitation.value));
const messages = computed(() => accepted.value ? meetupState.messages.filter(m => m.invitationId === id.value) : []);
watch(id, () => { draft.value = ""; });
watch(() => messages.value.length, async () => { scrollTo.value = ""; await nextTick(); scrollTo.value = "chat-bottom"; }, { immediate: true });
function send() { try { if (!enabled.value) return; sendMessage(id.value, draft.value); draft.value = ""; } catch (error) { flowError(error); } }
function time(value: number) { const date = new Date(value); return `${String(date.getHours()).padStart(2, "0")}:${String(date.getMinutes()).padStart(2, "0")}`; }
</script>
<template>
  <FlowShell :title="peer?.nickname || '聊天'" fallback="/pages/messages/index" :scroll-to="scrollTo">
    <template v-if="invitation && accepted" #top><view class="chat-context"><ActivityContext :group-id="invitation.groupId" :label="enabled ? '查看局' : '聊天已关闭'" /></view></template>
    <view v-if="invitation && accepted" class="chat-body">
      <text class="chat-system">已接受聊天邀请</text>
      <view v-for="message in messages" :key="message.id" class="chat-message" :class="{ 'chat-message--mine': message.sender === me.id }"><FlowAvatar :person="people[message.sender]" /><view class="chat-message__copy"><text class="chat-bubble">{{ message.text }}</text><text class="chat-time">{{ time(message.time) }}</text></view></view>
      <text v-if="!messages.length" class="chat-system">围绕本局的安排，开始交流吧。</text>
      <text v-if="!enabled" class="chat-system">本局或报名关系已结束，聊天已关闭</text><view id="chat-bottom" />
    </view>
    <view v-else class="flow-empty"><text>{{ invitation ? '聊天尚未开启' : '未找到这段会话' }}</text><text>接受邀请后才可以查看和发送消息。</text><view v-if="invitation?.to === me.id && invitation.status === 'pending'" class="flow-actions"><button class="flow-button" @tap="openFlow('chat-invitation', id)">查看邀请</button></view></view>
    <template #footer><view v-if="enabled" class="chat-composer"><input v-model="draft" class="chat-input" placeholder="输入消息" :maxlength="500" confirm-type="send" :cursor-spacing="24" @confirm="send" /><button class="flow-button action-accent chat-send" :disabled="!draft.trim()" @tap="send">发送</button></view><view v-else class="chat-closed"><UiIcon name="lock" tone="muted" /><text>{{ accepted ? '聊天已关闭，历史消息仍可查看' : '接受邀请后开启聊天' }}</text></view></template>
  </FlowShell>
</template>
<style scoped lang="scss">
.chat-context { padding: 10rpx 28rpx 20rpx; }
.chat-body { padding: 16rpx 28rpx 28rpx; }
.chat-system { display: block; margin: 26rpx 0; text-align: center; color: $color-textMuted; font-size: $font-size-v23rpx; line-height: $line-height-v1p5; }
.chat-message { display: flex; align-items: flex-start; gap: 16rpx; margin: 28rpx 0; }
.chat-message :deep(.flow-avatar) { flex-basis: 72rpx; width: 72rpx; height: 72rpx; }
.chat-message__copy { min-width: 0; max-width: 78%; }
.chat-message--mine { flex-direction: row-reverse; }
.chat-message--mine .chat-message__copy { text-align: right; }
.chat-bubble { display: block; padding: 20rpx 24rpx; border-radius: $radius-v20rpx; background: $color-card; font-size: $font-size-v28rpx; line-height: $line-height-v1p6; overflow-wrap: anywhere; white-space: pre-wrap; text-align: left; }
.chat-message--mine .chat-bubble { background: $color-tag; }
.chat-time { display: block; margin-top: 8rpx; color: $color-textMuted; font-size: $font-size-v21rpx; line-height: $line-height-v1p4; }
.chat-composer { display: grid; grid-template-columns: minmax(0, 1fr) 152rpx; align-items: center; gap: 18rpx; width: 100%; }
.chat-input { box-sizing: border-box; width: 100%; min-width: 0; height: 96rpx; min-height: 44px; padding: 16rpx 22rpx; border: 1rpx solid $color-border; border-radius: $radius-v24rpx; font-size: $font-size-v28rpx; color: $color-text; background: $color-white; }
.chat-send { width: 100%; height: 96rpx; min-height: 44px; padding: 0; border-radius: $radius-v24rpx; font-size: $font-size-v28rpx; }
.chat-closed { box-sizing: border-box; display: flex; align-items: center; justify-content: center; gap: 14rpx; width: 100%; min-height: 112rpx; padding: 22rpx 20rpx; border-radius: $radius-v24rpx; background: $color-card; color: $color-textMuted; text-align: center; font-size: $font-size-v24rpx; line-height: $line-height-v1p5; }
.chat-closed :deep(.ui-icon) { width: 32rpx; height: 32rpx; }
.chat-closed text { min-width: 0; }
</style>
