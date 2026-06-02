package main

import "fmt"

type mahasiswa struct {
	nama string
	nim  string
	ipk  float64
}

func selectionSort(mahasiswaList []mahasiswa) {
	n := len(mahasiswaList)

	for i := 0; i < n-1; i++ {
		idxMin := i

		for j := i + 1; j < n; j++ {
			if mahasiswaList[j].ipk < mahasiswaList[idxMin].ipk {
				idxMin = j
			}
		}

		if idxMin != i {
			mahasiswaList[i], mahasiswaList[idxMin] =
				mahasiswaList[idxMin], mahasiswaList[i]
		}
	}
}

func printData(mahasiswaList []mahasiswa, status string) {
	fmt.Printf("\n=== %s ===\n", status)

	for i, mhs := range mahasiswaList {
		fmt.Printf("\n--- Data Mahasiswa Ke-%d ---\n", i+1)
		fmt.Printf("Nama : %s\n", mhs.nama)
		fmt.Printf("NIM  : %s\n", mhs.nim)
		fmt.Printf("IPK  : %.2f\n", mhs.ipk)
	}

	fmt.Println("=======================")
}

func main() {
	structMHS := make([]mahasiswa, 5)

	for i := 0; i < len(structMHS); i++ {
		fmt.Printf("\n--- Masukkan Data Mahasiswa Ke-%d ---\n", i+1)

		fmt.Print("Nama : ")
		fmt.Scan(&structMHS[i].nama)

		fmt.Print("NIM  : ")
		fmt.Scan(&structMHS[i].nim)

		fmt.Print("IPK  : ")
		fmt.Scan(&structMHS[i].ipk)
	}

	printData(structMHS, "SEBELUM SORTING")

	selectionSort(structMHS)

	printData(structMHS, "SETELAH SORTING")
}