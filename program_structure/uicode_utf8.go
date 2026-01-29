package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// "世界" means "World". Each of these characters is 3 bytes.
	// Total length: 7 (Hello, ) + 6 (世界) = 13 bytes.
	s := "Hello, 世界"

	// len(s) returns the number of BYTES
	fmt.Printf("Byte length: %d\n", len(s))

	// RuneCountInString returns the number of UNICODE CODE POINTS
	fmt.Printf("Rune count: %d\n", utf8.RuneCountInString(s))

	fmt.Println("\nIndex\tRune")
	fmt.Println("-----\t----")

	// Manual iteration using DecodeRuneInString
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
}
