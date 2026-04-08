//go:build darwin

package keymap

type CGKeyCode uint16

/*
 * 来源：Carbon.framework 中的 Events.h
 * 概述：
 *   虚拟键码定义
 *
 * 说明：
 *   这些常量是定义于《Inside Mac Volume V》第 V-191 页的虚拟键码。
 *   它们用于标识键盘上的物理按键。
 *   名称中包含 "ANSI" 的常量，表示基于 ANSI 标准美式键盘的按键位置进行标注。
 *   例如，kVK_ANSI_A 表示美式键盘布局中字母 'A' 所在物理按键对应的虚拟键码。
 *   其他键盘布局可能会将 'A' 键的标签放在不同的物理按键上；
 *   在这种情况下，按下显示为 'A' 的按键将会生成不同的虚拟键码。
 */
const (
	kVK_ANSI_A              CGKeyCode = 0x00
	kVK_ANSI_S              CGKeyCode = 0x01
	kVK_ANSI_D              CGKeyCode = 0x02
	kVK_ANSI_F              CGKeyCode = 0x03
	kVK_ANSI_H              CGKeyCode = 0x04
	kVK_ANSI_G              CGKeyCode = 0x05
	kVK_ANSI_Z              CGKeyCode = 0x06
	kVK_ANSI_X              CGKeyCode = 0x07
	kVK_ANSI_C              CGKeyCode = 0x08
	kVK_ANSI_V              CGKeyCode = 0x09
	kVK_ANSI_B              CGKeyCode = 0x0B
	kVK_ANSI_Q              CGKeyCode = 0x0C
	kVK_ANSI_W              CGKeyCode = 0x0D
	kVK_ANSI_E              CGKeyCode = 0x0E
	kVK_ANSI_R              CGKeyCode = 0x0F
	kVK_ANSI_Y              CGKeyCode = 0x10
	kVK_ANSI_T              CGKeyCode = 0x11
	kVK_ANSI_1              CGKeyCode = 0x12
	kVK_ANSI_2              CGKeyCode = 0x13
	kVK_ANSI_3              CGKeyCode = 0x14
	kVK_ANSI_4              CGKeyCode = 0x15
	kVK_ANSI_6              CGKeyCode = 0x16
	kVK_ANSI_5              CGKeyCode = 0x17
	kVK_ANSI_Equal          CGKeyCode = 0x18
	kVK_ANSI_9              CGKeyCode = 0x19
	kVK_ANSI_7              CGKeyCode = 0x1A
	kVK_ANSI_Minus          CGKeyCode = 0x1B
	kVK_ANSI_8              CGKeyCode = 0x1C
	kVK_ANSI_0              CGKeyCode = 0x1D
	kVK_ANSI_RightBracket   CGKeyCode = 0x1E
	kVK_ANSI_O              CGKeyCode = 0x1F
	kVK_ANSI_U              CGKeyCode = 0x20
	kVK_ANSI_LeftBracket    CGKeyCode = 0x21
	kVK_ANSI_I              CGKeyCode = 0x22
	kVK_ANSI_P              CGKeyCode = 0x23
	kVK_ANSI_L              CGKeyCode = 0x25
	kVK_ANSI_J              CGKeyCode = 0x26
	kVK_ANSI_Quote          CGKeyCode = 0x27
	kVK_ANSI_K              CGKeyCode = 0x28
	kVK_ANSI_Semicolon      CGKeyCode = 0x29
	kVK_ANSI_Backslash      CGKeyCode = 0x2A
	kVK_ANSI_Comma          CGKeyCode = 0x2B
	kVK_ANSI_Slash          CGKeyCode = 0x2C
	kVK_ANSI_N              CGKeyCode = 0x2D
	kVK_ANSI_M              CGKeyCode = 0x2E
	kVK_ANSI_Period         CGKeyCode = 0x2F
	kVK_ANSI_Grave          CGKeyCode = 0x32
	kVK_ANSI_KeypadDecimal  CGKeyCode = 0x41
	kVK_ANSI_KeypadMultiply CGKeyCode = 0x43
	kVK_ANSI_KeypadPlus     CGKeyCode = 0x45
	kVK_ANSI_KeypadClear    CGKeyCode = 0x47
	kVK_ANSI_KeypadDivide   CGKeyCode = 0x4B
	kVK_ANSI_KeypadEnter    CGKeyCode = 0x4C
	kVK_ANSI_KeypadMinus    CGKeyCode = 0x4E
	kVK_ANSI_KeypadEquals   CGKeyCode = 0x51
	kVK_ANSI_Keypad0        CGKeyCode = 0x52
	kVK_ANSI_Keypad1        CGKeyCode = 0x53
	kVK_ANSI_Keypad2        CGKeyCode = 0x54
	kVK_ANSI_Keypad3        CGKeyCode = 0x55
	kVK_ANSI_Keypad4        CGKeyCode = 0x56
	kVK_ANSI_Keypad5        CGKeyCode = 0x57
	kVK_ANSI_Keypad6        CGKeyCode = 0x58
	kVK_ANSI_Keypad7        CGKeyCode = 0x59
	kVK_ANSI_Keypad8        CGKeyCode = 0x5B
	kVK_ANSI_Keypad9        CGKeyCode = 0x5C
	// 独立于键盘布局的按键的按键代码
	kVK_Return        CGKeyCode = 0x24
	kVK_Tab           CGKeyCode = 0x30
	kVK_Space         CGKeyCode = 0x31
	kVK_Delete        CGKeyCode = 0x33
	kVK_Escape        CGKeyCode = 0x35
	kVK_Command       CGKeyCode = 0x37
	kVK_Shift         CGKeyCode = 0x38
	kVK_CapsLock      CGKeyCode = 0x39
	kVK_Option        CGKeyCode = 0x3A
	kVK_Control       CGKeyCode = 0x3B
	kVK_RightCommand  CGKeyCode = 0x36
	kVK_RightShift    CGKeyCode = 0x3C
	kVK_RightOption   CGKeyCode = 0x3D
	kVK_RightControl  CGKeyCode = 0x3E
	kVK_Function      CGKeyCode = 0xb3
	kVK_F17           CGKeyCode = 0x40
	kVK_VolumeUp      CGKeyCode = 0x48
	kVK_VolumeDown    CGKeyCode = 0x49
	kVK_Mute          CGKeyCode = 0x4A
	kVK_F18           CGKeyCode = 0x4F
	kVK_F19           CGKeyCode = 0x50
	kVK_F20           CGKeyCode = 0x5A
	kVK_F5            CGKeyCode = 0x60
	kVK_F6            CGKeyCode = 0x61
	kVK_F7            CGKeyCode = 0x62
	kVK_F3            CGKeyCode = 0x63
	kVK_F8            CGKeyCode = 0x64
	kVK_F9            CGKeyCode = 0x65
	kVK_F11           CGKeyCode = 0x67
	kVK_F13           CGKeyCode = 0x69
	kVK_F16           CGKeyCode = 0x6A
	kVK_F14           CGKeyCode = 0x6B
	kVK_F10           CGKeyCode = 0x6D
	kVK_F12           CGKeyCode = 0x6F
	kVK_F15           CGKeyCode = 0x71
	kVK_Help          CGKeyCode = 0x72
	kVK_Home          CGKeyCode = 0x73
	kVK_PageUp        CGKeyCode = 0x74
	kVK_ForwardDelete CGKeyCode = 0x75
	kVK_F4            CGKeyCode = 0x76
	kVK_End           CGKeyCode = 0x77
	kVK_F2            CGKeyCode = 0x78
	kVK_PageDown      CGKeyCode = 0x79
	kVK_F1            CGKeyCode = 0x7A
	kVK_LeftArrow     CGKeyCode = 0x7B
	kVK_RightArrow    CGKeyCode = 0x7C
	kVK_DownArrow     CGKeyCode = 0x7D
	kVK_UpArrow       CGKeyCode = 0x7E
	// 仅限ISO键盘
	kVK_ISO_Section CGKeyCode = 0x0A
	// 仅限JIS键盘
	kVK_JIS_Yen         CGKeyCode = 0x5D
	kVK_JIS_Underscore  CGKeyCode = 0x5E
	kVK_JIS_KeypadComma CGKeyCode = 0x5F
	kVK_JIS_Eisu        CGKeyCode = 0x66
	kVK_JIS_Kana        CGKeyCode = 0x68
)

