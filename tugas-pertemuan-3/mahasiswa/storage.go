package mahasiswa

var Version string = "v1.0.0"
var MaxNilai = 100


func Create(nama string, umur int, nilai ...int) Mahasiswa {
    return Mahasiswa{
        Nama:  nama,
        Umur:  umur,
        Nilai: nilai,
    }
}

func GetMaxNilai() int {
	return MaxNilai
}