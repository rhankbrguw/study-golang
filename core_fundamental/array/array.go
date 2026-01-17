package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Galih"
	names[1] = "Adhi"
	names[2] = "Kusuma"

	fmt.Println("Hasil Array string adalah = ", names[0], names[1], names[2])

	values := [...]int{
		100,
		90,
		80,
	}

	fmt.Println("Hasil Array = ", values)
	fmt.Println(len(values))
}
