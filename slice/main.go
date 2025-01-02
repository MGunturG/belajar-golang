package main

import "fmt"

func main() {
	// inisialisasi slice
	var slice_buah = []string{"apel", "jeruk", "pisang", "melon"}
	fmt.Println(slice_buah[0])

	var fruitsA = []string{"apel", "anggur"} // slice
	var fruitsB = [2]string{"apel", "anggur"} // array
	var fruitsC = [...]string{"apel", "anggur"} // array

	// print elemen, variabel harus terpakai 
	fmt.Println(fruitsA[1])
	fmt.Println(fruitsB[0])
	fmt.Println(fruitsC[1])

	// slice bisa dibentuk dari array yang udah ada, misal
	var buah = []string{"melon", "pisang", "salak", "jambu"}
	var new_buah = buah[0:2]
	fmt.Println(new_buah)


	// slice itu tipe data reference, jadi kalo ada slice lama datanya diubah
	// di slice baru, data di slice lama bakal berubah mengikuti slice baru
	var fruits = []string{"apple", "grape", "banana", "melon"}

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]
	

	// fungsi len()
	// berguna buat ngitung jumlah elemen slice yang ada

	// fungsi cap()
	// dipakai buat ngitung lebar/kapasitas maksimum slice

	var f1_team = []string{"ferrari", "mclaren", "redbull", "mercedes"}
	fmt.Println(len(f1_team)) // len 4
	fmt.Println(cap(f1_team)) // len 4

	var f1_team_a = f1_team[0:3]
	fmt.Println(len(f1_team_a)) // len 3
	fmt.Println(cap(f1_team_a)) // cap 4

	var f1_team_b = f1_team[1:4]
	fmt.Println(len(f1_team_b)) // len 3
	fmt.Println(cap(f1_team_b)) // cap 3

	// explanation
	// kode			output		len() 	cap()
	// f1_team[0:4]		[x x x x]	4	4
	// f1_team[0:3]		[x x x -]	3	4
	// f1_team[1:4]		- [x x x]	3 	3

	// TL;DR dari web asli. jadi kalo ngeslice dari index 0
	// len akan mengikuti berapa elemen yang kita slice tapi capacity 
	// ngikutin slice asli
	//
	// kalo dari index (> 0), misal x, si nilai x jadi elemen pertama
	// hasil slicenya, dan cap mengikuti jumlah elemen yang di slice


	// fungsi append()

}
