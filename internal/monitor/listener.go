package monitor

import (
	"log"
	"sync"

	hook "github.com/robotn/gohook"
	"github.com/zxc7563598/key-heat/pkg/keymap"
)

type KeyEventListener struct {
	keyChan  chan<- string
	stopChan chan struct{}
	mapper   keymap.KeyMapper

	once sync.Once
}

func NewKeyEventListener(keyChan chan<- string) *KeyEventListener {
	return &KeyEventListener{
		keyChan:  keyChan,
		stopChan: make(chan struct{}),
		mapper:   keymap.GetGlobalMapper(),
	}
}

// 开始监听
func (l *KeyEventListener) Start() error {
	log.Println("键盘监听启动中...")
	evChan := hook.Start()
	log.Println("键盘监听已启动")
	for {
		select {
		case <-l.stopChan:
			log.Println("键盘监听收到停止信号")
			hook.End() // 主动结束 hook
			return nil
		case ev, ok := <-evChan:
			if !ok {
				log.Println("键盘事件通道关闭")
				return nil
			}
			if ev.Kind == hook.KeyDown {
				keyName := l.mapper.Normalize(ev.Rawcode)
				log.Printf("按键: %s", keyName)
				select {
				case l.keyChan <- keyName:
				default:
					log.Printf("通道满，丢弃: %s", keyName)
				}
			}
		}
	}
}

// 停止监听
func (l *KeyEventListener) Stop() {
	l.once.Do(func() {
		close(l.stopChan)
	})
}
