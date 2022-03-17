// Package isogram implements checking of uniqueness of characters in a word
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram checks if the word passed as an argument has unique characters in it
func IsIsogram(word string) bool {
	buff := ""

	for _, ch := range strings.ToUpper(word) {
		if unicode.IsLetter(ch) {
			if strings.Contains(buff, string(ch)) {
				return false
			}
			buff += string(ch)
		}
	}
	return true
}
