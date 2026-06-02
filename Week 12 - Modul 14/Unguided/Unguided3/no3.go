package main

import "fmt"

func main() {
	var input int
	var data []int

	for {
		_, err := fmt.Scan(&input)

		if err != nil {
			break
		}

		if input < 0 {
			break
		}

		data = append(data, input)
	}

	if len(data) == 0 {
		return
	}

	insertionSort(data)

	for i, val := range data {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(val)
	}
	fmt.Println()

	statusJarak(data)
}

func insertionSort(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}

func statusJarak(arr []int) {
	n := len(arr)

	if n <= 1 {
		fmt.Println("Data berjarak 0")
		return
	}

	jarakAwal := arr[1] - arr[0]
	isiTetap := true

	for i := 1; i < n-1; i++ {
		if arr[i+1]-arr[i] != jarakAwal {
			isiTetap = false
			break
		}
	}

	if isiTetap {
		fmt.Printf("Data berjarak %d\n", jarakAwal)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}