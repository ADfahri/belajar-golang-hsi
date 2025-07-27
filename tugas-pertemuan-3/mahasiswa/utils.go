package mahasiswa

//Menambahkan jumlah umur dari beberapa mahasiswa yang di data
func HitungTotalUmur(mahasiswa ...Mahasiswa) int { 
	total := 0
	for _, s := range mahasiswa {
		total += s.Umur
	}
	return total
}
//memakai struct mahasiswa dari model.go untuk dipanggil ke main.go nantinya