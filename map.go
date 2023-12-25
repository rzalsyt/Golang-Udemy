package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "eko"
	// person["alamat"] = "Bekasi"

	person := map[string]string{
		"name":   "eko",
		"alamat": "Bekasi",
	}
	fmt.Println(person)
	fmt.Println(person["name"])

	//function Map
	// -> len(map) -> mendapatkan jumalh dimao
	// -> map[key] - > mengambil data di map
	// -> map[key] = value -> mengubah data dimap
	// -> make(map[Typekey]TypeValue) -> membuat map baru
	// -> delete(map,key) -> menghapus map

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Rizal"
	book["Terbit"] = "2022"

	fmt.Println(book)

	delete(book, "Terbit")

	fmt.Println(book)
}
