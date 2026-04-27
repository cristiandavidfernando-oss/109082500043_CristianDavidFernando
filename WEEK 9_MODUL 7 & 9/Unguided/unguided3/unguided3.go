package main

import "fmt"

const MAX = 1000

type tabString [MAX]string

func rekapPertandingan(klubA string, klubB string, pemenang *tabString, n *int) {
	var skorA, skorB int
	*n = 0

	for {
		fmt.Printf("Pertandingan %d : ", *n+1)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang[*n] = klubA
		} else if skorB > skorA {
			pemenang[*n] = klubB
		} else {
			pemenang[*n] = "Draw"
		}
		*n++
	}
}

func cetakHasil(pemenang tabString, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Hasil %d: %s\n", i+1, pemenang[i])
	}
	fmt.Println("Pertandingan selesai")
}

func main() {
	var klubA, klubB string
	var arrPemenang tabString
	var jumlahPertandingan int

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	rekapPertandingan(klubA, klubB, &arrPemenang, &jumlahPertandingan)
	cetakHasil(arrPemenang, jumlahPertandingan)
}
