<script setup lang="ts">
import { computed, ref, watch } from "vue";
import BottomSheet from "./BottomSheet.vue";
import UiIcon from "./UiIcon.vue";
import { authState, isAuthenticated, saveMyProfile } from "../stores/auth";
const props = defineProps<{ visible: boolean }>();
const emit = defineEmits<{ close: []; saved: [] }>();
const nickname = ref(""); const avatar = ref(""); const picked = ref(""); const busy = ref(false); const failed = ref(false); const error = ref("");
const validName = computed(() => { const name = nickname.value.trim(); const length = Array.from(name).length; return length >= 2 && length <= 20 && name !== "微信用户" && !/[\u0000-\u001f\u007f-\u009f]/.test(name); });
const canSave = computed(() => validName.value && !!avatar.value && !busy.value && !failed.value);
watch(() => props.visible, visible => {
  if (!visible) return;
  const user = authState.session?.user;
  nickname.value = user?.nickname === "微信用户" ? "" : user?.nickname || "";
  avatar.value = user?.avatarUrl || ""; picked.value = ""; failed.value = false; error.value = "";
}, { immediate: true });
watch(isAuthenticated, value => { if (!value) emit('close'); });
function close() { if (!busy.value) emit('close'); }
function onChooseAvatar(event: { detail: { avatarUrl?: string } }) {
  const tempPath = event.detail.avatarUrl;
  if (busy.value || !tempPath) return;
  picked.value = tempPath; avatar.value = tempPath; failed.value = false; error.value = "";
}
function chooseImage() {
  uni.chooseImage({ count: 1, sizeType: ["compressed"], success: result => { if (result.tempFilePaths[0]) onChooseAvatar({ detail: { avatarUrl: result.tempFilePaths[0] } }); } });
}
function nicknameBlur(event: unknown) {
  const value = (event as { detail?: { value?: unknown } }).detail?.value;
  if (typeof value === 'string') nickname.value = value;
}
async function submit(event: unknown) {
  if (busy.value) return;
  const value = (event as { detail?: { value?: { nickname?: unknown } } }).detail?.value?.nickname;
  if (typeof value === 'string') nickname.value = value;
  if (!canSave.value) { error.value = failed.value ? "头像加载失败，请重新选择" : "请选择头像，并填写2–20字昵称"; return; }
  busy.value = true; error.value = "";
  try { await saveMyProfile(nickname.value.trim(), picked.value || undefined); emit('saved'); }
  catch (value) { error.value = value instanceof Error ? value.message : "保存失败，请重试"; }
  finally { busy.value = false; }
}
</script>

<template>
  <BottomSheet class="profile-sheet" :visible="visible && isAuthenticated" title="确认头像与昵称" @close="close">
      <form @submit="submit"><view class="profile-form">
        <text class="profile-form__description">选择头像和昵称，用于展示你的个人资料。</text>
        <view class="avatar-field">
          <!-- #ifdef MP-WEIXIN -->
          <button class="avatar-picker" open-type="chooseAvatar" :disabled="busy" aria-label="选择头像" @chooseavatar="onChooseAvatar">
            <image v-if="avatar && !failed" :src="avatar" mode="aspectFill" @error="failed = true" />
            <UiIcon v-else name="people" />
          </button>
          <!-- #endif -->
          <!-- #ifndef MP-WEIXIN -->
          <button class="avatar-picker" :disabled="busy" aria-label="选择头像" @tap="chooseImage">
            <image v-if="avatar && !failed" :src="avatar" mode="aspectFill" @error="failed = true" />
            <UiIcon v-else name="people" />
          </button>
          <!-- #endif -->
          <text class="avatar-field__label">{{ avatar ? '点击更换头像' : '点击选择头像' }}</text>
          <text class="avatar-field__hint">JPG / PNG，2MB以内</text>
        </view>
        <label class="nickname-field">
          <text>昵称</text>
          <!-- #ifdef MP-WEIXIN -->
          <input v-model="nickname" name="nickname" type="nickname" :disabled="busy" :maxlength="20" placeholder="填写你的昵称" confirm-type="done" @blur="nicknameBlur" @change="nicknameBlur" />
          <!-- #endif -->
          <!-- #ifndef MP-WEIXIN -->
          <input v-model="nickname" name="nickname" :disabled="busy" :maxlength="20" placeholder="填写你的昵称" confirm-type="done" />
          <!-- #endif -->
        </label>
        <text class="nickname-hint">点击昵称输入框，可选用微信昵称，也可以修改。</text>
        <text v-if="error" class="profile-error" role="alert">{{ error }}</text>
        <button class="profile-save action-accent" form-type="submit" :disabled="busy">确认并继续</button>
        <text class="profile-privacy">头像与昵称由你主动提供，可在「我的 → 编辑资料」修改。</text>
      </view></form>
  </BottomSheet>
</template>

<style scoped lang="scss">
.profile-sheet :deep(.sheet) { border-top: 0; background: $color-white; }
.profile-sheet :deep(.sheet__close) { margin: 0; }
.profile-form { padding: 8rpx 0 12rpx; color: $color-text; }
.profile-form__description { display: block; color: $color-textSecondary; font-size: $font-size-v26rpx; line-height: $line-height-v1p6; }
.avatar-field { display: flex; flex-direction: column; align-items: center; margin: 30rpx 0; gap: 12rpx; }
.avatar-picker { display: flex; align-items: center; justify-content: center; width: 156rpx; height: 156rpx; padding: 0; margin: 0; background: $color-input; border-radius: $radius-v50percent; overflow: hidden; }
.avatar-picker::after { border: 0; }
.avatar-picker image { width: 100%; height: 100%; }
.avatar-picker :deep(.ui-icon) { width: 68rpx; height: 68rpx; }
.avatar-field__label { color: $color-primary; font-size: $font-size-v28rpx; font-weight: $font-weight-v600; line-height: $line-height-v1p4; }
.avatar-field__hint, .nickname-hint, .profile-privacy { color: $color-textMuted; font-size: $font-size-v24rpx; line-height: $line-height-v1p6; }
.nickname-field { display: flex; align-items: center; gap: 28rpx; padding: 26rpx 24rpx; background: $color-card; border-radius: $radius-v16rpx; font-size: $font-size-v30rpx; line-height: $line-height-v1p4; }
.nickname-field input { flex: 1; min-width: 0; height: 52rpx; color: $color-text; font-size: $font-size-v30rpx; }
.nickname-hint { display: block; margin-top: 16rpx; }
.profile-error { display: block; margin-top: 24rpx; color: $color-text; font-size: $font-size-v26rpx; line-height: $line-height-v1p6; }
.profile-save { display: flex; align-items: center; justify-content: center; width: 100%; min-height: 96rpx; margin-top: 28rpx; padding: 22rpx; border-radius: $radius-v24rpx; font-size: $font-size-v30rpx; font-weight: $font-weight-v600; line-height: $line-height-v1p4; }
.profile-privacy { display: block; margin-top: 24rpx; }
.avatar-picker:focus-visible, .profile-save:focus-visible { outline: 3rpx solid $color-primary; outline-offset: 4rpx; }
</style>
