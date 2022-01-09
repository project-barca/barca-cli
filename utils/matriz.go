package utils

import (
	"fmt"
)

func FoundCommonItems(a []string, b []string) {
	// compare a to be and find common items
	// and store into c
	c := []string{}

	for i := range a {
		if a[i] == b[i] {
			c = append(c, a[i])
		}
	}
	fmt.Println("Ítens em comum : ", c)
	// compare a to be and find item not in a
	// and store into d
	d := []string{}

	for i := range a {
		if a[i] != b[i] {
			d = append(d, a[i])
		}
	}
	fmt.Println("Item que 'a' tem sobre 'b'", d)
	// compare a to be and find item not in b
	// and store into e
	e := []string{}

	for i := range b { // iterate b instead of a
		if a[i] != b[i] {
			e = append(e, b[i])
		}
	}
	fmt.Println("Item que 'b' tem sobre 'a'  ", e)
}
