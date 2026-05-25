package main
import "fmt"

const NMAX = 1001

type Pemain struct {
	namaDepan  string
	namaBelakang string
	gol        int
	assist     int
}

type arrPemain [NMAX]Pemain

func SelectionSort(T *arrPemain, n int) {
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if T[j].gol > T[maxIdx].gol {
				maxIdx = j
			} else if T[j].gol == T[maxIdx].gol && T[j].assist > T[maxIdx].assist {
				maxIdx = j
			}
		}
		T[i], T[maxIdx] = T[maxIdx], T[i]
	}
}

func main() {
	var n int
	var A arrPemain

	fmt.Println("Masukkan Data Input :")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&A[i].namaDepan, &A[i].namaBelakang, &A[i].gol, &A[i].assist)
	}

	SelectionSort(&A, n)

	fmt.Println("\nHasil Sorting :")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %d %d\n", A[i].namaDepan, A[i].namaBelakang, A[i].gol, A[i].assist)
	}
}