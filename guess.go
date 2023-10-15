package main

import (
	"fmt"
	"math/rand"
	"time"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)


func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number betwwen 1 and 100")
	fmt.Println("Can you guess it?")
	fmt.Println(target)

	reader := bufil.NewReader(os.Stdin)

	fmt.Print("Make a guess: ")
	input, err := refer==
}
