package main

import "fmt"

// function parameters (firstName string, lastName string,)

func sayHello(firstName string, lastName string) {
	fmt.Println("hallo", firstName, lastName)
}

func main() {
	// harus mengisi nilai paramatersnya
	sayHello("Rizal ", "Susmiyanti")
}
