package main

import (
	"fmt"
)

// Define a constant and data type const NMAX = 100
const NMAX = 100

// Define a struct type for a room
type Kamar struct {
	NoKamar   int    // Room number
	TipeKamar string // Room type
	HargaSewa int    // Room rental price
	Status    int    // Room status: 1 = kosong, 2 = dipesan, 3 = terisi
	Fasilitas string // Room facilities
}

// Define a struct type for a tenant
type Penghuni struct {
	Nama        string // Tenant name
	NoKamar     int    // Tenant room number
	NoIdentitas int    // Tenant ID number
	NoTelp      string // Tenant phone number
	Alamat      string // Tenant address
	TglMasuk    string // Tenant check-in date
	Durasi      int    // Duration of stay
}

type daftarKamar [NMAX]Kamar
type daftarPenghuni [NMAX]Penghuni

// Define a function for adding a tenant
func tambahPenghuni(p *daftarPenghuni, k *daftarKamar, n *int, m int) {
	var ada bool
	if *n >= NMAX {
		fmt.Println("Data penghuni sudah penuh")
		return
	}
	var noKamar int
	fmt.Println("Masukkan data penghuni:")
	fmt.Print("Nama: ")
	fmt.Scan(&p[*n].Nama)

	// Periksa ketersediaan kamar
	for !ada {
		fmt.Print("Nomor Kamar: ")
		fmt.Scan(&noKamar)
		for i := 0; i < m && !ada; i++ {
			if k[i].NoKamar == noKamar && (k[i].Status == 1 || k[i].Status == 2) {
				p[*n].NoKamar = noKamar
				k[i].Status = 3
				ada = true
			}
		}
		if !ada {
			fmt.Println("Kamar tidak tersedia atau sudah terisi. Coba lagi.")
		}
	}

	fmt.Print("Nomor Identitas: ")
	fmt.Scan(&p[*n].NoIdentitas)
	fmt.Print("Nomor Telepon: ")
	fmt.Scan(&p[*n].NoTelp)
	fmt.Print("Alamat: ")
	fmt.Scan(&p[*n].Alamat)
	fmt.Print("Tanggal Masuk: ")
	fmt.Scan(&p[*n].TglMasuk)
	fmt.Print("Durasi: ")
	fmt.Scan(&p[*n].Durasi)
	*n += 1
}

// Define a function for adding a room
func tambahKamar(k *daftarKamar, n *int) {
	if *n >= NMAX {
		fmt.Println("Data kamar sudah penuh")
		return
	}
	fmt.Println("Masukkan data kamar:")
	fmt.Print("Nomor Kamar: ")
	fmt.Scan(&k[*n].NoKamar)
	fmt.Print("Tipe Kamar: ")
	fmt.Scan(&k[*n].TipeKamar)
	fmt.Print("Harga Sewa: ")
	fmt.Scan(&k[*n].HargaSewa)
	fmt.Print("Status (1 = kosong/2 = dipesan): ")
	fmt.Scan(&k[*n].Status)
	fmt.Print("Fasilitas: ")
	fmt.Scan(&k[*n].Fasilitas)
	*n += 1
}

// Define a function for deleting a room
func hapusKamar(k *daftarKamar, p *daftarPenghuni, noKamar int, nk *int, np *int) {
	found := false
	for i := 0; i < *nk; i++ {
		if k[i].NoKamar == noKamar {
			found = true
			// Menghapus penghuni yang ada di kamar tersebut
			for j := 0; j < *np; j++ {
				if p[j].NoKamar == noKamar {
					hapusPenghuni(p, p[j].NoIdentitas, np)
					j-- // Kembali ke indeks sebelumnya setelah penghapusan
				}
			}
			// Menggeser elemen-elemen setelah kamar yang dihapus ke posisi sebelumnya
			for j := i; j < *nk-1; j++ {
				k[j] = k[j+1]
			}
			*nk -= 1
			i-- // Kembali ke indeks sebelumnya untuk memeriksa elemen yang digeser
		}
	}
	if !found {
		fmt.Println("Kamar tidak ditemukan!")
	}
}

// Define a function for deleting a tenant
func hapusPenghuni(p *daftarPenghuni, noIdentitas int, n *int) {
	found := false
	for i := 0; i < *n; i++ {
		if p[i].NoIdentitas == noIdentitas {
			found = true
			// Menggeser elemen-elemen setelah penghuni yang dihapus ke posisi sebelumnya
			for j := i; j < *n-1; j++ {
				p[j] = p[j+1]
			}
			*n -= 1
			i-- // Kembali ke indeks sebelumnya untuk memeriksa elemen yang digeser
		}
	}
	if !found {
		fmt.Println("Penghuni tidak ditemukan!")
	}
}

