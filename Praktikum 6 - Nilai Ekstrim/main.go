package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var pilihanMenu int

	fmt.Println("")
	fmt.Println("-=-=-=- Welcome to Jordan's Praktikum 6 - Pencarian Nilai Ekstrim -=-=-=-")
	fmt.Println("")
	fmt.Println("Menu:")
	fmt.Println("1. Nilai Ekstrim Array Dasar")
	fmt.Println("2. Nilai Ekstrim Array Terstrukur")
	fmt.Println("3. Latihan 1")
	fmt.Println("4. Latihan 2")
	fmt.Println("5. Latihan 3")

	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihanMenu)

	reader := bufio.NewReader(os.Stdin) // Membuat reader baru dimana Stdin adalah input dari terminal
	reader.ReadString('\n')             // Membuang karakter newline

	fmt.Println("")

	switch pilihanMenu {
	case 1:
		ne_arraydasar()
	case 2:
		ne_arrayterstruktur()
	case 3:
		latihan1()
	case 4:
		latihan2()
	case 5:
		latihan3()
	default:
		fmt.Println("Huh? What number did you just type?")
	}
}
