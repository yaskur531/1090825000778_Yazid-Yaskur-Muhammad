package main

import "fmt"

const nMax int = 100

type arrString [nMax]string

func insertionSortString(T *arrString, n int) {
	var temp string
	var i, j int

	i = 1
	for i <= n-1 {
		j = i
		temp = T[j]

		// Ascending (A-Z)
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}

		T[j] = temp
		i = i + 1
	}
}

func main() {
	var data arrString
	var n int

	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Data ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	fmt.Println("\nSebelum sorting:")
	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}

	insertionSortString(&data, n)

	fmt.Println("\n\nSetelah sorting (ascending):")
	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()
}