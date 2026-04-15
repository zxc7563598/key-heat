<template>
    <Transition name="tooltip-fade">
        <div v-if="visible && data" class="tooltip" :style="style">
            <div class="title">按键: {{ data.Code }}</div>
            <div class="line">筛选日期: {{ start != "" ? start : "最初" }} - {{ end != "" ? end : "至今" }}</div>
            <div class="line">点击次数: {{ data.Count }}</div>
        </div>
    </Transition>
</template>

<script setup>
import { computed, ref, watch } from 'vue'

const props = defineProps({
    data: Object,
    x: Number,
    y: Number,
    start: String,
    end: String
})

// 控制显示（用于动画 & 延迟消失）
const visible = ref(false)

// 延迟出现 / 延迟消失
let showTimer = null
let hideTimer = null

watch(() => props.data, (val) => {
    clearTimeout(showTimer)
    clearTimeout(hideTimer)
    if (val) {
        // 延迟出现
        showTimer = setTimeout(() => {
            visible.value = true
        }, 80)
    } else {
        // 延迟消失（更丝滑）
        hideTimer = setTimeout(() => {
            visible.value = false
        }, 120)
    }
})

// 跟随鼠标 + 边界处理
const style = computed(() => {
    let x = props.x + 14
    let y = props.y + 14
    const width = 150
    const height = 80
    if (x + width > window.innerWidth) {
        x -= width + 28
    }
    if (y + height > window.innerHeight) {
        y -= height + 28
    }
    return {
        position: 'fixed',
        left: x + 'px',
        top: y + 'px',
        pointerEvents: 'none',
        zIndex: 9999
    }
})

const formatHot = (v) => {
    if (v == null) return '-'
    return (v * 100).toFixed(1) + '%'
}
</script>

<style scoped>
.tooltip {
    background: rgba(28, 28, 30, 0.85);
    color: #fff;
    padding: 10px 12px;
    border-radius: 12px;
    font-size: 12px;
    backdrop-filter: blur(14px);
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.35);
    min-width: 130px;

    transform-origin: top left;
}

/* ✨ macOS 风格动画 */
.tooltip-fade-enter-active,
.tooltip-fade-leave-active {
    transition:
        opacity 0.14s ease,
        transform 0.14s cubic-bezier(0.2, 0.8, 0.2, 1);
}

.tooltip-fade-enter-from {
    opacity: 0;
    transform: scale(0.94) translateY(4px);
}

.tooltip-fade-enter-to {
    opacity: 1;
    transform: scale(1) translateY(0);
}

.tooltip-fade-leave-from {
    opacity: 1;
    transform: scale(1) translateY(0);
}

.tooltip-fade-leave-to {
    opacity: 0;
    transform: scale(0.96) translateY(2px);
}

.title {
    font-weight: 600;
    margin-bottom: 4px;
}

.line {
    opacity: 0.75;
}
</style>