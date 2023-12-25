package main

import "fmt"

// mengembalikan string
func getHello(name string) string {
	hello := "hello " + name

	return hello

}

func main() {
	result := getHello("eko")

	fmt.Println(result)
}
