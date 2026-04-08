// internal/monitor/counter.go
package monitor

import (
	"sync"
	"time"
)

// KeyCounter 内存计数器（按天和按键名统计）
type KeyCounter struct {
	mu          sync.RWMutex
	currentDate string         // 当前统计日期 YYYY-MM-DD
	counts      map[string]int // 按键名 -> 当前批次的累加次数
	totalCount  int            // 总按键次数（当前批次）
}

// NewKeyCounter 创建内存计数器
func NewKeyCounter() *KeyCounter {
	return &KeyCounter{
		currentDate: today(),
		counts:      make(map[string]int),
		totalCount:  0,
	}
}

// Add 添加按键计数（线程安全 + 跨天自动切换）
func (c *KeyCounter) Add(keyName string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	t := today()
	// 跨天：立即切换新的一天
	if c.currentDate != t {
		c.currentDate = t
		c.counts = make(map[string]int)
		c.totalCount = 0
	}
	c.counts[keyName]++
	c.totalCount++
}

// GetAndReset 获取当前所有计数并重置（用于批量写入）
// 返回: (date string, counts map[string]int, total int)
func (c *KeyCounter) GetAndReset() (string, map[string]int, int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// 先取“旧数据”
	date := c.currentDate
	countsCopy := make(map[string]int, len(c.counts))
	for k, v := range c.counts {
		countsCopy[k] = v
	}
	total := c.totalCount
	// reset，并更新为最新日期
	c.currentDate = today()
	c.counts = make(map[string]int)
	c.totalCount = 0
	return date, countsCopy, total
}

// IsEmpty 检查当前是否有数据
func (c *KeyCounter) IsEmpty() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.totalCount == 0
}

// ShouldFlush 检查是否需要刷新（根据总次数是否达到阈值）
func (c *KeyCounter) ShouldFlush(threshold int) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.totalCount >= threshold
}

// GetCurrentStats 获取当前统计（用于实时展示，不清空数据）
func (c *KeyCounter) GetCurrentStats() (string, map[string]int, int) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	date := c.currentDate
	countsCopy := make(map[string]int, len(c.counts))
	for k, v := range c.counts {
		countsCopy[k] = v
	}
	return date, countsCopy, c.totalCount
}

// today 获取当前日期字符串（统一格式）
func today() string {
	return time.Now().Format("2006-01-02")
}
