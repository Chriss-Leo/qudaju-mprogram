<script setup lang="ts">
import { computed, ref } from "vue";
import FlowShell from "../../components/FlowShell.vue";
import RegistrationCard from "../../components/RegistrationCard.vue";
import { effectiveStatus, me, meetupState, simulatePayment, withdraw, type Registration } from "../../stores/meetups";
import { theme } from "../../theme";
import { flowError } from "../../utils/flow-navigation";
const tabs = ["全部", "待处理", "进行中", "已结束"]; const tab = ref("全部");
const records = computed(() => meetupState.registrations.filter(r => r.userId === me.id).filter(r => {
  const status = effectiveStatus(r); return tab.value === "全部" || (tab.value === "待处理" ? ["review", "payment"] : tab.value === "进行中" ? ["joined", "ongoing"] : ["ended", "withdrawn", "rejected", "cancelled"]).includes(status);
}));
function leave(r: Registration) { uni.showModal({ title: "取消这次报名？", content: "退出后，本局相关聊天将关闭发送，历史记录保留。", confirmColor: theme.colors.primary, success: result => { if (result.confirm) try { withdraw(r.id); } catch (error) { flowError(error); } } }); }
function pay(r: Registration) { uni.showModal({ title: "本地付款流程体验", content: "当前没有接入支付，本次不会收款。确认仅模拟付款完成与占名额状态。", confirmText: "模拟完成", confirmColor: theme.colors.primary, success: result => { if (result.confirm) try { simulatePayment(r.id); } catch (error) { flowError(error); } } }); }
</script>
<template><FlowShell title="我的报名"><template #top><view class="flow-tabs"><button v-for="item in tabs" :key="item" class="flow-tab" :class="{ 'flow-tab--active': tab === item }" @tap="tab = item">{{ item }}</button></view></template><view class="flow-content"><RegistrationCard v-for="r in records" :key="r.id" :registration="r" @withdraw="leave(r)" @pay="pay(r)" /><view v-if="!records.length" class="flow-empty"><text>这里还没有报名记录</text><text>去发现同城好局，找到想参加的活动。</text></view></view></FlowShell></template>
