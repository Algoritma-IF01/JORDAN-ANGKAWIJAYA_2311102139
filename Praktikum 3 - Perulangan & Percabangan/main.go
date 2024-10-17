package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	var pilihanMenu int

	fmt.Println("")
	fmt.Println("-=-=-=- Welcome to Jordan's Praktikum 3 - Perulangan & Percabangan -=-=-=-")
	fmt.Println("")
	fmt.Println("Menu:")
	fmt.Println("1. Perulangan Latihan 1")
	fmt.Println("2. Perulangan Latihan 2")
	fmt.Println("3. Perulangan Latihan 3")
	fmt.Println("4. Perulangan Latihan 4")
	fmt.Println("5. Percabangan Latihan 1")
	fmt.Println("6. Percabangan Latihan 2")
	fmt.Println("7. Percabangan Latihan 3") 

	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihanMenu)

	reader := bufio.NewReader(os.Stdin) // Membuat reader baru dimana Stdin adalah input dari terminal
	reader.ReadString('\n') // Membuang karakter newline

	fmt.Println("")

	switch pilihanMenu {
	case 1:
		perulanganLatihan1()
	case 2:
		perulanganLatihan2()
	case 3:
		perulanganLatihan3()
	case 4:
		perulanganLatihan4()
	case 5:
		percabanganLatihan1()
	case 6:
		percabanganLatihan2()
	case 7:
		percabanganLatihan3()
	default:
		fmt.Println("Huh? What number did you just type?")
	}
}