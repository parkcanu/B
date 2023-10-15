package main

import (
	"fmt"
)

func main() {
	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Printf("A float: %d\n", 15)
	fmt.Printf("A float: %s\n", "hello")
	fmt.Printf("A float: %t\n", false)
	fmt.Printf("A float: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("A float: %v %v %v\n", 1.2, "\t", true)
	fmt.Printf("A float: %#v %#v %#v\n", 1.2, "\t", true)
	fmt.Printf("A float: %%\n")
}
