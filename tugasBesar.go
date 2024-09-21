package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const NMAX int = 100

type user struct {
	usn         string
	pw          string
	nama        string
	umur        int
	jumlahTeman int
	teman       [NMAX]friend
	statusUsr   string
}

type friend struct {
	usn    string
	status string
	komen  string
	umur   int
}

type tabUser [NMAX]user

func main() {
	var data tabUser
	var nData int

	// Inisialisasi data user
	data[0] = user{"nesya", "123", "nesya", 18, 2, [NMAX]friend{}, "apa kabar Semuanya?"}
	data[1] = user{"dante", "456", "dante", 19, 1, [NMAX]friend{}, "login"}
	data[2] = user{"nomi", "789", "nomi", 20, 1, [NMAX]friend{}, "ayo Beli Tiket"}
	nData = 3

	// Inisialisasi data teman untuk setiap user
	data[0].teman[0] = friend{usn: data[1].usn, status: data[1].statusUsr, umur: data[1].umur}
	data[0].teman[1] = friend{usn: data[2].usn, status: data[2].statusUsr, umur: data[2].umur}
	data[1].teman[0] = friend{usn: data[0].usn, status: data[0].statusUsr, umur: data[0].umur}
	data[2].teman[0] = friend{usn: data[0].usn, status: data[0].statusUsr, umur: data[0].umur}

	halamanPertama(&data, &nData)
}

func halamanPertama(A *tabUser, n *int) {
	/*menampilkan halaman pertama yang akan di lihat user,
	memberikan pilihan kepada user untuk registrasi/login, kemudian mengarahkan ke procedure registrasi/login*/

	var pilih int

	fmt.Println("--------------------------------")
	fmt.Println("    SELAMAT DATANG DI TYPE")
	fmt.Println("--------------------------------")
	fmt.Println("Apa yang ingin anda lakukan?")
	fmt.Println("1. Registrasi")
	fmt.Println("2. Login")
	fmt.Println("3. Keluar")
	fmt.Println("--------------------------------")
	fmt.Print("Pilih(1/2/3): ")

	fmt.Scan(&pilih)
	if pilih == 1 {
		clearScreen()
		registrasi(A, n)
	} else if pilih == 2 {
		clearScreen()
		login(A, n)
	} else if pilih == 3 {
		clearScreen()
		fmt.Println("Terima kasih telah menggunakan Type!")
	}
}

func registrasi(A *tabUser, n *int) {
	/*IS : Array A dan  n banyak data terdefinisi
	Proses: meminta inputan username dan password yang diinginkan oleh user,
	dan memeriksa apakah username sudah di gunakan atau belum dengan menggunakan function sequential search.
	jika username telah digunakan, meminta inputan ulang
	FS : array tabUser bertambah 1*/

	var usn, pw string
	var idx, pilih int

	fmt.Print("Masukan Username: ")
	fmt.Scan(&usn)
	idx = sequentialSearch(*A, *n, usn)

	if idx != -1 {
		clearScreen()
		fmt.Println("Username telah digunakan!")
		fmt.Println("--------------------------------")
		fmt.Println("1. Registrasi ulang")
		fmt.Println("2. Keluar")
		fmt.Println("--------------------------------")
		fmt.Print("Pilih(1/2): ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			clearScreen()
			registrasi(A, n)
		} else {
			clearScreen()
			fmt.Println("Terima kasih telah menggunakan Type!")
		}
	} else {
		fmt.Print("Masukan Password: ")
		fmt.Scan(&pw)
		A[*n].usn = usn
		A[*n].pw = pw
		*n++

		clearScreen()
		fmt.Println("Registrasi berhasil!")
		fmt.Println("--------------------------------")
		fmt.Println("1. Login")
		fmt.Println("2. Keluar")
		fmt.Println("--------------------------------")
		fmt.Print("Pilih(1/2): ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			clearScreen()
			login(A, n)
		} else if pilih == 2 {
			clearScreen()
			fmt.Println("Terima kasih telah menggunakan Type!")
		}
	}
}

