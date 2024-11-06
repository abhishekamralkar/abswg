package main

import "fmt"

func main() {
	fmt.Println("Enter your name:")
	var n string
	fmt.Scan(&n)
	switch {
	case n == "Abhishek":
		fmt.Printf("The author is %s\n", n)
	default:
		for i := 0; i < 1000; i++ {
			fmt.Printf("%d. Thanks %s for reading\n", i, n)
		}
	}
}
