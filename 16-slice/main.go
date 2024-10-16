// package main

// import "fmt"

// func main() {
// 	s:=[]int{1,2,3}
// 	fmt.Println(s)
// 	fmt.Println("len(s)",len(s))
// 	fmt.Println("cap(s)",cap(s))
// }
//---------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	c:=10			 	// capasity	
// 	s:=make([]int,c) 	// make(type,capacity)
// 	fmt.Println(s)
// 	fmt.Println("len(s)",len(s))
// 	fmt.Println("cap(s)",cap(s))
// }
//---------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	s:=make([]int,0,10) 	// make(type,length,capacity)
// 	fmt.Println(s)
// 	fmt.Println("len(s)",len(s))
// 	fmt.Println("cap(s)",cap(s))
// }
//---------------------------------------------------------------------------
package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s==nil)
	fmt.Println("len(s)",len(s))
	fmt.Println("cap(s)",cap(s))
}