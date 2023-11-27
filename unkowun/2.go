package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "F!!K !OU"
	r := strings.NewReplacer("!", "U")
	fmt.Println(str)
	new := r.Replace(str)
	fmt.Println(new)
}
