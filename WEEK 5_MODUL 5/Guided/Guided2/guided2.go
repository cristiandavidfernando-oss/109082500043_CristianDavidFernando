package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(penjumlah(n))
}

func penjumlah(bilangan int) int {
	if bilangan == 1 {
		return 1
	} else {
		return bilangan + penjumlah(bilangan-1)
	}
}