func login(A *tabUser, n *int) {
	/*IS : Array A, n banyak data, dan idx sebagai indeks user terdefinisi
	FS : meminta inputan username dari user. kemudian menggunakan sequential search untuk mencari tau indeks dari username agar bisa mengetahui password
	yang benar. atau jika indeks masih -1 maka username tidak valid.*/

	var usn, pw string
	var idx int

	for i := 0; i < 3; i++ {
		fmt.Println("--------------------------------")
		fmt.Print("Masukkan Username: ")
		fmt.Scan(&usn)
		fmt.Print("Masukkan Password: ")
		fmt.Scan(&pw)

		idx = sequentialSearch(*A, *n, usn)

		if idx != -1 && A[idx].pw == pw {
			clearScreen()
			fmt.Println("--------------------------------")
			fmt.Printf("Selamat datang %s\n", A[idx].usn)
			home(A, n, idx)
			break
		} else {
			fmt.Println("Username/Password tidak valid!")
		}
	}
}

func home(A *tabUser, n *int, idx int) {
	/*IS : Array A, n banyak data, dan idx sebagai indeks user terdefinisi
	FS : menampilkan menu utama setelah login. Kemudian memberikan beberapa pilihan kepada user, kemudian mengarahkan sesuai dengan inputan yang di pilih oleh user.*/

	var pilih int
	var choose int

	fmt.Println("--------------------------------")
	fmt.Println("Apa yang ingin anda lakukan?")
	fmt.Println("1. Lihat status teman")
	fmt.Println("2. Tambah/hapus teman")
	fmt.Println("3. Edit Profile")
	fmt.Println("4. Daftar teman")
	fmt.Println("5. Cari user")
	fmt.Println("6. Buat status")
	fmt.Println("7. Keluar")
	fmt.Println("--------------------------------")
	fmt.Print("Pilih(1/2/3/4/5/6/7): ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		clearScreen()
		MenampilkanStatus(A, n, idx)
	} else if pilih == 2 {
		clearScreen()
		fmt.Println("--------------------------------")
		fmt.Println("1. Tambah teman")
		fmt.Println("2. Hapus teman")
		fmt.Println("--------------------------------")
		fmt.Print("Pilih(1/2): ")
		fmt.Scan(&choose)

		if choose == 1 {
			clearScreen()
			MenambahkanTeman(A, n, idx)
		} else if choose == 2 {
			clearScreen()
			menghapusTeman(A, n, idx)
		}

	} else if pilih == 3 {
		clearScreen()
		mengeditProfile(A, n, idx)
	} else if pilih == 4 {
		clearScreen()
		menampilkanDaftarTeman(A, n, idx)
	} else if pilih == 5 {
		clearScreen()
		mencariUser(A, n, idx)
	} else if pilih == 6 {
		clearScreen()
		membuatStatus(A, n, idx)
	} else {
		clearScreen()
		fmt.Println("Terima kasih telah menggunakan Type!")
	}
}

func MenampilkanStatus(A *tabUser, n *int, idx int) {
	/* IS : Array A, n banyak data, dan idx sebagai indeks user terdefinisi
	FS : Menampilkan status teman jika jumlah teman user > 0. Jika = 0 maka user akan di arahkan ke procedure tambah teman.
	Jika jumlah teman > 0 maka user bisa memilih username teman yang ingin di komentari dan memberikan komentar.*/

	var pilih, choose int
	var usnTeman, komentar string

	if A[idx].jumlahTeman == 0 {
		fmt.Println("--------------------------------")
		fmt.Println("Anda belum memiliki teman!")
		fmt.Println("--------------------------------")
		fmt.Println("1. Tambah teman")
		fmt.Println("2. Kembali")
		fmt.Println("--------------------------------")
		fmt.Print("Pilih(1/2): ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			clearScreen()
			MenambahkanTeman(A, n, idx)
		} else {
			clearScreen()
			home(A, n, idx)
		}
	} else {
		fmt.Println("--------------------------------")
		fmt.Println("Status Teman:")

		for i := 0; i < A[idx].jumlahTeman; i++ {
			fmt.Printf("%s: %s\n", A[idx].teman[i].usn, A[idx].teman[i].status)
		}
		fmt.Print("Masukkan username teman yang ingin dikomentari: ")
		fmt.Scan(&usnTeman)

		var temanDitemukan bool
		for j := 0; j < A[idx].jumlahTeman; j++ {
			if usnTeman == A[idx].teman[j].usn {
				temanDitemukan = true
				clearScreen()
				fmt.Printf("Masukkan Komentar untuk %s:\n ", A[idx].teman[j].usn)
				fmt.Println("(Note: akhiri dengan \".\")")
				komentar = inputString()
				clearScreen()
				fmt.Printf("%s mengomentari status %s: '%s'\n", A[idx].usn, A[idx].teman[j].usn, komentar)

				fmt.Println("--------------------------------")
				fmt.Println("1. Kembali")
				fmt.Println("2. Keluar")
				fmt.Println("--------------------------------")
				fmt.Print("Pilih(1/2): ")
				fmt.Scan(&choose)

				if choose == 1 {
					clearScreen()
					home(A, n, idx)
				} else {
					clearScreen()
					fmt.Println("Terima kasih telah menggunakan Type!")
					return
				}
			}
		}

		if !temanDitemukan {
			fmt.Println("Username teman tidak ditemukan dalam daftar teman Anda.")
			fmt.Println("--------------------------------")
			fmt.Println("1. Kembali")
			fmt.Println("2. Keluar")
			fmt.Println("--------------------------------")
			fmt.Print("Pilih(1/2): ")
			fmt.Scan(&choose)

			if choose == 1 {
				clearScreen()
				home(A, n, idx)
			} else {
				clearScreen()
				fmt.Println("Terima kasih telah menggunakan Type!")
			}
		}
	}
}

