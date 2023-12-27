package main

import "fmt"

// digolang bisa function menjadi value

func getGoodBy(name string) string {
	return "goodby " + name

}

func main() {

	contoh := getGoodBy
	misal := getGoodBy
	fmt.Println(contoh("eko"))
	fmt.Println(misal("eko"))
}
