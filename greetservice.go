package main

import (
	"github.com/zxc7563598/key-heat/internal/storage"
	"github.com/zxc7563598/key-heat/pkg/keymap"
)

type GreetService struct {
	repo storage.Repository
}

func (g *GreetService) GetKeyLayout(layoutType keymap.LayoutType) keymap.Layout {
	return keymap.GetKeyLayout(layoutType)
}
