//go:build windows

package keymap

type VirtualKeyCode uint16

// 来源 https://learn.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes
const (
	VK_LBUTTON                         VirtualKeyCode = 0x01
	VK_RBUTTON                         VirtualKeyCode = 0x02
	VK_CANCEL                          VirtualKeyCode = 0x03
	VK_MBUTTON                         VirtualKeyCode = 0x04
	VK_XBUTTON1                        VirtualKeyCode = 0x05
	VK_XBUTTON2                        VirtualKeyCode = 0x06
	VK_BACK                            VirtualKeyCode = 0x08
	VK_TAB                             VirtualKeyCode = 0x09
	VK_CLEAR                           VirtualKeyCode = 0x0C
	VK_RETURN                          VirtualKeyCode = 0x0D
	VK_SHIFT                           VirtualKeyCode = 0x10
	VK_CONTROL                         VirtualKeyCode = 0x11
	VK_MENU                            VirtualKeyCode = 0x12
	VK_PAUSE                           VirtualKeyCode = 0x13
	VK_CAPITAL                         VirtualKeyCode = 0x14
	VK_KANA                            VirtualKeyCode = 0x15
	VK_HANGUL                          VirtualKeyCode = 0x15
	VK_IME_ON                          VirtualKeyCode = 0x16
	VK_JUNJA                           VirtualKeyCode = 0x17
	VK_FINAL                           VirtualKeyCode = 0x18
	VK_HANJA                           VirtualKeyCode = 0x19
	VK_KANJI                           VirtualKeyCode = 0x19
	VK_IME_OFF                         VirtualKeyCode = 0x1A
	VK_ESCAPE                          VirtualKeyCode = 0x1B
	VK_CONVERT                         VirtualKeyCode = 0x1C
	VK_NONCONVERT                      VirtualKeyCode = 0x1D
	VK_ACCEPT                          VirtualKeyCode = 0x1E
	VK_MODECHANGE                      VirtualKeyCode = 0x1F
	VK_SPACE                           VirtualKeyCode = 0x20
	VK_PRIOR                           VirtualKeyCode = 0x21
	VK_NEXT                            VirtualKeyCode = 0x22
	VK_END                             VirtualKeyCode = 0x23
	VK_HOME                            VirtualKeyCode = 0x24
	VK_LEFT                            VirtualKeyCode = 0x25
	VK_UP                              VirtualKeyCode = 0x26
	VK_RIGHT                           VirtualKeyCode = 0x27
	VK_DOWN                            VirtualKeyCode = 0x28
	VK_SELECT                          VirtualKeyCode = 0x29
	VK_PRINT                           VirtualKeyCode = 0x2A
	VK_EXECUTE                         VirtualKeyCode = 0x2B
	VK_SNAPSHOT                        VirtualKeyCode = 0x2C
	VK_INSERT                          VirtualKeyCode = 0x2D
	VK_DELETE                          VirtualKeyCode = 0x2E
	VK_HELP                            VirtualKeyCode = 0x2F
	VK_0                               VirtualKeyCode = 0x30
	VK_1                               VirtualKeyCode = 0x31
	VK_2                               VirtualKeyCode = 0x32
	VK_3                               VirtualKeyCode = 0x33
	VK_4                               VirtualKeyCode = 0x34
	VK_5                               VirtualKeyCode = 0x35
	VK_6                               VirtualKeyCode = 0x36
	VK_7                               VirtualKeyCode = 0x37
	VK_8                               VirtualKeyCode = 0x38
	VK_9                               VirtualKeyCode = 0x39
	VK_A                               VirtualKeyCode = 0x41
	VK_B                               VirtualKeyCode = 0x42
	VK_C                               VirtualKeyCode = 0x43
	VK_D                               VirtualKeyCode = 0x44
	VK_E                               VirtualKeyCode = 0x45
	VK_F                               VirtualKeyCode = 0x46
	VK_G                               VirtualKeyCode = 0x47
	VK_H                               VirtualKeyCode = 0x48
	VK_I                               VirtualKeyCode = 0x49
	VK_J                               VirtualKeyCode = 0x4A
	VK_K                               VirtualKeyCode = 0x4B
	VK_L                               VirtualKeyCode = 0x4C
	VK_M                               VirtualKeyCode = 0x4D
	VK_N                               VirtualKeyCode = 0x4E
	VK_O                               VirtualKeyCode = 0x4F
	VK_P                               VirtualKeyCode = 0x50
	VK_Q                               VirtualKeyCode = 0x51
	VK_R                               VirtualKeyCode = 0x52
	VK_S                               VirtualKeyCode = 0x53
	VK_T                               VirtualKeyCode = 0x54
	VK_U                               VirtualKeyCode = 0x55
	VK_V                               VirtualKeyCode = 0x56
	VK_W                               VirtualKeyCode = 0x57
	VK_X                               VirtualKeyCode = 0x58
	VK_Y                               VirtualKeyCode = 0x59
	VK_Z                               VirtualKeyCode = 0x5A
	VK_LWIN                            VirtualKeyCode = 0x5B
	VK_RWIN                            VirtualKeyCode = 0x5C
	VK_APPS                            VirtualKeyCode = 0x5D
	VK_SLEEP                           VirtualKeyCode = 0x5F
	VK_NUMPAD0                         VirtualKeyCode = 0x60
	VK_NUMPAD1                         VirtualKeyCode = 0x61
	VK_NUMPAD2                         VirtualKeyCode = 0x62
	VK_NUMPAD3                         VirtualKeyCode = 0x63
	VK_NUMPAD4                         VirtualKeyCode = 0x64
	VK_NUMPAD5                         VirtualKeyCode = 0x65
	VK_NUMPAD6                         VirtualKeyCode = 0x66
	VK_NUMPAD7                         VirtualKeyCode = 0x67
	VK_NUMPAD8                         VirtualKeyCode = 0x68
	VK_NUMPAD9                         VirtualKeyCode = 0x69
	VK_MULTIPLY                        VirtualKeyCode = 0x6A
	VK_ADD                             VirtualKeyCode = 0x6B
	VK_SEPARATOR                       VirtualKeyCode = 0x6C
	VK_SUBTRACT                        VirtualKeyCode = 0x6D
	VK_DECIMAL                         VirtualKeyCode = 0x6E
	VK_DIVIDE                          VirtualKeyCode = 0x6F
	VK_F1                              VirtualKeyCode = 0x70
	VK_F2                              VirtualKeyCode = 0x71
	VK_F3                              VirtualKeyCode = 0x72
	VK_F4                              VirtualKeyCode = 0x73
	VK_F5                              VirtualKeyCode = 0x74
	VK_F6                              VirtualKeyCode = 0x75
	VK_F7                              VirtualKeyCode = 0x76
	VK_F8                              VirtualKeyCode = 0x77
	VK_F9                              VirtualKeyCode = 0x78
	VK_F10                             VirtualKeyCode = 0x79
	VK_F11                             VirtualKeyCode = 0x7A
	VK_F12                             VirtualKeyCode = 0x7B
	VK_F13                             VirtualKeyCode = 0x7C
	VK_F14                             VirtualKeyCode = 0x7D
	VK_F15                             VirtualKeyCode = 0x7E
	VK_F16                             VirtualKeyCode = 0x7F
	VK_F17                             VirtualKeyCode = 0x80
	VK_F18                             VirtualKeyCode = 0x81
	VK_F19                             VirtualKeyCode = 0x82
	VK_F20                             VirtualKeyCode = 0x83
	VK_F21                             VirtualKeyCode = 0x84
	VK_F22                             VirtualKeyCode = 0x85
	VK_F23                             VirtualKeyCode = 0x86
	VK_F24                             VirtualKeyCode = 0x87
	VK_NUMLOCK                         VirtualKeyCode = 0x90
	VK_SCROLL                          VirtualKeyCode = 0x91
	VK_LSHIFT                          VirtualKeyCode = 0xA0
	VK_RSHIFT                          VirtualKeyCode = 0xA1
	VK_LCONTROL                        VirtualKeyCode = 0xA2
	VK_RCONTROL                        VirtualKeyCode = 0xA3
	VK_LMENU                           VirtualKeyCode = 0xA4
	VK_RMENU                           VirtualKeyCode = 0xA5
	VK_BROWSER_BACK                    VirtualKeyCode = 0xA6
	VK_BROWSER_FORWARD                 VirtualKeyCode = 0xA7
	VK_BROWSER_REFRESH                 VirtualKeyCode = 0xA8
	VK_BROWSER_STOP                    VirtualKeyCode = 0xA9
	VK_BROWSER_SEARCH                  VirtualKeyCode = 0xAA
	VK_BROWSER_FAVORITES               VirtualKeyCode = 0xAB
	VK_BROWSER_HOME                    VirtualKeyCode = 0xAC
	VK_VOLUME_MUTE                     VirtualKeyCode = 0xAD
	VK_VOLUME_DOWN                     VirtualKeyCode = 0xAE
	VK_VOLUME_UP                       VirtualKeyCode = 0xAF
	VK_MEDIA_NEXT_TRACK                VirtualKeyCode = 0xB0
	VK_MEDIA_PREV_TRACK                VirtualKeyCode = 0xB1
	VK_MEDIA_STOP                      VirtualKeyCode = 0xB2
	VK_MEDIA_PLAY_PAUSE                VirtualKeyCode = 0xB3
	VK_LAUNCH_MAIL                     VirtualKeyCode = 0xB4
	VK_LAUNCH_MEDIA_SELECT             VirtualKeyCode = 0xB5
	VK_LAUNCH_APP1                     VirtualKeyCode = 0xB6
	VK_LAUNCH_APP2                     VirtualKeyCode = 0xB7
	VK_OEM_1                           VirtualKeyCode = 0xBA
	VK_OEM_PLUS                        VirtualKeyCode = 0xBB
	VK_OEM_COMMA                       VirtualKeyCode = 0xBC
	VK_OEM_MINUS                       VirtualKeyCode = 0xBD
	VK_OEM_PERIOD                      VirtualKeyCode = 0xBE
	VK_OEM_2                           VirtualKeyCode = 0xBF
	VK_OEM_3                           VirtualKeyCode = 0xC0
	VK_GAMEPAD_A                       VirtualKeyCode = 0xC3
	VK_GAMEPAD_B                       VirtualKeyCode = 0xC4
	VK_GAMEPAD_X                       VirtualKeyCode = 0xC5
	VK_GAMEPAD_Y                       VirtualKeyCode = 0xC6
	VK_GAMEPAD_RIGHT_SHOULDER          VirtualKeyCode = 0xC7
	VK_GAMEPAD_LEFT_SHOULDER           VirtualKeyCode = 0xC8
	VK_GAMEPAD_LEFT_TRIGGER            VirtualKeyCode = 0xC9
	VK_GAMEPAD_RIGHT_TRIGGER           VirtualKeyCode = 0xCA
	VK_GAMEPAD_DPAD_UP                 VirtualKeyCode = 0xCB
	VK_GAMEPAD_DPAD_DOWN               VirtualKeyCode = 0xCC
	VK_GAMEPAD_DPAD_LEFT               VirtualKeyCode = 0xCD
	VK_GAMEPAD_DPAD_RIGHT              VirtualKeyCode = 0xCE
	VK_GAMEPAD_MENU                    VirtualKeyCode = 0xCF
	VK_GAMEPAD_VIEW                    VirtualKeyCode = 0xD0
	VK_GAMEPAD_LEFT_THUMBSTICK_BUTTON  VirtualKeyCode = 0xD1
	VK_GAMEPAD_RIGHT_THUMBSTICK_BUTTON VirtualKeyCode = 0xD2
	VK_GAMEPAD_LEFT_THUMBSTICK_UP      VirtualKeyCode = 0xD3
	VK_GAMEPAD_LEFT_THUMBSTICK_DOWN    VirtualKeyCode = 0xD4
	VK_GAMEPAD_LEFT_THUMBSTICK_RIGHT   VirtualKeyCode = 0xD5
	VK_GAMEPAD_LEFT_THUMBSTICK_LEFT    VirtualKeyCode = 0xD6
	VK_GAMEPAD_RIGHT_THUMBSTICK_UP     VirtualKeyCode = 0xD7
	VK_GAMEPAD_RIGHT_THUMBSTICK_DOWN   VirtualKeyCode = 0xD8
	VK_GAMEPAD_RIGHT_THUMBSTICK_RIGHT  VirtualKeyCode = 0xD9
	VK_GAMEPAD_RIGHT_THUMBSTICK_LEFT   VirtualKeyCode = 0xDA
	VK_OEM_4                           VirtualKeyCode = 0xDB
	VK_OEM_5                           VirtualKeyCode = 0xDC
	VK_OEM_6                           VirtualKeyCode = 0xDD
	VK_OEM_7                           VirtualKeyCode = 0xDE
	VK_OEM_8                           VirtualKeyCode = 0xDF
	VK_OEM_102                         VirtualKeyCode = 0xE2
	VK_PROCESSKEY                      VirtualKeyCode = 0xE5
	VK_PACKET                          VirtualKeyCode = 0xE7
	VK_ATTN                            VirtualKeyCode = 0xF6
	VK_CRSEL                           VirtualKeyCode = 0xF7
	VK_EXSEL                           VirtualKeyCode = 0xF8
	VK_EREOF                           VirtualKeyCode = 0xF9
	VK_PLAY                            VirtualKeyCode = 0xFA
	VK_ZOOM                            VirtualKeyCode = 0xFB
	VK_NONAME                          VirtualKeyCode = 0xFC
	VK_PA1                             VirtualKeyCode = 0xFD
	VK_OEM_CLEAR                       VirtualKeyCode = 0xFE
)

