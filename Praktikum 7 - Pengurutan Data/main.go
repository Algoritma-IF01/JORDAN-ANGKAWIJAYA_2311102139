package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var pilihanMenu int

	fmt.Println("")
	fmt.Println("-=-=-=- Welcome to Jordan's Praktikum 7 - Pengurutan Data -=-=-=-")
	fmt.Println("")
	fmt.Println("Menu:")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Selection Sort Struct")
	fmt.Println("3. Insertion Sort")
	fmt.Println("4. Insertion Sort Struct")
	fmt.Println("5. Latihan Selection Sort 1")
	fmt.Println("6. Latihan Selection Sort 2")
	fmt.Println("7. Latihan Selection Sort 3")
	fmt.Println("8. Latihan Insertion Sort 1")
	fmt.Println("9. Latihan Insertion Sort 2")

	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihanMenu)

	reader := bufio.NewReader(os.Stdin) // Membuat reader baru dimana Stdin adalah input dari terminal
	reader.ReadString('\n')             // Membuang karakter newline

	fmt.Println("")

	switch pilihanMenu {
	case 1:
		selectionSort()
	case 2:
		selectionStruct()
	case 3:
		insertionSort()
	case 4:
		insertionStruct()
	case 5:
		latihanSelection1()
	case 6:
		latihanSelection2()
	case 7:
		latihanSelection3()
	case 8:
		latihanInsertion1()
	case 9:
		latihanInsertion2()
	default:
		fmt.Println("Huh? What number did you just type?")
	}
}
