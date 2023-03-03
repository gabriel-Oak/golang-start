package main

import "fmt"

const Token = "IJHDUIDUHIDUIHEJIHBGYUPUAÇXNVPJKFEIOJDISDFOIHSF"

func main() {
	fmt.Println("Variables")

	var username string = "gabs"
	fmt.Printf("username is %v \n", username)
	fmt.Printf("username type is %T \n", username)

	var isUser bool = true
	fmt.Printf("isUser is %v \n", isUser)
	fmt.Printf("isUser type is %T \n", isUser)

	var userPoints uint8 = 255
	fmt.Printf("userPoints is %v \n", userPoints)
	fmt.Printf("userPoints type is %T \n", userPoints)

	var userHeight float32 = 1.7857445464654646376
	fmt.Printf("userHeight is %v \n", userHeight)
	fmt.Printf("userHeight type is %T \n", userHeight)

	// default values and some aliases
	var something int
	fmt.Println(something)
	fmt.Printf("Something is type %T \n", something)

	// implicit type
	var host = "pazuzu.com.br"
	fmt.Println(host)
	fmt.Printf("host type is %T \n", host)

	// no var styles
	usersAmount := 43234543
	fmt.Println(usersAmount)

	fmt.Println(Token)
}
