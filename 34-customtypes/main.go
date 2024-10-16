// package main

// import (
// 	"fmt"
// 	"math"
// )
// type Point struct{
// 	X float64
// 	Y float64
// }
// func (p Point) Abs() float64 {
// 	return math.Sqrt(p.X*p.X+p.Y*p.Y)
// }

// func (p Point) SetX(newX float64)  {
// 	p.X=newX
// }
// func main() {
// 	p:=Point{1,2}
// 	p.SetX(100)
// 	fmt.Println(p.X) //выводить 1 так как p.SetX работает со значением а не ссылкой
// }
//-------------------------------------------------------------------------------------------------
package main

import (
	"fmt"
	"math"
)
type Point struct{
	X float64
	Y float64
}
// послу того как мы поменя тип получателся на ссылочный тепер наши действие окажет влияние на оригинальную структуру 
func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X+p.Y*p.Y)
}

func (p Point) SetX(newX float64)  {
	p.X=newX
}
func main() {
	p:=Point{1,2}

	p.SetX(100) // здесь у нас неявно будет братся указатель, так как в мы поменяли тип получателя в ссылочный в методе

	fmt.Println(p.X)
}