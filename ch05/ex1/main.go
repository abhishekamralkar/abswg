package main

import "fmt"

func main() {
	evens := [5]int{2, 4, 6, 8, 10}

	var s []int = evens[1:4]
	fmt.Println(s)
}
