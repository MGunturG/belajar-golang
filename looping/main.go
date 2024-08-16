package main

import "fmt"

func main() {
	// for for for for for for fors
	// di go hanya punya for, yes

	// varian 1
	for i := 0; i < 5; i++ {
		fmt.Println("Itersasi ke -", i)
	}

	// varian 2 (hanya kondisi)
	var j = 0

	for j < 5 { // kita hanya memasukan kondisinya di sini
		fmt.Println("Iterasi ke -", j) // hasil output akan sama seperti varian 1
		j++
	}

	// varian 3
	// tanpa argumen, conditionalnnya di dalam sarang for
	var k = 0

	for {
		fmt.Println("Angka", k)

		k++
		if k == 5 {
			break
		}
	}

	// varian 4
	// penggunaan for - range
	// cara ini biasa digunakan untuk me-looping data gabungan (misalnya string, array, slice, map)

	var xs = "123"
	for i, v := range xs {
		fmt.Println("Index=", i, "Value=", v)
		// oh ya, variabel v menghasilkan output
		// angka 1 dalam ascii
	}

	var ys = [5]int{10, 20, 30, 40, 50} // array
	for _, v := range ys {
		fmt.Println("Value=", v)
	}

	var zs = ys[0:2] // slice, data yang diambil 10,20
	for _, v := range zs {
		fmt.Println("Value=", v)
	}

	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} // map
	for k, v := range kvs {
		fmt.Println("Key=", k, "Value=", v)
	}

	// boleh juga baik k dan atau v nya diabaikan
	for range kvs {
		fmt.Println("Done")
	}

	// selain itu, bisa juga dengan cukup menentukan nilai numerik perulangan
	for i := range 5 {
		fmt.Print(i) // 01234
	}

	// penggunaan continue dan break
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	// nested loop
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	// label di loop
	// TODO
}
