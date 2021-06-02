package main

import "fmt"

// Structs and interfaces are our main building blocks. No inheritance, only composition.
// Convention:
// - lowercase - internal for a package, to a struct for a field
// - uppercase - exported for a package, public for a struct.
// Suggested: https://gobyexample.com/structs
type Square struct {
	Side float64
}

// Defined a method for a struct.
func (s Square) Area() float64 {
	return s.Side * s.Side
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (s Rectangle) Area() float64 {
	return s.Height * s.Width
}

// Interfaces are structural. Defined where used, not where implemented.
// We will talk a bit more about interfaces next time.
type shape interface {
	Area() float64
}

func sumArea(shapes []shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}
	return area
}

func main() {
	shapes := []shape{
		Rectangle{Width: 3.0, Height: 2.0},
		Square{Side: 2.0},
	}
	fmt.Printf("The total area is: %f\n", sumArea(shapes))
}
