package main

import "fmt"

type Point struct{
	X,Y int
}

func main() {
	p:=Point{}
	fill(p)

	fmt.Println(p)
}

// структура тоже как массив, не передается по значению, и не является ссылочным типом, поэтому наши внутри функции получается мы модифицируеи копия структуры
func fill(p Point)  {
	p.X=1
	p.Y=2
}