func MenambahkanTeman(A *tabUser, n *int, idx int) {
	/*IS : Array A, n banyak data, dan idx sebagai indeks user terdefinisi
	FS : User diminta untuk memasukan username teman yang ingin di tambahkan.
	jika username teman terdaftar di array user, maka dapat di tambahkan sebagai teman.
	jika tidak maka akan tercetak di layar "Pengguna tidak ditemukan" */

	var namaTeman string
	var friendIdx, choose int

	fmt.Print("Masukkan username teman yang ingin ditambahkan: ")
	fmt.Scan(&namaTeman)

	friendIdx = -1
	for i := 0; i < *n; i++ {
		if (*A)[i].usn == namaTeman {
			friendIdx = i
			break
		}
	}

	if friendIdx != -1 {
		if (*A)[idx].jumlahTeman < NMAX && (*A)[friendIdx].jumlahTeman < NMAX {
			// Tambahkan friendIdx ke teman user idx
			(*A)[idx].teman[(*A)[idx].jumlahTeman] = friend{
				usn:    (*A)[friendIdx].usn,
				status: (*A)[friendIdx].statusUsr,
				umur:   (*A)[friendIdx].umur,
			}
			(*A)[idx].jumlahTeman++

			// Tambahkan idx ke teman user friendIdx
			(*A)[friendIdx].teman[(*A)[friendIdx].jumlahTeman] = friend{
				usn:    (*A)[idx].usn,
				status: (*A)[idx].statusUsr,
				umur:   (*A)[idx].umur,
			}
			(*A)[friendIdx].jumlahTeman++

			clearScreen()
			fmt.Printf("%s berhasil ditambahkan sebagai teman!\n", (*A)[friendIdx].usn)
		} else {
			fmt.Println("Salah satu dari kalian sudah memiliki jumlah teman maksimal.")
		}

		fmt.Println("--------------------------------")
		fmt.Println("1. Kembali")
		fmt.Println("2. Keluar")
		fmt.Println("--------------------------------")
		fmt.Print("Pilih(1/2): ")
		fmt.Scan(&choose)

		if choose == 1 {
			clearScreen()
			home(A, n, idx)
		} else {
			clearScreen()
			fmt.Println("Terima kasih telah menggunakan Type!")
		}

	} else {
		fmt.Println("Pengguna tidak ditemukan.")
		fmt.Println("--------------------------------")
		fmt.Println("1. Kembali")
		fmt.Println("2. Keluar")
		fmt.Println("--------------------------------")
		fmt.Print("Pilih(1/2): ")
		fmt.Scan(&choose)

		if choose == 1 {
			clearScreen()
			home(A, n, idx)
		} else {
			clearScreen()
			fmt.Println("Terima kasih telah menggunakan Type!")
		}
	}
}

