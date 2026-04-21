package main

import "fmt"

type suhu = float64

func CelciusToReamur(celcius suhu) suhu {
	return celcius * 4 / 5
}

func CelciusToFahrenheit(celcius suhu) suhu {
	return celcius*9/5 + 32
}

func CelciusToKelvin(celcius suhu) suhu {
	return celcius + 273.15
}

func main() {
	var input suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius) : ")
	fmt.Scan(&input)

	fmt.Println()
	fmt.Printf("%.4g celcius = %.4g reamur\n", input, CelciusToReamur(input))
	fmt.Printf("%.4g celcius = %.4g fahrenheit\n", input, CelciusToFahrenheit(input))
	fmt.Printf("%.4g celcius = %.4g kelvin\n", input, CelciusToKelvin(input))
}