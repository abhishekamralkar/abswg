package main

import "fmt"

func pickUp(toys []string) []string {
	for i, _ := range toys {
		toys[i] = ""
	}
	return toys
}

func main() {
	toys := []string{"Winnie the Pooh", "ring tower", "wooden blocks"}
	done := pickUp(toys)
	fmt.Printf("All done: %v", done)
}
