package main

import "github.com/zxc7563598/key-heat/internal/storage"

type GreetService struct {
	repo storage.Repository
}

func (g *GreetService) Greet(name string) string {
	return "Hello " + name + "!"
}
