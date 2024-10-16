// package main

// import "fmt"

// func main() {
// 	allNewsPosts := []string{
// 		"new title 1",
// 		"new title 2",
// 		"new title 3",
// 		"new title 4",
// 		"new title 5",
// 	}

// 	fmt.Println("<mainpage>")
// 	showMainPage(allNewsPosts[:3]) // [:3:3]
// 	fmt.Println("<mainpage>")

// 	fmt.Println("<all>")
// 	showMainPage(allNewsPosts) // [:3:3]
// 	fmt.Println("<all>")
// }

// func showMainPage(posts []string)  {
// 	postsWithAds:=append(posts, "CLICK HERE TO BUY THIS THING!!!")

// 	showPosts(postsWithAds)
// }

// func showPosts(posts []string)  {
// 	for _, post := range posts {
// 		fmt.Println(post)
// 	}
// }
//-------------------------------------------------------------------------------------------------
package main

import "fmt"

func main() {
	allNewsPosts := []string{
		"new title 1",
		"new title 2",
		"new title 3",
		"new title 4",
		"new title 5",
	}

	fmt.Println("<mainpage>")
	showMainPage(allNewsPosts[:3:3]) // allNewsPosts[:3:3] дает нам первые три элемента [1, 2, 3], и мы не сможем добавить больше элементов в этот срез, чем уже есть.Если вы попытаетесь добавить еще элементы, Go создаст новый срез, который будет содержать дополнительные элементы, так как максимальная длина была достигнута.
	fmt.Println("<mainpage>")

	fmt.Println("<all>")
	showMainPage(allNewsPosts) // [:3:3]
	fmt.Println("<all>")
}

func showMainPage(posts []string)  {
	postsWithAds:=append(posts, "CLICK HERE TO BUY THIS THING!!!")

	showPosts(postsWithAds)
}

func showPosts(posts []string)  {
	for _, post := range posts {
		fmt.Println(post)
	}
}


