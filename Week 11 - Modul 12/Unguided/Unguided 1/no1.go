package main

import "fmt"

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

	for i := 1; i <= 20; i++ {
		if hitungSuara[i] > 0 {
			fmt.Printf("%d: %d\n", i, hitungSuara[i])
		}
	}
}