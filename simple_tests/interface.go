package main

import (
	"fmt"
	"math"
)

type figure interface {
	area() float64
	perimeter() float64
}

type square struct {
	width float64
}

func (sq square) area() float64 {
	return sq.width * sq.width 
}

func (sq square) perimeter() float64 {
	return 4 * sq.width
}

type rectangle struct {
	width, height float64 
}

func (rect rectangle) area() float64 {
	return rect.width * rect.height
}

func (rect rectangle) perimeter() float64 {
	return 2 * (rect.width + rect.height)
}

type circle struct {
	radius float64
}

func (circ circle) area() float64 {
	return 2 * circ.radius * circ.radius
}

func (circ circle) perimeter() float64 {
	return 2 * math.Pi * circ.radius
}

func describeFigure(fig figure) {
	fmt.Println(fig)
	fmt.Printf("Area: %f\n", fig.area())
	fmt.Printf("Perimeter: %f\n", fig.perimeter())
}


func main() {
	sq := square{width: 2}
	rect := rectangle{width: 2, height: 4}
	circ := circle{radius: 3}

	describeFigure(sq)
	describeFigure(rect)
	describeFigure(circ)
}