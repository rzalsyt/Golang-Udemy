package main

import "fmt"

func main() {
	name := "rizal"
	switch name {
	case "rizal":
		fmt.Println("Hallo", name)
	case "budi":
		fmt.Println("Halo", name)
	default:
		fmt.Println("halo kamu belum Kenalan dengan aku")
	}

	//switch shrot statment
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah cocok")
	}

}
