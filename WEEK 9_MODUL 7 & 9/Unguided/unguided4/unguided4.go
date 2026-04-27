package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]byte

func isiArray(t *tabel, n *int) {
	var char byte
	*n = 0
	for *n < NMAX {
		fmt.Scanf("%c", &char)
		if char == '.' {
			break
		}
		if char != '\n' && char != '\r' && char != ' ' {
			t[*n] = char
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	var temp tabel = t
	balikanArray(&temp, n)
	for i := 0; i < n; i++ {
		if t[i] != temp[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	isPalindrom := palindrom(tab, m)

	balikanArray(&tab, m)

	cetakArray(tab, m)

	fmt.Println(isPalindrom)
}
