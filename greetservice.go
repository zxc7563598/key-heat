package main

import (
	"fmt"
	"log"

	"github.com/zxc7563598/key-heat/internal/autostart"
	"github.com/zxc7563598/key-heat/internal/prompt"
	"github.com/zxc7563598/key-heat/internal/storage"
	"github.com/zxc7563598/key-heat/pkg/keymap"
	"github.com/zxc7563598/key-heat/pkg/permissions"
)

type GreetService struct {
	repo storage.Repository
	t    *TrayApp
}

// 获取键盘布局
func (g *GreetService) GetKeyLayout(layoutType keymap.LayoutType) keymap.Layout {
	return keymap.GetKeyLayout(layoutType)
}

// 根据日期查询热力数据
func (g *GreetService) GetHeatmap(start, end string) map[string]int {
	if start == "" {
		start = "2000-01-01"
	}
	if end == "" {
		end = "2100-01-01"
	}
	data, err := g.repo.GetStatsByDateRange(start, end)
	if err != nil {
		log.Printf("获取热力图数据失败: %v", err)
		return map[string]int{}
	}
	return data
}

// 获取当前是否拥有监听权限
func (g *GreetService) GetPermission() bool {
	return permissions.HasAccessibilityPermission()
}

// 获取当前是否监听
func (g *GreetService) GetMonitorStatus() bool {
	return g.t.isActive
}

// 启动监听
func (g *GreetService) SetMonitorStart() {
	if !g.t.isActive {
		g.t.startListening()
	}
}

// 停止监听
func (g *GreetService) SetMonitorClose() {
	if g.t.isActive {
		g.t.stopListening()
	}
}

// 获取开机自启动状态
func (g *GreetService) GetBootStartup() bool {
	return autostart.IsEnabled()
}

// 设置开机自启动
func (g *GreetService) OpenBootStartup() bool {
	err := autostart.Enable()
	if err != nil {
		log.Printf("开机自启动设置失败: %v", err)
		return false
	}
	return true
}

// 关闭开机自启动
func (g *GreetService) CloseBootStartup() bool {
	err := autostart.Disable()
	if err != nil {
		log.Printf("取消自启动设置失败: %v", err)
		return false
	}
	return true
}

func (g *GreetService) GetPrompt(start, end string) string {
	startText := start
	endText := end
	if start == "" {
		start = "2000-01-01"
		startText = "最初"
	}
	if end == "" {
		end = "2100-01-01"
		endText = "此刻"
	}
	data, err := g.repo.GetStatsByDateRange(start, end)
	if err != nil {
		log.Printf("获取热力图数据失败: %v", err)
		return ""
	}
	prompt, err := prompt.BuildKeyboardReportPrompt(data, fmt.Sprintf("%s - %s", startText, endText))
	if err != nil {
		log.Printf("提示词数据获取失败: %v", err)
		return ""
	}
	return prompt
}
