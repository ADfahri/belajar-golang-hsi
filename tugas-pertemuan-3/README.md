=========================================

Judul : Sistem Informasi Mahasiswa Modular

=========================================

- Struktur folder :

tugas-pertemuan-3/
main.go
go.mod
README.md
mahasiswa/
model.go
logic.go
storage.go
utils.go

==========================================

- Ringkasan Files :

model.go ---> tempat struct mahasiswa dan deskripsi interface
storage.go ---> variabel global (ex :version, maxnilai) dan factory pattern
(pembuatan object mahasiswa)
logic.go ---> implementasi dari deskripsi interface (di model.go)
utils.go ---> menghitung jumlah umur mahasiswa
(menggunakan struct model.go dan di panggil ke main.go)
main.go ---> tempat executable dari semua fitur package mahasiswa

==========================================

- Alur kerja Program :

* Membuat Object Mahasiswa di main.go via Create() dari storage.go.

* TampilkanInfo() dari main.go memanggil method Info(), RataRata(), GetUmur() dari logic.go.

* tampilkanSistem() mengakses variabel global dan fungsi GetMaxNilai() dari storage.go,
  serta HitungTotalUmur() dari utils.go.

===========================================
