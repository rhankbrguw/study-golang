package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)

	fmt.Println("Hasil semua nilai adalah = ", nilai32, ",", nilai64, ", dan", nilai16)
}
