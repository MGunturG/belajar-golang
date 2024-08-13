package main

import "fmt"

func main() {
	// reading :
	//	- https://www.geeksforgeeks.org/go-operators/#Relational%20Operators
	// 	- https://www.programiz.com/c-programming/bitwise-operators

	// operator aritmatika : + (jumlah), -(pengurangan), /(pembagian), *(perkalian), %(modulo atau sisa pembagian)

	// ========== operator perbandingan =========		============== operator bitwise =================
	// |----------------------------------------|		|-----------------------------------------------|
	// | operator 	| fungsi					|		| 	operator 		| fungsi					|
	// |----------------------------------------|		|-----------------------------------------------|
	// |	==		| sama dengan				|		| 	& (and) 		|
	// |	!=		| tidak sama dengan			|		| 	| (or)			|
	// |	<		| kurang dari				| 		| 	^ (xor)			|
	// |	<=		| kurang dari sama dengan	| 		| 	<<(left shift)	|
	// |	>		| lebih dari				|		|	>>(right shift)	|
	// |	>=		| lebih dari sama dengan	|		|	&^ (and not)	|
	// ------------------------------------------		------------------------------------------------

	// ============= operator logika ============
	// |----------------------------------------|
	// | operator 	| fungsi					|
	// |----------------------------------------|
	// | 	&&		| and						|
	// |	||		| or						|
	// |	!		| not						|
	// ------------------------------------------

	// ========== contoh penggunaan operator aritmatika dan logika ========
	var value = (((2+6)%3)*4 - 2) / 3
	var isEqual = (value == 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	// =========== contoh penggunaan operator bitwise ==========
	p := 12
	q := 25

	// & (bitwise AND)
	// 12 = 00001100 (In Binary)
	// 25 = 00011001 (In Binary)

	// Bit Operation of 12 and 25
	// 00001100
	// & 00011001
	// ________
	// 00001000  = 8 (In decimal)
	result1 := p & q
	fmt.Printf("Result of p & q = %d", result1)

	// | (bitwise OR)
	// 12 = 00001100 (In Binary)
	// 25 = 00011001 (In Binary)

	// Bitwise OR Operation of 12 and 25
	// 00001100
	// | 00011001
	// ________
	// 00011101  = 29 (In decimal)
	result2 := p | q
	fmt.Printf("\nResult of p | q = %d", result2)

	// ^ (bitwise XOR)
	// 12 = 00001100 (In Binary)
	// 25 = 00011001 (In Binary)

	// Bitwise XOR Operation of 12 and 25
	// 00001100
	// ^ 00011001
	// ________
	// 00010101  = 21 (In decimal)
	result3 := p ^ q
	fmt.Printf("\nResult of p ^ q = %d", result3)

	// << (left shift)
	// 212 = 11010100 (In binary)
	// 212<<1 = 110101000 (In binary) [Left shift by one bit]
	// 212<<0 = 11010100 (Shift by 0)
	// 212<<4 = 110101000000 (In binary) =3392(In decimal)
	result4 := 34 << 1
	fmt.Printf("\nResult of p << 1 = %d", result4)

	// >> (right shift)
	// 212 = 11010100 (In binary)
	// 212 >> 2 = 00110101 (In binary) [Right shift by two bits]
	// 212 >> 7 = 00000001 (In binary)
	// 212 >> 8 = 00000000
	// 212 >> 0 = 11010100 (No Shift)
	result5 := 20 >> 1
	fmt.Printf("\nResult of p >> 1 = %d", result5)

	// &^ (AND NOT)
	// i dont have any idea what is this
	// TODO : nyari tau ini apaan dan penjelasannya
	result6 := 34 &^ 20
	fmt.Printf("\nResult of p &^ q = %d", result6)

	// =========== misc =============
	a := 4

	// Using address of operator(&) and
	// pointer indirection(*) operator
	b := &a
	fmt.Println(*b)
	*b = 7
	fmt.Println(a)
}
