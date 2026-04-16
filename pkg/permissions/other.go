//go:build !darwin

package permissions

func hasAccessibilityPermission() bool {
	return true // 非 macOS 直接放行
}
