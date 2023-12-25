package main

import "fmt"

func main() {

	// + , - , * , % module, /
	var a = 10
	var b = 10
	var c = a * b
	var e = a / b

	fmt.Println(c)
	fmt.Println(e)

	//augmented Assignments
	var i = 10
	i += 10
	fmt.Println(i) // 10 + 10
	i += 5
	fmt.Println(i) // 20 + 5

	//Unery operator
	// ++ a = a+ 1

	var j = 1
	j++
	fmt.Println(j) // 1 + 1
	j++
	fmt.Println(j) //2 + 1

}
