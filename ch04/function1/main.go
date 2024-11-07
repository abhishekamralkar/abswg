package main

import "fmt"

func sayHello(r []string) {
	for _, n := range r {
		fmt.Printf("Hello %s\n", n)
	}
}

func main() {
	rdrs := []string{
		"Anay",
		"Vinu",
		"Vinayaka",
		"Jhampudi",
	}

	sayHello(rdrs)
}
