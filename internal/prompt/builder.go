package prompt

import "fmt"

func BuildKeyboardReportPrompt(
	keyCounts map[string]int,
	timeRange string,
) (string, error) {
	jsonStr, err := formatKeyCounts(keyCounts, timeRange, 50)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(getTemplate("prompt"), jsonStr), nil
}