func sequentialSearchFriend(teman [NMAX]friend, jumlahTeman int, usnTeman string) int {
	/*Menggembalikan idx dari username jika terdapat di array teman. Jika tidak, maka mengembalikan -1*/

	for i := 0; i < jumlahTeman; i++ {
		if teman[i].usn == usnTeman {
			return i
		}
	}
	return -1
}

func menghapusTeman(A *tabUser, n *int, idx int) {
	/*IS : Array A, n banyak data, dan idx sebagai indeks user tedefinisi
	FS : Menghapus username yang diinginkan user dari Array teman. jumlah teman berkurang 1*/

	var namaTeman string
	var idxT, choose int

	fmt.Println("--------------------------------")
	fmt.Print("Masukkan username yang ingin dihapus: ")
	fmt.Scan(&namaTeman)

	idxT = sequentialSearchFriend(A[idx].teman, A[idx].jumlahTeman, namaTeman)

	if idxT != -1 {
		for i := idxT; i < A[idx].jumlahTeman-1; i++ {
			A[idx].teman[i] = A[idx].teman[i+1]
		}
		A[idx].jumlahTeman--
		clearScreen()
		fmt.Printf("%s telah dihapus dari daftar teman.\n", namaTeman)
	} else {
		fmt.Println("Teman tidak ditemukan dalam daftar teman Anda.")
	}

	fmt.Println("--------------------------------")
	fmt.Println("1. Kembali")
	fmt.Println("2. Keluar")
	fmt.Println("--------------------------------")
	fmt.Print("Pilih(1/2): ")
	fmt.Scan(&choose)

	if choose == 1 {
		clearScreen()
		home(A, n, idx)
	} else {
		clearScreen()
		fmt.Println("Terima kasih telah menggunakan Type!")
	}
}

func mengeditProfile(A *tabUser, n *int, idx int) {
	/*IS : Array A, n banyak data, dan idx sebagai indeks user tedefinisi
	FS : memperbaharui username, password, nama, atau umur sesuai keinginan user.
	(Hanya bisa memnggunakan username yang belum digunakan pengguna lainnya)*/

	var pilih, choose int

	var usn, pw, nn string
	var umr int

	fmt.Println("--------------------------------")
	fmt.Println("1. Username")
	fmt.Println("2. Password")
	fmt.Println("3. Nama")
	fmt.Println("4. Umur")
	fmt.Println("5. Kembali")
	fmt.Println("--------------------------------")
	fmt.Print("Pilih(1/2/3/4/5): ")
	fmt.Scan(&pilih)

	clearScreen()

	if pilih == 1 {
		fmt.Print("Masukkan username baru: ")
		fmt.Scan(&usn)

		// Cari apakah username sudah digunakan oleh pengguna lain
		existingIdx := sequentialSearch(*A, *n, usn)
		if existingIdx != -1 && existingIdx != idx {
			fmt.Println("Username telah digunakan oleh pengguna lain!")
		} else {
			A[idx].usn = usn
			fmt.Printf("Username berhasil diperbarui menjadi %s\n", A[idx].usn)
		}
		fmt.Println("--------------------------------")
		fmt.Println("1. Kembali")
		fmt.Println("2. Keluar")
		fmt.Println("--------------------------------")
		fmt.Print("Pilih(1/2): ")
		fmt.Scan(&choose)

		if choose == 1 {
			clearScreen()
			mengeditProfile(A, n, idx)
		} else {
			clearScreen()
			fmt.Println("Terima kasih telah menggunakan Type!")
		}
	} else if pilih == 2 {
		fmt.Print("Masukkan password baru: ")
		fmt.Scan(&pw)
		A[idx].pw = pw
		fmt.Println("Password berhasil diperbarui!")
		fmt.Println("--------------------------------")
		fmt.Println("1. Kembali")
		fmt.Println("2. Keluar")
		fmt.Println("--------------------------------")
		fmt.Print("Pilih(1/2): ")
		fmt.Scan(&choose)

		if choose == 1 {
			clearScreen()
			mengeditProfile(A, n, idx)
		} else {
			clearScreen()
			fmt.Println("Terima kasih telah menggunakan Type!")
		}
	} else if pilih == 3 {
		fmt.Print("Masukkan Nama: ")
		fmt.Scan(&nn)
		A[idx].nama = nn
		fmt.Printf("Nama berhasil diperbarui menjadi %s\n", A[idx].nama)
		fmt.Println("--------------------------------")
		fmt.Println("1. Kembali")
		fmt.Println("2. Keluar")
		fmt.Println("--------------------------------")
		fmt.Print("Pilih(1/2): ")
		fmt.Scan(&choose)

		if choose == 1 {
			clearScreen()
			mengeditProfile(A, n, idx)
		} else {
			clearScreen()
			fmt.Println("Terima kasih telah menggunakan Type!")
		}
	} else if pilih == 4 {
		fmt.Print("Masukkan umur: ")
		fmt.Scan(&umr)
		A[idx].umur = umr
		fmt.Printf("Umur berhasil diperbarui menjadi %d\n", A[idx].umur)
		fmt.Println("--------------------------------")
		fmt.Println("1. Kembali")
		fmt.Println("2. Keluar")
		fmt.Println("--------------------------------")
		fmt.Print("Pilih(1/2): ")
		fmt.Scan(&choose)

		if choose == 1 {
			clearScreen()
			mengeditProfile(A, n, idx)
		} else {
			clearScreen()
			fmt.Println("Terima kasih telah menggunakan Type!")
		}
	} else if pilih == 5 {
		clearScreen()
		home(A, n, idx)
	}
}

