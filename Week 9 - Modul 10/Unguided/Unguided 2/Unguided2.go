package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	berat := make([]float64, x)
	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := int(math.Ceil(float64(x) / float64(y)))

	totalPerWadah := make([]float64, jumlahWadah)
	grandTotal := 0.0

	for w := 0; w < jumlahWadah; w++ {
		mulai := w * y
		akhir := mulai + y

		if akhir > x {
			akhir = x
		}

		for i := mulai; i < akhir; i++ {
			totalPerWadah[w] += berat[i]
		}

		grandTotal += totalPerWadah[w]
	}

	for w := 0; w < jumlahWadah; w++ {
		if w > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%.2f", totalPerWadah[w])
	}
	fmt.Println()

	rataRata := grandTotal / float64(jumlahWadah)
	fmt.Printf("%.2f\n", rataRata)
}