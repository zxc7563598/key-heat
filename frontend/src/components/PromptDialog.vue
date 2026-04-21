<template>
    <div>
        <div v-if="visible" class="overlay" @click.self="close">
            <div class="card">
                <div class="header">
                    <h3>生成你的键盘使用报告</h3>
                    <span class="close" @click="close">×</span>
                </div>
                <p class="desc">
                    我们已经帮你生成了一段分析提示词<br />
                    复制后粘贴到 ChatGPT / DeepSeek 即可生成报告
                </p>
                <div class="prompt-box">
                    <pre v-if="expanded">{{ prompt }}</pre>
                    <pre v-else>{{ shortPrompt }}</pre>
                </div>
                <div class="toggle" @click="expanded = !expanded">
                    {{ expanded ? "收起" : "展开完整提示词" }}
                </div>
                <div class="actions">
                    <button class="primary" @click="copyPrompt">
                        复制提示词
                    </button>
                </div>
                <div class="hint">
                    支持 ChatGPT / DeepSeek / Gemini 等 AI 工具
                </div>
            </div>
        </div>
    </div>
    <div v-if="toast.show" class="toast">
        {{ toast.text }}
    </div>
</template>

<script setup>
import { ref, computed } from "vue"

const props = defineProps({
    modelValue: Boolean,
    prompt: String
})

const emit = defineEmits(["update:modelValue"])

const visible = computed({
    get: () => props.modelValue,
    set: (val) => emit("update:modelValue", val)
})

const expanded = ref(false)

// 截断版本（前500字符）
const shortPrompt = computed(() => {
    if (!props.prompt) return ""
    return props.prompt.slice(0, 500) + "\n...\n（点击展开查看完整内容）"
})

// 关闭
const close = () => {
    visible.value = false
}

// 复制
const copyPrompt = async () => {
    try {
        await navigator.clipboard.writeText(props.prompt)
        showToast("已复制，快去生成你的报告吧")
    } catch (err) {
        // fallback
        const textarea = document.createElement("textarea")
        textarea.value = props.prompt
        document.body.appendChild(textarea)
        textarea.select()
        document.execCommand("copy")
        document.body.removeChild(textarea)
        showToast("已复制，快去生成你的报告吧")
    }
}

const toast = ref({
    show: false,
    text: ""
})

const showToast = (text) => {
    toast.value.text = text
    toast.value.show = true
    setTimeout(() => {
        toast.value.show = false
    }, 2000)
}
</script>

<style scoped>
.overlay {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.4);
    display: flex;
    align-items: center;
    justify-content: center;
    backdrop-filter: blur(6px);
}

.card {
    width: 600px;
    max-height: 80vh;
    background: #fff;
    border-radius: 16px;
    padding: 20px;
    box-shadow: 0 10px 40px rgba(0, 0, 0, 0.2);
    display: flex;
    flex-direction: column;
}

.header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.close {
    cursor: pointer;
    font-size: 20px;
}

.desc {
    margin: 10px 0;
    color: #666;
    font-size: 14px;
}

.prompt-box {
    background: #f6f6f6;
    border-radius: 10px;
    padding: 10px;
    overflow: auto;
    max-height: 300px;
    font-size: 12px;
}

.toggle {
    margin-top: 6px;
    color: #409eff;
    cursor: pointer;
    font-size: 13px;
}

.actions {
    display: flex;
    gap: 10px;
    margin-top: 15px;
}

button {
    flex: 1;
    padding: 10px;
    border-radius: 10px;
    cursor: pointer;
    border: none;
}

.primary {
    background: #409eff;
    color: white;
}

.secondary {
    background: #eee;
}

.hint {
    margin-top: 10px;
    font-size: 12px;
    color: #999;
    text-align: center;
}

.toast {
    position: fixed;
    bottom: 40px;
    left: 50%;
    transform: translateX(-50%);
    background: rgba(0, 0, 0, 0.8);
    color: #fff;
    padding: 10px 16px;
    border-radius: 20px;
    font-size: 13px;
    z-index: 9999;
    animation: fadeInOut 2s ease;
}

@keyframes fadeInOut {
    0% {
        opacity: 0;
        transform: translate(-50%, 10px);
    }

    10% {
        opacity: 1;
        transform: translate(-50%, 0);
    }

    90% {
        opacity: 1;
    }

    100% {
        opacity: 0;
        transform: translate(-50%, 10px);
    }
}
</style>