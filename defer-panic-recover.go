package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")

}

func runApplication() {
	defer logging() //logging memanggil di akhir, yang dijalankan awal run aplikation
	fmt.Println("run Aplikation")
}

func main() {
	//defer -> function yang bisa kita jadwalkan untuk di eksekusi setelah function selesai di eksekusi
	//defer function akan dieksekusi walaupun terjadi error di function yang di eksekusi

	runApplication()
}
