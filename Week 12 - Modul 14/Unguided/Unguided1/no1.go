package main

import (
	"bufio"
	"fmt"
	"os"
)

func selectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		idxMin := i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}

		if idxMin != i {
			arr[i], arr[idxMin] = arr[idxMin], arr[i]
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int

	_, err := fmt.Fscan(reader, &n)
	if err != nil || n <= 0 || n >= 1000 {
		return
	}

	for i := 0; i < n; i++ {
		var m int
		fmt.Fscan(reader, &m)

		rumah := make([]int, m)

		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &rumah[j])
		}

		selectionSort(rumah)

		for j, nilai := range rumah {
			if j > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, nilai)
		}
		fmt.Fprintln(writer)
	}
}