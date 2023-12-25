package main

import "fmt"

func main() {

	//variable = tempat untuk menyimpan data
	// wajib menyebutkan tipe data variable / langsung bisa (opsional)
	//

	var name string

	name = "rizal Susmiyanto"
	fmt.Println(name)

	name = "rizal keren"
	fmt.Println(name)

	//:= hanya awalan untuk mendeklarasikan variable , tidak bisa mengubahnya

	umur := 7
	fmt.Println(umur)

	umur = 10
	fmt.Println(umur)

	//deklarasi multiple variable
	var (
		namaDepan    = "rizal"
		namaBelakang = "Susmiyanto"
		lastName     = "Ijol"
	)

	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)
	fmt.Println(lastName)
}
