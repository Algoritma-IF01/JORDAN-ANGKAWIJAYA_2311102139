package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik (a, b) dan (c, d)
func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

// Fungsi untuk memeriksa apakah titik (x, y) berada di dalam lingkaran dengan pusat (cx, cy) dan radius r
func diDalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func titikLingkaran() {
	var cx1, cy1, r1 float64 // Lingkaran 1
	var cx2, cy2, r2 float64 // Lingkaran 2
	var x, y float64         // Titik sembarang

	// Input untuk Lingkaran 1
	fmt.Print("Latihan 3 Fungsi - Titik di dalam Lingkaran\n")
	fmt.Print("Masukkan pusat lingkaran 1 dan radius (cx1 cy1 r1): ")
	fmt.Scan(&cx1, &cy1, &r1)

	// Input untuk Lingkaran 2
	fmt.Print("Masukkan pusat lingkaran 2 dan radius (cx2 cy2 r2): ")
	fmt.Scan(&cx2, &cy2, &r2)

	// Input untuk Titik sembarang
	fmt.Print("Masukkan titik sembarang (x y): ")
	fmt.Scan(&x, &y)

	// Cek apakah titik berada di dalam lingkaran 1 dan/atau lingkaran 2
	dalamLingkaran1 := diDalam(cx1, cy1, r1, x, y) // titik berada di dalam lingkaran 1 jika jarak titik ke pusat lingkaran 1 <= radius lingkaran 1, := berfungsi untuk mendeklarasikan variabel dalamLingkaran1
	dalamLingkaran2 := diDalam(cx2, cy2, r2, x, y) // titik berada di dalam lingkaran 2 jika jarak titik ke pusat lingkaran 2 <= radius lingkaran 2

	// Tentukan keluaran berdasarkan posisi titik
	if dalamLingkaran1 && dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2.")
	} else if dalamLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1.")
	} else if dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2.")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2.")
	}
}
