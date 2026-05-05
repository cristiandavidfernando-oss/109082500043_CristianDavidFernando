package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const nProv = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func inputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("--- Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ---")

	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		scanner.Scan()
		line := scanner.Text()

		data := strings.Fields(line)

		nama := strings.Join(data[:len(data)-2], " ")
		jumlahPopulasi, _ := strconv.Atoi(data[len(data)-2])
		angkaTumbuh, _ := strconv.ParseFloat(data[len(data)-1], 64)

		prov[i] = nama
		pop[i] = jumlahPopulasi
		tumbuh[i] = angkaTumbuh
	}
}

func provinsiTercepat(tumbuh TumbuhProv) int {
	idxMax := 0

	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[idxMax] {
			idxMax = i
		}
	}

	return idxMax
}

func indeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}

	return -1
}

func prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("--- Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ---")

	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksiPenduduk := int((1 + tumbuh[i]) * float64(pop[i]))
			fmt.Println(prov[i], prediksiPenduduk)
		}
	}
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv
	var cari string

	inputData(&prov, &pop, &tumbuh)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan nama provinsi yang dicari : ")
	scanner.Scan()
	cari = scanner.Text()

	idxCepat := provinsiTercepat(tumbuh)
	idxCari := indeksProvinsi(prov, cari)

	fmt.Println()
	fmt.Println("Provinsi dengan angka pertumbuhan tercepat :", prov[idxCepat])
	fmt.Println()

	if idxCari != -1 {
		fmt.Println("Indeks provinsi yang dicari :", idxCari)
		fmt.Println("Data provinsi yang dicari :", prov[idxCari])
	} else {
		fmt.Println("Provinsi tidak ditemukan")
	}

	fmt.Println()
	prediksi(prov, pop, tumbuh)
}
