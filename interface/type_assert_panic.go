package main

import (
	"fmt"
)

type MyInt int

func main() {
	typeAssertPanicWrongType()
}

func typeAssertPanicWrongType() {
	// we are using recover to allow us to run through the
	// failing type assertions. recover is explained in chapter 9.
	defer func() {
		if m := recover(); m != nil {
			fmt.Println(m) // prints out because a panic happens
		}
	}()
	var i any
	var mine MyInt = 20
	i = mine
	i2 := i.(string)
	fmt.Println(i2)
}
