package main

import "fmt"

const NMAX = 1000000

type partai struct {
	nama  int
	suara int
}

type tabPartai [NMAX]partai

func posisi(t tabPartai, n int, nama int) int {
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1
}

func main() {
	var p tabPartai
	var n int = 0
	var input_partai int

	fmt.Println("Masukkan proses input suara :")
	for {
		fmt.Scan(&input_partai)

		if input_partai == -1 {
			break
		}

		idx := posisi(p, n, input_partai)

		if idx == -1 {
			p[n].nama = input_partai
			p[n].suara = 1
			n++
		} else {
			p[idx].suara++
		}
	}

	for i := 1; i < n; i++ {
		temp := p[i]
		j := i - 1

		for j >= 0 && p[j].suara < temp.suara {
			p[j+1] = p[j]
			j--
		}
		p[j+1] = temp
	}

	fmt.Println("\nHasil Perhitungan suara :")
	for i := 0; i < n; i++ {
		fmt.Printf("%d(%d) ", p[i].nama, p[i].suara)
	}
	fmt.Println()
}
