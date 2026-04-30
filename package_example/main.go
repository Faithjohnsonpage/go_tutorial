package main

import (
	"fmt"
	format "package_example/do-format"
	"package_example/math"
)

func main() {
	num := math.Double(2)
	output := format.Number(num)
	fmt.Println(output)
}
