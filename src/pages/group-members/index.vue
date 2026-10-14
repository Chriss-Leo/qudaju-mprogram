<script setup lang="ts">
import { computed, ref } from "vue";
import { useFlowId } from "../../utils/use-flow-id";
import FlowShell from "../../components/FlowShell.vue";
import FlowAvatar from "../../components/FlowAvatar.vue";
import ActivityContext from "../../components/ActivityContext.vue";
import BottomSheet from "../../components/BottomSheet.vue";
import { effectiveStatus, hostId, me, meetupState, people, requestInvitation, reviewMember, statusLabel, type Registration } from "../../stores/meetups";
import { flowError, openFlow } from "../../utils/flow-navigation";
const id = useFlowId("group-members"); const tab = ref("已报名"); const selected = ref<Registration>(); const note = ref("");
const authorized = computed(() => hostId(id.value) === me.id);
const registrations = computed(() => authorized.value ? meetupState.registrations.filter(r => r.groupId === id.value) : []);
const members = computed(() => registrations.value.filter(r => tab.value === "待审核" ? effectiveStatus(r) === "review" : ["joined", "ongoing", "ended"].includes(effectiveStatus(r))));
function invitation(r: Registration) { return meetupState.invitations.slice().reverse().find(i => i.groupId === r.groupId && i.to === r.userId && (effectiveStatus(r) === 'ended' || !i.closed)); }
function buttonLabel(r: Registration) { const i = invitation(r); return i?.status === "accepted" ? "进入聊天" : i?.status === "pending" ? "等待回应" : i?.status === "rejected" ? "已拒绝" : "发起聊天"; }
function start(r: Registration) { const i = invitation(r); if (i?.status === "accepted") openFlow("chat", i.id); else { selected.value = r; note.value = "想和你聊聊这次活动。"; } }
function send() { if (!selected.value) return; try { requestInvitation(id.value, selected.value.userId, note.value); selected.value = undefined; uni.showToast({ title: "邀请已发送，等待回应", icon: "none" }); } catch (error) { flowError(error); } }
function review(r: Registration, approve: boolean) { uni.showModal({ title: approve ? "通过报名申请？" : "拒绝报名申请？", content: approve ? "通过后成员将获得本局正式名额。" : "拒绝后该报名关系结束。", success: result => { if (result.confirm) try { reviewMember(r.id, approve); } catch (error) { flowError(error); } } }); }
</script>
<template>
  <FlowShell title="报名成员" fallback="/pages/my-hosted/index">
    <template v-if="authorized" #top><view class="members-context"><ActivityContext :group-id="id" /></view><view class="flow-tabs"><button v-for="item in ['已报名','待审核']" :key="item" class="flow-tab" :class="{ 'flow-tab--active': tab === item }" @tap="tab = item">{{ item }} {{ registrations.filter(r => item === '待审核' ? effectiveStatus(r) === 'review' : ['joined','ongoing','ended'].includes(effectiveStatus(r))).length }}</button></view></template>
    <view v-if="authorized" class="flow-content">
      <view v-for="r in members" :key="r.id" class="flow-card"><view class="flow-row"><FlowAvatar :person="people[r.userId]" /><view class="flow-copy"><text class="flow-name">{{ people[r.userId].nickname }}</text><text class="flow-muted">{{ statusLabel[effectiveStatus(r)] }}</text></view><button v-if="tab === '已报名'" class="flow-button" :disabled="invitation(r)?.status === 'pending' || invitation(r)?.status === 'rejected' || effectiveStatus(r) === 'ended' && invitation(r)?.status !== 'accepted'" @tap="start(r)">{{ buttonLabel(r) }}</button></view><text class="flow-muted member-note">报名留言：{{ r.note || '没有留下留言' }}</text><view v-if="tab === '待审核'" class="flow-actions"><button class="flow-button flow-button--muted" @tap="review(r, false)">拒绝</button><button class="flow-button action-accent" @tap="review(r, true)">通过申请</button></view></view>
      <view v-if="!members.length" class="flow-empty"><text>{{ tab === '待审核' ? '没有待审核的申请' : '还没有报名成员' }}</text><text>成员报名后会显示在这里。</text></view><view class="flow-notice">仅可邀请报名本局的成员，接受后开启聊天。</view>
    </view><view v-else class="flow-empty"><text>仅局主可管理本局成员</text><text>请从「我发起的」进入自己的活动。</text></view>
  </FlowShell>
  <BottomSheet :visible="!!selected" title="成员信息与聊天邀请" @close="selected = undefined"><view v-if="selected" class="member-sheet"><view class="flow-row"><FlowAvatar :person="people[selected.userId]" /><view class="flow-copy"><text class="flow-name">{{ people[selected.userId].nickname }}</text><text class="flow-muted">已报名本局 · 不公开联系方式</text></view></view><text class="flow-muted">报名留言：{{ selected.note || '未留言' }}</text><textarea v-model="note" class="member-invite-input" :maxlength="120" placeholder="写一句邀请，说明想聊什么" /><button class="flow-button action-accent" @tap="send">发送聊天邀请</button></view></BottomSheet>
</template>
<style scoped lang="scss">
.members-context { padding: 12rpx 28rpx 20rpx; }
.member-note { margin-top: 22rpx; }
.member-sheet { padding-top: 22rpx; }
.member-invite-input { box-sizing: border-box; width: 100%; height: 170rpx; padding: 24rpx; margin: 26rpx 0; border-radius: $radius-v16rpx; background: $color-input; color: $color-text; font-size: $font-size-v28rpx; line-height: $line-height-v1p6; }
@media (max-width: 350px) { .flow-row { flex-wrap: wrap; } .flow-row .flow-button { margin-left: auto; } }
</style>
