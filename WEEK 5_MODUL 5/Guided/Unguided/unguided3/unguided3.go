package main

import (
	"fmt"
)

func cetakGanjil(n int, current int) {
	if current > n {
		return
	}

	if current%2 != 0 {
		fmt.Printf("%d ", current)
	}

	cetakGanjil(n, current+1)
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
	cetakGanjil(n, 1)
	fmt.Println()
}
