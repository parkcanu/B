package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:]) // Args == Argument
	fmt.Println(os.Args[2])
}
