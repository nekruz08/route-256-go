package main

import "fmt"

func main() {
	// var a int // int8, uint8, int16, int32, uint32 int64, byte, rune
	// fmt.Println(a)

	// var f float64 	// float32
	// fmt.Println(f)

	// var b bool		// false, true
	// fmt.Println(b)

	// var s string
	// fmt.Println(s)

	// const (
	// 	a = iota
	// 	b
	// 	c
	// )
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	const (
		a = 1 <<iota
		b
		c
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}