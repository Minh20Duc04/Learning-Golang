package main

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	Area() float64      // Diện tích
	Perimeter() float64 // Chu vi
}

// Circle struct
type Circle struct {
	radius float64
}

// Rectangle 
type Rectangle struct {
	length, width float64
}

// Triển khai các phương thức của Circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Triển khai các phương thức của Rectangle
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func main() {
	// Sử dụng Shape interface
	var s Shape

	// Test với Circle
	c := Circle{radius: 5}
	s = c
	fmt.Println("Circle:")
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())

	// Test với Rectangle
	r := Rectangle{length: 4, width: 3}
	s = r
	fmt.Println("Rectangle:")
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}
