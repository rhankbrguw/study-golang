package main

import "fmt"

func main() {
	name := "Fahri Arkan"

	// if name == "Fahri Arkan" {
	// 	fmt.Println("Halo, Petani Cupang")
	// }

	// IF-ELSE
	// if name == "Fahri Arkan" {
	// 	fmt.Println("Halo, Petani Cupang")
	// } else {
	// 	fmt.Println("Kamu bukan petani cupang")
	// }

	// 	ELSE-IF
	// 	if name == "Fahri Arkan" {
	// 		fmt.Println("Halo, Petani Cupang")
	// 	} else if name == "Mcboy" {
	// 		fmt.Println("Halo, CEO")
	// 	} else {
	// 		fmt.Println("Dapet ga?")
	// 	}
	// }

	// SHORT STATEMENTS
	length := len(name)
	if length > 5 {
		fmt.Println("Kepanjangan nama lu")
	} else {
		fmt.Println("Pendek amat")
	}
}
