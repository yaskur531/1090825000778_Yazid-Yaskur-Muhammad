package main

import "fmt"

func main() {
	correct := [4]string{"merah", "kuning", "hijau", "ungu"}
	berhasil := true

	for i := 1; i <= 5; i++ {
		var w1, w2, w3, w4 string
		fmt.Printf("Percobaan %d : ", i)
		fmt.Scan(&w1, &w2, &w3, &w4)

		input := [4]string{w1, w2, w3, w4}
		if input != correct {
			berhasil = false
		}
	}

	if berhasil {
		fmt.Println("BERHASIL : true")
	} else {
		fmt.Println("BERHASIL : false")
	}
}