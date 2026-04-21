package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	if tujuan == "domestik" {
		if jumlahHari > 3 {
			return 3
		}
		return jumlahHari
	} else {
		if jumlahHari > 8 {
			return 8
		}
		return jumlahHari
	}
}

func biayaPerHari() int {
	return 620000
}

func perhitunganBiaya(jumlahMhs, lama int, tujuan string, totalBiaya *float64) {
	hari := tanggunganHari(lama, tujuan)
	biaya := biayaPerHari()

	total := float64(hari * biaya * jumlahMhs)

	if tujuan == "mancanegara" {
		total = total * 1.5
	}

	*totalBiaya = total
}

func main() {
	var jumlah, lama int
	var tujuan string
	var total float64

	fmt.Print("Masukkan jumlah mahasiswa : ")
	fmt.Scan(&jumlah)

	fmt.Print("Masukkan lama hari study tour : ")
	fmt.Scan(&lama)

	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara) : ")
	fmt.Scan(&tujuan)

	perhitunganBiaya(jumlah, lama, tujuan, &total)

	fmt.Println("Biaya perjalanan yang harus dikeluarkan Tel-U : Rp.", int(total))
}
