package main

import "fmt"

func main() {
	numbers := [...]int{
		100,
		90,
		80,
		70,
		60,
		50,
		40,
		30,
		20,
		10,
	}

	slice1 := numbers[6:8]
	fmt.Println("Hasil dari slice pertama adalah = ", slice1)

	slice2 := numbers[3:]
	fmt.Println("Hasil dari slice kedua adalah = ", slice2)

	slice3 := numbers[:3]
	fmt.Println("Hasil dari slice ketiga adalah = ", slice3)

	slice4 := numbers[:]
	fmt.Println("Hasil dari slice keempat adalah = ", slice4)

	numbersSlice1 := numbers[5:]
	numbersSlice1[0] = 55
	numbersSlice1[1] = 50
	numbersSlice1[2] = 40
	numbersSlice1[3] = 30
	numbersSlice1[4] = 20

	fmt.Println("Hasil dari slice kelima adalah = ", numbers)

	// FOR ADD DATA INSIDE ARRAY
	numbersSlice2 := append(numbersSlice1, 10)
	fmt.Println("Hasil dari slice kelima adalah = ", numbersSlice2, "dan", numbers)

	var newSlice1 []string = make([]string, 2, 5)
	newSlice1[0] = "Fahri"
	newSlice1[1] = "Arkan"

	fmt.Println("Hasil dari new slice pertama adalah = ", newSlice1, ",", len(newSlice1), ", dan", cap(newSlice1))

	newSlice2 := append(newSlice1, "Akamsi")
	fmt.Println("Hasil dari new slice kedua adalah = ", newSlice2, ",", len(newSlice2), ", dan", cap(newSlice2))

	// COPY SLICE
	fromSlice := numbers[:]
	toSlice := make([]int, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// COMPARISON BETWEEN
	thisArray := [...]int{1, 2, 3, 4, 5}
	thisSlice := []int{1, 2, 3, 4, 5}

	fmt.Println("Perbandingan array dan slice, ini array = ", thisArray, ", dan ini slice = ", thisSlice)
}
