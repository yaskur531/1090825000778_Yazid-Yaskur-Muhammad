package main

import "fmt"

func TahunKabisat(tahun int) bool {
	return tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0)
}

func main() {
	var tahun int
	fmt.Print("Masukkan tahun : ")
	fmt.Scan(&tahun)
	
	if TahunKabisat(tahun) {
		fmt.Println("Kabisat : true")
	} else {
		fmt.Println("Kabisat : false")
	}
}