type darwinMapper struct {
	codeToName map[CGKeyCode]string
}

func init() {
	mapper := &darwinMapper{
		codeToName: make(map[CGKeyCode]string),
	}
	mapper.codeToName[kVK_ANSI_A] = Key_A
	mapper.codeToName[kVK_ANSI_S] = Key_S
	mapper.codeToName[kVK_ANSI_D] = Key_D
	mapper.codeToName[kVK_ANSI_F] = Key_F
	mapper.codeToName[kVK_ANSI_H] = Key_H
	mapper.codeToName[kVK_ANSI_G] = Key_G
	mapper.codeToName[kVK_ANSI_Z] = Key_Z
	mapper.codeToName[kVK_ANSI_X] = Key_X
	mapper.codeToName[kVK_ANSI_C] = Key_C
	mapper.codeToName[kVK_ANSI_V] = Key_V
	mapper.codeToName[kVK_ANSI_B] = Key_B
	mapper.codeToName[kVK_ANSI_Q] = Key_Q
	mapper.codeToName[kVK_ANSI_W] = Key_W
	mapper.codeToName[kVK_ANSI_E] = Key_E
	mapper.codeToName[kVK_ANSI_R] = Key_R
	mapper.codeToName[kVK_ANSI_Y] = Key_Y
	mapper.codeToName[kVK_ANSI_T] = Key_T
	mapper.codeToName[kVK_ANSI_1] = Key_1
	mapper.codeToName[kVK_ANSI_2] = Key_2
	mapper.codeToName[kVK_ANSI_3] = Key_3
	mapper.codeToName[kVK_ANSI_4] = Key_4
	mapper.codeToName[kVK_ANSI_6] = Key_6
	mapper.codeToName[kVK_ANSI_5] = Key_5
	mapper.codeToName[kVK_ANSI_Equal] = Key_Equal
	mapper.codeToName[kVK_ANSI_9] = Key_9
	mapper.codeToName[kVK_ANSI_7] = Key_7
	mapper.codeToName[kVK_ANSI_Minus] = Key_Minus
	mapper.codeToName[kVK_ANSI_8] = Key_8
	mapper.codeToName[kVK_ANSI_0] = Key_0
	mapper.codeToName[kVK_ANSI_RightBracket] = Key_RightBracket
	mapper.codeToName[kVK_ANSI_O] = Key_O
	mapper.codeToName[kVK_ANSI_U] = Key_U
	mapper.codeToName[kVK_ANSI_LeftBracket] = Key_LeftBracket
	mapper.codeToName[kVK_ANSI_I] = Key_I
	mapper.codeToName[kVK_ANSI_P] = Key_P
	mapper.codeToName[kVK_ANSI_L] = Key_L
	mapper.codeToName[kVK_ANSI_J] = Key_J
	mapper.codeToName[kVK_ANSI_Quote] = Key_Quote
	mapper.codeToName[kVK_ANSI_K] = Key_K
	mapper.codeToName[kVK_ANSI_Semicolon] = Key_Semicolon
	mapper.codeToName[kVK_ANSI_Backslash] = Key_Backslash
	mapper.codeToName[kVK_ANSI_Comma] = Key_Comma
	mapper.codeToName[kVK_ANSI_Slash] = Key_Slash
	mapper.codeToName[kVK_ANSI_N] = Key_N
	mapper.codeToName[kVK_ANSI_M] = Key_M
	mapper.codeToName[kVK_ANSI_Period] = Key_Period
	mapper.codeToName[kVK_ANSI_Grave] = Key_Grave
	mapper.codeToName[kVK_ANSI_KeypadDecimal] = Key_KeypadDecimal
	mapper.codeToName[kVK_ANSI_KeypadMultiply] = Key_KeypadMultiply
	mapper.codeToName[kVK_ANSI_KeypadPlus] = Key_KeypadPlus
	mapper.codeToName[kVK_ANSI_KeypadClear] = Key_KeypadClear
	mapper.codeToName[kVK_ANSI_KeypadDivide] = Key_KeypadDivide
	mapper.codeToName[kVK_ANSI_KeypadEnter] = Key_KeypadEnter
	mapper.codeToName[kVK_ANSI_KeypadMinus] = Key_KeypadMinus
	mapper.codeToName[kVK_ANSI_KeypadEquals] = Key_KeypadEquals
	mapper.codeToName[kVK_ANSI_Keypad0] = Key_Keypad0
	mapper.codeToName[kVK_ANSI_Keypad1] = Key_Keypad1
	mapper.codeToName[kVK_ANSI_Keypad2] = Key_Keypad2
	mapper.codeToName[kVK_ANSI_Keypad3] = Key_Keypad3
	mapper.codeToName[kVK_ANSI_Keypad4] = Key_Keypad4
	mapper.codeToName[kVK_ANSI_Keypad5] = Key_Keypad5
	mapper.codeToName[kVK_ANSI_Keypad6] = Key_Keypad6
	mapper.codeToName[kVK_ANSI_Keypad7] = Key_Keypad7
	mapper.codeToName[kVK_ANSI_Keypad8] = Key_Keypad8
	mapper.codeToName[kVK_ANSI_Keypad9] = Key_Keypad9
	mapper.codeToName[kVK_Return] = Key_Return
	mapper.codeToName[kVK_Tab] = Key_Tab
	mapper.codeToName[kVK_Space] = Key_Space
	mapper.codeToName[kVK_Delete] = Key_Delete
	mapper.codeToName[kVK_Escape] = Key_Escape
	mapper.codeToName[kVK_Command] = Key_Command
	mapper.codeToName[kVK_Shift] = Key_Shift
	mapper.codeToName[kVK_CapsLock] = Key_CapsLock
	mapper.codeToName[kVK_Option] = Key_Option
	mapper.codeToName[kVK_Control] = Key_Control
	mapper.codeToName[kVK_RightCommand] = Key_RightCommand
	mapper.codeToName[kVK_RightShift] = Key_RightShift
	mapper.codeToName[kVK_RightOption] = Key_RightOption
	mapper.codeToName[kVK_RightControl] = Key_RightControl
	mapper.codeToName[kVK_Function] = Key_Function
	mapper.codeToName[kVK_F17] = Key_F17
	mapper.codeToName[kVK_VolumeUp] = Key_VolumeUp
	mapper.codeToName[kVK_VolumeDown] = Key_VolumeDown
	mapper.codeToName[kVK_Mute] = Key_Mute
	mapper.codeToName[kVK_F18] = Key_F18
	mapper.codeToName[kVK_F19] = Key_F19
	mapper.codeToName[kVK_F20] = Key_F20
	mapper.codeToName[kVK_F5] = Key_F5
	mapper.codeToName[kVK_F6] = Key_F6
	mapper.codeToName[kVK_F7] = Key_F7
	mapper.codeToName[kVK_F3] = Key_F3
	mapper.codeToName[kVK_F8] = Key_F8
	mapper.codeToName[kVK_F9] = Key_F9
	mapper.codeToName[kVK_F11] = Key_F11
	mapper.codeToName[kVK_F13] = Key_F13
	mapper.codeToName[kVK_F16] = Key_F16
	mapper.codeToName[kVK_F14] = Key_F14
	mapper.codeToName[kVK_F10] = Key_F10
	mapper.codeToName[kVK_F12] = Key_F12
	mapper.codeToName[kVK_F15] = Key_F15
	mapper.codeToName[kVK_Help] = Key_Help
	mapper.codeToName[kVK_Home] = Key_Home
	mapper.codeToName[kVK_PageUp] = Key_PageUp
	mapper.codeToName[kVK_ForwardDelete] = Key_ForwardDelete
	mapper.codeToName[kVK_F4] = Key_F4
	mapper.codeToName[kVK_End] = Key_End
	mapper.codeToName[kVK_F2] = Key_F2
	mapper.codeToName[kVK_PageDown] = Key_PageDown
	mapper.codeToName[kVK_F1] = Key_F1
	mapper.codeToName[kVK_LeftArrow] = Key_LeftArrow
	mapper.codeToName[kVK_RightArrow] = Key_RightArrow
	mapper.codeToName[kVK_DownArrow] = Key_DownArrow
	mapper.codeToName[kVK_UpArrow] = Key_UpArrow
	mapper.codeToName[kVK_ISO_Section] = Key_ISO_Section
	mapper.codeToName[kVK_JIS_Yen] = Key_JIS_Yen
	mapper.codeToName[kVK_JIS_Underscore] = Key_JIS_Underscore
	mapper.codeToName[kVK_JIS_KeypadComma] = Key_JIS_KeypadComma
	mapper.codeToName[kVK_JIS_Eisu] = Key_JIS_Eisu
	mapper.codeToName[kVK_JIS_Kana] = Key_JIS_Kana
	globalMapper = mapper
}

func (m *darwinMapper) Normalize(code any) string {
	// 支持直接传入 CGKeyCode 类型
	if c, ok := code.(CGKeyCode); ok {
		if name, ok := m.codeToName[c]; ok {
			return name
		}
	}
	// 兼容直接传入 uint16
	if c, ok := code.(uint16); ok {
		if name, ok := m.codeToName[CGKeyCode(c)]; ok {
			return name
		}
	}
	return "Unknown"
}
