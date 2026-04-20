//go:build darwin

package autostart

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const appName = "com.keyheat.app"

func getPlistPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Library", "LaunchAgents", appName+".plist"), nil
}

func getExecutablePath() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Abs(exe)
}

func ensureLaunchAgentsDir(plistPath string) error {
	dir := filepath.Dir(plistPath)
	return os.MkdirAll(dir, 0755)
}

func buildPlist(exePath string) string {
	logDir := filepath.Join(os.TempDir(), "keyheat")

	_ = os.MkdirAll(logDir, 0755)

	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN">
<plist version="1.0">
<dict>
	<key>Label</key>
	<string>%s</string>

	<key>ProgramArguments</key>
	<array>
		<string>%s</string>
	</array>

	<key>RunAtLoad</key>
	<true/>

	<key>KeepAlive</key>
	<false/>

	<key>StandardOutPath</key>
	<string>%s/out.log</string>

	<key>StandardErrorPath</key>
	<string>%s/err.log</string>
</dict>
</plist>`, appName, exePath, logDir, logDir)
}

func execCommand(name string, arg ...string) (string, error) {
	out, err := exec.Command(name, arg...).CombinedOutput()
	return string(out), err
}

func contains(s, sub string) bool {
	return strings.Contains(s, sub)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func Enable() error {
	plistPath, err := getPlistPath()
	if err != nil {
		return err
	}
	exePath, err := getExecutablePath()
	if err != nil {
		return err
	}
	// 确保目录存在
	if err := ensureLaunchAgentsDir(plistPath); err != nil {
		return err
	}
	content := buildPlist(exePath)
	if err := os.WriteFile(plistPath, []byte(content), 0644); err != nil {
		return err
	}
	uid := os.Getuid()
	target := fmt.Sprintf("gui/%d", uid)
	// 先卸载（避免重复）
	execCommand("launchctl", "bootout", target, plistPath)
	// 加载
	_, err = execCommand("launchctl", "bootstrap", target, plistPath)
	return err
}

func Disable() error {
	plistPath, err := getPlistPath()
	if err != nil {
		return err
	}
	uid := os.Getuid()
	target := fmt.Sprintf("gui/%d", uid)
	// 卸载
	execCommand("launchctl", "bootout", target, plistPath)
	// 移除注册（关键）
	execCommand("launchctl", "remove", appName)
	// 删除 plist
	if err := os.Remove(plistPath); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

func IsEnabled() bool {
	plistPath, err := getPlistPath()
	if err != nil {
		return false
	}
	uid := os.Getuid()
	target := fmt.Sprintf("gui/%d", uid)
	out, err := execCommand("launchctl", "print", target)
	if err != nil {
		return false
	}
	// 判断 label 是否存在
	return contains(out, appName) && fileExists(plistPath)
}
