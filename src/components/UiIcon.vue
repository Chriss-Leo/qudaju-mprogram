<script setup lang="ts">
withDefaults(defineProps<{ name: string; tone?: "blue" | "muted" | "black" | "white" }>(), { tone: "blue" });
</script>

<template><image class="ui-icon" :src="`/static/ui-icons/${name}-${tone}.svg`" mode="aspectFit" aria-hidden="true" /></template>

<style scoped lang="scss">
.ui-icon { display: block; width: 44rpx; height: 44rpx; flex: 0 0 auto; }
</style>
