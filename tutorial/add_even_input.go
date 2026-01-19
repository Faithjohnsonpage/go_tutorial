package main

import "fmt"

func main() {
	fmt.Println("Write the number we want to add even numbers between:")

	var num int
	fmt.Scan(&num)

	if num <= 0 {
		fmt.Println("Invalid value, number must be greater than zero")
		return
	}

	sum := 0

	for i := num; i > 0; i-- {
		if i%2 == 0 {
			sum += i
		}
	}

	fmt.Printf("The final value is: %d\n", sum)
}
