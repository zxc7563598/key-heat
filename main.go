package main

import (
	"embed"
	_ "embed"
	"fmt"
	"log"
	"os/exec"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
	"github.com/wailsapp/wails/v3/pkg/services/dock"
	"github.com/zxc7563598/key-heat/internal/monitor"
	"github.com/zxc7563598/key-heat/internal/storage"
	"github.com/zxc7563598/key-heat/pkg/permissions"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed assets/iconTemplate.png
var iconTemplate []byte

//go:embed assets/iconTemplateWhite.png
var iconTemplateWhite []byte

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
	greet := &GreetService{
		repo: repo,
	}
	dockService := dock.New()
	app := application.New(application.Options{
		Name:        "Key Heat",
		Description: "A demo of using raw HTML & CSS",
		Services: []application.Service{
			application.NewService(greet),
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
	greet.t = trayApp
	trayApp.setup()
	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (t *TrayApp) setup() {
	t.isActive = false
	if permissions.HasAccessibilityPermission() {
		t.isActive = true
		t.mon = monitor.NewMonitor(t.repo, nil, t.app)
		go t.mon.Start()
	}
	t.app.Event.Emit("active:switch", t.isActive)
	// 创建托盘
	t.systray = t.app.SystemTray.New()
	t.systray.SetIcon(iconTemplate)
	t.systray.SetDarkModeIcon(iconTemplateWhite)
	t.systray.SetTemplateIcon(iconTemplate)
	// 创建菜单
	t.createMenu()
	// 打开窗口
	t.openWindow()
}

func (t *TrayApp) createMenu() {
	t.menu = t.app.NewMenu()
	t.menu.Add("KeyHeat").SetEnabled(false)
	// 切换监听状态
	if t.isActive {
		t.menu.Add("● 正在监听").SetEnabled(false)
		// Toggle active
		t.menu.Add("停止监听").OnClick(func(ctx *application.Context) {
			t.stopListening()
		})
	} else {
		t.menu.Add("○ 已停止").SetEnabled(false)
		// Toggle active
		t.menu.Add("启动监听").OnClick(func(ctx *application.Context) {
			t.startListening()
		})
	}
	t.menu.AddSeparator()
	// 打开主界面
	t.menu.Add("打开主界面").OnClick(func(ctx *application.Context) {
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
		Title:  "Key Heat",
		Width:  1200,
		Height: 600,
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

func (t *TrayApp) startListening() {
	log.Println("startListening开始调用")
	if permissions.HasAccessibilityPermission() {
		t.mon = monitor.NewMonitor(t.repo, nil, t.app)
		go t.mon.Start()
		t.isActive = true
		t.app.Event.Emit("active:switch", t.isActive)
	} else {
		t.showPermissionDialog()
	}
	t.createMenu()
}
func (t *TrayApp) stopListening() {
	log.Println("stopListening开始调用")
	t.mon.Stop()
	t.isActive = false
	t.app.Event.Emit("active:switch", t.isActive)
	t.createMenu()
}

func (t *TrayApp) showPermissionDialog() {
	dialog := t.app.Dialog.Question().
		SetTitle("需要权限").
		SetMessage("需要开启辅助功能权限才能监听键盘，是否前往设置？")
	goSetting := dialog.AddButton("前往设置")
	goSetting.OnClick(func() {
		exec.Command("open", "x-apple.systempreferences:com.apple.preference.security?Privacy_Accessibility").Run()
	})
	goCancel := dialog.AddButton("取消")
	goCancel.OnClick(func() {
		t.app.Quit()
	})
	dialog.Show()
}
