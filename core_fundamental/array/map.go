package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Fahri Arkan",
		"address": "Jl. Sahabat Baru No.17A, RT.8/RW.1, Duri Kepa, Kec. Kb. Jeruk, Kota Jakarta Barat, Daerah Khusus Ibukota Jakarta 11510",
	}

	fmt.Println("Identitas lengkap petani cupang = ", person)
	fmt.Println("Nama petani cupang = ", person["name"])
	fmt.Println("Alamat petani cupang = ", person["address"])

	// FUNCTION MAP (delete map index)
	petaniCupang := make(map[string]string)
	petaniCupang["name"] = "Fahri Arkan"
	petaniCupang["age"] = "20"
	petaniCupang["address"] = "Jl. Sahabat Baru No.17A, RT.8/RW.1, Duri Kepa, Kec. Kb. Jeruk, Kota Jakarta Barat, Daerah Khusus Ibukota Jakarta 11510"
	petaniCupang["ytta(ups)"] = "di bandung"

	delete(petaniCupang, "ytta(ups)")

	fmt.Println(petaniCupang)
	fmt.Println(len(petaniCupang["address"]))
}
