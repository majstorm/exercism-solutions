// Package luhn implements luhn checksum algo
package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}
	sum := 0
	id_length := len(id)

	for i := id_length - 1; i >= 0; i-- {
		if num, err := strconv.Atoi(string(id[i])); err == nil {
			if i%2 == (id_length)%2 {
				num *= 2
				if num > 9 {
					num -= 9
				}
			}
			sum += num
		} else {
			return false
		}
	}
	return sum%10 == 0
}
