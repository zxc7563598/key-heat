<script setup>
import { Events } from "@wailsio/runtime";
import { ref, onMounted, onBeforeUnmount, nextTick, computed } from 'vue'
import { GetKeyLayout } from "../../bindings/github.com/zxc7563598/key-heat/greetservice";
const layout = ref([])

const keyHeight = ref(52)

const standardKeyRef = ref(null)

const setStandardKeyRef = (el, key) => {
  if (key.W === 1 && el && !standardKeyRef.value) {
    standardKeyRef.value = el
  }
}

const activeKeys = ref(new Set())
const handleBlur = () => {
  activeKeys.value = new Set()
}

const getKeyContentStyle = (key) => {
  return {
    flex: key.W,
    height: Math.floor(keyHeight.value) + "px",
  }
}

const baseKeyStyle = {
  background: "#f5f5f7",
  boxShadow: "3px 3px 6px rgba(0, 0, 0, 0.08), -3px -3px 6px rgba(255, 255, 255, 0.9)",
  textAlign: "center"
}

const keyStyleCache = computed(() => {
  const h = keyHeight.value
  const gap = Math.floor(h / 9)
  const fontSize = Math.floor(h / 4)
  const size = `calc(100% - ${gap * 2}px)`

  return {
    fontSize: `${fontSize}px`,
    margin: `${gap}px`,
    width: size,
    height: size,
  }
})

let observer

const eventKey = ref("")
const eventRaw = ref("")
const eventType = ref("")

onMounted(() => {
  GetKeyLayout("ANSI")
    .then(async (result) => {
      layout.value = result

      // 👇 等 DOM 渲染完成
      await nextTick()

      const el = standardKeyRef.value
      if (!el) return

      observer = new ResizeObserver(([entry]) => {
        const width = entry.contentRect.width
        if (width > 0) {
          keyHeight.value = width
        }
      })

      observer.observe(el)

      window.addEventListener("blur", handleBlur)
    })
    .catch((err) => {
      console.log(err)
    })

  // 键盘事件监听（这个可以独立）
  Events.On("key:pressed", (event) => {
    const { key, type } = event.data
    const next = new Set(activeKeys.value)

    if (type === "down") {
      next.add(key)
    } else if (type === "up") {
      next.delete(key)
    }

    activeKeys.value = next
  })
})

onBeforeUnmount(() => {
  observer?.disconnect()
  window.removeEventListener("blur", handleBlur)
})

</script>

<template>
  <div class="keyboard">
    <div v-for="(row, rowIndex) in layout" :key="rowIndex" class="row">
      <div v-for="key in row" :key="key.Code" class="key-content" :style="getKeyContentStyle(key)"
        :ref="el => setStandardKeyRef(el, key)">
        <div class="key" :class="{ active: activeKeys.has(key.Code) }"
          :style="key.Code === 'None' ? null : [baseKeyStyle, keyStyleCache]">
          {{ key.Label }}
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.keyboard {
  width: 85%;
  padding: 15px;
  border-radius: 24px;
  background: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(20px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.08), inset 0 1px 0 rgba(255, 255, 255, 0.6);
  margin: 50px auto;
}

.row {
  box-sizing: border-box;
  display: flex;
}

.key-content {
  box-sizing: border-box;
  display: flex;
}

.key {
  white-space: pre-wrap;
  box-sizing: border-box;
  display: flex;
  align-items: center;
  cursor: pointer;
  border-radius: 10px;
  justify-content: center;
  transition: all 0.08s ease;
  color: #333;
  font-weight: 500;
}

.key:active,
.key.active {
  transform: translateY(2px) scale(0.96) !important;
  box-shadow: inset 2px 2px 5px rgba(0, 0, 0, 0.15), inset -2px -2px 5px rgba(255, 255, 255, 0.8) !important;
}
</style>
