package main

import (
	"fmt"
	"log"
	"net/http"
)

// 1. Define the custom types
type dollars float32

type database map[string]dollars

// 2. Add methods to those types to satisfy interfaces
// This satisfies the fmt.Stringer interface
func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

// This satisfies the http.Handler interface
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func main() {
	// 3. Initialize the data
	db := database{"shoes": 50, "socks": 5}

	// 4. Start the server
	// We pass 'db' as the handler because it now has the ServeHTTP method
	fmt.Println("Server starting at http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
