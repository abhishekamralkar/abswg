package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(100)
	if n%2 == 0 {
		fmt.Printf("Number is even: %d\n", n)
		return
	}
	fmt.Printf("Number is odd: %d\n", n)
}