package main

import "fmt"

const nMax int = 51

type mahasiswa struct {
	nim   int
	nama  string
	nilai int
}

type arrMahasiswa [nMax]mahasiswa

func cariNilaiPertama(data arrMahasiswa, n int, nimCari int) int {
	for i := 0; i < n; i++ {
		if data[i].nim == nimCari {
			return data[i].nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(data arrMahasiswa, n int, nimCari int) int {
	maxNilai := -1

	for i := 0; i < n; i++ {
		if data[i].nim == nimCari {
			if data[i].nilai > maxNilai {
				maxNilai = data[i].nilai
			}
		}
	}

	return maxNilai
}

func main() {
	var data arrMahasiswa
	var n, nimCari int

	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data ke-%d: ", i+1)
		fmt.Scan(&data[i].nim, &data[i].nama, &data[i].nilai)
	}

	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari: ")
	fmt.Scan(&nimCari)

	nilaiPertama := cariNilaiPertama(data, n, nimCari)
	nilaiTerbesar := cariNilaiTerbesar(data, n, nimCari)

	if nilaiPertama == -1 {
		fmt.Println("NIM tidak ditemukan")
	} else {
		fmt.Println("Nilai pertama dari NIM", nimCari, "adalah", nilaiPertama)
		fmt.Println("Nilai terbesar dari NIM", nimCari, "adalah", nilaiTerbesar)
	}
}