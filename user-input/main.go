package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	fmt.Println("Enter the rating for our pizza:")
	reader := bufio.NewReader(os.Stdin)
	rating, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Oh my, something went wrong %v", err.Error())
		return
	}

	fmt.Println("Thanks for rating:", rating)
	fmt.Printf("The type of this rating is %T\n", rating)
}
