package main

import "fmt"

func diffReturn(n string, a int) (string, int) {
	return n, a
}

func main() {
	name, age := diffReturn("Anay", 9)

	fmt.Printf("%s is now %d\n", name, age)
}
