package main

import (
	"fmt"
)

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	var arr []int
	var num int

	for {
		_, err := fmt.Scan(&num)
		if err != nil {
			break
		}

		if num < 0 {
			break
		}

		arr = append(arr, num)
	}

	if len(arr) == 0 {
		return
	}

	insertionSort(arr)

	for i, val := range arr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(val)
	}
	fmt.Println()

	if len(arr) < 2 {
		fmt.Println("Data berjarak tidak tetap")
		return
	}

	jarak := arr[1] - arr[0]
	jarakTetap := true

	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != jarak {
			jarakTetap = false
			break
		}
	}

	if jarakTetap {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
