package main

import "fmt"

func main() {
	fmt.Println("welcome to array in golang")

	var fruitList [4]string
	fruitList[0] = "tomato"
	fruitList[1] = "apple"
	fruitList[3] = "orange"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list length is: ", len(fruitList))

	vegList := [3]string{"beans", "mushroom", "potato"}
	fmt.Println("Vegetable list is: ", vegList)
	fmt.Println("Vegetable list length is: ", len(vegList))
}
