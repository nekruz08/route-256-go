//like C programming language
// package main

// import "fmt"

// func main() {
// 	var sum int = 0

// 	for i := 0; i < 10; i++ {
// 		if i%2==1{
// 			fmt.Println(i,"is odd")
// 		}else{
// 			fmt.Println(i,"is even")
// 		}

// 		sum+=i
// 	}
// 	fmt.Println("sum is",sum)
// }
//----------------------------------------------------------------------------------------------------------------
// like while
// package main

// import "fmt"

// func main() {
// 	var sum int = 0

// 	i := 0;
// 	for i < 10{
// 		if i%2==1{
// 			fmt.Println(i,"is odd")
// 		}else{
// 			fmt.Println(i,"is even")
// 		}

//			sum+=i
//			i++
//		}
//		fmt.Println("sum is",sum)
//	}
//
// ----------------------------------------------------------------------------------------------------------------
// бесконечный цикл
// package main

// import "fmt"

// func main() {
// 	var sum int = 0

// 	i := 0
// 	for {
// 		if i >= 10 {
// 			break //прерывает цыкл
// 		}

// 		if i%2 == 1 {
// 			fmt.Println(i, "is odd")
// 		} else {
// 			fmt.Println(i, "is even")
// 		}

// 		sum += i
// 		i++
// 	}
// 	fmt.Println("sum is", sum)
// }
// ----------------------------------------------------------------------------------------------------------------
// continue пропускает итерацию когда i==3, и перейти к следующему итерацию
// package main

// import "fmt"

// func main() {
// 	var sum int = 0

// 	for i := 0; i < 10; i++ {
// 		if i==3{
// 			continue //пропускает итерацию когда i==3, и перейти к следующему итерацию
// 		}		
// 		if i%2==1{
// 			fmt.Println(i,"is odd")
// 		}else{
// 			fmt.Println(i,"is even")
// 		}

// 		sum+=i
// 	}
// 	fmt.Println("sum is",sum)
// }
// ----------------------------------------------------------------------------------------------------------------
package main

import "fmt"

func main() {
	var sum int = 0

	outer:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			continue outer
		}		
		if i%2==1{
			fmt.Println(i,"is odd")
		}else{
			fmt.Println(i,"is even")
		}

		sum+=i
	}
	fmt.Println("sum is",sum)
}

