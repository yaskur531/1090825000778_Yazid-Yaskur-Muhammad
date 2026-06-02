package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func main() {
	var pustaka DaftarBuku
	var nPustaka int
	var ratingDicari int

	reader := bufio.NewReader(os.Stdin)

	fmt.Scanln(&nPustaka)

	BacaDaftarBuku(&pustaka, nPustaka, reader)

	inputRating, _ := reader.ReadString('\n')
	inputRating = strings.TrimSpace(inputRating)
	ratingDicari, _ = strconv.Atoi(inputRating)

	CetakTerfavorit(pustaka, nPustaka)

	UrutBuku(&pustaka, nPustaka)

	Cetak5Terbaru(pustaka, nPustaka)

	CariBuku(pustaka, nPustaka, ratingDicari)
}

func BacaDaftarBuku(pustaka *DaftarBuku, n int, reader *bufio.Reader) {
	for i := 0; i < n; i++ {
		var b Buku

		b.id, _ = reader.ReadString('\n')
		b.id = strings.TrimSpace(b.id)

		b.judul, _ = reader.ReadString('\n')
		b.judul = strings.TrimSpace(b.judul)

		b.penulis, _ = reader.ReadString('\n')
		b.penulis = strings.TrimSpace(b.penulis)

		b.penerbit, _ = reader.ReadString('\n')
		b.penerbit = strings.TrimSpace(b.penerbit)

		txtEksemplar, _ := reader.ReadString('\n')
		b.eksemplar, _ = strconv.Atoi(strings.TrimSpace(txtEksemplar))

		txtTahun, _ := reader.ReadString('\n')
		b.tahun, _ = strconv.Atoi(strings.TrimSpace(txtTahun))

		txtRating, _ := reader.ReadString('\n')
		b.rating, _ = strconv.Atoi(strings.TrimSpace(txtRating))

		pustaka[i] = b
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		return
	}

	maxIdx := 0

	for i := 1; i < n; i++ {
		if pustaka[i].rating > pustaka[maxIdx].rating {
			maxIdx = i
		}
	}

	fav := pustaka[maxIdx]

	fmt.Printf("%s, %s, %s, %d\n",
		fav.judul,
		fav.penulis,
		fav.penerbit,
		fav.tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1

		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}

		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	limit := 5

	if n < 5 {
		limit = n
	}

	var judulList []string

	for i := 0; i < limit; i++ {
		judulList = append(judulList, pustaka[i].judul)
	}

	fmt.Println(strings.Join(judulList, " "))
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	left := 0
	right := n - 1
	foundIdx := -1

	for left <= right {
		mid := left + (right-left)/2

		if pustaka[mid].rating == r {
			foundIdx = mid
			break
		} else if pustaka[mid].rating < r {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if foundIdx != -1 {
		b := pustaka[foundIdx]

		fmt.Printf("%s, %s, %s, %d, %d, %d\n",
			b.judul,
			b.penulis,
			b.penerbit,
			b.tahun,
			b.eksemplar,
			b.rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}