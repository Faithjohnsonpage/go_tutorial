package main

import (
	"fmt"
	"io"
	"os"
)

// ByteCounter implements ONLY io.Writer (Write method).
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	// --- Example 1: Successful assertion to an interface type ---
	var w io.Writer = os.Stdout

	// *os.File implements BOTH Read and Write, so it satisfies io.ReadWriter.
	// The type assertion succeeds and changes the static type to io.ReadWriter,
	// exposing both Read() and Write() methods.
	rw, ok := w.(io.ReadWriter)
	if ok {
		fmt.Printf("Success! Type of rw is io.ReadWriter, holding concrete type: %T\n", rw)
		// rw now gives access to both Read and Write methods
		rw.Write([]byte("Hello to Stdout via rw!\n"))
	}

	// --- Example 2: Failing assertion with recover ---
	w = new(ByteCounter)

	// We recover to gracefully capture the panic when trying to assert ByteCounter to io.ReadWriter.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\nPanicked as expected:", r)
		}
	}()

	fmt.Println("\nAttempting to assert *ByteCounter to io.ReadWriter...")
	// *ByteCounter does NOT have a Read method, so this panics:
	_ = w.(io.ReadWriter)
}
