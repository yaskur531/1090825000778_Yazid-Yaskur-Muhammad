package main

import "fmt"

func main() {
	var totalGram int
	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&totalGram)

	kg := totalGram / 1000
	gram := totalGram % 1000

	biayaKg := kg * 10000
	var biayaGram int

	if totalGram > 10000 {
    biayaGram = 0
} else if gram >= 500 {  
    biayaGram = gram * 5
} else {                  
    biayaGram = gram * 15
}

	total := biayaKg + biayaGram

	fmt.Println()
	fmt.Println("===== Detail Perhitungan =====")
	fmt.Printf("Detail berat  : %d kg + %d gram\n", kg, gram)
	fmt.Printf("Detail biaya  : Rp. %d + Rp. %d\n", biayaKg, biayaGram)
	fmt.Printf("Total biaya   : Rp %d\n", total)
}