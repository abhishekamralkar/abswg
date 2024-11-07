package main

import (
	"errors"
	"fmt"
	"log"
)

func Name(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name cannot be empty")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func main() {
	message, err := Name("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
