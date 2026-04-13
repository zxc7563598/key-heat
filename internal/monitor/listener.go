package monitor

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	hook "github.com/robotn/gohook"
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/zxc7563598/key-heat/pkg/keymap"
)

type KeyEventListener struct {
	keyChan  chan<- string
	stopChan chan struct{}
	mapper   keymap.KeyMapper
	once     sync.Once
	app      *application.App
}

func NewKeyEventListener(keyChan chan<- string, app *application.App) *KeyEventListener {
	return &KeyEventListener{
		keyChan:  keyChan,
		stopChan: make(chan struct{}),
		mapper:   keymap.GetGlobalMapper(),
		app:      app,
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
				l.app.Event.Emit("key:pressed", map[string]any{
					"key":  keyName,
					"raw":  ev.Rawcode,
					"type": "down",
				})

				// 测试写入
				configDir, _ := os.UserConfigDir()
				// 跨平台应用数据目录
				// Windows: %APPDATA%/KeyHeat/
				// macOS:   ~/Library/Application Support/KeyHeat/
				// Linux:   ~/.config/KeyHeat/
				appDir := filepath.Join(configDir, "KeyHeat")
				textPath := filepath.Join(appDir, "keyheat.text")
				line := fmt.Sprintf("按压Keycode: %v, 按压Rawcode: %v, 按压Keychar: %v", ev.Keycode, ev.Rawcode, ev.Keychar)
				appendLineToFile(textPath, line)

				select {
				case l.keyChan <- keyName:
				default:
					log.Printf("通道满，丢弃: %s", keyName)
				}
			}
			if ev.Kind == hook.KeyUp {
				keyName := l.mapper.Normalize(ev.Rawcode)
				l.app.Event.Emit("key:pressed", map[string]any{
					"key":  keyName,
					"raw":  ev.Rawcode,
					"type": "up",
				})
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

func appendLineToFile(filePath string, line string) error {
	// 确保目录存在
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return err
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString(line + "\n"); err != nil {
		return err
	}

	return writer.Flush() // 确保数据写入磁盘
}
