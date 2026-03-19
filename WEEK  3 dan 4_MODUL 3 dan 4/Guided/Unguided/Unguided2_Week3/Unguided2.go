package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var x, y, z int

	fmt.Print("Masukan nilai x : ")
	fmt.Scan(&x)

	fmt.Print("Masukan nilai y : ")
	fmt.Scan(&y)

	fmt.Print("Masukan nilai z : ")
	fmt.Scan(&z)

	hasil1 := f(g(h(x)))
	hasil2 := g(h(f(y)))
	hasil3 := h(f(g(z)))

	fmt.Println("f(g(h(", x, "))) : ", hasil1)
	fmt.Println("g(h(f(", y, "))) : ", hasil2)
	fmt.Println("h(f(g(", z, "))) : ", hasil3)
}
