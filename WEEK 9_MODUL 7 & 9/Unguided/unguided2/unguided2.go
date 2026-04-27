package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah N elemen: ")
	fmt.Scan(&n)

	var arr [1000]int
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Print("a. Keseluruhan isi array: ")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("b. Elemen indeks ganjil: ")
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	fmt.Print("c. Elemen indeks genap: ")
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var x int
	fmt.Print("d. Masukkan nilai x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Printf("   Elemen indeks kelipatan %d: ", x)
	for i := 0; i < n; i++ {
		if x != 0 && i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var hapusIdx int
	fmt.Print("e. Masukkan indeks yang akan dihapus: ")
	fmt.Scan(&hapusIdx)
	if hapusIdx >= 0 && hapusIdx < n {
		for i := hapusIdx; i < n-1; i++ {
			arr[i] = arr[i+1]
		}
		n--
	}
	fmt.Print("   Isi array setelah dihapus: ")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var sum float64
	for i := 0; i < n; i++ {
		sum += float64(arr[i])
	}
	avg := sum / float64(n)
	fmt.Printf("f. Rata-rata: %.2f\n", avg)

	var sumSq float64
	for i := 0; i < n; i++ {
		diff := float64(arr[i]) - avg
		sumSq += diff * diff
	}
	stdDev := math.Sqrt(sumSq / float64(n))
	fmt.Printf("g. Standar deviasi: %.2f\n", stdDev)

	var cariAngka int
	fmt.Print("h. Masukkan bilangan untuk dihitung frekuensinya: ")
	fmt.Scan(&cariAngka)
	var count int
	for i := 0; i < n; i++ {
		if arr[i] == cariAngka {
			count++
		}
	}
	fmt.Printf("   Frekuensi bilangan %d: %d\n", cariAngka, count)
}
