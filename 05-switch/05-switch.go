// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	var input = readLine("Enter yes or no: ")
	
// 	switch input {
// 	case "yes","да":
// 		fmt.Println("You've agreed")
// 	case "no","нет":	
// 		fmt.Println("You've disagreed")
// 	default:
// 		fmt.Println("I don't understand")	
// 	}
// }

// func readLine(greeting string) string {
// 	fmt.Println(greeting)
// 	reader:=bufio.NewReader(os.Stdin)
// 	text,_,_:=reader.ReadLine()
// 	return string(text)
// }
//-------------------------------------------------------------------
// использование fallthrough
package main

import (
    "fmt"
)

func main() {
    day := 3

    switch day {
    case 1:
        fmt.Println("Понедельник")
    case 2:
        fmt.Println("Вторник")
    case 3:
        fmt.Println("Среда")
        fallthrough		// позволяет продолжить выполнение следующего случая в switch, даже если условие не соответствует
    case 4:
        fmt.Println("Четверг")
    case 5:
        fmt.Println("Пятница")
    default:
        fmt.Println("Выходные")
    }
}
