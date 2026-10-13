<script setup lang="ts">
import { ref } from "vue";
import { onShow } from "@dcloudio/uni-app";
import { theme } from "../../theme";
import AppTabBar from "../../components/AppTabBar.vue";
import SurfaceHeader from "../../components/SurfaceHeader.vue";
import UiIcon from "../../components/UiIcon.vue";

const avatarFailed = ref(false);
const primaryLinks = [{ label: "我的报名", icon: "calendar" }, { label: "我发起的", icon: "launch" }];
const compactLinks = [{ label: "我的收藏", icon: "favorite" }, { label: "费用记录", icon: "wallet" }];
const utilityLinks = [{ label: "客服与帮助", icon: "support" }, { label: "设置", icon: "settings" }];
function pending(label: string) { uni.showToast({ title: label + "尚未接入，当前为界面预览", icon: "none" }); }
function publish() { uni.showToast({ title: "发布组局功能开发中", icon: "none" }); }
onShow(() => uni.setNavigationBarColor({ frontColor: "#000000", backgroundColor: theme.colors.white }));
</script>

<template>
  <view class="profile-page" data-design-contract="THESIS: layered identity, no service-shop grid. OWN-WORLD: white, blue utilities, orange primary. STORY: identity then participation. FIRST VIEWPORT: blue back card, curved white intro, meetup actions. FORM: user-approved A 2026-10-13. FINISH: unreviewed and undocumented is unfinished; this build ends with the finish review, the verdict, and DESIGN.md">
    <!-- THESIS: A layered identity card, open meetup actions, no legacy service-shop grid.
         OWN-WORLD: white, blue ordinary controls, orange primary actions; real text plus independent raster assets.
         STORY: identity, participation, creation, utilities. FIRST VIEWPORT: blue back card, curved white intro, clear action rows.
         FORM: User-approved A, 2026-10-13, .impeccable/mocks/A-profile-detail.png.
         FINISH: unreviewed and undocumented is unfinished; this build ends with the finish review, the verdict, and DESIGN.md -->
    <SurfaceHeader title="我的" messages @messages="pending('消息')" />
    <view class="profile-body">
      <view class="identity-card">
        <view class="identity-card__back">
          <view class="identity-card__person">
            <image v-if="!avatarFailed" class="identity-card__avatar" src="/static/profile-a/profile-avatar.jpg" mode="aspectFill" aria-label="示例个人头像" @error="avatarFailed = true" />
            <view v-else class="identity-card__avatar identity-card__fallback">周</view>
            <view class="identity-card__name">
              <text>周末探险家</text>
              <button class="edit-profile" @tap="pending('编辑资料')">编辑资料</button>
            </view>
          </view>
          <image class="identity-card__tickets" src="/static/profile-a/decorative-tickets.png" mode="aspectFit" aria-hidden="true" />
        </view>
        <view class="identity-card__intro">
          <image class="identity-card__seam" src="/static/ui-icons/identity-seam.svg" mode="scaleToFill" aria-hidden="true" />
          <text class="identity-card__bio">喜欢散步、桌游，也想认识新朋友。</text>
          <view class="interest-tags"><text v-for="tag in ['桌游', '城市漫步', '陪拍']" :key="tag">{{ tag }}</text></view>
        </view>
      </view>
      <view class="profile-links">
        <button v-for="item in primaryLinks" :key="item.label" class="profile-link profile-link--primary" @tap="pending(item.label)">
          <UiIcon :name="item.icon" /><text>{{ item.label }}</text><UiIcon class="profile-link__chevron" name="chevron" tone="muted" />
        </button>
        <view class="compact-links">
          <button v-for="item in compactLinks" :key="item.label" class="profile-link" @tap="pending(item.label)">
            <UiIcon :name="item.icon" /><text>{{ item.label }}</text><UiIcon class="profile-link__chevron" name="chevron" tone="muted" />
          </button>
        </view>
        <button v-for="item in utilityLinks" :key="item.label" class="profile-link" @tap="pending(item.label)">
          <UiIcon :name="item.icon" /><text>{{ item.label }}</text><UiIcon class="profile-link__chevron" name="chevron" tone="muted" />
        </button>
      </view>
      <text class="demo-note">示例资料 · 当前为界面预览</text>
    </view>
    <AppTabBar active="profile" owner-route="pages/profile/index" @discover="publish" />
  </view>
