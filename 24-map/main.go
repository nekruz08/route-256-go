// package main

// import "fmt"

//	func main() {
//		fmt.Println(`"Hello"`)
//	}
//
// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	var m map[string]int  = make(map[string]int)
// 	m["one"] = 1
// 	fmt.Println(m)
// }
// -------------------------------------------------------------------------------------------------
// give cap for map
package main

import "fmt"

func main() {
	var m map[string]int  = make(map[string]int,10)
	m["one"] = 1
	fmt.Println(len(m))
}
