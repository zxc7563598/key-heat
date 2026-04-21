<script setup>
import { Events } from "@wailsio/runtime";
import { ref, onMounted, onBeforeUnmount, nextTick, computed, watch, watchEffect } from 'vue'
import { GetHeatmap, GetKeyLayout, GetPermission, SetMonitorClose, SetMonitorStart, GetMonitorStatus, GetBootStartup, OpenBootStartup, CloseBootStartup, GetPrompt } from "../../bindings/github.com/zxc7563598/key-heat/greetservice";
import KeyTooltip from "./KeyTooltip.vue";
import PromptDialog from "./PromptDialog.vue"

// 当前监听状态
const monitor = ref(false)
const permission = ref(false)
const bootStartup = ref(false)

// 启动监听
const setMonitorStart = () => {
  console.log("启动监听", "启动")
  SetMonitorStart().then(() => {
    GetMonitorStatus().then(async (result) => {
      monitor.value = result
    }).catch((err) => {
      console.log(err)
    })
  }).catch((err) => {
    console.log(err)
  })
}

// 关闭监听
const setMonitorClose = () => {
  console.log("关闭监听", "启动")
  SetMonitorClose().then(() => {
    GetMonitorStatus().then(async (result) => {
      monitor.value = result
    }).catch((err) => {
      console.log(err)
    })
  }).catch((err) => {
    console.log(err)
  })
}

// 开启自启动
const openBootStartup = () => {
  OpenBootStartup().then(() => {
    GetBootStartup().then(async (result) => {
      bootStartup.value = result
    }).catch((err) => {
      console.log(err)
    })
  }).catch((err) => {
    console.log(err)
  })
}

