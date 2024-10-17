package main

import (
	"fmt"
)

func latihan3() {
	var r, luas, volume float64

	fmt.Print("Jejari: ")
	fmt.Scanln(&r)
	luas = 4 * 3.14 * r * r
	volume = (4.0 / 3.0) * 3.14 * r * r * r
	fmt.Println("Bola dengan jejari", r, "memiliki volume", volume, "dan luas kulit", luas)
}