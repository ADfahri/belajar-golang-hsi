package main

import (
	"fmt"
	"tugas-pertemuan-3/mahasiswa"
)

//function ini tersambung ke utils.go
//yang di mana Nama, Umur, Nilai tersambung agar meringkas penulisan kode
func main() {
    mhs1 := mahasiswa.Create("Ali", 20, 80, 85, 90, 67)
    mhs2 := mahasiswa.Create("Yono", 24, 85, 90, 83, 87)

	tampilkanInfo(mhs1)
	tampilkanInfo(mhs2)

	tampilkanSistem(mhs1, mhs2)
}

func tampilkanInfo(s mahasiswa.Mahasiswa) {
	fmt.Printf("Nama: %s, Umur: %d\n", s.Info(), s.Umur)
	fmt.Printf("Rata-rata Nilai: %.2f\n", s.RataRata())
	fmt.Println("--------")
}

func tampilkanSistem(m1, m2 mahasiswa.Mahasiswa) {
    fmt.Printf("Versi Package: %s\n", mahasiswa.Version)
    fmt.Printf("Nilai Maksimum: %d\n", mahasiswa.GetMaxNilai())

	totalUmur := mahasiswa.HitungTotalUmur(m1, m2)
	fmt.Printf("Total Umur Mahasiswa: %d\n", totalUmur)
}