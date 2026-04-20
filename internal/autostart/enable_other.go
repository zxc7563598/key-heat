//go:build !windows && !darwin

package autostart

import "fmt"

func Enable() error {
	return fmt.Errorf("当前环境不支持")
}

func Disable() error {
	return fmt.Errorf("当前环境不支持")
}

func IsEnabled() bool {
	return false
}
