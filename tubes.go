package main

import "fmt"

const nmax = 100
const maxMatkul = 10

type MataKuliah struct {
	Nama  string
	SKS   int
	UTS   float64
	UAS   float64
	Quiz  float64
	Total float64
	Grade string
}

type tabMatkul [maxMatkul]MataKuliah

type Mahasiswa struct {
	NIM          string
	Nama         string
	Matkul       tabMatkul
	JumlahMatkul int
}

type tabMahasiswa [nmax]Mahasiswa

func hitungTotal(uts, uas, quiz float64) float64 {
	return (uts + uas + quiz) / 3
}

func hitungGrade(total float64) string {
	if total >= 85 {
		return "A"
	} else if total >= 80 {
		return "AB"
	} else if total >= 70 {
		return "B"
	} else if total >= 65 {
		return "BC"
	} else if total >= 60 {
		return "C"
	} else if total >= 50 {
		return "D"
	}
	return "E"
}

func bobotGrade(grade string) float64 {
	if grade == "A" {
		return 4
	} else if grade == "AB" {
		return 3.5
	} else if grade == "B" {
		return 3
	} else if grade == "BC" {
		return 2.5
	} else if grade == "C" {
		return 2
	} else if grade == "D" {
		return 1
	}
	return 0
}

func hitungIPK(m Mahasiswa) float64 {
	var totalMutu float64
	var totalSKS, i int

	for i = 0; i < m.JumlahMatkul; i++ {
		totalMutu += bobotGrade(m.Matkul[i].Grade) * float64(m.Matkul[i].SKS)
		totalSKS += m.Matkul[i].SKS
	}

	if totalSKS == 0 {
		return 0
	}

	return totalMutu / float64(totalSKS)
}

func hitungSKS(m Mahasiswa) int {
	var i, total int
	for i = 0; i < m.JumlahMatkul; i++ {
		total += m.Matkul[i].SKS
	}
	return total
}

func sequentialSearch(data tabMahasiswa, n int, cari string) int {
	var i int
	for i = 0; i < n; i++ {
		if data[i].Nama == cari {
			return i
		}
	}
	return -1
}

