package main

import "fmt"

func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * faktorial(n-1)
}

func permutasi(n int, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n int, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Print("Masukan nilai a : ")
	fmt.Scan(&a)

	fmt.Print("Masukan nilai b : ")
	fmt.Scan(&b)

	fmt.Print("Masukan nilai c : ")
	fmt.Scan(&c)

	fmt.Print("Masukan nilai d : ")
	fmt.Scan(&d)

	fmt.Println("hasil permutasi", a, "dan", c, "adalah : ", permutasi(a, c))
	fmt.Println("hasil kombinasi", a, "dan", d, "adalah : ", kombinasi(a, c))
	fmt.Println("hasil permutasi", b, "dan", d, "adalah : ", permutasi(b, d))
	fmt.Println("hasil kombinasi", b, "dan", d, "adalah : ", kombinasi(b, d))
}
