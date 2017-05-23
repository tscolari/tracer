package tracer

import "strings"

var closingActions = map[string]bool{
	"done":   true,
	"ending": true,
	"finish": true,
	"end":    true,
}

func isClosing(action string) bool {
	return closingActions[strings.ToLower(action)]
}
