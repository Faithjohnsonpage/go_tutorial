// Package format provides functions for converting data types into formatted strings.
package format

import "fmt"

// Number takes an integer and returns a string formatted as "The number is %d".
func Number(num int) string {
	return fmt.Sprintf("The number is %d", num)
}
