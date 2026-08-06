<script setup lang="ts">
defineProps<{ type: string; tone?: string }>();
</script>

<template>
  <view class="glyph" :class="[`glyph--${type}`, `glyph--${tone || 'mint'}`]" aria-hidden="true">
    <view class="shape shape-a" />
    <view class="shape shape-b" />
    <view class="shape shape-c" />
  </view>
</template>

<style scoped lang="scss">
.glyph {
  position: relative;
  width: 84rpx;
  height: 84rpx;
  border-radius: 26rpx;
  color: #173f34;
  background: #e2eee7;
  overflow: hidden;
}
.glyph--peach { background: #f9e3d8; color: #c7553f; }
.glyph--blue { background: #dfecef; color: #376872; }
.glyph--lilac { background: #ebe4f4; color: #64517c; }
.glyph--gold { background: #f7ebcd; color: #a26c27; }
.shape { position: absolute; border: 5rpx solid currentColor; }
.glyph--game .shape-a { width: 48rpx; height: 30rpx; left: 18rpx; top: 27rpx; border-radius: 13rpx; }
.glyph--game .shape-b { width: 17rpx; height: 5rpx; left: 27rpx; top: 41rpx; border: 0; background: currentColor; }
.glyph--game .shape-c { width: 5rpx; height: 17rpx; left: 33rpx; top: 35rpx; border: 0; background: currentColor; }
.glyph--map .shape-a { width: 36rpx; height: 46rpx; left: 24rpx; top: 18rpx; border-radius: 50% 50% 55% 55%; transform: rotate(45deg); }
.glyph--map .shape-b { width: 8rpx; height: 8rpx; left: 38rpx; top: 31rpx; border-radius: 50%; background: currentColor; }
.glyph--map .shape-c { display: none; }
.glyph--camera .shape-a { width: 50rpx; height: 35rpx; left: 17rpx; top: 29rpx; border-radius: 10rpx; }
.glyph--camera .shape-b { width: 16rpx; height: 16rpx; left: 34rpx; top: 38rpx; border-radius: 50%; }
.glyph--camera .shape-c { width: 20rpx; height: 8rpx; left: 29rpx; top: 21rpx; border-radius: 6rpx; }
.glyph--bulb .shape-a { width: 34rpx; height: 34rpx; left: 25rpx; top: 17rpx; border-radius: 50%; }
.glyph--bulb .shape-b { width: 24rpx; height: 12rpx; left: 30rpx; top: 50rpx; border-radius: 3rpx 3rpx 10rpx 10rpx; }
.glyph--bulb .shape-c { display: none; }
.glyph--dice .shape-a { width: 44rpx; height: 44rpx; left: 20rpx; top: 20rpx; border-radius: 12rpx; transform: rotate(-8deg); }
.glyph--dice .shape-b, .glyph--dice .shape-c { width: 6rpx; height: 6rpx; border-radius: 50%; background: currentColor; border: 0; }
.glyph--dice .shape-b { left: 32rpx; top: 33rpx; box-shadow: 17rpx 17rpx 0 currentColor; }
.glyph--dice .shape-c { left: 49rpx; top: 33rpx; box-shadow: -17rpx 17rpx 0 currentColor; }
.glyph--mountain .shape-a { width: 39rpx; height: 39rpx; left: 16rpx; top: 30rpx; border-width: 5rpx 0 0 5rpx; transform: rotate(45deg); }
.glyph--mountain .shape-b { width: 27rpx; height: 27rpx; left: 42rpx; top: 40rpx; border-width: 5rpx 0 0 5rpx; transform: rotate(45deg); }
.glyph--mountain .shape-c { display: none; }
.glyph--music .shape-a { width: 5rpx; height: 42rpx; left: 48rpx; top: 18rpx; border: 0; background: currentColor; transform: rotate(-9deg); }
.glyph--music .shape-b, .glyph--music .shape-c { width: 19rpx; height: 14rpx; border-radius: 50%; background: currentColor; border: 0; }
.glyph--music .shape-b { left: 27rpx; top: 52rpx; }
.glyph--music .shape-c { left: 45rpx; top: 47rpx; }
.glyph--grid .shape-a { width: 16rpx; height: 16rpx; left: 22rpx; top: 22rpx; border-radius: 6rpx; box-shadow: 26rpx 0 0 -5rpx transparent, 0 26rpx 0 -5rpx transparent, 26rpx 26rpx 0 -5rpx transparent; }
.glyph--grid .shape-b { width: 16rpx; height: 16rpx; right: 22rpx; top: 22rpx; border-radius: 6rpx; }
.glyph--grid .shape-c { width: 16rpx; height: 16rpx; left: 22rpx; bottom: 22rpx; border-radius: 6rpx; box-shadow: 26rpx 0 0 -5rpx transparent; }
</style>
