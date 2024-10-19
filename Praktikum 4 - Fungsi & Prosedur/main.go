package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	var pilihanMenu int

	fmt.Println("")
	fmt.Println("-=-=-=- Welcome to Jordan's Praktikum 4 - Fungsi & Prosedur -=-=-=-")
	fmt.Println("")
	fmt.Println("Menu:")
	fmt.Println("1. Contoh Fungsi")
	fmt.Println("2. Latihan 1 Fungsi")
	fmt.Println("3. Latihan 2 Fungsi")
	fmt.Println("4. Latihan 3 Fungsi")
	fmt.Println("5. Contoh Prosedur")
	fmt.Println("6. Latihan 1 Prosedur")
	fmt.Println("7. Latihan 2 Prosedur")
	fmt.Println("8. Latihan 3 Prosedur")

	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihanMenu)

	reader := bufio.NewReader(os.Stdin) // Membuat reader baru dimana Stdin adalah input dari terminal
	reader.ReadString('\n') // Membuang karakter newline

	fmt.Println("")

	switch pilihanMenu {
	case 1:
		fungsi()
	case 2:
		perDanKomFungsi()
	case 3:
		matematika()
	case 4:
		titikLingkaran()
	case 5:
		prosedur()
	case 6:
		perDanKomProsedur()
	case 7:
		soalWaktu()
	case 8:
		deretBilangan()
	default:
		fmt.Println("Huh? What number did you just type?")
	}
}