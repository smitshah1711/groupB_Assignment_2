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
	// Change the value of 'str' to test different strings
	fmt.Println("-------------------------------------------------------------------")
	fmt.Print("3.multiple \n")
	fmt.Print("author name:Nishita Manojbhai Raviya   Student ID:500229175 \n")
	var result int
	result = mul(3, 4)
	fmt.Println("The sum is:", result)
	fmt.Println("-------------------------------------------------------------------")

	// Fibonacci
	fmt.Print("2.Fibonacci  \n")
	fmt.Print("author name:Axay dilipbhai Narigara  student ID:500227623 \n")
	fibonacciLength := 10
	fibonacciResult := fibonacci(fibonacciLength)
	fmt.Printf("Fibonacci Series of length %d: %v\n", fibonacciLength, fibonacciResult)
	fmt.Println("-------------------------------------------------------------------")

	// Change the value of 'year' to test different years
	fmt.Print("4.check whether year is Leap year \n")
	fmt.Print("author name:Vishakhakumari Sureshbhai Patel  Student ID:500226405 \n")
	year := 2023
	leapyear := isLeapYear(year)
	if leapyear {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}

	fmt.Println("-------------------------------------------------------------------")
	//check whether number is even or odd.
	fmt.Print("5.Check whether number is even or odd  \n")
	fmt.Print("author name:Mehak Kapoor  Student ID:500231273 \n")
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	evenOrodd(number)
	fmt.Println("-----------------------------------------------------------------------")

	fmt.Print("7. Reverse String \n")
	fmt.Print("author name:Asfaq Hussain  Student ID:500228114 \n")
	var input string

	fmt.Print("Enter a string: ")
	fmt.Scan(&input)

	reversedString := reverseString(input)

	fmt.Printf("Reversed String: %s\n", reversedString)
	fmt.Println("-----------------------------------------------------------------------")

	fmt.Print("8.check whether the value of variable is greater than or less than 5 \n")
	fmt.Print("author name:Hirenkumar Savani   Student ID:500226947 \n")
	// conditional statements to check whether the value of variable is greater than or less than 5
	var number1 int
	var number2 int

	fmt.Print("Enter a number: ")
	fmt.Scan(&number1)
	fmt.Scan(&number2)
	greaterOrLess(number1, number2)
	fmt.Println("-------------------------------------------------------------------")

	checkSign(10)

	fmt.Print("9. Swap two Numbers: \n")
	fmt.Print("author name:Dhruv Ashvinbhai Jobanputra Student ID:500227799 \n")
	x, y := 3, 7
	fmt.Printf("The initial value of x and y: %d and %d \n", x, y)
	x, y = swap(x, y)
	fmt.Printf("After swapping: x=  %d  and y = %d \n ", x, y)
	fmt.Println("-------------------------------------------------------------------------")

	//author name:Dilpreet Singh Hardoi
	//Student ID:500218797
	//Creating a map
	fmt.Print("10. Mapping \n")
	fmt.Print("author name:Dilpreet Singh Hardoi  Student ID:500218797 \n")
	mapping()
}

// author name:Smit Sandeepkumar Shah
// Student ID:500221322
func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

// author name:Axaykumar Narigara
// Student ID:500227623
func fibonacci(n int) []int {
	fibSeries := make([]int, n)
	fibSeries[0], fibSeries[1] = 0, 1

	for i := 2; i < n; i++ {
		fibSeries[i] = fibSeries[i-1] + fibSeries[i-2]
	}

	return fibSeries
}

// author name:Vishakhakumari Sureshbhai Patel
// Student ID:500226405
func isLeapYear(year int) bool {
	// A leap year is divisible by 4,
	// except for years that are divisible by 100,
	// unless they are also divisible by 400.
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func greaterOrLess(number1 int, number2 int) {
	if number1 > number2 {
		fmt.Printf(" %d is greater than %d \n", number1, number2)
	} else {
		fmt.Printf(" %d is less than %d \n", number1, number2)
	}
}

//author name:Nishita Manojbhai Raviya
//Student ID:500229175

func mul(a, b int) int {
	return a * b
}

// author name:Asfaq Hussain
// Student ID:500228114
func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// author name: Dhruv Jobanputra
// Student ID: 500227799
func swap(a, b int) (int, int) {
	return b, a
}

//author name:Mehak Kapoor
//Student ID:500231273

func evenOrodd(number int) {
	if number%2 == 0 {
		fmt.Printf("%d is an even number.\n", number)
	} else {
		fmt.Printf("%d is an odd number.\n", number)
	}
}

// author name:Akshit Malik
// Student ID:500228578
func checkSign(num int) {
	if num > 0 {
		println("The given number is positive")
	} else if num < 0 {
		println("The given number is negative")
	} else {
		println("The given number is zero")
	}
}

// Auther - Dilpreet Singh Hardoi
// Student ID - 500218797
// This function illustrate the use of map object.
func mapping() {
	grades := map[string]int{

		"Dilpreet": 90,
		"Bob":      85,
		"Charlie":  95,
	}
	fmt.Println("Grades list before updating: ", grades)

	grades["Jarvis"] = 75
	fmt.Println("All Grades after updating: ", grades)
	fmt.Println("The Value of Dilpreet is : ", grades["Dilpreet"])
}
