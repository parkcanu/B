package main

import(
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input := reader.ReadString('\n')
	fmt.Println(input)
}

//pass_fail
/*
dddd */
