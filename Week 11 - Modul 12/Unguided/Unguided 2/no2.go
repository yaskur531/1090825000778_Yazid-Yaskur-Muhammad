package main

import "fmt"

func cariKetuaWakil(hitungSuara [21]int) (int, int) {
	// Cari suara terbanyak pertama dan kedua
	// Jika sama, pilih nomor peserta terkecil
	var ketua int = -1
	var wakil int = -1

	for i := 1; i <= 20; i++ {
		if hitungSuara[i] > 0 {
			if ketua == -1 {
				ketua = i
			} else if hitungSuara[i] > hitungSuara[ketua] {
				wakil = ketua
				ketua = i
			} else if wakil == -1 {
				wakil = i
			} else if hitungSuara[i] > hitungSuara[wakil] {
				wakil = i
			}
		}
	}
	return ketua, wakil
}

func main() {
	var suara int
	var totalMasuk int = 0
	var totalSah int = 0
	var hitungSuara [21]int

	for {
		fmt.Scan(&suara)
		if suara == 0 {
			break
		}
		totalMasuk++
		if suara >= 1 && suara <= 20 {
			totalSah++
			hitungSuara[suara]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalMasuk)
	fmt.Printf("Suara sah: %d\n", totalSah)

	ketua, wakil := cariKetuaWakil(hitungSuara)
	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}