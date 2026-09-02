package main

import (
	"fmt"
	"io"
	"strings"
)

type MyInt int

func doThings(i any) {
	switch j := i.(type) {
	case nil:
		// i is nil, type of j is any
		fmt.Printf("case nil: type of j = %T, value = %v\n", j, j)
	case int:
		// j is bound to type int
		fmt.Printf("case int: type of j = %T, value = %d (math: %d)\n", j, j, j+10)
	case MyInt:
		// j is bound to type MyInt
		fmt.Printf("case MyInt: type of j = %T, value = %d (math: %d)\n", j, j, j*2)
	case io.Reader:
		// j is bound to interface type io.Reader
		fmt.Printf("case io.Reader: dynamic type = %T\n", j)
		buf := make([]byte, 13)
		j.Read(buf) // Reader method is directly accessible on j
		fmt.Printf("  Read output: %s\n", string(buf))
	case string:
		// j is bound to type string
		fmt.Printf("case string: type of j = %T, value = %q (len: %d)\n", j, j, len(j))
	case bool, rune:
		// Multiple types in one case: j remains type 'any'
		fmt.Printf("case bool, rune: static type of j = %T, value = %v\n", j, j)
	default:
		// Fallback: j remains type 'any'
		fmt.Printf("default: static type of j = %T, value = %v\n", j, j)
	}
}

func main() {
	doThings(nil)
	doThings(42)
	doThings(MyInt(100))
	doThings(strings.NewReader("Hello, World!"))
	doThings("Gopher")
	doThings(true)
	doThings('A')
	doThings(3.14159)
}
