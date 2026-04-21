package prompt

import (
	"bytes"
	"encoding/json"
	"sort"
	"strings"
)

type KeyStat struct {
	Key   string  `json:"key"`
	Count int     `json:"count"`
	Ratio float64 `json:"ratio"` // 占比
	Group string  `json:"group"` // 分类
}

type FormattedResult struct {
	Total     int            `json:"total"`
	TopKeys   []KeyStat      `json:"top_keys"`
	Groups    map[string]int `json:"groups"`
	TimeRange string         `json:"time_range"`
}

func formatKeyCounts(
	keyCounts map[string]int,
	timeRange string,
	topN int,
) (string, error) {
	// 计算总数
	total := 0
	for _, v := range keyCounts {
		total += v
	}
	// 转 slice
	var stats []KeyStat
	for k, v := range keyCounts {
		stats = append(stats, KeyStat{
			Key:   k,
			Count: v,
			Group: classifyKey(k),
		})
	}
	// 排序（降序）
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Count > stats[j].Count
	})
	// Top N 截断
	if topN > 0 && len(stats) > topN {
		stats = stats[:topN]
	}
	// 计算占比
	for i := range stats {
		if total > 0 {
			stats[i].Ratio = float64(stats[i].Count) / float64(total)
		}
	}
	// 分组统计
	groupStats := make(map[string]int)
	for k, v := range keyCounts {
		group := classifyKey(k)
		groupStats[group] += v
	}
	// 组装结果
	result := FormattedResult{
		Total:     total,
		TopKeys:   stats,
		Groups:    groupStats,
		TimeRange: timeRange,
	}
	// 美化 JSON
	raw, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	var pretty bytes.Buffer
	err = json.Indent(&pretty, raw, "", "  ")
	if err != nil {
		return "", err
	}
	return pretty.String(), nil
}

func classifyKey(key string) string {
	// 字母
	if len(key) == 1 && key[0] >= 'A' && key[0] <= 'Z' {
		return "letter"
	}
	// 数字
	if len(key) == 1 && key[0] >= '0' && key[0] <= '9' {
		return "number"
	}
	// 小键盘（优先判断）
	if strings.HasPrefix(key, "NumPad") || key == "NumLock" {
		return "numpad"
	}
	// 修饰键（快捷键核心）
	switch key {
	case "Ctrl", "Control", "RightCtrl",
		"Cmd", "Command", "RightCommand",
		"Shift", "RightShift",
		"Alt", "Option", "RightAlt", "RightOption",
		"Fn":
		return "modifier"
	}
	// 控制键（输入控制）
	switch key {
	case "Enter", "Tab", "Space", "Esc", "Escape",
		"Delete", "Del", "Backspace", "CapsLock":
		return "control"
	}
	// 导航键（光标 / 页面）
	switch key {
	case "←", "→", "↑", "↓",
		"Home", "End", "PageUp", "PageDown", "Insert":
		return "navigation"
	}
	// 功能键 F1-F20
	if strings.HasPrefix(key, "F") {
		return "function"
	}
	// 系统键（系统行为）
	switch key {
	case "VolumeUp", "VolumeDown", "Mute",
		"Print", "Scroll", "Pause":
		return "system"
	}
	// 输入法 / 特殊区域（JIS / ISO）
	switch key {
	case "Section", "Yen", "Underscore",
		"KeypadComma", "Eisu", "Kana":
		return "ime"
	}
	// 符号（主键盘）
	switch key {
	case "=", "-", "]", "[", "'", ";", "\\",
		",", "/", ".", "`":
		return "symbol"
	}
	return "other"
}
