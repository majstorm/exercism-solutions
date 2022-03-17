// Package twofer implements simple sharing message
package twofer

import "fmt"

// ShareWith takes name as a parameter and return corresponding message on how to share between.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
