// package main

// import "fmt"

// func main() {
// 	fmt.Println(sum(2,3,4))
// }
// func sum(items ...int) int {
// 	result:=0
// 	for _, value := range items {
// 		result+=value
// 	}
// 	return result
// }

// items to mention:
// - multiple argument
// - named arguments
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	s := []int{3, 4}
// 	fmt.Println(sum(2, s...))
// }
// func sum(firstItems int, items ...int) int {
// 	result := firstItems
// 	for _, value := range items {
// 		result += value
// 	}
// 	return result
// }
// -------------------------------------------------------------------------------------------------
// multiple return argument
// package main

// import "fmt"

// func main() {
// 	s := []int{3, 4}
// 	fmt.Println(sum(2, s...))
// }
// func sum(firstItems int, items ...int) (int,bool) {
// 	result := firstItems
// 	for _, value := range items {
// 		result += value
// 	}
// 	return result,true
// }
// -------------------------------------------------------------------------------------------------
// named multiple return argument
// package main

// import "fmt"

// func main() {
// 	s := []int{3, 4}
// 	fmt.Println(sum(2, s...))
// }
// func sum(firstItems int, items ...int) (result int,ok bool) {
// 	result = firstItems
// 	for _, value := range items {
// 		result += value
// 	}
// 	return //можно не указать возвращаемые элементы, так как мы указалаи их в секций возвращаемых значений, в цедом подход не рекомендуется даже если можно так как она уменьшаеь читаемось кода
// }
// -------------------------------------------------------------------------------------------------
// named multiple return argument
package main

import "fmt"

func main() {
	s := []int{3, 4}
	fmt.Println(sum(2, s...))
}
func sum(firstItems int, items ...int) (result int,ok bool) {
	result = firstItems
	for _, value := range items {
		result += value
	}
	return 29, true //если явно указать что возвращаем, то функция возвращает этих значений, и функции пофиг что ты указал в секции возвращаемых значений, result или другое
}