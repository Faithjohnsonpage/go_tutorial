package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	// Step 1: Assign *os.File to an io.Writer interface.
	var w io.Writer = os.Stdout

	fmt.Println("--- Before Type Assertion ---")
	fmt.Printf("Interface w holding dynamic type : %T\n", w) // *os.File
	fmt.Printf("Dynamic value inside w           : %v\n", w)

	// w.Read([]byte{}) // ❌ Compile error: w (io.Writer) only has Write()

	// Step 2: Assert to another interface type (io.ReadWriter).
	rw := w.(io.ReadWriter)

	fmt.Println("\n--- After Type Assertion to io.ReadWriter ---")
	fmt.Printf("Interface rw holding dynamic type: %T\n", rw) // Still *os.File!
	fmt.Printf("Dynamic value inside rw          : %v\n", rw)

	// rw now exposes the Read method because the interface type changed to io.ReadWriter
	fmt.Println("\nMethod set expanded: rw can now call both Read() and Write().")

	// Step 3: Failing interface assertion (ByteCounter lacks Read method)
	w = new(ByteCounter)

	fmt.Println("\n--- Swapping w to *ByteCounter ---")
	fmt.Printf("New dynamic type inside w        : %T\n", w)

	rw2, ok := w.(io.ReadWriter)
	fmt.Printf("Does *ByteCounter satisfy io.ReadWriter? %t\n", ok)
	fmt.Printf("Value of rw2 after failed assertion      : %v\n", rw2)
}
