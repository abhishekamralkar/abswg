package main

import "fmt"

func OddEven(n []int) ([]int, []int) {
	var odd, even []int

	for _, num := range n {
		if num%2 == 0 {
			even = append(even, num)
		} else if num%2 != 0 {
			odd = append(odd, num)
		}
	}
	return odd, even
}

func main() {
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	odd, even := OddEven(num)
	fmt.Println("Odd: ", odd)
	fmt.Println("Even: ", even)
}
