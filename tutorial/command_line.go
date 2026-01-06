package main

import (
	"fmt"
	"os"
)

func main() {
	var s, seg string
	for i := 1; i < len(os.Args); i++ {
		s += seg + os.Args[i]
		seg = " "
	}

	fmt.Println(s)
}
