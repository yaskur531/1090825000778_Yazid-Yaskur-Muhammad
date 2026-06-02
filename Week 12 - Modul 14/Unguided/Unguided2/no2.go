package main

import (
	"bufio"
	"fmt"
	"os"
)

func selectionSortAsc(arr []int) {
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

func selectionSortDesc(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		idxMax := i

		for j := i + 1; j < n; j++ {
			if arr[j] > arr[idxMax] {
				idxMax = j
			}
		}

		if idxMax != i {
			arr[i], arr[idxMax] = arr[idxMax], arr[i]
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

		var ganjil []int
		var genap []int

		for j := 0; j < m; j++ {
			var nilai int
			fmt.Fscan(reader, &nilai)

			if nilai%2 == 0 {
				genap = append(genap, nilai)
			} else {
				ganjil = append(ganjil, nilai)
			}
		}

		selectionSortAsc(ganjil)
		selectionSortDesc(genap)

		isFirst := true

		for _, nilai := range ganjil {
			if !isFirst {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, nilai)
			isFirst = false
		}

		for _, nilai := range genap {
			if !isFirst {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, nilai)
			isFirst = false
		}

		fmt.Fprintln(writer)
	}
}