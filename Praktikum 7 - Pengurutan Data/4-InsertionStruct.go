package main

import "fmt"

type mahasiswa2 struct {
	nama, nim, kelas, jurusan string
	ipk                       float64
}

type arrMhs2 [2023]mahasiswa2

func insertionSort2(T *arrMhs2, n int) {
	/* I.S. terdefinisi array T yang berisi n data mahasiswa
	   F.S. array T terurut secara mengecil (descending) berdasarkan nama
	   dengan menggunakan algoritma INSERTION SORT */
	var temp mahasiswa2
	var i, j int

	for i = 1; i < n; i++ {
		temp = T[i] // Simpan elemen ke-i
		j = i       // Inisialisasi indeks pembanding

		// Geser elemen-elemen sebelumnya
		for j > 0 && temp.nama > T[j-1].nama {
			T[j] = T[j-1]
			j--
		}

		// Tempatkan temp pada posisi yang sesuai
		T[j] = temp
	}
}

func insertionStruct() {
	// Contoh data mahasiswa
	var T arrMhs2
	T[0] = mahasiswa2{"Charlie", "125", "A", "Teknik Informatika", 3.5}
	T[1] = mahasiswa2{"Alice", "123", "A", "Teknik Informatika", 3.8}
	T[2] = mahasiswa2{"Bob", "124", "B", "Sistem Informasi", 3.2}
	T[3] = mahasiswa2{"Diana", "126", "B", "Sistem Informasi", 3.9}
	n := 4

	fmt.Println("Data mahasiswa sebelum diurutkan:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s - %s - %s - %s - %.2f\n", T[i].nama, T[i].nim, T[i].kelas, T[i].jurusan, T[i].ipk)
	}

	insertionSort2(&T, n)

	fmt.Println("\nData mahasiswa setelah diurutkan berdasarkan nama (descending):")
	for i := 0; i < n; i++ {
		fmt.Printf("%s - %s - %s - %s - %.2f\n", T[i].nama, T[i].nim, T[i].kelas, T[i].jurusan, T[i].ipk)
	}
}