// 关闭自启动
const closeBootStartup = () => {
  CloseBootStartup().then(() => {
    GetBootStartup().then(async (result) => {
      bootStartup.value = result
    }).catch((err) => {
      console.log(err)
    })
  }).catch((err) => {
    console.log(err)
  })
}

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
  if (value == "heat") {
    applyRange()
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

// 提交热力图筛选日期
const startDate = ref("");
const endDate = ref("");
const formatDate = (date) => {
  if (!date) return "";
  return date;
};
const applyRange = () => {
  const start = formatDate(startDate.value) ?? "";
  const end = formatDate(endDate.value) ?? "";
  GetHeatmap(start, end).then(async (result) => {
    // 找最大值
    const values = Object.values(result)
    const max = Math.max(...values, 0)
    // 映射 + 写回
    layout.value = layout.value.map(row =>
      row.map(key => {
        const raw = result[key.Code] || 0
        const t = ref(0)
        if (raw != 0) {
          const logVal = Math.log(raw + 1) / Math.log(max + 1) // log 压缩
          const normalized = Math.pow(logVal, 1) // 再轻微调整分布
          t.value = Math.pow(normalized, 1.4) // 视觉映射
        }
        return {
          ...key,
          Hot: t.value,
          Count: raw
        }
      })
    )
  }).catch((err) => {
    console.log("热力图获取失败", err)
  })
};
watch([startDate, endDate], ([s, e]) => {
  applyRange();
});

// 鼠标移入按键
const hoverKey = ref(null)
const mouseX = ref(0)
const mouseY = ref(0)
let moveRAF = null
const updateMouse = (e) => {
  if (moveRAF) cancelAnimationFrame(moveRAF)
  moveRAF = requestAnimationFrame(() => {
    mouseX.value = e.clientX
    mouseY.value = e.clientY
  })
}
const onKeyEnter = (key, e) => {
  if (tab.value == "heat") {
    hoverKey.value = key
    updateMouse(e)
  }
}
const onKeyMove = (e) => {
  if (tab.value == "heat") {
    updateMouse(e)
  }
}
const onKeyLeave = () => {
  hoverKey.value = null
}

// 导出数据
const showDialog = ref(false)
const promptText = ref("")

const getPrompt = () => {
  const start = formatDate(startDate.value) ?? ""
  const end = formatDate(endDate.value) ?? ""
  GetPrompt(start, end).then((result) => {
    promptText.value = result
    showDialog.value = true
  }).catch((err) => {
    console.log("提示词获取失败", err)
  })
}

let observer
const layout = ref([])
onMounted(async () => {
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
  Events.On("active:switch", (event) => {
    monitor.value = event.data
  })
  // 按钮初始化
  const active = document.querySelector(".button.active");
  if (active) {
    updateIndicator(active);
  }
  // 检查权限
  await checkPermission()
  document.addEventListener("visibilitychange", () => {
    if (document.visibilityState === "visible") {
      checkPermission()
    }
  })
  window.addEventListener("focus", checkPermission)
  if (!permission.value) {
    startPermissionWatch()
  }
})

onBeforeUnmount(() => {
  observer?.disconnect()
  window.removeEventListener("blur", handleBlur)
})

const checkPermission = () => {
  // 确认监听权限
  GetPermission().then(async (result) => {
    permission.value = result
  }).catch((err) => {
    console.log(err)
  })
  // 确认监听状态
  GetMonitorStatus().then(async (result) => {
    monitor.value = result
  }).catch((err) => {
    console.log(err)
  })
  // 确认开机自启动
  GetBootStartup().then(async (result) => {
    bootStartup.value = result
  }).catch((err) => {
    console.log(err)
  })
}

let timer = null
const startPermissionWatch = () => {
  if (timer) return
  timer = setInterval(async () => {
    checkPermission()
    if (permission.value) {
      clearInterval(timer)
      timer = null
    }
  }, 1500)
}

</script>

<template>
  <div class="status-bar">
    <div class="status-left">
      <div class="action-btn primary" v-if="!monitor" @click="setMonitorStart">
        启动监听
      </div>
      <div class="action-btn danger" v-if="monitor" @click="setMonitorClose">
        停止监听
      </div>
      <div class="action-btn primary" v-if="!bootStartup" @click="openBootStartup">
        开机自启动
      </div>
      <div class="action-btn danger" v-if="bootStartup" @click="closeBootStartup">
        关闭开机自启动
      </div>
    </div>
    <div class="status-right">
      <div class="status-item">
        <div class="dot-wrapper">
          <span class="dot" :class="{ ok: bootStartup, bad: !bootStartup }"></span>
          <div class="tooltip-status">
            {{ permission ? "已开启开机自启动" : "未开启开机自启动" }}
          </div>
        </div>
        <span class="label">自启动</span>
      </div>
      <div class="status-item">
        <div class="dot-wrapper">
          <span class="dot" :class="{ ok: permission, bad: !permission }"></span>
          <div class="tooltip-status">
            {{ permission ? "已授予辅助功能权限" : "未授予辅助功能权限" }}
          </div>
        </div>
        <span class="label">权限</span>
      </div>
      <div class="status-item">
        <div class="dot-wrapper">
          <span class="dot monitor" :class="{ ok: monitor, bad: !monitor }"></span>
          <div class="tooltip-status">
            {{ monitor ? "正在监听键盘输入" : "监听已停止" }}
          </div>
        </div>
        <span class="label">监听</span>
      </div>
    </div>
  </div>
  <div class="keyboard">
    <div v-for="(row, rowIndex) in layout" :key="rowIndex" class="row">
      <div v-for="key in row" :key="key.Code" class="key-content" :style="getKeyContentStyle(key)"
        :ref="el => setStandardKeyRef(el, key)">
        <div class="key" :class="{ active: activeKeys.has(key.Code), hot: key.Hot > 0.6 }"
          :style="key.Code === 'None' ? null : [keyStyleCache, { '--t': key.Hot }]"
          @mouseenter="onKeyEnter(key, $event)" @mousemove="onKeyMove" @mouseleave="onKeyLeave">
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
    <div class="date-range" v-if="tab === 'heat'">
      <input type="date" v-model="startDate" class="date-input" />
      <span class="separator">—</span>
      <input type="date" v-model="endDate" class="date-input" />
      <div class="apply-btn" @click="startDate = endDate = ''">
        清空
      </div>
      <div class="apply-btn" @click="getPrompt">
        生成报告
      </div>
    </div>
  </div>
  <KeyTooltip :data="hoverKey" :x="mouseX" :y="mouseY" :start="formatDate(startDate)" :end="formatDate(endDate)" />
  <PromptDialog v-model="showDialog" :prompt="promptText" />
</template>

<style scoped>
.keyboard {
  width: 85%;
  padding: 15px;
  border-radius: 24px;
  background: rgba(255, 255, 255, 0.7);
  backdrop-filter: blur(20px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.08), inset 0 1px 0 rgba(255, 255, 255, 0.6);
  margin: 10px auto 30px auto;
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
  background: hsl(0, calc(8% + var(--t) * 72%), calc(97% - var(--t) * 47%));
  color: #333;
  box-shadow: 3px 3px 6px rgba(0, 0, 0, calc(0.08 + var(--t) * 0.12)), -3px -3px 6px rgba(255, 255, 255, calc(0.9 - var(--t) * 0.7));
  transition: transform 0.08s ease, box-shadow 0.12s ease, background 0.2s ease, color 0.2s ease;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

.key:active,
.key.active {
  transform: translateY(2px) scale(0.96) !important;
  box-shadow: inset 2px 2px 5px rgba(0, 0, 0, calc(0.15 + var(--t) * 0.2)), inset -2px -2px 5px rgba(255, 255, 255, calc(0.8 - var(--t) * 0.6)) !important;
}

.key.hot {
  color: #fff;
}

.functional {
  display: flex;
  width: 85%;
}

.group {
  position: relative;
  display: flex;
  background: #ffffff;
  border-radius: 999px;
  padding: 4px;
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.06), inset 0 1px 0 rgba(255, 255, 255, 0.8);
  left: -15px;
}

.indicator {
  position: absolute;
  top: 4px;
  left: 0px;
  height: calc(100% - 8px);
  border-radius: 999px;
  background: linear-gradient(to bottom, #f7f8fa, #eef0f3);
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.08), 0 2px 4px rgba(0, 0, 0, 0.06), inset 0 1px 2px rgba(255, 255, 255, 0.9);
  transform: translateY(-0.5px);
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
  z-index: 0;
}

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

.button:hover {
  color: #222;
  transform: translateY(-0.5px);
}

.button.active {
  color: #111;
}

.button:active {
  transform: scale(0.96);
}

.date-range {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-left: auto;
  position: relative;
  right: -15px;
}

.date-input {
  appearance: none;
  -webkit-appearance: none;
  padding: 6px 12px;
  border-radius: 999px;
  border: none;
  background: #f7f8fa;
  color: #333;
  font-size: 13px;
  box-shadow: inset 0 1px 2px rgba(0, 0, 0, 0.06), 0 1px 1px rgba(255, 255, 255, 0.8);
  transition: all 0.2s ease;
}

.date-input:hover {
  background: #f0f1f3;
}

.date-input:focus {
  outline: none;
  background: #eef0f3;
}

.separator {
  color: #999;
  font-size: 12px;
}

.apply-btn {
  padding: 6px 14px;
  border-radius: 999px;
  background: #ffffff;
  cursor: pointer;
  font-size: 13px;
  color: #444;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.08), inset 0 1px 0 rgba(255, 255, 255, 0.8);
  transition: all 0.2s ease;
}

.apply-btn:hover {
  color: #111;
  transform: translateY(-0.5px);
}

.apply-btn:active {
  transform: scale(0.96);
}

.status-bar {
  width: 85%;
  display: flex;
  align-items: center;
  justify-content: space-between;
}


.status-left {
  display: flex;
  align-items: center;
  gap: 18px;
  position: relative;
  left: -15px;
}

.status-item {
  display: flex;
  align-items: center;
  gap: 6px;
  position: relative;
}

.label {
  font-size: 13px;
  color: #777;
  letter-spacing: 0.2px;
}

.dot-wrapper {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
}

.dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #c7c7cc;
  transition: all 0.25s ease;
}

