package main

import "fmt"

type suhu float64

func CelciusToReamur(celcius suhu) suhu {
	return celcius * 4 / 5
}

func CelciusToFahrenheit(celcius suhu) suhu {
	return (celcius * 9 / 5) + 32
}

func CelciusToKelvin(celcius suhu) suhu {
	return celcius + 273.15
}

func main() {
	var celcius suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius) : ")
	fmt.Scan(&celcius)

	fmt.Printf("%.2f celcius = %.2f reamur\n", celcius, CelciusToReamur(celcius))
	fmt.Printf("%.2f celcius = %.2f fahrenheit\n", celcius, CelciusToFahrenheit(celcius))
	fmt.Printf("%.2f celcius = %.2f kelvin\n", celcius, CelciusToKelvin(celcius))
}
