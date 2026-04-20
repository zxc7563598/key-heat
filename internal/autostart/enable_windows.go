//go:build windows

package autostart

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

const appName = "key-heat"

// 获取当前 exe 路径
func getExecutablePath() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Abs(exe)
}

// 启用开机自启
func Enable() error {
	exePath, err := getExecutablePath()
	if err != nil {
		return err
	}
	value := fmt.Sprintf(`"%s"`, exePath)
	key, _, err := registry.CreateKey(
		registry.CURRENT_USER,
		`Software\Microsoft\Windows\CurrentVersion\Run`,
		registry.SET_VALUE,
	)
	if err != nil {
		return err
	}
	defer key.Close()
	return key.SetStringValue(appName, value)
}

// 关闭开机自启
func Disable() error {
	key, err := registry.OpenKey(
		registry.CURRENT_USER,
		`Software\Microsoft\Windows\CurrentVersion\Run`,
		registry.SET_VALUE,
	)
	if err != nil {
		return err
	}
	defer key.Close()
	err = key.DeleteValue(appName)
	// 不存在也算成功
	if err == registry.ErrNotExist {
		return nil
	}
	return err
}

// 是否已开启
func IsEnabled() bool {
	key, err := registry.OpenKey(
		registry.CURRENT_USER,
		`Software\Microsoft\Windows\CurrentVersion\Run`,
		registry.QUERY_VALUE,
	)
	if err != nil {
		return false
	}
	defer key.Close()
	val, _, err := key.GetStringValue(appName)
	if err != nil {
		return false
	}
	exePath, err := getExecutablePath()
	if err != nil {
		return false
	}
	expected := fmt.Sprintf(`"%s"`, exePath)
	return val == expected
}
