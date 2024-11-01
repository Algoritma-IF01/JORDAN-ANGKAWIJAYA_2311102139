package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	var pilihanMenu int

	fmt.Println("")
	fmt.Println("-=-=-=- Welcome to Jordan's Praktikum 5 - Struct & Array -=-=-=-")
	fmt.Println("")
	fmt.Println("Menu:")
	fmt.Println("1. Alias")
	fmt.Println("2. Struct")
	fmt.Println("3. Array")
	fmt.Println("4. Map")
	fmt.Println("5. Latihan 1")
	fmt.Println("6. Latihan 2")
	fmt.Println("7. Latihan 3")
	fmt.Println("8. Latihan 4")

	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihanMenu)

	reader := bufio.NewReader(os.Stdin) // Membuat reader baru dimana Stdin adalah input dari terminal
	reader.ReadString('\n') // Membuang karakter newline

	fmt.Println("")

	switch pilihanMenu {
	case 1:
		alias()
	case 2:
		structParkir()
	case 3:
		array()
	case 4:
		mapName()
	case 5:
		latihan1()
	case 6:
		latihan2()
	case 7:
		latihan3()
	case 8:
		latihan4()
	default:
		fmt.Println("Huh? What number did you just type?")
	}
}