package main

import (
	"fmt"
)

func main() {
	var price int = 100
	fmt.Println("Price is", price, "dollars")
	fprice := float64(price)

	var taxRate float64 = 0.08
	var tax float64 = fprice * taxRate
	fmt.Println("tax is", tax, "dollars")

	total := fprice + tax
	fmt.Println("total is", total, "dollars")
}
