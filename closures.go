package main

import "fmt"

func main() {

	//closures adalah kemampuan sebuah function berinteraksi dengan data data disekitarnya dalaaha scope yang sama
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	increment()

	fmt.Println(counter)

}
