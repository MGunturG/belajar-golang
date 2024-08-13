package main

import "fmt"

func main() {
	// variabel konstanta (const) itu variabel yang hanya bisa
	// sekali dideclare. setelah dideclare valuenya gak bisa diganti
	// contohnya : pi (3.14...), g (gravitasi 9.8m/s), dan lainnya

	const warna_sapi string = "hitam putih"
	fmt.Print(warna_sapi)

	// deklarasi multi konstanta
	// contoh 1
	const (
		square          = "kotak"
		int_num   uint8 = 3
		float_num       = 6.9
	)
	fmt.Print(square, int_num, float_num)

	// contoh 2
	const (
		today    string = "selasa"
		hari_ini        // konstanta hari_ini akan memiliki value "selasa"
		// jika deklarasi konstanta tidak dimasukan, maka akan mengikuti
		// value konstanta diatasnya
	)
	fmt.Print(today, hari_ini)

	// contoh 3
	const satu, dua = 1, 2
	const tiga, empat string = "tiga", "empat"
	fmt.Print(satu, dua, tiga, empat)
}
