<script setup lang="ts">
import { computed } from "vue";
import FlowAvatar from "./FlowAvatar.vue";
import UiIcon from "./UiIcon.vue";
import { canChat, effectiveStatus, getGroup, hostId, myConversation, people, statusLabel, type Registration } from "../stores/meetups";
import { openFlow } from "../utils/flow-navigation";
const props = defineProps<{ registration: Registration }>();
const emit = defineEmits<{ (event: "withdraw"): void; (event: "pay"): void }>();
const group = computed(() => getGroup(props.registration.groupId)!);
const status = computed(() => effectiveStatus(props.registration));
const history = computed(() => ["ended", "withdrawn", "rejected", "cancelled"].includes(status.value));
const conversation = computed(() => myConversation(props.registration.groupId));
const steps = computed(() => status.value === "review" ? ["申请已提交", "局主审核", "报名成功"] : status.value === "payment" ? ["已通过审核", "完成付款", "报名成功"] : ["报名成功", "等待开始", "活动结束"]);
const currentStep = computed(() => status.value === "ended" ? 2 : 1);
const notice = computed(() => ({ review: "审核通过后才占名额", payment: "完成付款才占正式名额", joined: "活动即将开始，记得按时到达", ongoing: "活动进行中，享受一起参与的时光", ended: "聊天已关闭，历史消息仍可查看", withdrawn: "报名关系已结束，聊天已关闭", rejected: "申请未通过，报名关系已结束", cancelled: "活动报名已取消" }[status.value]));
function meeting() { uni.showModal({ title: "集合地点", content: `${group.value.area} · 本地流程示例集合点，实际集合细节由活动方提供。`, showCancel: false }); }
</script>
<template>
  <view class="flow-card registration-card"><view class="registration-card__heading"><image :src="group.photos[0]" mode="aspectFill" /><view class="flow-copy"><text class="registration-card__title">{{ group.title }}</text><view class="registration-card__owner"><FlowAvatar :person="people[hostId(registration.groupId)]" /><text>{{ group.host }}</text><text class="flow-badge">{{ group.category }}</text></view></view></view>
    <view class="registration-card__meta"><view><UiIcon name="clock" tone="black" /><text>{{ group.time }}</text></view><view><UiIcon name="location" tone="black" /><text>{{ group.area }}</text></view><text class="flow-badge" :class="{ 'flow-badge--neutral': history }">{{ statusLabel[status] }}</text></view>
    <view v-if="['review','payment','joined','ended'].includes(status)" class="registration-steps"><view v-for="(label,index) in steps" :key="label" class="registration-step" :class="{ 'registration-step--done': index < currentStep, 'registration-step--current': index === currentStep, 'registration-step--payment': status === 'payment' && index === currentStep, 'registration-step--history': history }"><view class="registration-step__dot" /><text>{{ label }}</text></view></view>
    <view class="flow-notice registration-notice">{{ notice }}</view>
    <text v-if="status === 'payment'" class="registration-price">¥{{ group.fund }} / 人</text>
    <view class="flow-actions"><template v-if="status === 'review' || status === 'payment'"><button class="flow-button flow-button--muted" @tap="emit('withdraw')">{{ status === 'review' ? '撤回申请' : '取消报名' }}</button><button v-if="status === 'payment'" class="flow-button action-accent" @tap="emit('pay')">去付款</button></template><template v-else-if="status === 'joined'"><button v-if="conversation && canChat(conversation)" class="flow-button" @tap="openFlow('chat', conversation.id)">联系局主</button><button class="flow-button" @tap="meeting">查看集合地点</button><button class="flow-button flow-button--muted" @tap="emit('withdraw')">退出报名</button></template><button v-else-if="status === 'ongoing'" class="flow-button" @tap="openFlow('service-detail', registration.groupId)">查看详情</button><button v-else class="flow-button flow-button--muted" @tap="openFlow(conversation ? 'chat' : 'service-detail', conversation ? conversation.id : registration.groupId)">查看记录</button></view>
  </view>
</template>
<style scoped lang="scss">
.registration-card__heading { display: flex; align-items: center; gap: 22rpx; }
.registration-card__heading > image { flex: 0 0 120rpx; width: 120rpx; height: 120rpx; border-radius: $radius-v16rpx; }
.registration-card__title { display: block; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; font-size: $font-size-v29rpx; font-weight: $font-weight-v600; line-height: $line-height-v1p4; }
.registration-card__owner { display: flex; align-items: center; flex-wrap: wrap; gap: 12rpx; margin-top: 18rpx; color: $color-textMuted; font-size: $font-size-v23rpx; }
.registration-card__owner :deep(.flow-avatar) { flex-basis: 40rpx; width: 40rpx; height: 40rpx; }
.registration-card__meta { display: grid; grid-template-columns: minmax(0,1fr) auto; align-items: center; gap: 14rpx; margin-top: 26rpx; font-size: $font-size-v25rpx; line-height: $line-height-v1p5; }
.registration-card__meta > view { display: flex; align-items: center; gap: 14rpx; }
.registration-card__meta > view:nth-child(2) { grid-row: 2; }
.registration-card__meta > .flow-badge { grid-column: 2; grid-row: 1 / 3; }
.registration-card__meta :deep(.ui-icon) { width: 30rpx; height: 30rpx; }
.registration-steps { display: flex; margin: 34rpx 0; }
.registration-step { position: relative; flex: 1; min-width: 0; display: flex; flex-direction: column; align-items: center; gap: 16rpx; color: $color-textMuted; font-size: $font-size-v23rpx; line-height: $line-height-v1p4; }
.registration-step:not(:last-child)::before { content: ""; position: absolute; top: 15rpx; left: calc(50% + 22rpx); width: calc(100% - 44rpx); height: 2rpx; background: $color-border; }
.registration-step__dot { width: 32rpx; height: 32rpx; border: 3rpx solid $color-borderStrong; border-radius: $radius-v50percent; background: $color-white; }
.registration-step--done .registration-step__dot { border-color: $color-primary; background: $color-primary; }
.registration-step--current { color: $color-primary; }
.registration-step--current .registration-step__dot { border-color: $color-primary; }
.registration-step--payment { color: $color-accent; }
.registration-step--payment .registration-step__dot { border-color: $color-accent; }
.registration-step--history { color: $color-textMuted; }
.registration-step--history .registration-step__dot { background: $color-textMuted; border-color: $color-textMuted; }
.registration-notice { margin-top: 22rpx; }
.registration-price { display: block; margin-top: 22rpx; color: $color-accent; font-size: $font-size-v30rpx; font-weight: $font-weight-v600; }
</style>
