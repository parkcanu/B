package main

import "fmt"

func main() {
	var price int = 100
	fmt.Println("Price is", price, "dollars")

	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println(tax)
}
