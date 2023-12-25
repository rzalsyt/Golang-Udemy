package main

import "fmt"

func getCompalteName() (firsttName, middleName, lastName string) {
	firsttName = "rizal"
	middleName = "sus"
	lastName = "Miyanto"

	return firsttName, middleName, lastName
}

func main() {

	namaDepan, middleName, namaAkhir := getCompalteName()
	// namaDepan, _, namaAkhir := getCompalteName()
	fmt.Println(namaDepan, middleName, namaAkhir)

}
