package main

import "fmt"

func main() {
	// && harus 2 2 nya benar.
	// || harus ada nilai yang bener
	// ! kebalikan

	var uts = 90
	var uas = 81

	var lulusUts bool = uts > 80
	var lulusUas bool = uas > 80

	var lulus bool = lulusUts && lulusUas

	fmt.Println(lulus)

}
