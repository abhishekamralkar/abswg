package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	n := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(n)

	choices := []string{"rock", "paper", "scissors"}

	var userChoice string
	fmt.Print("Enter rock, paper or scissors: ")
	fmt.Scan(&userChoice)

	computerChoice := choices[randomizer.Intn(3)]
	fmt.Println("Computer choice:", computerChoice)

	if computerChoice == userChoice {
		fmt.Println("It's a tie!")
	} else if (userChoice == "rock" && computerChoice == "scissors") ||
		(userChoice == "scissors" && computerChoice == "paper") ||
		(userChoice == "paper" && computerChoice == "rock") {
		fmt.Println("You win!")
	} else {
		fmt.Println("Computer wins!")
	}
}
