//go:build !darwin

package keymap

type stubMapper struct{}

func init() {
	globalMapper = &stubMapper{}
}

func (m *stubMapper) Normalize(code any) string {
	return "Unknown"
}

func GetKeyLayout(l LayoutType) Layout {
	return Layout{}
}
