package main

import (
	"fmt"
)

type Pemain struct {
	NamaDepan    string
	NamaBelakang string
	Gol          int
	Assist       int
}

func main() {
	var n int

	fmt.Println("Masukkan Data Input :")
	fmt.Scan(&n)

	dataPemain := make([]Pemain, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&dataPemain[i].NamaDepan, &dataPemain[i].NamaBelakang, &dataPemain[i].Gol, &dataPemain[i].Assist)
	}

	for i := 0; i < n-1; i++ {
		maxIdx := i

		for j := i + 1; j < n; j++ {
			if dataPemain[j].Gol > dataPemain[maxIdx].Gol {
				maxIdx = j
			} else if dataPemain[j].Gol == dataPemain[maxIdx].Gol {
				if dataPemain[j].Assist > dataPemain[maxIdx].Assist {
					maxIdx = j
				}
			}
		}

		dataPemain[i], dataPemain[maxIdx] = dataPemain[maxIdx], dataPemain[i]
	}

	fmt.Println("\nHasil Sorting :")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %d %d\n", dataPemain[i].NamaDepan, dataPemain[i].NamaBelakang, dataPemain[i].Gol, dataPemain[i].Assist)
	}
}
