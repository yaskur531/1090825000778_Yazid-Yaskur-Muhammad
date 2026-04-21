package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type angka = int
type kata = string

type Buku struct {
	judul         kata
	penulis       kata
	penerbit      kata
	tahunTerbit   angka
	jumlahHalaman angka
}

var scanner = bufio.NewScanner(os.Stdin)

func readLine() string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func main() {
	var b Buku

	fmt.Println("=== INPUT BIODATA BUKU ===")

	fmt.Print("Masukkan judul buku : ")
	b.judul = readLine()

	fmt.Print("Masukkan nama penulis : ")
	b.penulis = readLine()

	fmt.Print("Masukkan penerbit : ")
	b.penerbit = readLine()

	fmt.Print("Masukkan tahun terbit : ")
	tahun, _ := strconv.Atoi(readLine())
	b.tahunTerbit = tahun

	fmt.Print("Masukkan jumlah halaman: ")
	halaman, _ := strconv.Atoi(readLine())
	b.jumlahHalaman = halaman

	fmt.Println()
	fmt.Println("=== BIODATA BUKU ===")
	fmt.Println("Judul Buku :", b.judul)
	fmt.Println("Penulis :", b.penulis)
	fmt.Println("Penerbit :", b.penerbit)
	fmt.Println("Tahun Terbit :", b.tahunTerbit)
	fmt.Println("Jumlah Halaman :", b.jumlahHalaman)
}