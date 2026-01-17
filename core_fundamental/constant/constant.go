package main

import "fmt"

func main() {
	const namaSekarang = "Fahri Arkan"
	const namaKampus = "Akamsi"

	// ERROR (constant can't be changed)
	// const namaSekarang = "Galih Adhi Kusuma"
	// const namaKampus = "Semong"

	fmt.Println("Nama kamu sekarang dan dikampus adalah = ", namaSekarang, "dan", namaKampus)
}

// MULTIPLE VARIABLE
// package main
//
// import "fmt"
//
// func main() {
// 	const (
// 		namaSekarang = "Fahri Arkan"
// 		namaKampus   = "Akamsi"
// 	)
//
// 	fmt.Println("Nama kamu sekarang dan dikampus adalah = ", namaSekarang, "dan", namaKampus)
// }
