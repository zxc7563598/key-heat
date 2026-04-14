<script setup>
import { Events } from "@wailsio/runtime";
import { ref, onMounted, onBeforeUnmount, nextTick, computed } from 'vue'
import { GetHeatmap, GetKeyLayout } from "../../bindings/github.com/zxc7563598/key-heat/greetservice";

// 标记基础单位的按键
const keyHeight = ref(52)
const standardKeyRef = ref(null)
const setStandardKeyRef = (el, key) => {
  if (key.W === 1 && el && !standardKeyRef.value) {
    standardKeyRef.value = el
  }
}

// 按键按压状态
const activeKeys = ref(new Set())
const handleBlur = () => {
  activeKeys.value = new Set()
}

// 按键边界范围控制
const getKeyContentStyle = (key) => {
  return {
    flex: key.W,
    height: Math.floor(keyHeight.value) + "px",
  }
}

// 按键动态样式
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

// 切换组件
const tab = ref("base");
const indicatorStyle = ref({});

const setTab = async (value, event) => {
  tab.value = value;
  await nextTick();
  updateIndicator(event.target);
  // 获取热力图数据
  console.log("按钮切换", value)
  if (value == "heat") {
    GetHeatmap("", "").then(async (result) => {
      // 找最大值
      const values = Object.values(result)
      const max = Math.max(...values, 0)
      // 防止全 0
      if (max === 0) return layout.value
      // 映射 + 写回
      layout.value = layout.value.map(row =>
        row.map(key => {
          const raw = result[key.Code] || 0
          // log 压缩
          const logVal = Math.log(raw + 1) / Math.log(max + 1)
          // 再轻微调整分布
          const normalized = Math.pow(logVal, 1)
          // 视觉映射
          const t = Math.pow(normalized, 1.4)
          return {
            ...key,
            Hot: t,
            Count: raw
          }
        })
      )
    }).catch((err) => {
      console.log("热力图获取失败", err)
    })
  } else {
    layout.value = layout.value.map(row =>
      row.map(key => ({
        ...key,
        Hot: Math.pow(0 / 100, 1.4),
        Count: 0
      }))
    )
  }
};

const updateIndicator = (el) => {
  const rect = el.getBoundingClientRect();
  const parentRect = el.parentNode.getBoundingClientRect();
  indicatorStyle.value = {
    width: rect.width + "px",
    transform: `translateX(${rect.left - parentRect.left}px)`,
  };
};

let observer
const layout = ref([])
onMounted(() => {
  // 获取键盘布局
  GetKeyLayout("ANSI").then(async (result) => {
    layout.value = result.map(row =>
      row.map(key => ({
        ...key,
        Hot: Math.pow(0 / 100, 1.4),
        Count: 0
      }))
    )
    // 等 DOM 渲染完成
    await nextTick()
    // 计算按键高度
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
  }).catch((err) => {
    console.log(err)
  })
  // 键盘事件监听
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
  // 按钮初始化
  const active = document.querySelector(".button.active");
  if (active) {
    updateIndicator(active);
  }
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
        <div class="key" :class="{ active: activeKeys.has(key.Code), hot: key.Hot > 0.6 }"
          :style="key.Code === 'None' ? null : [keyStyleCache, { '--t': key.Hot }]">
          {{ key.Label }}
        </div>
      </div>
    </div>
  </div>
  <div class="functional">
    <div class="group">
      <div class="indicator" :style="indicatorStyle"></div>
      <div class="button" :class="{ active: tab === 'base' }" @click="setTab('base', $event)">
        基础
      </div>
      <div class="button" :class="{ active: tab === 'heat' }" @click="setTab('heat', $event)">
        热力图
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
  --t: 0;
  text-align: center;
  white-space: pre-wrap;
  box-sizing: border-box;
  display: flex;
  align-items: center;
  cursor: pointer;
  border-radius: 10px;
  justify-content: center;
  font-weight: 500;
  background: hsl(0,
      calc(8% + var(--t) * 72%),
      calc(97% - var(--t) * 47%));
  color: #333;
  box-shadow:
    3px 3px 6px rgba(0, 0, 0, calc(0.08 + var(--t) * 0.12)),
    -3px -3px 6px rgba(255, 255, 255, calc(0.9 - var(--t) * 0.7));
  transition:
    transform 0.08s ease,
    box-shadow 0.12s ease,
    background 0.2s ease,
    color 0.2s ease;
}

.key:active,
.key.active {
  transform: translateY(2px) scale(0.96) !important;
  box-shadow:
    inset 2px 2px 5px rgba(0, 0, 0, calc(0.15 + var(--t) * 0.2)),
    inset -2px -2px 5px rgba(255, 255, 255, calc(0.8 - var(--t) * 0.6)) !important;
}

.key.hot {
  color: #fff;
}

/* 居中 */
.functional {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 10px;
}

/* 胶囊容器 */
.group {
  position: relative;
  display: flex;
  background: #ffffff;
  border-radius: 999px;
  padding: 4px;
  box-shadow:
    0 6px 16px rgba(0, 0, 0, 0.06),
    inset 0 1px 0 rgba(255, 255, 255, 0.8);
}

/* 滑块（核心） */
.indicator {
  position: absolute;
  top: 4px;
  left: 4px;
  height: calc(100% - 8px);
  border-radius: 999px;
  background: #f0f1f3;
  box-shadow:
    0 2px 6px rgba(0, 0, 0, 0.08),
    inset 0 1px 2px rgba(255, 255, 255, 0.9);

  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
  z-index: 0;
}

/* 按钮 */
.button {
  position: relative;
  z-index: 1;
  padding: 8px 20px;
  font-size: 14px;
  color: #666;
  border-radius: 999px;
  cursor: pointer;
  user-select: none;
  transition: all 0.25s ease;
  white-space: nowrap;
}

/* hover */
.button:hover {
  color: #222;
}

/* 激活文字 */
.button.active {
  color: #111;
}

/* 点击反馈 */
.button:active {
  transform: scale(0.96);
}
</style>
