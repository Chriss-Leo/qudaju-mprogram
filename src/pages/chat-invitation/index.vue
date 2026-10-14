<script setup lang="ts">
import { computed, watch } from "vue";
import { useFlowId } from "../../utils/use-flow-id";
import { isAuthenticated } from "../../stores/auth";
import FlowShell from "../../components/FlowShell.vue";
import FlowAvatar from "../../components/FlowAvatar.vue";
import ActivityContext from "../../components/ActivityContext.vue";
import UiIcon from "../../components/UiIcon.vue";
import { me, meetupState, people, relationshipActive, respondInvitation } from "../../stores/meetups";
import { flowError, openFlow } from "../../utils/flow-navigation";
const id = useFlowId("chat-invitation");
const invitation = computed(() => meetupState.invitations.find(i => i.id === id.value && i.to === me.id));
const pending = computed(() => invitation.value?.status === "pending" && relationshipActive(invitation.value));
watch([invitation, isAuthenticated], ([value, authenticated]) => { if (value && authenticated) value.read = true; });
function respond(accept: boolean) { try { respondInvitation(id.value, accept); if (accept) uni.redirectTo({ url: `/pages/chat/index?id=${encodeURIComponent(id.value)}` }); } catch (error) { flowError(error); } }
</script>
<template>
  <FlowShell title="聊天邀请" fallback="/pages/messages/index">
    <view v-if="invitation" class="invite-body">
      <view class="invite-identity"><FlowAvatar :person="people[invitation.from]" large /><text class="invite-name">{{ people[invitation.from].nickname }}</text><text class="flow-muted">局主 · 发来聊天邀请</text></view>
      <ActivityContext :group-id="invitation.groupId" />
      <text class="invite-quote">“{{ invitation.note }}”</text><text class="invite-description">{{ pending ? '接受后开启文字聊天' : !relationshipActive(invitation) ? '报名关系已结束，此邀请已失效' : invitation.status === 'accepted' ? '你已接受聊天邀请' : '你已拒绝聊天邀请' }}</text>
    </view>
    <view v-else class="flow-empty"><text>未找到这条邀请</text><text>返回消息，查看其他邀请。</text></view>
    <template #footer>
      <view v-if="pending" class="invite-controls"><view><button class="invite-circle invite-circle--reject" aria-label="拒绝聊天邀请" @tap="respond(false)"><UiIcon name="close" tone="white" /></button><text>拒绝</text></view><view><button class="invite-circle invite-circle--accept" aria-label="接受聊天邀请" @tap="respond(true)"><UiIcon name="check" tone="white" /></button><text>接受</text></view></view>
      <button v-else-if="invitation?.status === 'accepted'" class="flow-button" @tap="openFlow('chat', id)">查看聊天</button>
      <text class="invite-privacy">不公开你的电话或微信号</text>
    </template>
  </FlowShell>
</template>
<style scoped lang="scss">
.invite-body { padding: 40rpx 36rpx 20rpx; }
.invite-identity { display: flex; flex-direction: column; align-items: center; margin: 28rpx 0 40rpx; }
.invite-name { display: block; margin-top: 24rpx; font-size: $font-size-v38rpx; font-weight: $font-weight-v600; line-height: $line-height-v1p4; }
.invite-quote { display: block; margin-top: 44rpx; text-align: center; font-size: $font-size-v30rpx; line-height: $line-height-v1p6; }
.invite-description { display: block; margin-top: 16rpx; color: $color-textMuted; text-align: center; font-size: $font-size-v26rpx; line-height: $line-height-v1p6; }
.invite-controls { display: flex; justify-content: space-around; padding: 22rpx 24rpx; }
.invite-controls > view { display: flex; flex-direction: column; align-items: center; gap: 20rpx; font-size: $font-size-v28rpx; line-height: $line-height-v1p4; }
.invite-circle { display: flex; align-items: center; justify-content: center; width: 136rpx; height: 136rpx; margin: 0; padding: 0; border-radius: $radius-v50percent; }
.invite-circle--reject { background: $color-inviteReject; }
.invite-circle--accept { background: $color-inviteAccept; }
.invite-circle :deep(.ui-icon) { width: 72rpx; height: 72rpx; }
.invite-circle:active { opacity: .75; }
.invite-circle::after { border: 0; }
.invite-circle:focus-visible { outline: 3rpx solid $color-primary; outline-offset: 5rpx; }
.invite-privacy { display: block; padding: 22rpx 0 14rpx; text-align: center; color: $color-textMuted; font-size: $font-size-v23rpx; line-height: $line-height-v1p4; }
</style>
