package main

import (
	"fmt"
)

func sortAscending(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func sortDescending(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	var hasilAkhir []string

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		var ganjil []int
		var genap []int

		for j := 0; j < m; j++ {
			var num int
			fmt.Scan(&num)
			if num%2 != 0 {
				ganjil = append(ganjil, num)
			} else {
				genap = append(genap, num)
			}
		}

		sortAscending(ganjil)
		sortDescending(genap)

		barisHasil := ""
		for j, v := range ganjil {
			if j > 0 {
				barisHasil += " "
			}
			barisHasil += fmt.Sprintf("%d", v)
		}

		for _, v := range genap {
			if len(barisHasil) > 0 {
				barisHasil += " "
			}
			barisHasil += fmt.Sprintf("%d", v)
		}

		hasilAkhir = append(hasilAkhir, barisHasil)
	}

	for _, hasil := range hasilAkhir {
		fmt.Println(hasil)
	}
}
