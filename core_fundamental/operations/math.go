package main

import "fmt"

func main() {
	a := 20
	b := 30
	c := 40
	d := 70
	e := 50
	f := a + b - c*d/e

	fmt.Println("Hasil penjumlahan adalah = ", f)
}

// AUGMENTED ASSIGNMENT
// package main
//
// import "fmt"
//
// func main() {
// 	i := 10
// 	i += 10
// 	fmt.Println("Hasil dari i = "i)
//
// 	i += 30
// 	fmt.Println("Hasil dari i = ", i)
// }

// UNARY OPERATOR
// package main
//
// import "fmt"
//
// func main() {
// 	j := 1
//
// 	j++
// 	fmt.Println(j)
//
// 	j--
// 	fmt.Println(j)
// }
