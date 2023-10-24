package main

import "fmt"

func main() {

	quailty := 4
	length, width := 1.2, 1.4
	customerName := "Chanwoo Park"

	fmt.Println(customerName)
	fmt.Println("has ordered", quailty, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")
}
