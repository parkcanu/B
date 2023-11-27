package main

import (
	"fmt"
)

func main() {
	var orignalCount int = 10
	var eatenCount int = 4

	fmt.Println("I started with", orignalCount, "apples.")
	fmt.Println("some jerk ate", eatenCount, "apples.")
	fmt.Println("There are", orignalCount-eatenCount, "apples.")
}
