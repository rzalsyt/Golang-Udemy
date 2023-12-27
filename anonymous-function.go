package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are blocked", name)
	} else {
		fmt.Println("welcme", name)
	}
}

func main() {

	blacklist := func(name string) bool {
		return name == "anjgn"
	}
	registerUser("rizal", blacklist)
}
