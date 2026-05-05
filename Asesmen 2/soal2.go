package main

import "fmt"

const nMax = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func inputData(T *arrayMahasiswa, N *int) {
	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(N)

	if *N > nMax {
		*N = nMax
	}

	for i := 0; i < *N; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&T[i].NIM, &T[i].nama, &T[i].nilai)
	}
}

func nilaiPertama(T arrayMahasiswa, N int, nim string) int {
	for i := 0; i < N; i++ {
		if T[i].NIM == nim {
			return T[i].nilai
		}
	}

	return -1
}

func nilaiTerbesar(T arrayMahasiswa, N int, nim string) int {
	var max int
	ketemu := false

	for i := 0; i < N; i++ {
		if T[i].NIM == nim {
			if !ketemu || T[i].nilai > max {
				max = T[i].nilai
			}
			ketemu = true
		}
	}

	if ketemu {
		return max
	}

	return -1
}

func main() {
	var data arrayMahasiswa
	var N int
	var nimCari string
	var pertama, terbesar int

	inputData(&data, &N)

	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&nimCari)

	pertama = nilaiPertama(data, N, nimCari)
	terbesar = nilaiTerbesar(data, N, nimCari)

	if pertama == -1 {
		fmt.Println("NIM tidak ditemukan")
	} else {
		fmt.Println("Nilai pertama dari NIM", nimCari, "adalah", pertama)
		fmt.Println("Nilai terbesar dari NIM", nimCari, "adalah", terbesar)
	}
}
