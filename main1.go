package main

import (
	"fmt"
)

func main() {
	//1
	// var games map[int]string
	// games = make(map[int]string)
	//append
	//2
	// games := make(map[int]string)
	//3
	// games := map[int]string{456: "성기훈", 218: "박해수", 065: "강새벽", 001: "오일남", 199: "알리"}
	//3.replace
	games := map[int]string{
		456: "성기훈",
		218: "박해수",
		065: "강새벽",
		001: "오일남",
		199: "알리",
	}

	// games[456] = "성기훈"
	// games[218] = "박해수"
	// games[065] = "강새벽"
	// games[001] = "오일남"
	// games[199] = "오일이"
	// games[101] = "알리"

	fmt.Println(games[065])

	for _, v := range games {
		fmt.Println(v)
	}

	//update
	games[101] = "장덕수"
	//delete
	delete(games, 199)
	for k, v := range games {
		fmt.Println(k, v)
	}
}
