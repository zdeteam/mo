<template>
    <div class="skeleton-item">
    </div>
</template>

<script setup lang="ts">
const props = withDefaults(defineProps<{
    num: number,
}>(), {
    num: 1
});
</script>

<style lang="less" scoped>
.skeleton-item {
    padding: 12px;
    display: flex;

    .user {
        width: 42px;
    }

    .content {
        width: calc(100% - 42px);
    }
}
</style>