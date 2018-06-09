package main

import (
	"fmt"
)

func main() {
	var n int
	var fact = 1
	fmt.Println("Enter a number to calculate its factorial: \n")
	// Scanning an integer input
	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {
		fact = fact * i
	}

	fmt.Println("\nThe factorial of ", n, " is ", fact)
}
