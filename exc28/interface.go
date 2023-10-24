package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

type shape interface {
	Area() float64
}

func main() {

	sq := square{length: 2, width: 3}
	circ := circle{radius: 2.5}

	info(sq)
	info(circ)
}

func info(s shape) {
	fmt.Println("The area of the shape is", s.Area())
}

func (sq square) Area() float64 {
	return sq.width * sq.length
}

func (circ circle) Area() float64 {
	return math.Pi * math.Pow(circ.radius, 2)
}
