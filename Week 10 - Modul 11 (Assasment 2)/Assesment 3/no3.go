package main

import "fmt"

const nProv = 10

type NamaProv   [nProv]string
type PopProv    [nProv]int
type TumbuhProv [nProv]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d (nama populasi pertumbuhan): ", i+1)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	idxMax := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[idxMax] {
			idxMax = i
		}
	}
	return idxMax
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("\n=== Prediksi Jumlah Penduduk Tahun Depan ===")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksi := int((1 + tumbuh[i]) * float64(pop[i]))
			fmt.Println(prov[i], prediksi)
		}
	}
}

func main() {
	var prov   NamaProv
	var pop    PopProv
	var tumbuh TumbuhProv
	var cari   string

	InputData(&prov, &pop, &tumbuh)

	idxCepat := ProvinsiTercepat(tumbuh)
	fmt.Println("\nProvinsi dengan angka pertumbuhan tercepat:", prov[idxCepat])

	fmt.Print("\nMasukkan nama provinsi yang dicari: ")
	fmt.Scan(&cari)

	idx := IndeksProvinsi(prov, cari)
	if idx != -1 {
		fmt.Println("Provinsi ditemukan di index:", idx)
	} else {
		fmt.Println("Provinsi tidak ditemukan")
	}

	Prediksi(prov, pop, tumbuh)
}