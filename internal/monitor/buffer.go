package monitor

import (
	"log"
	"sync"
	"time"

	"github.com/zxc7563598/key-heat/internal/storage"
)

type Config struct {
	FlushInterval  time.Duration
	FlushThreshold int
	KeyChanSize    int
}

func DefaultConfig() *Config {
	return &Config{
		FlushInterval:  10 * time.Second,
		FlushThreshold: 500,
		KeyChanSize:    1000,
	}
}

type Monitor struct {
	config    *Config
	repo      storage.Repository
	listener  *KeyEventListener
	counter   *KeyCounter
	keyChan   chan string
	flushChan chan struct{}
	stopChan  chan struct{}
	wg        sync.WaitGroup
	once      sync.Once
}

func NewMonitor(repo storage.Repository, cfg *Config) *Monitor {
	if cfg == nil {
		cfg = DefaultConfig()
	}
	m := &Monitor{
		config:    cfg,
		repo:      repo,
		counter:   NewKeyCounter(),
		keyChan:   make(chan string, cfg.KeyChanSize),
		flushChan: make(chan struct{}, 1),
		stopChan:  make(chan struct{}),
	}
	m.listener = NewKeyEventListener(m.keyChan)
	return m
}

// 启动监听
func (m *Monitor) Start() {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		if err := m.listener.Start(); err != nil {
			log.Printf("键盘监听启动失败: %v", err)
		}
	}()
	m.wg.Add(1)
	go m.keyProcessor()
	// flush worker（统一写库）
	m.wg.Add(1)
	go m.flushWorker()
	log.Println("键盘监控器已启动")
}

// 按键处理器
func (m *Monitor) keyProcessor() {
	defer m.wg.Done()
	for {
		select {
		case keyName, ok := <-m.keyChan:
			if !ok {
				return
			}
			// 将按键加入队列
			m.counter.Add(keyName)
			// 判断队列是否达到阈值
			if m.counter.ShouldFlush(m.config.FlushThreshold) {
				m.triggerFlush()
			}
		case <-m.stopChan:
			log.Println("keyProcessor 停止")
			return
		}
	}
}

// 处理监听信息写入
func (m *Monitor) flushWorker() {
	defer m.wg.Done()
	ticker := time.NewTicker(m.config.FlushInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C: // 定时
			m.flush()
		case <-m.flushChan: // 阈值
			m.flush()
		case <-m.stopChan: // 关闭
			log.Println("flushWorker 收到停止信号，执行最后 flush")
			m.flush()
			return
		}
	}
}

// 达到阈值触发存储
func (m *Monitor) triggerFlush() {
	select {
	case m.flushChan <- struct{}{}:
	default:
		// 已经有 flush 在排队了，忽略
	}
}

// 监听数据写入
func (m *Monitor) flush() {
	date, counts, total := m.counter.GetAndReset()
	if total == 0 {
		return
	}
	for keyName, delta := range counts {
		if err := m.repo.UpsertKeyCount(date, keyName, delta); err != nil {
			log.Printf("写入失败 key=%s count=%d err=%v", keyName, delta, err)
		}
	}
}

// 停止监听
func (m *Monitor) Stop() {
	close(m.keyChan)
	m.once.Do(func() {
		log.Println("停止监控器...")
		// 停 listener（不再写 keyChan）
		if m.listener != nil {
			m.listener.Stop()
		}
		// 发停止信号
		close(m.stopChan)
		// 等所有 worker 结束
		m.wg.Wait()
		log.Println("监控器已完全停止")
	})
}
