//go:build darwin

package permissions

/*
#cgo LDFLAGS: -framework ApplicationServices
#include <ApplicationServices/ApplicationServices.h>
*/
import "C"

func hasAccessibilityPermission() bool {
	return C.AXIsProcessTrusted() == 1
}
