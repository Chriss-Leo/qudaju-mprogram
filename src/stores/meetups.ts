import { reactive, computed, watch } from "vue";
import { groupCards, type GroupCardData, type GroupParticipant } from "../mocks/group-cards";

// Local interaction model only. A server must independently enforce every rule.
export type RegistrationStatus = "review" | "payment" | "joined" | "ongoing" | "ended" | "withdrawn" | "rejected" | "cancelled";
export interface Registration { id: string; groupId: string; userId: string; status: RegistrationStatus; note: string }
export interface Invitation { id: string; groupId: string; from: string; to: string; status: "pending" | "accepted" | "rejected"; note: string; read: boolean; closed?: boolean }
export interface ChatMessage { id: string; invitationId: string; sender: string; text: string; time: number }
interface State { version: 1; registrations: Registration[]; invitations: Invitation[]; messages: ChatMessage[]; ended: string[]; notificationsRead: boolean }
export const me: GroupParticipant = { id: "me", nickname: "周末探险家", avatar: "/static/profile-a/profile-avatar.jpg" };
export const people: Record<string, GroupParticipant> = { me };
Object.entries(groupCards).forEach(([id, group]) => {
  people[`owner-${id}`] = { id: `owner-${id}`, nickname: group.host, avatar: group.hostAvatar };
  group.participants.forEach(person => { if (!person.id.startsWith("host-")) people[person.id] = person; });
});
const activities: Record<string, GroupCardData> = {
  ...groupCards,
  "my-game": { ...groupCards["game-01"], title: "周末桌游，一起轻松玩", host: me.nickname, hostAvatar: me.avatar, participants: [me, ...groupCards["game-01"].participants.slice(1)], isJoinedByMe: false },
  "past-game": { ...groupCards["game-01"], title: "上次桌游，快乐回顾", time: "上周六 14:00" },
};
export const hostId = (id: string) => id === "my-game" ? me.id : id === "past-game" ? "owner-game-01" : `owner-${id}`;
const fresh = (): State => ({
  version: 1, ended: ["past-game"], notificationsRead: false,
  registrations: [
    { id: "r-review", groupId: "game-01", userId: me.id, status: "review", note: "我是新手，想一起学玩。" },
    { id: "r-payment", groupId: "photo-01", userId: me.id, status: "payment", note: "带相机一起互拍。" },
    { id: "r-joined", groupId: "outdoor-01", userId: me.id, status: "joined", note: "周日见！" },
    { id: "r-ongoing", groupId: "city-guide-01", userId: me.id, status: "ongoing", note: "想了解街区故事。" },
    { id: "r-ended", groupId: "past-game", userId: me.id, status: "ended", note: "很开心认识大家。" },
    { id: "r-withdrawn", groupId: "guide-02", userId: me.id, status: "withdrawn", note: "" },
    ...["sample-1", "sample-2", "sample-3"].map((userId, index) => ({ id: `host-member-${index}`, groupId: "my-game", userId, status: "joined" as const, note: ["我是新手，想一起学玩。", "一起玩桌游呀！", "很喜欢策略类游戏。"][index] })),
    { id: "host-review", groupId: "my-game", userId: "sample-4", status: "review", note: "周末有空，想参加。" },
  ],
  invitations: [
    { id: "invite-me", groupId: "outdoor-01", from: "owner-outdoor-01", to: me.id, status: "pending", note: "想确认一下你的徒步经验。", read: false },
    { id: "invite-active", groupId: "city-guide-01", from: "owner-city-guide-01", to: me.id, status: "accepted", note: "一起聊聊街区路线。", read: true },
    { id: "invite-history", groupId: "past-game", from: "owner-game-01", to: me.id, status: "accepted", note: "欢迎加入桌游局。", read: true },
    { id: "invite-host-pending", groupId: "my-game", from: me.id, to: "sample-2", status: "pending", note: "想确认一下你的桌游经验。", read: true },
    { id: "invite-host-accepted", groupId: "my-game", from: me.id, to: "sample-3", status: "accepted", note: "周末见！", read: true },
  ],
  messages: [
    { id: "m1", invitationId: "invite-active", sender: "owner-city-guide-01", text: "你好，路线以街区慢走为主，可以边走边聊。", time: Date.now() - 60000 },
    { id: "m2", invitationId: "invite-active", sender: me.id, text: "好呀，想听听这些建筑的故事。", time: Date.now() - 30000 },
    { id: "m3", invitationId: "invite-history", sender: "owner-game-01", text: "欢迎！我们会先讲规则，不用担心。", time: Date.now() - 86400000 },
    { id: "m4", invitationId: "invite-history", sender: me.id, text: "谢谢，周六见！", time: Date.now() - 86000000 },
    { id: "m5", invitationId: "invite-host-accepted", sender: "sample-3", text: "周末见！我可以帮忙讲规则。", time: Date.now() - 30000 },
  ],
});
const key = "qudaju-local-meetups-v1";
function restore(): State {
  try {
    const s = uni.getStorageSync(key) as State;
    const statuses = ["review", "payment", "joined", "ongoing", "ended", "withdrawn", "rejected", "cancelled"];
    if (s?.version === 1 && Array.isArray(s.registrations) && Array.isArray(s.invitations) && Array.isArray(s.messages) && Array.isArray(s.ended)
      && s.registrations.every(r => r && typeof r.id === "string" && !!activities[r.groupId] && !!people[r.userId] && statuses.includes(r.status) && typeof r.note === "string")
      && s.invitations.every(i => i && typeof i.id === "string" && !!activities[i.groupId] && !!people[i.from] && !!people[i.to] && i.from === hostId(i.groupId) && i.to !== i.from && ["pending", "accepted", "rejected"].includes(i.status) && typeof i.note === "string" && typeof i.read === "boolean" && (i.closed === undefined || typeof i.closed === "boolean"))
      && s.messages.every(m => m && typeof m.id === "string" && typeof m.invitationId === "string" && !!people[m.sender] && typeof m.text === "string" && Number.isFinite(m.time))
      && s.ended.every(id => typeof id === "string" && !!activities[id]) && typeof s.notificationsRead === "boolean") return s;
  } catch { /* Malformed/older local data falls back to fixtures. */ }
  return fresh();
}
const initial = restore();
initial.invitations.forEach(i => {
  const r = initial.registrations.find(r => r.groupId === i.groupId && r.userId === i.to);
  if (!r || !["review", "payment", "joined", "ongoing"].includes(r.status) || initial.ended.includes(i.groupId)) i.closed = true;
});
export const meetupState = reactive<State>(initial);
watch(meetupState, () => { try { uni.setStorageSync(key, JSON.parse(JSON.stringify(meetupState))); } catch { uni.showToast({ title: "本地保存失败，请检查存储空间", icon: "none" }); } }, { deep: true });
export const myRegistration = (groupId: string) => meetupState.registrations.find(r => r.groupId === groupId && r.userId === me.id);
export const isFormal = (status?: RegistrationStatus) => status === "joined" || status === "ongoing";
export const effectiveStatus = (r: Registration): RegistrationStatus => meetupState.ended.includes(r.groupId) && isFormal(r.status) ? "ended" : r.status;
export const statusLabel: Record<RegistrationStatus, string> = { review: "待审核", payment: "待付款", joined: "已报名", ongoing: "活动中", ended: "已结束", withdrawn: "已退出", rejected: "未通过", cancelled: "已取消" };
export function getGroup(id: string): GroupCardData | undefined {
  const g = Object.prototype.hasOwnProperty.call(activities, id) ? activities[id] : undefined; if (!g) return undefined;
  const roster = g.participants.map(p => ({ ...p }));
  // Existing counts include the host. Local current-user state adjusts one seat.
  const baselineHasMe = !!g.isJoinedByMe;
  const r = myRegistration(id);
  const hasMe = isFormal(r?.status) || r?.status === "ended";
  if (baselineHasMe && roster.length > 1) roster[roster.length - 1] = me;
  if (hasMe && !baselineHasMe) roster.push(me);
  if (!hasMe && baselineHasMe) { const index = roster.findIndex(p => p.id === me.id); if (index >= 0) roster.splice(index, 1); }
  if (id === "my-game") {
    const members = meetupState.registrations.filter(r => r.groupId === id && isFormal(r.status)).map(r => people[r.userId]);
    return { ...g, participants: [me, ...members], joined: members.length + 1, isJoinedByMe: false };
  }
  return { ...g, participants: roster, joined: g.joined + Number(hasMe) - Number(baselineHasMe), isJoinedByMe: hasMe };
}
export const unreadCount = computed(() => meetupState.invitations.filter(i => i.to === me.id && i.status === "pending" && !i.read && relationshipActive(i)).length + Number(!meetupState.notificationsRead));
export function relationshipActive(i: Invitation) {
  if (i.closed) return false;
  const registration = meetupState.registrations.find(r => r.groupId === i.groupId && r.userId === i.to);
  return !!registration && ["review", "payment", "joined", "ongoing"].includes(registration.status) && !meetupState.ended.includes(i.groupId);
}
export function canChat(i?: Invitation) { return !!i && i.status === "accepted" && (i.from === me.id || i.to === me.id) && relationshipActive(i); }
export function myConversation(groupId: string) { return meetupState.invitations.slice().reverse().find(i => i.groupId === groupId && i.to === me.id && i.status === "accepted"); }
const uid = (prefix: string) => `${prefix}-${Date.now()}-${Math.random().toString(36).slice(2, 8)}`;
export function requestInvitation(groupId: string, userId: string, note: string): string {
  if (hostId(groupId) !== me.id || userId === me.id) throw new Error("仅局主可以向本局报名成员发起邀请");
  const r = meetupState.registrations.find(r => r.groupId === groupId && r.userId === userId);
  if (!r || !["review", "payment", "joined", "ongoing"].includes(r.status) || meetupState.ended.includes(groupId)) throw new Error("该报名关系已失效");
  const existing = meetupState.invitations.find(i => i.groupId === groupId && i.to === userId && !i.closed);
  if (existing) return existing.id;
  const id = uid("invite");
  meetupState.invitations.push({ id, groupId, from: me.id, to: userId, note: note.trim().slice(0, 120) || "想和你聊聊这次活动。", status: "pending", read: true });
  return id;
}
export function respondInvitation(id: string, accept: boolean) {
  const i = meetupState.invitations.find(i => i.id === id);
  if (!i || i.to !== me.id || i.status !== "pending" || !relationshipActive(i)) throw new Error("邀请已失效或已处理");
  i.status = accept ? "accepted" : "rejected"; i.read = true;
}
export function sendMessage(id: string, text: string) {
  const i = meetupState.invitations.find(i => i.id === id);
  if (!canChat(i)) throw new Error("聊天尚未开启或已经关闭");
  const content = text.trim(); if (!content) return;
  if (content.length > 500) throw new Error("消息最多500字");
  meetupState.messages.push({ id: uid("message"), invitationId: id, sender: me.id, text: content, time: Date.now() });
}
export function applyToGroup(groupId: string, note: string) {
  const g = getGroup(groupId); if (!g || meetupState.ended.includes(groupId) || hostId(groupId) === me.id) throw new Error("当前活动无法报名");
  const existing = myRegistration(groupId);
  if (existing && ["review", "payment", "joined", "ongoing"].includes(existing.status)) return existing;
  if (g.joined >= g.capacity) throw new Error("名额已满");
  const status: RegistrationStatus = g.admission === "局主审核" ? "review" : g.fund ? "payment" : "joined";
  const r: Registration = { id: existing?.id || uid("registration"), groupId, userId: me.id, status, note: note.trim().slice(0, 120) };
  if (existing) Object.assign(existing, r); else meetupState.registrations.push(r);
  return r;
}
export function withdraw(id: string) {
  const r = meetupState.registrations.find(r => r.id === id);
  if (!r || r.userId !== me.id || !["review", "payment", "joined", "ongoing"].includes(effectiveStatus(r))) throw new Error("当前报名状态不能退出");
  if (isFormal(r.status) && (getGroup(r.groupId)?.fund || 0) > 0) throw new Error("收费局退款规则尚待接入，请联系活动方处理");
  r.status = "withdrawn";
  meetupState.invitations.filter(i => i.groupId === r.groupId && i.to === r.userId).forEach(i => { i.closed = true; });
}
export function simulatePayment(id: string) {
  const r = meetupState.registrations.find(r => r.id === id); const g = r && getGroup(r.groupId);
  if (!r || !g || r.userId !== me.id || r.status !== "payment" || meetupState.ended.includes(r.groupId)) throw new Error("当前无法付款");
  if (g.joined >= g.capacity) throw new Error("名额已满，暂不能占正式名额");
  r.status = "joined";
}
export function reviewMember(id: string, approve: boolean) {
  const r = meetupState.registrations.find(r => r.id === id); const g = r && getGroup(r.groupId);
  if (!r || !g || hostId(r.groupId) !== me.id || r.status !== "review" || meetupState.ended.includes(r.groupId)) throw new Error("当前无法审核");
  if (approve && !g.fund && g.joined >= g.capacity) throw new Error("名额已满");
  r.status = approve ? g.fund ? "payment" : "joined" : "rejected";
}
export function endGroup(groupId: string) {
  if (hostId(groupId) !== me.id) throw new Error("仅局主可以结束本局");
  if (!meetupState.ended.includes(groupId)) meetupState.ended.push(groupId);
  meetupState.invitations.filter(i => i.groupId === groupId).forEach(i => { i.closed = true; });
  meetupState.registrations.filter(r => r.groupId === groupId && ["review", "payment"].includes(r.status)).forEach(r => { r.status = "cancelled"; });
}