func insertionSortNIM(data *tabMahasiswa, n int) {
	var i, j int
	var temp Mahasiswa

	for i = 1; i < n; i++ {
		temp = data[i]
		j = i - 1

		for j >= 0 && data[j].NIM > temp.NIM {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

func binarySearchNIM(data tabMahasiswa, n int, cari string) int {
	var kiri, kanan, tengah int
	kiri = 0
	kanan = n - 1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2

		if data[tengah].NIM == cari {
			return tengah
		} else if cari < data[tengah].NIM {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	return -1
}

func insertionSortIPKAsc(data *tabMahasiswa, n int) {
	var i, j int
	var temp Mahasiswa

	for i = 1; i < n; i++ {
		temp = data[i]
		j = i - 1

		for j >= 0 && hitungIPK(data[j]) > hitungIPK(temp) {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

func insertionSortIPKDesc(data *tabMahasiswa, n int) {
	var i, j int
	var temp Mahasiswa

	for i = 1; i < n; i++ {
		temp = data[i]
		j = i - 1

		for j >= 0 && hitungIPK(data[j]) < hitungIPK(temp) {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

func selectionSortSKSAsc(data *tabMahasiswa, n int) {
	var i, j, idx int
	var temp Mahasiswa

	for i = 0; i < n-1; i++ {
		idx = i
		for j = i + 1; j < n; j++ {
			if hitungSKS(data[j]) < hitungSKS(data[idx]) {
				idx = j
			}
		}
		temp = data[i]
		data[i] = data[idx]
		data[idx] = temp
	}
}

func selectionSortSKSDesc(data *tabMahasiswa, n int) {
	var i, j, idx int
	var temp Mahasiswa

	for i = 0; i < n-1; i++ {
		idx = i
		for j = i + 1; j < n; j++ {
			if hitungSKS(data[j]) > hitungSKS(data[idx]) {
				idx = j
			}
		}
		temp = data[i]
		data[i] = data[idx]
		data[idx] = temp
	}
}

func validNama(nama string) bool {
	if len(nama) == 0 {
		return false
	}

	for i := 0; i < len(nama); i++ {
		if !(nama[i] >= 'A' && nama[i] <= 'Z' ||
			nama[i] >= 'a' && nama[i] <= 'z' ||
			nama[i] == ' ') {
			return false
		}
	}

	return true
}

func validNIM(nim string) bool {
	if len(nim) != 12 {
		return false
	}

	for i := 0; i < len(nim); i++ {
		if nim[i] < '0' || nim[i] > '9' {
			return false
		}
	}

	return true
}

func tambahMahasiswa(data *tabMahasiswa, n *int) {
	var nama, nim string

	fmt.Print("NIM (12 digit) : ")
    fmt.Scan(&nim)

    for !validNIM(nim) {
	    fmt.Println("NIM harus terdiri dari 12 angka!")
	    fmt.Print("NIM (12 digit) : ")
	    fmt.Scan(&nim)
}

	data[*n].NIM = nim

	fmt.Print("Nama : ")
    fmt.Scan(&nama)

    for !validNama(nama) {
	    fmt.Println("Nama hanya boleh berisi huruf!")
	    fmt.Print("Nama : ")
	    fmt.Scan(&nama)
}

data[*n].Nama = nama

	*n = *n + 1

	insertionSortNIM(data, *n)

	fmt.Println("Mahasiswa berhasil ditambahkan")
}

func tambahMatkul(data *tabMahasiswa, n int) {
	var nim string
	var idx, j int

	insertionSortNIM(data, n)
	fmt.Print("NIM : ")
	fmt.Scan(&nim)

	idx = binarySearchNIM(*data, n, nim)

	if idx != -1 {
		j = data[idx].JumlahMatkul

		fmt.Print("Nama Mata Kuliah : ")
		fmt.Scan(&data[idx].Matkul[j].Nama)
		fmt.Print("SKS : ")
		fmt.Scan(&data[idx].Matkul[j].SKS)
		fmt.Print("UTS : ")
		fmt.Scan(&data[idx].Matkul[j].UTS)
		fmt.Print("UAS : ")
		fmt.Scan(&data[idx].Matkul[j].UAS)
		fmt.Print("Quiz : ")
		fmt.Scan(&data[idx].Matkul[j].Quiz)

		data[idx].Matkul[j].Total = hitungTotal(
			data[idx].Matkul[j].UTS,
			data[idx].Matkul[j].UAS,
			data[idx].Matkul[j].Quiz)

		data[idx].Matkul[j].Grade = hitungGrade(data[idx].Matkul[j].Total)
		data[idx].JumlahMatkul++
		fmt.Println("Mata kuliah berhasil ditambahkan")

	} else {
		fmt.Println("Mahasiswa tidak ditemukan")
	}
}

func cariMahasiswa(data tabMahasiswa, n int) {
	var nama string
	var idx int

	fmt.Print("Nama : ")
	fmt.Scan(&nama)

	idx = sequentialSearch(data, n, nama)

	if idx != -1 {
		fmt.Println("\n===== DATA DITEMUKAN =====")
		fmt.Println("NIM  :", data[idx].NIM)
		fmt.Println("Nama :", data[idx].Nama)
		fmt.Println("=========================")
	} else {
		fmt.Println("Data mahasiswa tidak ditemukan")
	}
}

func editMahasiswa(data *tabMahasiswa, n int) {
	var nim, namaBaru string
	var idx int

	fmt.Print("NIM : ")
	fmt.Scan(&nim)

	idx = binarySearchNIM(*data, n, nim)

	if idx != -1 {
		fmt.Println("Data ditemukan")
		fmt.Println("Nama lama :", (*data)[idx].Nama)

		fmt.Print("Nama Baru : ")
		fmt.Scan(&namaBaru)

		(*data)[idx].Nama = namaBaru

		fmt.Println("Data berhasil diubah")
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func hapusMahasiswa(data *tabMahasiswa, n *int) {
	var nim string
	var idx, i int

	insertionSortNIM(data, *n)

	fmt.Print("NIM : ")
	fmt.Scan(&nim)

	idx = binarySearchNIM(*data, *n, nim)

	if idx != -1 {
		for i = idx; i < *n-1; i++ {
			data[i] = data[i+1]
		}
		*n = *n - 1
	}
}

func tampilMahasiswa(data tabMahasiswa, n int) {
	var i int
	fmt.Println("==========================================")
	fmt.Printf("| %-3s | %-15s | %-20s |\n",
		"No", "NIM", "Nama")
	fmt.Println("==========================================")

	for i = 0; i < n; i++ {
		fmt.Printf("| %-3d | %-15s | %-20s |\n",
			i+1,
			data[i].NIM,
			data[i].Nama)
	}
	fmt.Println("==========================================")
}

func tampilTranskrip(data tabMahasiswa, n int) {
	var nim string
	var idx, i int

	fmt.Print("NIM : ")
	fmt.Scan(&nim)

	insertionSortNIM(&data, n)
	idx = binarySearchNIM(data, n, nim)

	if idx != -1 {

		fmt.Println("\n====================================")
		fmt.Println("NIM  :", data[idx].NIM)
		fmt.Println("Nama :", data[idx].Nama)
		fmt.Println("======================================")

		fmt.Println("================================================================")
		fmt.Printf("| %-15s | %-3s | %-7s | %-5s |\n",
			"Mata Kuliah", "SKS", "Total", "Grade")
		fmt.Println("================================================================")

		for i = 0; i < data[idx].JumlahMatkul; i++ {
			fmt.Printf("| %-15s | %-3d | %-7.2f | %-5s |\n",
				data[idx].Matkul[i].Nama,
				data[idx].Matkul[i].SKS,
				data[idx].Matkul[i].Total,
				data[idx].Matkul[i].Grade)
		}

		fmt.Println("================================================================")
		fmt.Printf("Total SKS : %d\n", hitungSKS(data[idx]))
		fmt.Printf("IPK       : %.2f\n", hitungIPK(data[idx]))

	} else {
		fmt.Println("Mahasiswa tidak ditemukan")
	}
}

func tampilUrut(data tabMahasiswa, n int) {
	var i int

	fmt.Println("===================================================================")
	fmt.Printf("| %-15s | %-15s | %-5s | %-5s |\n",
		"NIM", "Nama", "SKS", "IPK")
	fmt.Println("===================================================================")

	for i = 0; i < n; i++ {
		fmt.Printf("| %-15s | %-15s | %-5d | %-5.2f |\n",
			data[i].NIM,
			data[i].Nama,
			hitungSKS(data[i]),
			hitungIPK(data[i]))
	}

	fmt.Println("===================================================================")
}

func main() {
	var data tabMahasiswa
	var n, pilih int
	var jalan bool

	data[0].NIM = "103032500111"
	data[0].Nama = "Andi"

	data[0].Matkul[0] = MataKuliah{"Matdis", 3, 90, 88, 92, 90, "A"}
	data[0].Matkul[1] = MataKuliah{"Aljabar", 3, 85, 84, 86, 85, "A"}
	data[0].Matkul[2] = MataKuliah{"Alpro", 3, 88, 87, 89, 88, "A"}
	data[0].Matkul[3] = MataKuliah{"Agama", 3, 92, 90, 94, 92, "A"}
	data[0].JumlahMatkul = 4

	data[1].NIM = "103032500122"
	data[1].Nama = "Budi"

	data[1].Matkul[0] = MataKuliah{"Matdis", 2, 80, 82, 81, 81, "AB"}
	data[1].Matkul[1] = MataKuliah{"Aljabar", 3, 78, 80, 79, 79, "B"}
	data[1].Matkul[2] = MataKuliah{"Alpro", 3, 84, 85, 83, 84, "AB"}
	data[1].Matkul[3] = MataKuliah{"Agama", 2, 88, 87, 89, 88, "A"}
	data[1].JumlahMatkul = 4

	data[2].NIM = "103032500133"
	data[2].Nama = "Citra"

	data[2].Matkul[0] = MataKuliah{"Matdis", 4, 75, 76, 77, 76, "B"}
	data[2].Matkul[1] = MataKuliah{"Aljabar", 4, 70, 72, 71, 71, "B"}
	data[2].Matkul[2] = MataKuliah{"Alpro", 3, 82, 81, 83, 82, "AB"}
	data[2].Matkul[3] = MataKuliah{"Agama", 3, 90, 89, 91, 90, "A"}
	data[2].JumlahMatkul = 4

	data[3].NIM = "103032500144"
	data[3].Nama = "Deni"

	data[3].Matkul[0] = MataKuliah{"Matdis", 2, 68, 70, 69, 69, "BC"}
	data[3].Matkul[1] = MataKuliah{"Aljabar", 3, 72, 73, 71, 72, "B"}
	data[3].Matkul[2] = MataKuliah{"Alpro", 3, 65, 66, 67, 66, "BC"}
	data[3].Matkul[3] = MataKuliah{"Agama", 3, 80, 82, 81, 81, "AB"}
	data[3].JumlahMatkul = 4

	data[4].NIM = "103032500155"
	data[4].Nama = "Eka"

	data[4].Matkul[0] = MataKuliah{"Matdis", 4, 95, 94, 96, 95, "A"}
	data[4].Matkul[1] = MataKuliah{"Aljabar", 4, 92, 91, 93, 92, "A"}
	data[4].Matkul[2] = MataKuliah{"Alpro", 4, 90, 89, 91, 90, "A"}
	data[4].Matkul[3] = MataKuliah{"Agama", 4, 97, 96, 98, 97, "A"}
	data[4].JumlahMatkul = 4

	data[5].NIM = "103032500166"
	data[5].Nama = "Farhan"

	data[5].Matkul[0] = MataKuliah{"Matdis", 2, 60, 62, 61, 61, "C"}
	data[5].Matkul[1] = MataKuliah{"Aljabar", 2, 58, 59, 60, 59, "D"}
	data[5].Matkul[2] = MataKuliah{"Alpro", 2, 70, 68, 69, 69, "BC"}
	data[5].Matkul[3] = MataKuliah{"Agama", 2, 75, 76, 74, 75, "B"}
	data[5].JumlahMatkul = 4

	data[6].NIM = "103032500177"
	data[6].Nama = "Gita"

	data[6].Matkul[0] = MataKuliah{"Matdis", 3, 85, 86, 84, 85, "A"}
	data[6].Matkul[1] = MataKuliah{"Aljabar", 3, 83, 82, 84, 83, "AB"}
	data[6].Matkul[2] = MataKuliah{"Alpro", 4, 78, 80, 79, 79, "B"}
	data[6].Matkul[3] = MataKuliah{"Agama", 3, 88, 87, 89, 88, "A"}
	data[6].JumlahMatkul = 4

	data[7].NIM = "103032500188"
	data[7].Nama = "Hadi"

	data[7].Matkul[0] = MataKuliah{"Matdis", 4, 73, 74, 72, 73, "B"}
	data[7].Matkul[1] = MataKuliah{"Aljabar", 4, 77, 78, 76, 77, "B"}
	data[7].Matkul[2] = MataKuliah{"Alpro", 4, 81, 82, 80, 81, "AB"}
	data[7].Matkul[3] = MataKuliah{"Agama", 3, 86, 87, 85, 86, "A"}
	data[7].JumlahMatkul = 4

	data[8].NIM = "103032500199"
	data[8].Nama = "Intan"

	data[8].Matkul[0] = MataKuliah{"Matdis", 2, 55, 57, 56, 56, "D"}
	data[8].Matkul[1] = MataKuliah{"Aljabar", 2, 62, 61, 63, 62, "C"}
	data[8].Matkul[2] = MataKuliah{"Alpro", 3, 68, 67, 69, 68, "BC"}
	data[8].Matkul[3] = MataKuliah{"Agama", 2, 72, 73, 71, 72, "B"}
	data[8].JumlahMatkul = 4

	n = 9

	jalan = true

	for jalan {
		fmt.Println("\n======================================")
		fmt.Println("     APLIKASI NILAI MAHASISWA")
		fmt.Println("======================================")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Tambah Mata Kuliah")
		fmt.Println("3. Tampilkan Mahasiswa")
		fmt.Println("4. Tampilkan Transkrip")
		fmt.Println("5. Cari Mahasiswa")
		fmt.Println("6. Edit Mahasiswa")
		fmt.Println("7. Hapus Mahasiswa")
		fmt.Println("8. Urutkan Mahasiswa")
		fmt.Println("9. Keluar")
		fmt.Println("======================================")

		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)
		for pilih < 1 || pilih > 9 {
			fmt.Println("Input salah silahkan masukkan ulang!")
			fmt.Print("Pilih : ")
			fmt.Scan(&pilih)
		}

		if pilih == 1 {
			tambahMahasiswa(&data, &n)
		} else if pilih == 2 {
			tambahMatkul(&data, n)
		} else if pilih == 3 {
			tampilMahasiswa(data, n)
		} else if pilih == 4 {
			insertionSortNIM(&data, n)
			tampilTranskrip(data, n)
		} else if pilih == 5 {
			cariMahasiswa(data, n)
		} else if pilih == 6 {
			editMahasiswa(&data, n)
		} else if pilih == 7 {
			hapusMahasiswa(&data, &n)
		} else if pilih == 8 {

			var pilihUrut int

			fmt.Println("\n===== URUTKAN MAHASISWA =====")
			fmt.Println("1. IPK Terendah")
			fmt.Println("2. IPK Tertinggi")
			fmt.Println("3. SKS Tersedikit")
			fmt.Println("4. SKS Terbanyak")
			fmt.Print("Pilih : ")
			fmt.Scan(&pilihUrut)

			if pilihUrut == 1 {
				insertionSortIPKAsc(&data, n)
				tampilUrut(data, n)

			} else if pilihUrut == 2 {
				insertionSortIPKDesc(&data, n)
				tampilUrut(data, n)

			} else if pilihUrut == 3 {
				selectionSortSKSAsc(&data, n)
				tampilUrut(data, n)

			} else if pilihUrut == 4 {
				selectionSortSKSDesc(&data, n)
				tampilUrut(data, n)
			}
		} else if pilih == 9 {
			jalan = false
		}
	}
}
