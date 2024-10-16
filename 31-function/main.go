package main

import "fmt"

func main() {
	m:=map[string]int{}

	// карта тоже является ссылочным типом, передавая ее в функцию можно ее измеиться
	fill(m)

	fmt.Println(m)
}

func fill(m map[string]int)  {
	m["one"]=1
}