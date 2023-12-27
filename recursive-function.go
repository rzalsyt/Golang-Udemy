package main

import "fmt"

// recursive function ->function  yang memanggil dirinya sendiri
func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func facktorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * facktorialRecursive(value-1)
	}
}

func main() {
	fmt.Println(factorialLoop(10))
	fmt.Println(facktorialRecursive(10))
}
