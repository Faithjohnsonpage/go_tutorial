package main

import (
	"fmt"
)

type MyInt int

func main() {
	err := typeAssertCommaOK()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func typeAssertCommaOK() error {
	var i any
	var mine MyInt = 20
	i = mine
	i2, ok := i.(int)
	if !ok {
		// we are constructing a new error with fmt.Errorf.
		// fmt.Errorf is covered in chapter 9.
		return fmt.Errorf("unexpected type for %v", i)
	}
	fmt.Println(i2 + 1)
	return nil
}