.dot.ok {
  background: #34c759;
  box-shadow: 0 0 6px rgba(52, 199, 89, 0.6);
}

.dot.bad {
  background: #ff3b30;
  box-shadow: 0 0 6px rgba(255, 59, 48, 0.5);
}

.dot.monitor.ok {
  animation: pulse 1.8s ease-in-out infinite;
}

@keyframes pulse {
  0% {
    box-shadow: 0 0 4px rgba(52, 199, 89, 0.5);
  }

  50% {
    box-shadow: 0 0 10px rgba(52, 199, 89, 0.9);
  }

  100% {
    box-shadow: 0 0 4px rgba(52, 199, 89, 0.5);
  }
}

.tooltip-status {
  position: absolute;
  bottom: 140%;
  left: 50%;
  transform: translateX(-50%) scale(0.96);
  transform-origin: bottom center;
  padding: 6px 10px;
  border-radius: 8px;
  font-size: 12px;
  color: #fff;
  white-space: nowrap;
  background: rgba(60, 60, 67, 0.9);
  backdrop-filter: blur(8px);
  opacity: 0;
  pointer-events: none;
  transition: opacity 0.18s ease, transform 0.18s cubic-bezier(0.4, 0, 0.2, 1);
}

.tooltip-status::after {
  content: "";
  position: absolute;
  top: 100%;
  left: 50%;
  transform: translateX(-50%);
  border: 5px solid transparent;
  border-top-color: rgba(60, 60, 67, 0.9);
}

.dot-wrapper:hover .tooltip-status {
  opacity: 1;
  transform: translateX(-50%) scale(1);
}

.status-right {
  display: flex;
  gap: 10px;
  position: relative;
  right: -15px;
}

.action-btn {
  padding: 6px 14px;
  border-radius: 999px;
  font-size: 13px;
  cursor: pointer;
  background: #ffffff;
  color: #444;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.08), inset 0 1px 0 rgba(255, 255, 255, 0.85);
  transition: all 0.2s ease;
}

.action-btn:hover {
  transform: translateY(-0.5px);
  color: #111;
}

.action-btn:active {
  transform: scale(0.96);
}

.action-btn.primary {
  color: #1a7f37;
}

.action-btn.danger {
  color: #c5221f;
}
</style>
