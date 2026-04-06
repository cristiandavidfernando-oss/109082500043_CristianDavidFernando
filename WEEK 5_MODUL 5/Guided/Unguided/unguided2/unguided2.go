package main

import (
	"fmt"
)

func cetakFaktor(n int, current int) {
	if current > n {
		return
	}

	if n%current == 0 {
		fmt.Printf("%d ", current)
	}

	cetakFaktor(n, current+1)
}

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Harap masukkan bilangan bulat positif.")
		return
	}

	fmt.Print("Keluaran: ")
	cetakFaktor(n, 1)
	fmt.Println()
}
