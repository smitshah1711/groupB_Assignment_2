package main

import "fmt"


func evenOrodd(number int) {
	if number%2 == 0 {
		fmt.Printf("%d is an even number.\n", number)
	} else {
		fmt.Printf("%d is an odd number.\n", number)
	}
}
