package main

import "fmt"

func main() {

	//slice -> irisan dari array
	names := [...]string{"eko", "rizal", "budi", "rudi", "dani", "andini"}
	slice := names[4:6]

	fmt.Println(slice) // , 5 dan 6 di ambikl

	slice2 := names[:3]
	fmt.Println(slice2) // index 0 sampe 3

	slice3 := names[:]
	fmt.Println(slice3) // ini all

	//mrmbuat manual slice
	// var slice4 []string = [1:3]

	//fungsi array -> len(slice) -> mendapatkan panjang
	// -> cap(slice) -> mendapatkan capasitas
	// -> append(slice) -> membuat slice baru dengan posisi akhir slice, jika posisi kapaistas penuh
	// -> make(slice) -> membuat slice baru
	// -> copy -> menyalin slice dari source ke destination

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	daySlice1 := days[5:] // sabtu , minggu
	fmt.Println("day dari ", daySlice1)

	daySlice1[0] = "sabtu Baru"
	daySlice1[1] = "Minggu baru"
	fmt.Println("merubah data slice index ke 5 dan ke 6 ", daySlice1)
	fmt.Println(days)

	//menambahkan data append
	daySlice2 := append(daySlice1, "libur baru")
	// days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	//membuat slice baru

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "rizal"
	newSlice[1] = "budi"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice3 := append(newSlice, "kujang")

	newSlice3[1] = " ojan"
	fmt.Println(newSlice3)
	fmt.Println(newSlice)

	//copy slice

	fromSlice := days[:]
	toSlice := make([]string, len(days), cap(days))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// bikin array -> iniArray := [...]int{1,3,4,5,6,}
	//iniSlice -> :=[]int{1,2,3,4,5,}
}
