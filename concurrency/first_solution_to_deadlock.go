package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		inGoroutine := 1
		ch1 <- inGoroutine // 1. send to main
		fromMain := <-ch2  // 3. receive from main
		fmt.Println("goroutine:", inGoroutine, fromMain)
	}()

	inMain := 2
	fromGoroutine := <-ch1 // 2. receive from goroutine
	ch2 <- inMain          // 4. send to goroutine
	fmt.Println("main:", inMain, fromGoroutine)
}
