package main

import "fmt"

func main() {
	var angka int
	var totalSuara, suaraSah int
	var suara [21]int

	for {
		fmt.Scan(&angka)

		if angka == 0 {
			break
		}

		totalSuara++

		if angka >= 1 && angka <= 20 {
			suara[angka]++
			suaraSah++
		}
	}

	ketua := 1
	wakil := 1

	for i := 1; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			ketua = i
		}
	}

	for i := 1; i <= 20; i++ {
		if i != ketua {
			if wakil == ketua || suara[i] > suara[wakil] {
				wakil = i
			}
		}
	}

	fmt.Println("Suara masuk:", totalSuara)
	fmt.Println("Suara sah:", suaraSah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}
