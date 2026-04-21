package main
import "fmt"
	
type desimal float64
	
func main(){
	var luas, alas , tinggi desimal
    
	    fmt.Print("Masukan panjang alas segititiga : ")
	    fmt.Scan(&alas)
	
	    fmt.Print("Masukan Tinggi Segitiga : ")
	    fmt.Scan(&tinggi)
	
	    luas = 0.5 * alas * tinggi 
	
    fmt.Printf("Luas segitiga dengan alas %.2f dan tinggi %.2f adalah %.2f", alas, tinggi, luas)
	}
