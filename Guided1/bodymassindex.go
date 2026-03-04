package main
import "fmt"

func main(){
	var berat, tinggi, bmi float64

	fmt.Print("Masukan berat badan (kg) : ")
	fmt.Scanln(&berat)
	fmt.Print("Masukan tinggi badan (m) : ")
	fmt.Scanln(&tinggi)

	bmi = berat / (tinggi * tinggi)

	fmt.Print("Nilai BMI Anda adalah: 2f\n", bmi)
}