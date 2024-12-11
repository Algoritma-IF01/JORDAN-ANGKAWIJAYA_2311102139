package main

import (
	"fmt"
)

func twoOrangKetemuan(day, x, y, count int) int {
	if day > 365 {
		return count
	}
	if day%x == 0 && day%y != 0 {
		count++
	}
	return twoOrangKetemuan(day+1, x, y, count)
}

func soalA3() {
	var x, y int
	fmt.Print("\n-=-=- Jordan's Meeting Program (A3-3) -=-=-\n")
	fmt.Print("Hello! Now, I want you to masukkin nilai x pls: ")
	fmt.Scan(&x)
	fmt.Print("Nice! Now the nilai y pls: ")
	fmt.Scan(&y)

	jumlahPertemuan := twoOrangKetemuan(1, x, y, 0)
	fmt.Print("\n-=-=-=- RESULTS OF PERTEMUAN -=-=-=-\n")
	fmt.Printf("Ok calculated! So the jumlah pertemuan in a year is %d\n", jumlahPertemuan)
	fmt.Print("\n")
}