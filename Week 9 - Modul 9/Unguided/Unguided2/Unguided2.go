package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println(arr)

	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var x int
	fmt.Scan(&x)
	if x != 0 {
		for i := 0; i < n; i++ {
			if i%x == 0 {
				fmt.Print(arr[i], " ")
			}
		}
	}
	fmt.Println()

	var idx int
	fmt.Scan(&idx)
	if idx >= 0 && idx < len(arr) {
		arr = append(arr[:idx], arr[idx+1:]...)
	}
	fmt.Println(arr)

	if len(arr) > 0 {
		sum := 0
		for _, v := range arr {
			sum += v
		}
		mean := float64(sum) / float64(len(arr))
		fmt.Println(mean)

		var total float64
		for _, v := range arr {
			total += math.Pow(float64(v)-mean, 2)
		}
		fmt.Println(math.Sqrt(total / float64(len(arr))))
	} else {
		fmt.Println(0)
		fmt.Println(0)
	}

	var cari, count int
	fmt.Scan(&cari)
	for _, v := range arr {
		if v == cari {
			count++
		}
	}
	fmt.Println(count)
}