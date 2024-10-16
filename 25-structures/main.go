package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	p := Point{X: 1}
	fmt.Println(p)				//{1 0}
	fmt.Printf("%+v\n",p)		//{X:1 Y:0}
	fmt.Printf("%#+v\n",p)		//main.Point{X:1, Y:0}


}
