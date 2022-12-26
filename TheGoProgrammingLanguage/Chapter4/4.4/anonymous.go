package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle // this is an anonymous field
	Spokes int
}

// with the anonymuos fields, you can access the field as if you declared it
// directly inside the struct
func main() {
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20

	fmt.Printf("%#v\n", w)
}
