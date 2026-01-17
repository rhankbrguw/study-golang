package main

import "fmt"

func main() {
	var name string

	name = "Fahri Arkan"
	fmt.Println("Nama kamu sekarang = ", name)

	name = "Akamsi"
	fmt.Println("Nama kamu dikampus = ", name)
}

// SIMPLIFIED (no need to declare data type first)
// package main
//
// import "fmt"
//
// func main() {
// 	name := "Fahri Arkan"
// 	fmt.Println("Nama kamu sekarang = ", name)
//
// 	name = "Akamsi"
// 	fmt.Println("Nama kamu dikampus = ", name)
// }

// // MULTIPLE VARIABLE
// package main
//
// import "fmt"
//
// func main() {
// 	var (
// 		namaSekarang = "Fahri Arkan"
// 		namaKampus   = "Akamsi"
// 	)
//
// 	fmt.Println("Nama kamu sekarang dan dikampus adalah = ", namaSekarang, "dan", namaKampus)
// }
