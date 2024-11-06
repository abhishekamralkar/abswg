package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Println("Game: Guess the number between 1 to 100...")
	fmt.Println("you get 3 chances..")

	n := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(n)
	secretNumber := randomizer.Intn(10)
	reader := bufio.NewReader(os.Stdin)

	for try := 1; try <= 3; try++ {
		fmt.Printf("Attempt %d\n", try)
		fmt.Println("Enter your number: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input, please try again.")
			continue
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		if guess < secretNumber {
			fmt.Println("Too low. Try again!")
		} else if guess > secretNumber {
			fmt.Println("Too high. Try again!")
		} else {
			fmt.Printf("Congratulations! You guessed it in %d tries.\n", try)
			break
		}
	}

}
