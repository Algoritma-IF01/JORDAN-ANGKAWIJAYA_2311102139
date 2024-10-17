package main

import "fmt"

func latihan5() {
	fmt.Println("Masukkan 5 angka integer antara 32 sampai 127: ")

	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)

	fmt.Println("Masukkan 3 karakter:")
	
	var input string
	fmt.Scan(&input)

	fmt.Println("")
	fmt.Println("Output:")

	fmt.Printf("%c%c%c%c%c\n", a, b, c, d, e)

	if len(input) == 3 {
		fmt.Printf("%c%c%c\n", input[0]+1, input[1]+1, input[2]+1)
	} else {
		fmt.Println("Huh? Input karakter harus terdiri dari 3 karakter")
	}
}
