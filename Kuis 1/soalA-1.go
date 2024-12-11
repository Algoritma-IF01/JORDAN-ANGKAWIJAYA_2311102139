package main

import (
	"fmt"
)

func hitungBiayaSewa(jamDuration int, menitDuration int, isMember bool) int {
	var tarifPerJam int
	if isMember {
		tarifPerJam = 3500
	} else {
		tarifPerJam = 5000
	}

	totalJam := jamDuration
	if jamDuration < 1 {
		totalJam = 1
	} else if menitDuration >= 10 {
		totalJam += 1
	}

	totalBiaya := totalJam * tarifPerJam

	if totalJam > 3 {
		totalBiaya = int(float64(totalBiaya) * 0.9)
	}

	return totalBiaya
}

func soalA1() {
	var jam, menit int
	var isMember bool

	fmt.Print("\n-=-=- Jordan's Bicycle Rental Program (A1-1) -=-=-\n")
	fmt.Print("Ok first, berapa jam you rental bicyclenya: ")
	fmt.Scan(&jam)
	fmt.Print("Second, berapa menit you rental bicyclenya: ")
	fmt.Scan(&menit)
	fmt.Print("Ok now, you member or not?? (true/false) ")
	fmt.Scan(&isMember)

	biaya := hitungBiayaSewa(jam, menit, isMember)

	fmt.Print("\n-=-=-=- YOUR BICYCLE SEWA -=-=-=-\n")
	fmt.Printf("Hooray! From what you input, your bicycle biaya sewa is Rp%d\n", biaya)
	fmt.Print("\n")
}
