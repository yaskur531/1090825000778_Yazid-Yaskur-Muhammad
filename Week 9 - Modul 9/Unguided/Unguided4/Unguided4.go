package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimRight(input, "\r\n")

	*n = 0
	for _, ch := range input {
		if ch == '.' {
			break
		}
		if ch != ' ' {
			t[*n] = ch
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t tabel, n int) tabel {
	var hasil tabel
	for i := 0; i < n; i++ {
		hasil[i] = t[n-1-i]
	}
	return hasil
}

func palindrom(t tabel, n int) bool {
	balik := balikanArray(t, n)

	for i := 0; i < n; i++ {
		if t[i] != balik[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	isiArray(&tab, &n)
	cetakArray(tab, n)

	if palindrom(tab, n) {
		fmt.Println("Palindrom")
	} else {
		fmt.Println("Bukan palindrom")
	}
}