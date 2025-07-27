//implementasikan deskripsi interface dari model.go

package mahasiswa


func (m Mahasiswa) Info() string {
	return m.Nama
}

func (m Mahasiswa) GetUmur() int {
	return m.Umur
}

func (m Mahasiswa) RataRata() float64 {
	if len(m.Nilai) == 0 {
		return 0
	}

	//menjumlahkan total nilai, lalu di bagi dengan jumlah nilai
	total := 0
	for _, nilai := range m.Nilai {
		total += nilai
	}
	m.nilaiAvg = float64(total) / float64(len(m.Nilai))
	return m.nilaiAvg
}

