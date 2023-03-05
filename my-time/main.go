package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	timezone, _ := time.LoadLocation("America/Sao_Paulo")
	createdDate := time.Date(2021, time.April, 24, 13, 50, 23, 3, timezone)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("02-01-2006 15:04:05 Monday"))
}
