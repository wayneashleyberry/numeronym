package numeronym

import (
	"strconv"
	"strings"
	"unicode"
)

// FromString will create a numeronym from the given string
func FromString(str string) string {
	var last rune
	var setFirst bool
	var count int

	str = strings.ToLower(str)

	var b strings.Builder

	for _, ch := range str {
		if unicode.IsLetter(ch) {
			if setFirst == false {
				setFirst = true
				b.WriteRune(ch)
			}
			last = ch
			count++
		}
	}

	if count < 3 {
		return str
	}

	b.WriteString(strconv.Itoa(count - 2))
	b.WriteRune(last)

	return b.String()
}
