package main

import (
	"fmt"
	"math"
)

type Point struct{
	X float64
	Y float64
}

// получатель, receiver, это как неявный аргумен функции
func (p Point) Abs() float64 {
	return math.Sqrt(p.X*p.X+p.Y*p.Y)
}

func main() {
	p:=Point{1,2}
	fmt.Println(p.Abs()) //p. является не явным аргументом функции
	fmt.Println(Point.Abs(p))
	
}