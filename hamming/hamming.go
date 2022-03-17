// Package hamming calculates the hamming distance between two DNA strands
package hamming

import "errors"

// Distance function calculates the distance between two DNA strings and returns the value and an error
func Distance(a, b string) (int, error) {
	arune := []rune(a)
	brune := []rune(b)
	if len(arune) != len(brune) {
		return 0, errors.New("incorrect parameters")
	}
	cnt := 0
	for i := range arune {
		if arune[i] != brune[i] {
			cnt++
		}
	}
	return cnt, nil
}
