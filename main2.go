package main

import (
	"fmt"
)

func main() {
	balls := make(map[string]int) //값을 집어넣을 수 있음
	// var balls map[string]int //값을 집어넣을 수 없음
	fmt.Printf("%#v", balls)
	balls["성기훈"] = 20
	fmt.Println(balls["성기훈"])
	fmt.Println(balls["오일남"])
}
