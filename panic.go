package main

import "fmt"

// panic -> function yang bisa kita gunakan untuk menghentikan program
// ketita panic dipanggil, program akan terkhentim namun defer function akan di eksekusi

func endApp() {
	fmt.Println("end App")
	message := recover() // recover untuk menangani adanya panic terjadi
	fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
	defer endApp() // defer tetep keluar ketika error muncul
	if error {
		panic("ada kesalahan ")
	}

}

func main() {
	runApp(true)
	fmt.Println("ini rizal")

}
