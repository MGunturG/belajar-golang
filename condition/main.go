package main

import "fmt"

func main() {
	// kondisional atau condition di pemrogramman golang
	// hampir sama seperti di bahasa C, bedanya tidak perlu ada
	// parentheses()

	// ===== condition if, else if, else =======
	nilai := 70
	if nilai < 60 { // tidak perlu parentheses
		fmt.Println("perlu perbaikan")
	} else if nilai == 69 {
		fmt.Println("nice")
	} else {
		fmt.Println("lulus")
	}

	// =========== variable temporary ===========
	// variabel percent adalah var temporary
	// hanya bisa diakses oleh block if
	// dan kode di bawahnya. jika block if
	// selesai, var temporary tidak bisa
	// dipakai lagi
	var point = 8840.0

	if percent := point / 100; percent >= 100 { // variabel temporary, percent
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// Deklarasi variabel temporary hanya bisa dilakukan lewat metode type inference yang menggunakan tanda :=.
	// Penggunaan keyword var di situ tidak diperbolehkan karena menyebabkan error.

	// ======== switch - case ========
	// switch - case itu conditional yang fokusnya untuk perbandingan satu variabel

	var value = 6

	switch value {
	case 8, 9, 10: // bisa sekaligus, jadi tidak perlu buat satu per satu
		fmt.Println("good")
	case 7:
		fmt.Println("enough")
	case 6:
		fmt.Println("try again")
	default:
		fmt.Println("out of range")
		// jadi, jika case 8, case 7, dan 6 tidak terpenuhi
		// akan masuk ke bagian default
		// default sama seperti else
	}

	// penggunaan fallthrough
	// mirip seperti contiune pada c
	switch {
	case value == 8:
		fmt.Println("perfect")
	case (value < 8) && (value > 3):
		fmt.Println("awesome")
		fallthrough // sama seperti contiune
		// case di bawah fallthrough akan dijalankan
		// meskipun nilai value = 6
	case value < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// seleksi nested condition
	point = 10

	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
