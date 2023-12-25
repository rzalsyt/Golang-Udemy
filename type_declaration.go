package main

import "fmt"

func main() {
	//membuat alias
	type NoKTP string

	var ktpEko NoKTP = "111111"

	var contoh string = "222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpEko)
	fmt.Println(contohKtp)

}
