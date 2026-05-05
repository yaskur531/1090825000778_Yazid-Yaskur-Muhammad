package main

import "fmt"

const NMAX = 2022

type set [NMAX]int

func exist(T set, n int, val int) bool {
	for i := 0; i < n; i++ {
		if T[i] == val {
			return true
		}
	}
	return false
}

func inputSet(T *set, n *int) {
	var x int
	*n = 0

	for {
		fmt.Scan(&x)

		// stop kalau duplikat atau penuh
		if exist(*T, *n, x) || *n >= NMAX {
			break
		}

		T[*n] = x
		(*n)++
	}
}

func findIntersection(T1, T2 set, n1, n2 int, T3 *set, n3 *int) {
	*n3 = 0

	for i := 0; i < n1; i++ {
		if exist(T2, n2, T1[i]) && !exist(*T3, *n3, T1[i]) {
			T3[*n3] = T1[i]
			(*n3)++
		}
	}
}

func printSet(T set, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(T[i], " ")
	}
	fmt.Println()
}

func main() {
	var s1, s2, s3 set
	var n1, n2, n3 int

	inputSet(&s1, &n1)
	inputSet(&s2, &n2)

	findIntersection(s1, s2, n1, n2, &s3, &n3)

	printSet(s3, n3)
}