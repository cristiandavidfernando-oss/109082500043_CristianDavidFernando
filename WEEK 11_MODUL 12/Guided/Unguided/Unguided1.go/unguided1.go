package main

import "fmt"

func main() {
	var suara int
	var totalMasuk, suaraSah int
	var hasil [21]int

	for {
		fmt.Scan(&suara)

		if suara == 0 {
			break
		}

		totalMasuk++

		if suara >= 1 && suara <= 20 {
			suaraSah++
			hasil[suara]++
		}
	}

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", suaraSah)

	for i := 1; i <= 20; i++ {
		if hasil[i] > 0 {
			fmt.Printf("%d: %d\n", i, hasil[i])
		}
	}
}
