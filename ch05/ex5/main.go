package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11}

	s[1] = 2345
	fmt.Println(s)
}
