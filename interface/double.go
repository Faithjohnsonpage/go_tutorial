package main

import "fmt"

type Doubler interface {
	Double()
}
type DoubleInt int

func (d *DoubleInt) Double() {
	*d = *d * 2
}

type DoubleIntSlice []int

func (d DoubleIntSlice) Double() {
	for i := range d {
		d[i] = d[i] * 2
	}
}

func main() {
	// Initialize using your defined types directly
	var a DoubleInt = 4
	c := DoubleIntSlice{1, 2, 3}

	// DoubleInt uses a pointer receiver, so pass &a
	var d Doubler = &a
	d.Double()
	fmt.Println("a after Double():", a)

	// DoubleIntSlice uses a value receiver, so pass c directly
	d = c
	d.Double()
	fmt.Println("c after Double():", c)
}
