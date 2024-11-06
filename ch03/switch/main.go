package main

import (
	"fmt"
	"time"
)

func swtch() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before Noon")
	default:
		fmt.Println("It's after Noon")
	}

}

func ifelse() {
	t := time.Now()

	if t.Hour() < 12 {
		fmt.Println("It's before Noon")
	} else {
		fmt.Println("It's after Noon")
	}
}

func main() {
	fmt.Println("With Switch")
	swtch()
	fmt.Println("With IfElse")
	ifelse()

}
