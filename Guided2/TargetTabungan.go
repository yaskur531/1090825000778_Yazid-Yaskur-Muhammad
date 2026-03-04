package main
import "fmt"

func main(){
	var target, tabungan, total, hari int
	fmt.Print("Masukan target uang yang ingin dicapai :")
	fmt.Scanln(&target)

	total = 0
	hari = 0

	for total < target {
		hari++
		fmt.Printf("Masukan nominal tabungan hari ke-%d", hari)
		fmt.Scanln(&tabungan)

		total = total + tabungan
	}

	fmt.Print("Selamat Target Tercapai dalam %d hari.\n", hari)
	fmt.Print("Total tabungan anda terkumpul %d \n", total)

}