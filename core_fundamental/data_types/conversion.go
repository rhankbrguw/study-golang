// package main
//
// import "fmt"
//
// func main() {
// 	var nilai32 int32 = 32768
// 	var nilai64 int64 = int64(nilai32)
//
// 	var nilai16 int16 = int16(nilai32)
//
// 	fmt.Println("Hasil semua nilai adalah = ", nilai32, ",", nilai64, ", dan", nilai16)
// }

package main

import "fmt"

func main() {
	name := "Fahri Arkan"
	var e uint8 = name[0]
	eString := string(e)

	fmt.Println("Hasil semua nya adalah = ", name, ",", e, ", dan", eString)
}
