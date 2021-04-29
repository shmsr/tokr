package tokr

import (
	"strings"
	"unicode"
)

func tokenize(text string) document {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}
