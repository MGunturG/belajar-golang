package main

import "fmt"

func main() {
	// contoh penerapan array
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])
	// output : trafalgar d water law


	// contoh inisialisasi nilai awal array
	var array_buah = [4]string{"apel", "anggur", "pisang", "melon"}
	fmt.Println("Jumlah elemen array \t\t", len(array_buah))
	fmt.Println("Isi arraynya \t", array_buah)
	
	// output :
	// Jumlah elemen array	4
	// Isi arraynya		[apel anggur pisang melon]
	
	// inisialisasi juga bisa seperti ini
	// untuk inisialisasi seperti di bawah ini
	// tanda baca koma (,) harus diassign di setiap
	// akhir elemen
	var array_mobil = [4]string{
		"toyota",
		"honda",
		"nissan",
		"lexus",
	}
	
	fmt.Println(array_mobil[0])

	// atau begini
	var array_sepatu = make([]string, 2)
	array_sepatu[0] = "adidas"
	array_sepatu[1] = "nike"

	fmt.Println(array_sepatu[1])


	// Inisialisasi array tanpa jumlah elemen array
	// kalo sebelumnya kita harus ngasih tau jumlah elemen arraynya
	// misal : [4]string, berarti ada 4 string di dalam arraynya
	// nah, kita juga bisa inisiasi array tanpa assign jumlah elemen arraynya

	var array_angka = [...]int{1,2,3,4}
	fmt.Println("Data array \t:", array_angka)
	fmt.Println("Jumlah elemennya \t:", len(array_angka))

	// secara otomatis, compiler akan meng-assign jumlah elemen array angka
	// sebanyak 4 elemen. bisa dilihat dari outputnya:
	// Data array		: [1 2 3 4]
	// Jumlah elemennya	: 4
	

	// Array Multidimensi
	// > oh shit, here we go again. my first nightmare from learning C

	var numbers1 = [2][3]int{[3]int{3,2,3}, [3]int{3,4,5}} // ini ribet sih
	var numbers2 = [2][3]int{{3,2,3}, {3,4,5}} // pakai ini aja, mirip kaya C

	fmt.Println("Array numbers1", numbers1)
	fmt.Println("Array numbers2", numbers2)

	// looping array
	var array_hp = [4]string{"nokia", "sony", "pixel", "oneplus"}

	for i := 0; i < len(array_hp); i++ {
		fmt.Printf("elemen %d : %s\n", i, array_hp[i])
	}

	// atau

	for i, hp := range array_hp {
		fmt.Printf("elemen %d : %s\n", i, hp)
	}
}
