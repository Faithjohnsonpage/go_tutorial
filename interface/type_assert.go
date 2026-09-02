package main

import "fmt"

type MyInt int

func main() {
	var i any
	var mine MyInt = 20

	// 1. Box 'mine' inside the interface variable 'i'
	i = mine

	// 2. Type assertion: extracts the underlying MyInt value from 'i'
	i2 := i.(MyInt)

	// 3. Perform arithmetic and print (Output: 21)
	fmt.Println(i2 + 1)
}