type windowsMapper struct {
	codeToName map[VirtualKeyCode]string
}

func init() {
	mapper := &windowsMapper{
		codeToName: make(map[VirtualKeyCode]string),
	}

	mapper.codeToName[VK_ESCAPE] = Key_Escape
	mapper.codeToName[VK_F1] = Key_F1
	mapper.codeToName[VK_F2] = Key_F2
	mapper.codeToName[VK_F3] = Key_F3
	mapper.codeToName[VK_F4] = Key_F4
	mapper.codeToName[VK_F5] = Key_F5
	mapper.codeToName[VK_F6] = Key_F6
	mapper.codeToName[VK_F7] = Key_F7
	mapper.codeToName[VK_F8] = Key_F8
	mapper.codeToName[VK_F9] = Key_F9
	mapper.codeToName[VK_F10] = Key_F10
	mapper.codeToName[VK_F11] = Key_F11
	mapper.codeToName[VK_F12] = Key_F12
	mapper.codeToName[VK_F13] = Key_F13
	mapper.codeToName[VK_F14] = Key_F14
	mapper.codeToName[VK_F15] = Key_F15
	mapper.codeToName[VK_F16] = Key_F16
	mapper.codeToName[VK_F17] = Key_F17
	mapper.codeToName[VK_F18] = Key_F18
	mapper.codeToName[VK_F19] = Key_F19
	mapper.codeToName[VK_F20] = Key_F20
	mapper.codeToName[VK_PRINT] = Key_Print
	mapper.codeToName[VK_SCROLL] = Key_Scroll
	mapper.codeToName[VK_PAUSE] = Key_Pause

	mapper.codeToName[VK_OEM_3] = Key_Grave
	mapper.codeToName[VK_1] = Key_1
	mapper.codeToName[VK_2] = Key_2
	mapper.codeToName[VK_3] = Key_3
	mapper.codeToName[VK_4] = Key_4
	mapper.codeToName[VK_5] = Key_5
	mapper.codeToName[VK_6] = Key_6
	mapper.codeToName[VK_7] = Key_7
	mapper.codeToName[VK_8] = Key_8
	mapper.codeToName[VK_9] = Key_9
	mapper.codeToName[VK_0] = Key_0
	mapper.codeToName[VK_OEM_MINUS] = Key_Minus
	mapper.codeToName[VK_OEM_PLUS] = Key_Equal
	mapper.codeToName[VK_BACK] = Key_Delete
	mapper.codeToName[VK_INSERT] = Key_Insert
	mapper.codeToName[VK_HOME] = Key_Home
	mapper.codeToName[VK_PRIOR] = Key_PageUp

	mapper.codeToName[VK_TAB] = Key_Tab
	mapper.codeToName[VK_Q] = Key_Q
	mapper.codeToName[VK_W] = Key_W
	mapper.codeToName[VK_E] = Key_E
	mapper.codeToName[VK_R] = Key_R
	mapper.codeToName[VK_T] = Key_T
	mapper.codeToName[VK_Y] = Key_Y
	mapper.codeToName[VK_U] = Key_U
	mapper.codeToName[VK_I] = Key_I
	mapper.codeToName[VK_O] = Key_O
	mapper.codeToName[VK_P] = Key_P
	mapper.codeToName[VK_OEM_4] = Key_LeftBracket
	mapper.codeToName[VK_OEM_6] = Key_RightBracket
	mapper.codeToName[VK_OEM_5] = Key_Backslash
	mapper.codeToName[VK_DELETE] = Key_ForwardDelete
	mapper.codeToName[VK_END] = Key_End
	mapper.codeToName[VK_NEXT] = Key_PageDown

	mapper.codeToName[VK_CAPITAL] = Key_CapsLock
	mapper.codeToName[VK_A] = Key_A
	mapper.codeToName[VK_S] = Key_S
	mapper.codeToName[VK_D] = Key_D
	mapper.codeToName[VK_F] = Key_F
	mapper.codeToName[VK_G] = Key_G
	mapper.codeToName[VK_H] = Key_H
	mapper.codeToName[VK_J] = Key_J
	mapper.codeToName[VK_K] = Key_K
	mapper.codeToName[VK_L] = Key_L
	mapper.codeToName[VK_OEM_1] = Key_Semicolon
	mapper.codeToName[VK_OEM_7] = Key_Quote
	mapper.codeToName[VK_RETURN] = Key_Return

	mapper.codeToName[VK_LSHIFT] = Key_Shift
	mapper.codeToName[VK_Z] = Key_Z
	mapper.codeToName[VK_X] = Key_X
	mapper.codeToName[VK_C] = Key_C
	mapper.codeToName[VK_V] = Key_V
	mapper.codeToName[VK_B] = Key_B
	mapper.codeToName[VK_N] = Key_N
	mapper.codeToName[VK_M] = Key_M
	mapper.codeToName[VK_OEM_COMMA] = Key_Comma
	mapper.codeToName[VK_OEM_PERIOD] = Key_Period
	mapper.codeToName[VK_OEM_2] = Key_Slash
	mapper.codeToName[VK_RSHIFT] = Key_RightShift

	mapper.codeToName[VK_LCONTROL] = Key_Ctrl
	mapper.codeToName[VK_LWIN] = Key_Win
	mapper.codeToName[VK_LMENU] = Key_Alt
	mapper.codeToName[VK_SPACE] = Key_Space
	mapper.codeToName[VK_RMENU] = Key_RightAlt
	mapper.codeToName[VK_RWIN] = Key_RightWin
	mapper.codeToName[VK_APPS] = Key_RightMenu
	mapper.codeToName[VK_RCONTROL] = Key_RightCtrl

	mapper.codeToName[VK_UP] = Key_UpArrow
	mapper.codeToName[VK_DOWN] = Key_DownArrow
	mapper.codeToName[VK_LEFT] = Key_LeftArrow
	mapper.codeToName[VK_RIGHT] = Key_RightArrow

	mapper.codeToName[VK_NUMLOCK] = Key_KeypadClear
	mapper.codeToName[VK_DECIMAL] = Key_KeypadDecimal
	mapper.codeToName[VK_MULTIPLY] = Key_KeypadMultiply
	mapper.codeToName[VK_ADD] = Key_KeypadPlus
	mapper.codeToName[VK_DIVIDE] = Key_KeypadDivide
	mapper.codeToName[VK_SEPARATOR] = Key_KeypadEnter
	mapper.codeToName[VK_SUBTRACT] = Key_KeypadMinus
	mapper.codeToName[VK_NUMPAD0] = Key_Keypad0
	mapper.codeToName[VK_NUMPAD1] = Key_Keypad1
	mapper.codeToName[VK_NUMPAD2] = Key_Keypad2
	mapper.codeToName[VK_NUMPAD3] = Key_Keypad3
	mapper.codeToName[VK_NUMPAD4] = Key_Keypad4
	mapper.codeToName[VK_NUMPAD5] = Key_Keypad5
	mapper.codeToName[VK_NUMPAD6] = Key_Keypad6
	mapper.codeToName[VK_NUMPAD7] = Key_Keypad7
	mapper.codeToName[VK_NUMPAD8] = Key_Keypad8
	mapper.codeToName[VK_NUMPAD9] = Key_Keypad9

	globalMapper = mapper
}

