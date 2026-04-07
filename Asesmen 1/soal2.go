package main

import (
	"fmt"
)

func tanggunganHari(jumlahHari int, tujuan string) int {

	var maksHari int
	if tujuan == "domestik" {
		maksHari = 3
	} else if tujuan == "mancanegara" {
		maksHari = 8
	}

	if jumlahHari > maksHari {
		return maksHari
	}

	return jumlahHari
}

func biayaPerHari(jumlahMhs int) int {
	biayaPerMhs := 620000

	return jumlahMhs * biayaPerMhs
}

func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *float64) {
	hariDitanggung := tanggunganHari(lamaPerjalanan, tujuan)

	biayaDomestikHarian := biayaPerHari(jumlahMhs)

	if tujuan == "domestik" {
		*totalBiaya = float64(hariDitanggung * biayaDomestikHarian)
	} else if tujuan == "mancanegara" {
		*totalBiaya = float64(hariDitanggung) * float64(biayaDomestikHarian) * 1.5
	}
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	fmt.Print("Masukkan jumlah mahasiswa : ")
	fmt.Scanln(&jumlah)
	fmt.Print("Masukkan lama hari study tour : ")
	fmt.Scanln(&lama)
	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara) : ")
	fmt.Scanln(&tujuan)
	fmt.Println()

	perhitunganBiaya(jumlah, lama, tujuan, &biaya)

	fmt.Printf("Biaya perjalanan yang harus dikeluarkan Tel-U : Rp. %.0f\n", biaya)
}
