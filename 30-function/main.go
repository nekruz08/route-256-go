// package main

// import "fmt"

// func main() {
// 	// Слайс является ссылочным типом, поэтому, передавая его в функции, можно изменять его значения.
// 	s := []int{1, 2, 3}
// 	fmt.Println("before zerofy", s)
// 	zerofy(s)
// 	fmt.Println("after zerofy", s)

// }

// func zerofy(s []int) {
// 	for i := range s {
// 		s[i] = 0
// 	}
// }

// // - append slice
// // - array
//-------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	s := []int{1, 2, 3}
// 	fmt.Println("before zerofy", s)
// 	zerofy(s)
// 	fmt.Println("after zerofy", s)

// }

// func zerofy(s []int) {
// 	s=append(s, 4)	// происъодить релокация, поскольку у нас капасити не осталось, мы будем менять слайс который ссылается на новый память а не то что на передали в аргументе
// 	for i := range s {
// 		s[i] = 0
// 	}
// }
//-------------------------------------------------------------------------------------------------
package main

import "fmt"

func main() {
	s := [3]int{1, 2, 3}
	fmt.Println("before zerofy", s)
	zerofy(s)
	fmt.Println("after zerofy", s)

}
// массыв не является ссылочным типом, он копируетяся полностью все элементы этого массива, и мы будем модифицировать копию массива который уничтожится
func zerofy(s [3]int) {
	for i := range s {
		s[i] = 0
	}
}