func menampilkanDaftarTeman(A *tabUser, n *int, idx int) {
	/*IS : Array A, n banyak data, dan idx sebagai indeks user tedefinisi
	FS : Jika jumlahTeman dari user > 0 maka memberikan pilihan untuk menampilkan daftar teman berdasarkan umur(desscending atau asscending) maupun
	berdasarkan panjang username teman(desscending atau asscending).*/

	var pilih, pilih1, pilih2, choose int
	if A[idx].jumlahTeman == 0 {
		fmt.Println("--------------------------------")
		fmt.Println("Anda belum memiliki teman!")
		fmt.Println("--------------------------------")
		fmt.Println("1. Tambah teman")
		fmt.Println("2. Kembali")
		fmt.Println("--------------------------------")
		fmt.Print("Pilih(1/2): ")
		fmt.Scan(&pilih2)

		if pilih2 == 1 {
			clearScreen()
			MenambahkanTeman(A, n, idx)
		} else {
			clearScreen()
			home(A, n, idx)
		}
	} else {
		fmt.Println("----------------------------------------------------------")
		fmt.Println("          Menampilkan daftar teman berdasarkan")
		fmt.Println("----------------------------------------------------------")
		fmt.Println("1. Umur")
		fmt.Println("2. Panjang Username")
		fmt.Println("----------------------------------------------------------")
		fmt.Print("Pilih(1/2): ")
		fmt.Scan(&pilih1)

		if pilih1 == 1 {
			clearScreen()
			fmt.Println("----------------------------------------------------------")
			fmt.Println("1. Menampilkan daftar teman terurut dari umur tertua")
			fmt.Println("2. Menampilkan daftar teman terurut dari umur termuda")
			fmt.Println("----------------------------------------------------------")
			fmt.Print("Pilih(1/2): ")
			fmt.Scan(&pilih)

			if pilih == 1 {
				clearScreen()
				fmt.Println("Berikut adalah daftar teman: ")
				selectionSort(&A[idx], false) // Urutkan dari umur tertua
				cetak(&A[idx])
			} else if pilih == 2 {
				clearScreen()
				fmt.Println("Berikut adalah daftar teman: ")
				selectionSort(&A[idx], true) // Urutkan dari umur termuda
				cetak(&A[idx])
			}
		} else if pilih1 == 2 {
			clearScreen()
			fmt.Println("----------------------------------------------------------")
			fmt.Println("1. Menampilkan daftar teman terurut dari username terpanjang")
			fmt.Println("2. Menampilkan daftar teman terurut dari username terpendek")
			fmt.Println("----------------------------------------------------------")
			fmt.Print("Pilih(1/2): ")
			fmt.Scan(&pilih)

			if pilih == 1 {
				clearScreen()
				fmt.Println("Berikut adalah daftar teman: ")
				insertionSort(&A[idx], false) // Urutkan dari username terpanjang
				cetak(&A[idx])
			} else if pilih == 2 {
				clearScreen()
				fmt.Println("Berikut adalah daftar teman: ")
				insertionSort(&A[idx], true) // Urutkan dari username terpendek
				cetak(&A[idx])
			}
		}

		fmt.Println("--------------------------------")
		fmt.Println("1. Kembali")
		fmt.Println("2. Keluar")
		fmt.Println("--------------------------------")
		fmt.Print("Pilih(1/2): ")
		fmt.Scan(&choose)

		if choose == 1 {
			clearScreen()
			home(A, n, idx)
		} else {
			clearScreen()
			fmt.Println("Terima kasih telah menggunakan Type!")
		}
	}
}