func (m *windowsMapper) Normalize(code any) string {
	// 支持直接传入 VirtualKeyCode 类型
	if c, ok := code.(VirtualKeyCode); ok {
		if name, ok := m.codeToName[c]; ok {
			return name
		}
	}
	// 兼容直接传入 uint16
	if c, ok := code.(uint16); ok {
		if name, ok := m.codeToName[VirtualKeyCode(c)]; ok {
			return name
		}
	}
	return "Unknown"
}

func GetKeyLayout(l LayoutType) Layout {
	switch l {
	case LayoutANSI:
		return Layout{
			{
				{Code: "Esc", Label: "Esc", W: 1},
				{Code: "None", Label: "", W: 1},
				{Code: "F1", Label: "F1", W: 1},
				{Code: "F2", Label: "F2", W: 1},
				{Code: "F3", Label: "F3", W: 1},
				{Code: "F4", Label: "F4", W: 1},
				{Code: "None", Label: "", W: 0.5},
				{Code: "F5", Label: "F5", W: 1},
				{Code: "F6", Label: "F6", W: 1},
				{Code: "F7", Label: "F7", W: 1},
				{Code: "F8", Label: "F8", W: 1},
				{Code: "None", Label: "", W: 0.5},
				{Code: "F9", Label: "F9", W: 1},
				{Code: "F10", Label: "F10", W: 1},
				{Code: "F11", Label: "F11", W: 1},
				{Code: "F12", Label: "F12", W: 1},
				{Code: "None", Label: "", W: 0.2},
				{Code: "Print", Label: "PrtSc", W: 1},
				{Code: "Scroll", Label: "Scroll\nLock", W: 1},
				{Code: "Pause", Label: "Pause\nBreak", W: 1},
				{Code: "None", Label: "", W: 0.2},
				{Code: "None", Label: "", W: 1},
				{Code: "None", Label: "", W: 1},
				{Code: "None", Label: "", W: 1},
				{Code: "None", Label: "", W: 1},
			},
			{
				{Code: "`", Label: "`", W: 1},
				{Code: "1", Label: "1", W: 1},
				{Code: "2", Label: "2", W: 1},
				{Code: "3", Label: "3", W: 1},
				{Code: "4", Label: "4", W: 1},
				{Code: "5", Label: "5", W: 1},
				{Code: "6", Label: "6", W: 1},
				{Code: "7", Label: "7", W: 1},
				{Code: "8", Label: "8", W: 1},
				{Code: "9", Label: "9", W: 1},
				{Code: "0", Label: "0", W: 1},
				{Code: "-", Label: "-", W: 1},
				{Code: "=", Label: "=", W: 1},
				{Code: "Delete", Label: "⌫", W: 2},
				{Code: "None", Label: "", W: 0.2},
				{Code: "Insert", Label: "Ins", W: 1},
				{Code: "Home", Label: "Home", W: 1},
				{Code: "PageUp", Label: "PgUp", W: 1},
				{Code: "None", Label: "", W: 0.2},
				{Code: "NumLock", Label: "Num\nLock", W: 1},
				{Code: "NumPad /", Label: "/", W: 1},
				{Code: "NumPad *", Label: "*", W: 1},
				{Code: "NumPad -", Label: "-", W: 1},
			},
			{
				{Code: "Tab", Label: "⇥", W: 1.5},
				{Code: "Q", Label: "Q", W: 1},
				{Code: "W", Label: "W", W: 1},
				{Code: "E", Label: "E", W: 1},
				{Code: "R", Label: "R", W: 1},
				{Code: "T", Label: "T", W: 1},
				{Code: "Y", Label: "Y", W: 1},
				{Code: "U", Label: "U", W: 1},
				{Code: "I", Label: "I", W: 1},
				{Code: "O", Label: "O", W: 1},
				{Code: "P", Label: "P", W: 1},
				{Code: "[", Label: "[", W: 1},
				{Code: "]", Label: "]", W: 1},
				{Code: "\\", Label: "\\", W: 1.5},
				{Code: "None", Label: "", W: 0.2},
				{Code: "Del", Label: "Delete", W: 1},
				{Code: "End", Label: "End", W: 1},
				{Code: "PageDown", Label: "PgDn", W: 1},
				{Code: "None", Label: "", W: 0.2},
				{Code: "NumPad7", Label: "7", W: 1},
				{Code: "NumPad8", Label: "8", W: 1},
				{Code: "NumPad9", Label: "9", W: 1},
				{Code: "NumPad +", Label: "+", W: 1},
			},
			{
				{Code: "CapsLock", Label: "⇪", W: 1.75},
				{Code: "A", Label: "A", W: 1},
				{Code: "S", Label: "S", W: 1},
				{Code: "D", Label: "D", W: 1},
				{Code: "F", Label: "F", W: 1},
				{Code: "G", Label: "G", W: 1},
				{Code: "H", Label: "H", W: 1},
				{Code: "J", Label: "J", W: 1},
				{Code: "K", Label: "K", W: 1},
				{Code: "L", Label: "L", W: 1},
				{Code: ";", Label: ";", W: 1},
				{Code: "'", Label: "'", W: 1},
				{Code: "Enter", Label: "⏎", W: 2.25},
				{Code: "None", Label: "", W: 0.2},
				{Code: "None", Label: "", W: 1},
				{Code: "None", Label: "", W: 1},
				{Code: "None", Label: "", W: 1},
				{Code: "None", Label: "", W: 0.2},
				{Code: "NumPad4", Label: "4", W: 1},
				{Code: "NumPad5", Label: "5", W: 1},
				{Code: "NumPad6", Label: "6", W: 1},
				{Code: "NumPad +", Label: "+", W: 1},
			},
			{
				{Code: "Shift", Label: "⇧", W: 2.25},
				{Code: "Z", Label: "Z", W: 1},
				{Code: "X", Label: "X", W: 1},
				{Code: "C", Label: "C", W: 1},
				{Code: "V", Label: "V", W: 1},
				{Code: "B", Label: "B", W: 1},
				{Code: "N", Label: "N", W: 1},
				{Code: "M", Label: "M", W: 1},
				{Code: ",", Label: ",", W: 1},
				{Code: ".", Label: ".", W: 1},
				{Code: "/", Label: "/", W: 1},
				{Code: "RightShift", Label: "⇧", W: 2.75},
				{Code: "None", Label: "", W: 0.2},
				{Code: "None", Label: "", W: 1},
				{Code: "↑", Label: "▲", W: 1},
				{Code: "None", Label: "", W: 1},
				{Code: "None", Label: "", W: 0.2},
				{Code: "NumPad1", Label: "1", W: 1},
				{Code: "NumPad2", Label: "2", W: 1},
				{Code: "NumPad3", Label: "3", W: 1},
				{Code: "NumPad Enter", Label: "⏎", W: 1},
			},
			{
				{Code: "Ctrl", Label: "Fn", W: 1.25},
				{Code: "Win", Label: "⌃", W: 1.25},
				{Code: "Alt", Label: "⌥", W: 1.25},
				{Code: "Space", Label: "", W: 6.25},
				{Code: "RightCtrl", Label: "Fn", W: 1.25},
				{Code: "RightWin", Label: "⌃", W: 1.25},
				{Code: "RightMenu", Label: "⌥", W: 1.25},
				{Code: "RightAlt", Label: "⌥", W: 1.25},
				{Code: "None", Label: "", W: 0.2},
				{Code: "←", Label: "◀", W: 1},
				{Code: "↓", Label: "▼", W: 1},
				{Code: "→", Label: "▶", W: 1},
				{Code: "None", Label: "", W: 0.2},
				{Code: "NumPad0", Label: "0", W: 2},
				{Code: "NumPad .", Label: ".", W: 1},
				{Code: "NumPad Enter", Label: "⏎", W: 1},
			},
		}
	default:
		return Layout{}
	}
}
