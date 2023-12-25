package main

import "fmt"

func main() {

	//membuat 3 array
	var names [3]string

	names[0] = "eko"
	names[1] = "rizal"
	names[2] = "budi"

	fmt.Println(names[0])
	fmt.Println(names)

	var values = [3]int{
		90,
		80,
		70,
	}
	fmt.Println(values)

	//function / kemampuan array
	// len(array) jumlah panjang array
	//arry[index] mendapatkan data diposisi index
	//array[index] = value -> mengubah data posisi index

	// [...] tidak menentukan isi array
	var values1 = [...]int{
		90,
		80,
		70,
		110,
	}

	fmt.Println(values1)
	fmt.Println(len(values1))
	fmt.Println(values1[1])
	// fmt.Println(values1[2] = 2)
}
