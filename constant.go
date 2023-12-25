package main

import "fmt"

func main() {

	//constant = tidak bisa merubah nilai variable
	const umur = 20

	const (
		//multiple constanta
		firstName = "rizal"
		lastName  = "Susmiyanto"
	)
	fmt.Println("Umur : ", umur)
	fmt.Println("nama : ", firstName)

	// //error = merubah nilai sebelumnya
	// firstName = "Budi"
	// lastName = " kecil"

}
