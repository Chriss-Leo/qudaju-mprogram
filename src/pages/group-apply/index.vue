<script setup lang="ts">
import { computed, ref } from "vue";
import { useFlowId } from "../../utils/use-flow-id";
import FlowShell from "../../components/FlowShell.vue";
import ActivityContext from "../../components/ActivityContext.vue";
import { applyToGroup, getGroup } from "../../stores/meetups";
import { flowError } from "../../utils/flow-navigation";
const id = useFlowId("group-apply"); const note = ref(""); const submitting = ref(false);
const group = computed(() => getGroup(id.value));
function submit() { if (submitting.value) return; submitting.value = true; try { applyToGroup(id.value, note.value); uni.redirectTo({ url: "/pages/my-registrations/index" }); } catch (error) { flowError(error); } finally { submitting.value = false; } }
</script>
<template><FlowShell title="报名加入" fallback="/pages/groups/index"><view v-if="group" class="flow-content"><ActivityContext :group-id="id" /><view class="apply-copy"><text class="flow-name">留一句话，让局主认识你</text><text class="flow-muted">选填，可写兴趣、经验或参加活动的期待。</text><textarea v-model="note" class="apply-note" :maxlength="120" placeholder="例如：我是新手，想一起学玩。" /><text class="flow-muted">{{ note.length }} / 120</text></view><view class="flow-notice">{{ group.admission === '局主审核' ? group.fund ? '审核通过后付款，付款完成才占正式名额。' : '局主审核通过后，占正式名额。' : group.fund ? '付款完成才占正式名额。' : '提交报名后直接加入，占正式名额。' }}</view><text class="flow-muted apply-local">报名在本地保存，当前不会发起真实申请或收款。</text></view><view v-else class="flow-empty"><text>未找到这个活动</text><text>返回列表，选择其他同城活动。</text></view><template v-if="group" #footer><button class="flow-button action-accent" :disabled="submitting" @tap="submit">{{ group.admission === '局主审核' ? '提交报名申请' : '确认报名' }}</button></template></FlowShell></template>
<style scoped lang="scss">
.apply-copy { margin: 40rpx 0; }
.apply-note { box-sizing: border-box; width: 100%; height: 240rpx; margin-top: 24rpx; padding: 24rpx; color: $color-text; background: $color-input; border-radius: $radius-v24rpx; font-size: $font-size-v28rpx; line-height: $line-height-v1p6; }
.apply-local { margin-top: 24rpx; }
</style>
