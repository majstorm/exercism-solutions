// Package scrabble implements simple scrabble rules
package scrabble

import "strings"

var Points = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

// Score function calculates cummulative score for the given word
func Score(word string) int {
	word = strings.ToUpper(word)
	sum := 0
	for _, val := range word {
		for key, str := range Points {
			if strings.Contains(key, string(val)) {
				sum += str
			}
		}
	}
	return sum
}
