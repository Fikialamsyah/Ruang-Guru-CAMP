package main

import (
	"fmt"
)

// pada cekpoin ini kalian akan mencoba untuk menghapus data dengan tipe []map[string]string
// gabungan slice dan map.

func main() {
	var namaUmur = []map[string]string{
		{"name": "Socrates", "gender": "male"},
		{"name": "Plato", "gender": "male"},
		{"name": "Ada", "gender": "female"},
		{"name": "Leonhard Euler", "gender": "female"},
		{"name": "Blaise Pascal", "gender": "male"},
	}

	// terdapat kesalahan pada data gender tersebut dapatkan kalian memperbaiki nya ?
	// TODO: answer here
	for _, val := range namaUmur {
		if val["name"] == "Leonhard Euler" {
			val["gender"] = "male"
		}
		fmt.Println(val["name"], " ", val["gender"])
	}

	// Nah coba saatnya kalian menghapuskan key "gender" pada setiap data
	// delete data if key is equal "gender"

	for key, val := range namaUmur {
		fmt.Println(key,val)
	}
	// Output sebelum dihapus
	/*
		map[gender:male name:Socrates]
		map[gender:male name:Plato]
		map[gender:female name:Ada]
		map[gender:male name:Leonhard Euler]
		map[gender:male name:Blaise Pascal]
	*/

	// TODO: answer here
	for _,itemMap := range namaUmur{
		delete(namaUmur[0], "gender")
		for _,val := range itemMap{
			delete(itemMap, "gender")
			fmt.Println(val)
		}
		
	}
	// Output setelah dihapus
	/*
		map[name:Socrates]
		map[name:Plato]
		map[name:Ada]
		map[name:Leonhard Euler]
		map[name:Blaise Pascal]
	*/
	for _, val := range namaUmur {
		fmt.Println(val)
	}
}
