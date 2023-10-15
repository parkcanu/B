package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.1415))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello, Go!"))
}

// package main

// import "fmt"

// func main() {
// 	fmt.Println('A')
// }

// package main

// import (
// 	"math"
// 	"strings"
// 	"fmt"
// )

// func main() {
// 	fmt.Println(math.Floor(2.75))
// 	fmt.Println(strings.Title("Head first go"))

// }
