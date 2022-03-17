// Package raindrops is a simple implementation of raindrop sound implementation
package raindrops

import "strconv"

// Convert function transforms input number to corresponding string represented as sound
func Convert(number int) string {
	res := ""
	if number%3 == 0 {
		res += "Pling"
	}
	if number%5 == 0 {
		res += "Plang"
	}
	if number%7 == 0 {
		res += "Plong"
	}
	if res == "" {
		return strconv.Itoa(number)
	}
	return res
}
