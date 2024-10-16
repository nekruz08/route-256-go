package main

import "fmt"

type PointWithComments struct{
	Point Point
	Comments []string
}

type Point struct{
	X int
	Y int
}

func main() {
	var p PointWithComments
	fmt.Println(p)
	fmt.Println(p.Point)
	fmt.Println(p.Point.X)
	fmt.Println(p.Comments)
}