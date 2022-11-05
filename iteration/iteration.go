package iteration

import (
	"strings"
)

const REPEAT_COUNT = 5

func Repeat(st string, count int) string {
	var builder strings.Builder
	var repeatCount int
	if count == 0 {
		repeatCount = REPEAT_COUNT
	} else {
		repeatCount = count
	}

	for i := 0; i < repeatCount; i++ {
		builder.WriteString(st)
	}
	return builder.String()
}
