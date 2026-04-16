package main

import (
	"log"

	"github.com/zxc7563598/key-heat/internal/storage"
	"github.com/zxc7563598/key-heat/pkg/keymap"
	"github.com/zxc7563598/key-heat/pkg/permissions"
)

type GreetService struct {
	repo storage.Repository
	t    *TrayApp
}

func (g *GreetService) GetKeyLayout(layoutType keymap.LayoutType) keymap.Layout {
	return keymap.GetKeyLayout(layoutType)
}

func (g *GreetService) GetHeatmap(start, end string) map[string]int {
	if start == "" {
		start = "2000-01-01"
	}
	if end == "" {
		end = "2100-01-01"
	}
	data, err := g.repo.GetStatsByDateRange(start, end)
	if err != nil {
		log.Println("获取热力图数据失败")
		return map[string]int{}
	}
	return data
}

func (g *GreetService) GetPermission() bool {
	return permissions.HasAccessibilityPermission()
}

func (g *GreetService) GetMonitorStatus() bool {
	return g.t.isActive
}

func (g *GreetService) SetMonitorStart() {
	if !g.t.isActive {
		g.t.startListening()
	}
}

func (g *GreetService) SetMonitorClose() {
	if g.t.isActive {
		g.t.stopListening()
	}
}
