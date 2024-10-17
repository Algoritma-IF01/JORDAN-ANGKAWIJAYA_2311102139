package main

import (
	"fmt"
)

func latihan2() {
	var tahun int
	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&tahun)

	if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0) {
		fmt.Println("Tahun " + fmt.Sprint(tahun) + " adalah tahun kabisat")
	} else {
		fmt.Println("Tahun " + fmt.Sprint(tahun) + " bukan tahun kabisat")
	}
}