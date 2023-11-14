package fetch

import (
	"strings"
)

func IsStrEmpty(s string) bool {
	trimmed := strings.TrimSpace(s)
	return len(trimmed) == 0
}
