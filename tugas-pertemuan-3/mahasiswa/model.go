package mahasiswa // Deklarasi bagian dari package mahasiswa

type Mahasiswa struct { //structur mahasiswa sesuai yang di pakai
	Nama string
	Nilai []int
	Umur int
	nilaiAvg float64
}

type Deskripsi interface { //interface deskripsi
	Info() string
	RataRata() float64
	GetUmur() int
}


