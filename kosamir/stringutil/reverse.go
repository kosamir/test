// Package stringutil contains utility functions for working with strings.
package stringutil

import (
	"fmt"
)

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	fmt.Println("fdsfsfsfsf")
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
		fmt.Printf("%v\n", i)

	}
	return string(r)
}