</template>

<style scoped lang="scss">
.profile-page { min-height: 100vh; background: $color-white; color: $color-text; padding-bottom: calc(184rpx + env(safe-area-inset-bottom)); }
.profile-body { padding: 0 28rpx; }
.identity-card { position: relative; margin: 12rpx 0 36rpx; border-radius: $radius-v32rpx; box-shadow: 0 8rpx 24rpx $color-elevationSoft; }
.identity-card__back { position: relative; min-height: 278rpx; padding: 52rpx 30rpx 64rpx; border-radius: $radius-v32rpx; background: linear-gradient(145deg, $color-identityLight, $color-primary 72%, $color-identityBottom); }
.identity-card__person { display: flex; align-items: center; gap: 20rpx; position: relative; z-index: 1; max-width: 74%; }
.identity-card__avatar { width: 124rpx; height: 124rpx; flex: 0 0 124rpx; border: 5rpx solid $color-white; border-radius: $radius-v50percent; background: $color-tag; }
.identity-card__fallback { display: flex; align-items: center; justify-content: center; color: $color-primary; font-size: $font-size-v38rpx; }
.identity-card__name { min-width: 0; display: flex; flex-direction: column; align-items: flex-start; gap: 20rpx; color: $color-white; }
.identity-card__name > text { font-size: $font-size-v34rpx; font-weight: $font-weight-v700; line-height: $line-height-v1p3; }
.edit-profile { height: 52rpx; margin: 0; padding: 0 18rpx; display: flex; align-items: center; border: 1rpx solid $color-inverseBorder; border-radius: $radius-v999rpx; color: $color-white; background: transparent; font-size: $font-size-v24rpx; font-weight: $font-weight-v400; line-height: $line-height-v1p3; }
.identity-card__tickets { position: absolute; right: 6rpx; top: 20rpx; width: 244rpx; height: 244rpx; pointer-events: none; }
.identity-card__intro { position: relative; margin-top: -32rpx; padding: 46rpx 34rpx 30rpx; border-radius: $radius-v32rpx; background: $color-white; }
.identity-card__seam { position: absolute; top: -30rpx; left: 0; width: 100%; height: 48rpx; }
.identity-card__bio { position: relative; display: block; font-size: $font-size-v26rpx; line-height: $line-height-v1p6; }
.interest-tags { display: flex; flex-wrap: wrap; gap: 14rpx; margin-top: 22rpx; }
.interest-tags text { padding: 7rpx 20rpx; border-radius: $radius-v999rpx; color: $color-primary; background: $color-tag; font-size: $font-size-v23rpx; line-height: $line-height-v1p4; }
.profile-links { display: flex; flex-direction: column; gap: 22rpx; }
.profile-link { width: 100%; min-height: 126rpx; margin: 0; padding: 22rpx 26rpx; display: flex; align-items: center; gap: 22rpx; border-radius: $radius-v24rpx; color: $color-text; background: $color-card; font-size: $font-size-v29rpx; font-weight: $font-weight-v500; line-height: $line-height-v1p4; text-align: left; }
.profile-link > text { flex: 1; min-width: 0; }
.profile-link :deep(.ui-icon) { width: 54rpx; height: 54rpx; }
.profile-link__chevron { width: 28rpx !important; height: 28rpx !important; }
.compact-links { display: grid; grid-template-columns: repeat(2, minmax(0, 1fr)); gap: 22rpx; }
.compact-links .profile-link { gap: 16rpx; padding: 20rpx; font-size: $font-size-v26rpx; }
.compact-links .profile-link__chevron { display: none; }
.profile-link:active { background: $color-input; }
.profile-link { border: 0; }
.profile-link::after, .edit-profile::after { border: 0; }
.profile-link:focus-visible, .edit-profile:focus-visible { outline: 3rpx solid $color-primary; outline-offset: 3rpx; }
.demo-note { display: block; margin: 28rpx 0 12rpx; color: $color-textMuted; font-size: $font-size-v21rpx; text-align: center; line-height: $line-height-v1p5; }
@media (max-width: 350px) { .identity-card__person { max-width: 78%; gap: 12rpx; } .identity-card__tickets { width: 190rpx; right: 0; } .identity-card__name > text { font-size: $font-size-v30rpx; } }
</style>