// Define a function for editing a room
func editKamar(k *daftarKamar, noKamar int, n int) {
	var found bool
	for i := 0; i < n && !found; i++ {
		if k[i].NoKamar == noKamar {
			fmt.Println("Masukkan data baru untuk kamar")
			fmt.Print("Nomor Kamar: ")
			fmt.Scan(&k[i].NoKamar)
			fmt.Print("Tipe Kamar: ")
			fmt.Scan(&k[i].TipeKamar)
			fmt.Print("Harga Sewa: ")
			fmt.Scan(&k[i].HargaSewa)
			fmt.Print("Status (1 = kosong/2 = dipesan): ")
			fmt.Scan(&k[i].Status)
			fmt.Print("Fasilitas: ")
			fmt.Scan(&k[i].Fasilitas)
			fmt.Println("Data kamar berhasil diubah.")
			found = true
		}
	}
	if !found {
		fmt.Println("Kamar tidak ditemukan!")
	}
}

// Define a function for editing a tenant
func editPenghuni(p *daftarPenghuni, k *daftarKamar, noIdentitas int, np int, nk int) {
	var found bool = false
	for i := 0; i < np && !found; i++ {
		if p[i].NoIdentitas == noIdentitas {
			fmt.Println("Masukkan data baru untuk penghuni")
			fmt.Print("Nama: ")
			fmt.Scan(&p[i].Nama)
			// Periksa ketersediaan kamar
			for j := 0; j < nk && !found; j++ {
				if k[j].NoKamar == p[i].NoKamar && (k[j].Status == 1 || k[j].Status == 2) {
					k[j].Status = 3
					found = true
				}
			}
			if !found {
				fmt.Println("Kamar tidak tersedia atau sudah terisi. Coba lagi.")
			}
			fmt.Print("Nomor Identitas: ")
			fmt.Scan(&p[i].NoIdentitas)
			fmt.Print("Nomor Telepon: ")
			fmt.Scan(&p[i].NoTelp)
			fmt.Print("Alamat: ")
			fmt.Scan(&p[i].Alamat)
			fmt.Print("Tanggal Masuk: ")
			fmt.Scan(&p[i].TglMasuk)
			fmt.Print("Durasi: ")
			fmt.Scan(&p[i].Durasi)
			fmt.Println("Data penghuni berhasil diubah.")
		}
	}
	if !found {
		fmt.Println("Penghuni tidak ditemukan!")
	}
}

// Fungsi untuk menampilkan data kamar secara terurut
func tampilKamarTerurut(k *daftarKamar, n int) {
	// Mengurutkan data kamar
	selectionSortKamar(k, n)

	// Menampilkan data kamar yang sudah terurut
	fmt.Println("Data kamar:")
	for i := 0; i < n; i++ {
		status := ""
		if k[i].Status == 1 {
			status = "kosong"
		} else if k[i].Status == 2 {
			status = "dipesan"
		} else if k[i].Status == 3 {
			status = "terisi"
		}

		fmt.Printf("Nomor Kamar: %d, Tipe Kamar: %s, Harga Sewa: %d, Status: %s, Fasilitas: %s\n",
			k[i].NoKamar, k[i].TipeKamar, k[i].HargaSewa, status, k[i].Fasilitas)
	}
}

// Fungsi untuk menampilkan data penghuni secara terurut
func tampilPenghuniTerurut(p *daftarPenghuni, n int) {
	// Mengurutkan data penghuni
	insertionSortPenghuni(p, n)

	// Menampilkan data penghuni yang sudah terurut
	fmt.Println("Data penghuni:")
	for i := 0; i < n; i++ {
		fmt.Printf("Nama: %s, No. Identitas: %d, No. Telp: %s, Alamat: %s, Tgl. Masuk: %s, No. Kamar: %d, Durasi: %d\n",
			p[i].Nama, p[i].NoIdentitas, p[i].NoTelp, p[i].Alamat, p[i].TglMasuk, p[i].NoKamar, p[i].Durasi)
	}
}

