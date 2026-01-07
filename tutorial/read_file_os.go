package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := "data.txt"

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Convert bytes to string to see actual text
	fmt.Println(string(content))
}
