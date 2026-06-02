package main

import "fmt"

func selectionSortArray(angka *[5]int) {
	var idxMin, i, j int

	for i = 0; i < len(angka)-1; i++ { // loop luar (iterasi sorting)
		idxMin = i

		for j = i + 1; j < len(angka); j++ { // loop dalam (mencari nilai minimum)
			if angka[j] <= angka[idxMin] { // ascending
				idxMin = j
			}
		}

		if idxMin != i {
			angka[i], angka[idxMin] = angka[idxMin], angka[i]
		}
	}
}

func main() {
	var arrAngka [5]int

	for i := 0; i < len(arrAngka); i++ {
		fmt.Printf("Masukkan data angka ke-%d : ", i+1)
		fmt.Scan(&arrAngka[i])
	}

	fmt.Println()
	fmt.Println("=== SEBELUM DISORTING ===")
	for i := 0; i < len(arrAngka); i++ {
		fmt.Print(arrAngka[i], " ")
	}

	fmt.Println()
	fmt.Println("=== SETELAH DISORTING ===")

	selectionSortArray(&arrAngka)

	for i := 0; i < len(arrAngka); i++ {
		fmt.Print(arrAngka[i], " ")
	}
	fmt.Println()
}