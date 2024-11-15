package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func isValidEmail(email string) bool {

	//const emailReg = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	const emailReg = `[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,6}`

	re, _ := regexp.Compile(emailReg)

	return re.MatchString(email)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your email address:")

	input, _ := reader.ReadString('\n')

	email := strings.TrimSpace(input)

	if isValidEmail(email) {
		fmt.Println("Valid email address")
	} else {
		fmt.Println("Invalid email address")
	}
}
