package main

import "fmt"

func main() {
	// deklarasi variabel langsung assign tipe data
	var firstname string = "Deez"
	var lastname string = "Nuts"

	// print variabel ke terminal cara 1
	fmt.Printf("Hello %s %s!\n", firstname, lastname)

	// print variabel ke terminal cara 2
	fmt.Println("Hallo", firstname, lastname+"!")

	var nama_awal string = "John"
	// deklarasi variabel tanpa assign tipe data
	nama_akhir := "Wick" // --> ini otomatis diassign sama compiler tipe datanya
	// kalo kasus ini karena setelah := adalah tipe data string
	// berarti variabel nama_akhir adalah bertipe data string
	// atau bisa juga kaya di bawah ini
	var status = "tampan"

	fmt.Printf("Halo %s %s %s!\n", nama_awal, nama_akhir, status)

	// ======================= deklarasi multi variabel ============================
	var first, second, third string
	first, second, third = "pertama", "kedua", "ketiga"

	// atau bisa juga
	var fouth, fifth, sixth string = "empat", "lima", "enam"

	// kalo mau lebih ringkas
	ini_nomor, ini_string := 123, "satu dua tiga"

	fmt.Println(first, second, third, fouth, fifth, sixth, ini_nomor, ini_string)

	// 	====================== variabel underscore _ ===============================
	// !! Di Golang, variabel harus selalu dipakai. jika ada variabel tidak terpakai
	// !! dia gak akan bisa dicompile/run
	//
	// maka itu ada yang namanya reserved variable (_) yang bisa dipakai untuk menam-
	// pung variabel yang tidak dipakai. variabel (_) tidak bisa dipakai/tampilkan
	namaku, _ := "yoshikuni", 1234
	fmt.Printf("%s", namaku)
}
