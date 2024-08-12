package main

import "fmt"

func main() {
	// ============================= tipe data nunerik Go ==============================
	// | Unsigned 	| Signed 	| Decimal 	|
	// --------------------------------------
	// | uint8 		| int8		| float32 	|
	// | uint16		| int16		| float64 	|
	// | uint32		| int32		| 			|
	// | uint64		| int64		| 			|

	var positive_number uint8 = 98
	var negative_number = -1243423644 // ini otomatis diassign tipe datanya sama compiler
	fmt.Printf("Positive: %d | Negative: %d\n", positive_number, negative_number)

	var decimal_number = 6.93
	fmt.Printf("Decimal : %f\n", decimal_number)   // output normal
	fmt.Printf("Decimal : %.3f\n", decimal_number) // ini diset berapa angka belakang koma

	// =============================== tipe data boolean ==============================
	// boolean bernilai true atau false

	var exist bool = true
	fmt.Printf("exist? %t \n", exist) // Gunakan %t untuk memformat data bool
}
