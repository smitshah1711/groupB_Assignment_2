package main

import (
	"fmt"
)

func main() {
	// Factorial
	fmt.Print("1.Fictorial -   \n")
	fmt.Print("author name:Smit Sandeepkumar Shah Student ID:500221322 \n")
	num := 5
	factorialResult := factorial(num)
	fmt.Printf("Factorial of %d: %d\n", num, factorialResult)
	fmt.Println("-------------------------------------------------------------------")
}

// author name:Smit Sandeepkumar Shah
// Student ID:500221322
func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

