package main

import (
	"embed"
	_ "embed"
	"fmt"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
	"github.com/wailsapp/wails/v3/pkg/services/dock"
	"github.com/zxc7563598/key-heat/internal/monitor"
	"github.com/zxc7563598/key-heat/internal/storage"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed assets/icon.png
var icon []byte

type TrayApp struct {
	dockService *dock.DockService
	app         *application.App
	systray     *application.SystemTray
	window      *application.WebviewWindow
	menu        *application.Menu
	repo        storage.Repository
	mon         *monitor.Monitor
	isActive    bool
}

func main() {
	// 初始化数据库
	db, err := storage.NewDB(nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("数据库初始化成功:", db.GetPath())
	// 自动迁移表结构
	if err := storage.Run(db.DB); err != nil {
		log.Fatal(err)
	}
	fmt.Println("数据库自动迁移成功")
	// 注入 repository
	repo := storage.NewRepository(db.DB)
	// 构建APP
	dockService := dock.New()
	app := application.New(application.Options{
		Name:        "Key Heat",
		Description: "A demo of using raw HTML & CSS",
		Services: []application.Service{
			application.NewService(&GreetService{
				repo: repo,
			}),
			application.NewService(dockService),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: false,
		},
		Windows: application.WindowsOptions{
			DisableQuitOnLastWindowClosed: true,
		},
	})
	// 启动应用程序
	trayApp := &TrayApp{
		app:         app,
		dockService: dockService,
		repo:        repo,
	}
	trayApp.setup()
	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (t *TrayApp) setup() {
	t.isActive = true
	t.mon = monitor.NewMonitor(t.repo, nil)
	go t.mon.Start()
	// 创建托盘
	t.systray = t.app.SystemTray.New()
	t.systray.SetIcon(icon)
	// 创建菜单
	t.createMenu()
	// 打开窗口
	t.openWindow()
}

func (t *TrayApp) createMenu() {
	t.menu = t.app.NewMenu()
	// 切换监听状态
	if t.isActive {
		statusItem := t.menu.Add("当前状态: 已监听")
		statusItem.SetEnabled(false)
		// Toggle active
		t.menu.Add("停止监听").OnClick(func(ctx *application.Context) {
			t.mon.Stop()
			t.isActive = false
			t.createMenu()
		})
	} else {
		statusItem := t.menu.Add("当前状态: 未监听")
		statusItem.SetEnabled(false)
		// Toggle active
		t.menu.Add("启动监听").OnClick(func(ctx *application.Context) {
			t.mon = monitor.NewMonitor(t.repo, nil)
			go t.mon.Start()
			t.isActive = true
			t.createMenu()
		})
	}
	t.menu.AddSeparator()
	// 显示主界面
	t.menu.Add("显示主界面").OnClick(func(ctx *application.Context) {
		if t.window == nil {
			t.openWindow()
		}
		t.dockService.ShowAppIcon()
		t.window.Show()
		t.window.Focus()
	})
	t.menu.AddSeparator()
	// 退出软件
	t.menu.Add("退出").OnClick(func(ctx *application.Context) {
		t.app.Quit()
	})
	t.systray.SetMenu(t.menu)
}

func (t *TrayApp) openWindow() {
	t.window = t.app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title: "Key Heat",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/",
	})
	t.window.OnWindowEvent(events.Common.WindowClosing, func(e *application.WindowEvent) {
		t.dockService.HideAppIcon()
		t.window = nil
	})
}
