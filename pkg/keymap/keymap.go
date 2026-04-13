package keymap

// KeyMapper 按键映射器接口
type KeyMapper interface {
	Normalize(code any) string
}

var globalMapper KeyMapper

type Layout [][]Key
type Key struct {
	Code  string
	Label string
	W     float64
}

type LayoutType string

const (
	LayoutANSI LayoutType = "ANSI"
)

// GetGlobalMapper 获取当前平台的映射器
func GetGlobalMapper() KeyMapper {
	return globalMapper
}

// 统一按键名称常量（跨平台通用）
const (
	Key_A               = "A"
	Key_S               = "S"
	Key_D               = "D"
	Key_F               = "F"
	Key_H               = "H"
	Key_G               = "G"
	Key_Z               = "Z"
	Key_X               = "X"
	Key_C               = "C"
	Key_V               = "V"
	Key_B               = "B"
	Key_Q               = "Q"
	Key_W               = "W"
	Key_E               = "E"
	Key_R               = "R"
	Key_Y               = "Y"
	Key_T               = "T"
	Key_1               = "1"
	Key_2               = "2"
	Key_3               = "3"
	Key_4               = "4"
	Key_6               = "6"
	Key_5               = "5"
	Key_Equal           = "="
	Key_9               = "9"
	Key_7               = "7"
	Key_Minus           = "-"
	Key_8               = "8"
	Key_0               = "0"
	Key_RightBracket    = "]"
	Key_O               = "O"
	Key_U               = "U"
	Key_LeftBracket     = "["
	Key_I               = "I"
	Key_P               = "P"
	Key_L               = "L"
	Key_J               = "J"
	Key_Quote           = "'"
	Key_K               = "K"
	Key_Semicolon       = ";"
	Key_Backslash       = "\\"
	Key_Comma           = ","
	Key_Slash           = "/"
	Key_N               = "N"
	Key_M               = "M"
	Key_Period          = "."
	Key_Grave           = "`"
	Key_KeypadDecimal   = "NumPad ."
	Key_KeypadMultiply  = "NumPad *"
	Key_KeypadPlus      = "NumPad +"
	Key_KeypadClear     = "NumLock"
	Key_KeypadDivide    = "NumPad /"
	Key_KeypadEnter     = "NumPad Enter"
	Key_KeypadMinus     = "NumPad -"
	Key_KeypadEquals    = "NumPad ="
	Key_Keypad0         = "NumPad0"
	Key_Keypad1         = "NumPad1"
	Key_Keypad2         = "NumPad2"
	Key_Keypad3         = "NumPad3"
	Key_Keypad4         = "NumPad4"
	Key_Keypad5         = "NumPad5"
	Key_Keypad6         = "NumPad6"
	Key_Keypad7         = "NumPad7"
	Key_Keypad8         = "NumPad8"
	Key_Keypad9         = "NumPad9"
	Key_Return          = "Enter"
	Key_Tab             = "Tab"
	Key_Space           = "Space"
	Key_Delete          = "Delete"
	Key_Escape          = "Esc"
	Key_Command         = "Command"
	Key_Shift           = "Shift"
	Key_CapsLock        = "CapsLock"
	Key_Option          = "Option"
	Key_Control         = "Control"
	Key_RightCommand    = "RightCommand"
	Key_RightShift      = "RightShift"
	Key_RightOption     = "RightOption"
	Key_RightControl    = "RightControl"
	Key_Function        = "Fn"
	Key_F17             = "F17"
	Key_VolumeUp        = "VolumeUp"
	Key_VolumeDown      = "VolumeDown"
	Key_Mute            = "Mute"
	Key_F18             = "F18"
	Key_F19             = "F19"
	Key_F20             = "F20"
	Key_F5              = "F5"
	Key_F6              = "F6"
	Key_F7              = "F7"
	Key_F3              = "F3"
	Key_F8              = "F8"
	Key_F9              = "F9"
	Key_F11             = "F11"
	Key_F13             = "F13"
	Key_F16             = "F16"
	Key_F14             = "F14"
	Key_F10             = "F10"
	Key_F12             = "F12"
	Key_F15             = "F15"
	Key_Help            = "Help"
	Key_Home            = "Home"
	Key_PageUp          = "PageUp"
	Key_ForwardDelete   = "Del"
	Key_F4              = "F4"
	Key_End             = "End"
	Key_F2              = "F2"
	Key_PageDown        = "PageDown"
	Key_F1              = "F1"
	Key_LeftArrow       = "←"
	Key_RightArrow      = "→"
	Key_DownArrow       = "↓"
	Key_UpArrow         = "↑"
	Key_ISO_Section     = "Section"
	Key_JIS_Yen         = "Yen"
	Key_JIS_Underscore  = "Underscore"
	Key_JIS_KeypadComma = "KeypadComma"
	Key_JIS_Eisu        = "Eisu"
	Key_JIS_Kana        = "Kana"
)
