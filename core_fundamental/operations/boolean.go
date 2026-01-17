package main

import "fmt"

func main() {
	nilaiAkhir := 85
	absensi := 80

	var lulusNilaiAkhir bool = nilaiAkhir >= 75
	var lulusAbsensi bool = absensi >= 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulusNilaiAkhir)
	fmt.Println(lulusAbsensi)
	fmt.Println(lulus)
}