func selectionSort(u *user, ascending bool) {
	/*IS : Array u terdefinisi. terdefinisi pula boolean true untuk ascending dan false untuk descending
	FS : Array teman terurut berdasarkan umur dan ascending/descending sesuai boolean.*/

	for i := 0; i < u.jumlahTeman-1; i++ {
		var idx int
		if ascending {
			idx = findMin(u, i)
		} else {
			idx = findMax(u, i)
		}
		u.teman[i], u.teman[idx] = u.teman[idx], u.teman[i]
	}
}

func findMin(u *user, start int) int {
	/*Array u terdefinisi dan start terdefinisi sebagai indeks di mulainya pencarian nilai ekstrim.
	Mengembalikan indeks dari array teman dengan umur termuda*/

	minIdx := start
	for j := start + 1; j < u.jumlahTeman; j++ {
		if u.teman[j].umur < u.teman[minIdx].umur {
			minIdx = j
		}
	}
	return minIdx
}

func findMax(u *user, start int) int {
	/*Array u terdefinisi dan start terdefinisi sebagai indeks di mulainya pencarian nilai ekstrim.
	Mengembalikan indeks dari array teman dengan umur tertua*/

	maxIdx := start
	for j := start + 1; j < u.jumlahTeman; j++ {
		if u.teman[j].umur > u.teman[maxIdx].umur {
			maxIdx = j
		}
	}
	return maxIdx
}

func insertionSort(u *user, ascending bool) {
	/*IS : Array u terdefinisi. terdefinisi pula boolean true untuk ascending dan false untuk descending
	FS : Array teman terurut berdasarkan panjang username teman dari user secara ascending/descending sesuai boolean.*/

	for i := 1; i < u.jumlahTeman; i++ {
		key := u.teman[i]
		j := i - 1

		for j >= 0 && ((ascending && len(u.teman[j].usn) > len(key.usn)) || (!ascending && len(u.teman[j].usn) < len(key.usn))) {
			u.teman[j+1] = u.teman[j]
			j--
		}
		u.teman[j+1] = key
	}
}

func cetak(u *user) {
	/*IS : Array u terdefinisi
	FS : Tercetak di layar usename teman, status teman dan umur teman dari user*/

	for i := 0; i < u.jumlahTeman; i++ {
		fmt.Printf("Teman: %s, Status: %s, Umur: %d\n", u.teman[i].usn, u.teman[i].status, u.teman[i].umur)
	}
}

func mencariUser(A *tabUser, n *int, idx int) {
	/*IS : Array A, n banyak data, dan idx sebagai indeks user terdefinisi
	FS : meminta inputan username yang ingin di cari.
	kemudian mencetak keterangan bahwa username yang di cari ditemukan atau tidak.
	jika di temukan maka akan di cetak keterangan bahwa user saling berteman atau tidak*/

	var usn string
	var choose int

	fmt.Print("Masukkan username yang ingin dicari: ")
	fmt.Scan(&usn)

	userIdx := sequentialSearch(*A, *n, usn)
	if userIdx != -1 {
		clearScreen()
		fmt.Printf("User ditemukan: %s (Umur: %d)\n", (*A)[userIdx].usn, (*A)[userIdx].umur)

		isFriend := binarySearch(&(*A)[idx], usn)
		if isFriend {
			fmt.Println("Berteman")
		} else {
			fmt.Println("Tidak berteman")
		}

		fmt.Println("--------------------------------")
		fmt.Println("1. Kembali")
		fmt.Println("2. Keluar")
		fmt.Println("--------------------------------")
		fmt.Print("Pilih(1/2): ")
		fmt.Scan(&choose)

		if choose == 1 {
			clearScreen()
			home(A, n, idx)
		} else {
			clearScreen()
			fmt.Println("Terima kasih telah menggunakan Type!")
		}
	} else {
		clearScreen()
		fmt.Println("Pengguna tidak ditemukan!")
		fmt.Println("--------------------------------")
		fmt.Println("1. Kembali")
		fmt.Println("2. Keluar")
		fmt.Println("--------------------------------")
		fmt.Print("Pilih(1/2): ")
		fmt.Scan(&choose)

		if choose == 1 {
			clearScreen()
			home(A, n, idx)
		} else {
			clearScreen()
			fmt.Println("Terima kasih telah menggunakan Type!")
		}
	}
}

