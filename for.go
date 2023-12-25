package main

import "fmt"

func main() {
	// counter := 1

	// //kondisii for
	// for counter <= 10 ;  {
	// 	fmt.Println("perulangan ke ", counter)
	// 	counter++
	// }

	//for dengan statment -> init post statment , post statment

	// for init statment; Kondisi ; post statment
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke ", counter)

	}

	names := []string{"eko ", "kurniawan", "rizal"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	//for range
	for index, names := range names {
		fmt.Println("index", index, "=", names)
	}

	// _ menghilangakn index
	for _, names := range names {
		fmt.Println(names)
	}
}
