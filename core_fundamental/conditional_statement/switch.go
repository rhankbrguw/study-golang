package main

import "fmt"

func main() {
	name := "Fahri"

	switch name {
	case "Fahri Arkan":
		fmt.Println("Halo, Petani Cupang")
	case "Mcboy":
		fmt.Println("Halo, CEO")
	default:
		fmt.Println("Dapet ga?")
	}

	// SHORT STATEMENTS
	// 	switch length := len(name); length > 5 {
	// 	case true:
	// 		fmt.Println("Kepanjangan nama lu")
	// 	case false:
	// 		fmt.Println("Pendek amat")
	// 	}

	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Kepanjangan nama lu")
	case length < 5:
		fmt.Println("Pendek amat")
	default:
		fmt.Println("Dapet ga?")
	}
}
