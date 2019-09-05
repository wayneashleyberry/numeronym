package numeronym

import (
	"fmt"
	"strings"
	"unicode"
)

// FromString will create a numeronym from the given string
func FromString(s string) string {
	if len(s) < 3 {
		return s
	}
	s = trim(s)
	s = strings.ToLower(s)
	return fmt.Sprintf("%s%d%s", string(s[0]), len(s)-2, string(s[len(s)-1]))
}

func trim(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
