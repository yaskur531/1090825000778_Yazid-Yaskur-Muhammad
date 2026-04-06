package main
import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int

	fmt.Print("masukkan nilai a : ")
	fmt.Scan(&a)
	fmt.Print("masukkan nilai b : ")
	fmt.Scan(&b)
	fmt.Print("masukkan nilai c : ")
	fmt.Scan(&c)


	fogoh := f(g(h(a)))

	gohof := g(h(f(b)))

	hofog := h(f(g(c)))

	fmt.Printf("(f o g o h)(%d) = %d\n", a, fogoh)
	fmt.Printf("(g o h o f)(%d) = %d\n", b, gohof)
	fmt.Printf("(h o f o g)(%d) = %d\n", c, hofog)
}