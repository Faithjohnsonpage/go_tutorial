package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := "data.txt"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// Read the file character by character
	reader := bufio.NewReader(file)
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		fmt.Println(string(char))
	}
}
