package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	batas := 10

	fmt.Printf("Deret Fibonacci hingga suku ke-%d:\n", batas)
	fmt.Println("--------------------------------")
	fmt.Printf("%-5s | %-5s\n", "n", "Sn")
	fmt.Println("--------------------------------")

	for i := 0; i <= batas; i++ {
		fmt.Printf("%-5d | %-5d\n", i, fibonacci(i))
	}
	fmt.Println("--------------------------------")
}
