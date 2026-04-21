package main
import"fmt"

type kalimat string //tipe data alias 

type mahasiswa struct { //struct
	nama kalimat //field 
	nim int
	nilai float64
}

func main(){
	var m mahasiswa

	fmt.Print("Masukan Nama Mahasiswa :")
	fmt.Scan(&m.nama)
	fmt.Print("Masukan nim :")
	fmt.Scan(&m.nim)
	fmt.Print("Masukan nilai :")
	fmt.Scan(&m.nilai)

	fmt.Println("DATA MAHASISWA")
	fmt.Println("nama : ", m.nama)
	fmt.Println("nim : ", m.nim)
	fmt.Println("nilai : ", m.nilai)
}