// Define a function for searching for vacant rooms
func tampilkanKamarKosong(k *daftarKamar, n int) {
	fmt.Println("Kamar kosong:")
	for i := 0; i < n; i++ {
		if k[i].Status == 1 {
			fmt.Printf("Nomor Kamar: %d, Tipe Kamar: %s, Harga Sewa: %d, Fasilitas: %s\n",
				k[i].NoKamar, k[i].TipeKamar, k[i].HargaSewa, k[i].Fasilitas)
		}
	}
}

// Define a function for calculating total rental cost
func hitungTotalBiayaSewa(p *daftarPenghuni, k *daftarKamar, noIdentitas int, np int, nk int) {
	var totalBiaya int
	var found bool

	for i := 0; i < np && !found; i++ {
		if p[i].NoIdentitas == noIdentitas {
			for j := 0; j < nk && !found; j++ {
				if k[j].NoKamar == p[i].NoKamar {
					totalBiaya = k[j].HargaSewa * p[i].Durasi
					found = true
				}
			}
		}
	}

	fmt.Println("Total biaya sewa:", totalBiaya)
}

// Define a function for selection sort operation
func selectionSortKamar(k *daftarKamar, n int) {
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if k[j].NoKamar < k[minIndex].NoKamar {
				minIndex = j
			}
		}
		// Swap data
		k[i], k[minIndex] = k[minIndex], k[i]
	}
}

// Define a function for insertion sort operation
func insertionSortPenghuni(p *daftarPenghuni, n int) {
	for i := 1; i < n; i++ {
		key := p[i]
		j := i - 1
		for j >= 0 && p[j].NoIdentitas > key.NoIdentitas {
			p[j+1] = p[j]
			j--
		}
		p[j+1] = key
	}
}

func main() {
	// Loop utama program
	var banyakDataPenghuni, banyakDataKamar int
	var kamar daftarKamar
	var penghuni daftarPenghuni
	var pilihan int = -1
	for pilihan != 0 {
		fmt.Println("\nSistem Manajemen Kosan")
		fmt.Println("1. Tambah kamar")
		fmt.Println("2. Tambah penghuni")
		fmt.Println("3. Hapus kamar")
		fmt.Println("4. Hapus penghuni")
		fmt.Println("5. Edit kamar")
		fmt.Println("6. Edit penghuni")
		fmt.Println("7. Tampilkan data kamar")
		fmt.Println("8. Tampilkan data penghuni")
		fmt.Println("9. Cari kamar kosong")
		fmt.Println("10. Hitung total biaya sewa")
		fmt.Println("0. Keluar")
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahKamar(&kamar, &banyakDataKamar)
		} else if pilihan == 2 {
			tambahPenghuni(&penghuni, &kamar, &banyakDataPenghuni, banyakDataKamar)
		} else if pilihan == 3 {
			var noKamar int
			fmt.Print("Masukkan nomor kamar yang akan dihapus: ")
			fmt.Scan(&noKamar)
			hapusKamar(&kamar, &penghuni, noKamar, &banyakDataKamar, &banyakDataPenghuni)
		} else if pilihan == 4 {
			var noIdentitas int
			fmt.Print("Masukkan nomor identitas penghuni yang akan dihapus: ")
			fmt.Scan(&noIdentitas)
			hapusPenghuni(&penghuni, noIdentitas, &banyakDataPenghuni)
		} else if pilihan == 5 {
			var noKamar int
			fmt.Print("Masukkan nomor kamar yang akan diubah: ")
			fmt.Scan(&noKamar)
			editKamar(&kamar, noKamar, banyakDataKamar)
		} else if pilihan == 6 {
			var noIdentitas int
			fmt.Print("Masukkan nomor identitas penghuni yang akan diubah: ")
			fmt.Scan(&noIdentitas)
			editPenghuni(&penghuni, &kamar, noIdentitas, banyakDataPenghuni, banyakDataKamar)
		} else if pilihan == 7 {
			tampilKamarTerurut(&kamar, banyakDataKamar)
		} else if pilihan == 8 {
			tampilPenghuniTerurut(&penghuni, banyakDataPenghuni)
		} else if pilihan == 9 {
			tampilkanKamarKosong(&kamar, banyakDataKamar)
		} else if pilihan == 10 {
			var noIdentitas int
			fmt.Print("Masukkan nomor identitas penghuni: ")
			fmt.Scan(&noIdentitas)
			hitungTotalBiayaSewa(&penghuni, &kamar, noIdentitas, banyakDataPenghuni, banyakDataKamar)
		} else if pilihan == 0 {
			fmt.Println("Terima kasih telah menggunakan Sistem Manajemen Kosan!")
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
