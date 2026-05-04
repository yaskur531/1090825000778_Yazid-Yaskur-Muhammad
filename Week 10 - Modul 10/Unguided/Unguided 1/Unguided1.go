package main

import "fmt"

const NMAX = 1000

type arrKelinci [NMAX]float64

func minBerat(tab arrKelinci, n int) float64 {
	var min float64 = tab[0]
	var i int = 1
	for i < n {
		if tab[i] < min {
			min = tab[i]
		}
		i = i + 1
	}
	return min
}

func maxBerat(tab arrKelinci, n int) float64 {
	var max float64 = tab[0]
	var i int = 1
	for i < n {
		if tab[i] > max {
			max = tab[i]
		}
		i = i + 1
	}
	return max
}

func main() {
	var arr arrKelinci
	var n int

	fmt.Scan(&n)

	var i int = 0
	for i < n {
		fmt.Scan(&arr[i])
		i = i + 1
	}

	fmt.Printf("%.2f\n", minBerat(arr, n))
	fmt.Printf("%.2f\n", maxBerat(arr, n))
}