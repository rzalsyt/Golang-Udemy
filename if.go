package main

import "fmt"

func main() {

	name := "budi"

	// if
	if name == "rizal" {
		fmt.Println("hallo", name)
	} else if name == "budi" {
		fmt.Println(" hi ", name)
	} else {
		fmt.Println("Hi , Boleh kenalan")
	}

	// if short statment
	if lenght := len(name); lenght > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("nama sudah bener")
	}
}
