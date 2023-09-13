package main

import (
	"fmt"

	"stnc/controller"
)

func main() {
	fmt.Println("一天就學會 Go 語言")

	hi := controller.HelloWorld("appleboy")
	fmt.Println(hi)
}
