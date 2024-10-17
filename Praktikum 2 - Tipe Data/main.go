package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	var pilihanMenu int

	fmt.Println("")
	fmt.Println("-=-=-=- Welcome to Jordan's Praktikum 2 - Tipe Data -=-=-=-")
	fmt.Println("")
	fmt.Println("Menu:")
	fmt.Println("1. Hello")
	fmt.Println("2. Hipotenusa")
	fmt.Println("3. Latihan 1")
	fmt.Println("4. Latihan 2")
	fmt.Println("5. Latihan 3")
	fmt.Println("6. Latihan 4")
	fmt.Println("7. Latihan 5")

	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihanMenu)

	reader := bufio.NewReader(os.Stdin) // Membuat reader baru dimana Stdin adalah input dari terminal
	reader.ReadString('\n') // Membuang karakter newline

	fmt.Println("")

	switch pilihanMenu {
	case 1:
		hello()
	case 2:
		hipotenusa()
	case 3:
		latihan1()
	case 4:
		latihan2()
	case 5:
		latihan3()
	case 6:
		latihan4()
	case 7:
		latihan5()
	default:
		fmt.Println("Huh? What number did you just type?")
	}
}