func binarySearch(u *user, usn string) bool {
	/*IS : Array u terdefinisi. terdefinisi pula string username yang ingin di cari
	FS : mengembalikan boolean true jika username di temukan dan false jika tidak.*/

	var left, right, mid int
	var ketemu bool
	left = 0
	right = u.jumlahTeman - 1
	ketemu = false

	for left <= right && !ketemu {
		mid = (left + right) / 2
		if u.teman[mid].usn < usn {
			left = mid + 1
		} else if u.teman[mid].usn > usn {
			right = mid - 1
		} else {
			ketemu = true
		}
	}

	return ketemu
}

func membuatStatus(A *tabUser, n *int, idx int) {
	/*/*IS : Array A, n banyak data, dan idx sebagai indeks user terdefinisi
	FS : meminta inputan berupa string untuk di masukan kedalam A[i].statusUsr*/

	var sts string
	var y int
	var choose int

	fmt.Println("Apa yang anda ingin bagikan hari ini? ")
	fmt.Println("(Note: akhiri dengan \".\")")
	sts = inputString()
	fmt.Println()
	fmt.Println("--------------------------------")
	fmt.Println("Apakah anda yakin untuk posting?")
	fmt.Println("1. Yakin")
	fmt.Println("2. Tidak")
	fmt.Print("Pilih(1/2): ")
	fmt.Scan(&y)

	if y == 1 {
		clearScreen()
		A[idx].statusUsr = sts
		fmt.Println("Status sudah di upload!")
		fmt.Printf("%s: %s\n", A[idx].usn, A[idx].statusUsr)
		fmt.Println()
		fmt.Println("--------------------------------")
		fmt.Println("1. Kembali")
		fmt.Println("2. Keluar")
		fmt.Println("--------------------------------")
		fmt.Print("Pilih(1/2): ")
		fmt.Scan(&choose)

		if choose == 1 {
			clearScreen()
			home(A, n, idx)
		} else {
			clearScreen()
			fmt.Println("Terima kasih telah menggunakan Type!")
		}
	} else if y == 2 {
		clearScreen()
		fmt.Println("Status tidak di upload.")
		fmt.Println("--------------------------------")
		fmt.Println("1. Kembali")
		fmt.Println("2. Keluar")
		fmt.Println("--------------------------------")
		fmt.Print("Pilih(1/2): ")
		fmt.Scan(&choose)

		if choose == 1 {
			clearScreen()
			home(A, n, idx)
		} else {
			clearScreen()
			fmt.Println("Terima kasih telah menggunakan Type!")
		}
	}

}

func inputString() string {
	/*Menerima inputan string dan menyimpan di dalam satu variabel. imputan berhenti jika diberikan "." */
	var input, result string

	for {
		fmt.Scan(&input)
		if input == "." {
			break
		}
		if result == "" {
			result = input
		} else {
			result += " " + input
		}
	}

	return result
}

func clearScreen() {
	/*fungsi untuk membersihkan layar, digunakan untuk membersihkan layar di saat yang diinginkan*/

	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func sequentialSearch(A tabUser, n int, username string) int {
	/*array A dan n banyak data terdefinisi
	mengembalikan index dari array jika username di temukan dalam array A.
	jika tidak di temukan maka mengembalikan -1*/

	var i, index int
	i = 0
	index = -1

	// Melakukan iterasi hingga ditemukan atau mencapai akhir array
	for i < n && index == -1 {
		if A[i].usn == username {
			index = i
		}
		i++
	}

	return index
}
