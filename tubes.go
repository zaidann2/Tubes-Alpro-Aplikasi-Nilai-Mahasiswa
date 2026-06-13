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

// Menghitung nilai total
func hitungTotal(uts, uas, quiz float64) float64 {
	return (uts + uas + quiz) / 3
}

// Menentukan grade
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

// ubah grade ke bobot
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

// hitung IPK
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

// Menghitung total SKS
func hitungSKS(m Mahasiswa) int {
	var i, total int
	for i = 0; i < m.JumlahMatkul; i++ {
		total += m.Matkul[i].SKS
	}
	return total
}

// Sequential Search nama mahasiswa
func sequentialSearch(data tabMahasiswa, n int, cari string) int {
	var i int
	for i = 0; i < n; i++ {
		if data[i].Nama == cari {
			return i
		}
	}
	return -1
}

// Insertion Sort NIM Asc
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

// Binary Search NIM
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

// Insertion Sort IPK Asc
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

// Insertion Sort IPK Desc
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

// Selection Sort SKS Asc
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

// Selection Sort SKS Desc
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

// tambah mahasiswa
func tambahMahasiswa(data *tabMahasiswa, n *int) {
	fmt.Print("NIM : ")
	fmt.Scan(&data[*n].NIM)
	fmt.Print("Nama : ")
	fmt.Scan(&data[*n].Nama)
	*n = *n + 1
	insertionSortNIM(data, *n)
	fmt.Println("Mahasiswa berhasil ditambahkan")
}

// tambah mata kuliah
func tambahMatkul(data *tabMahasiswa, n int) {
	var nim string
	var idx, j int

	insertionSortNIM(data, n)
	fmt.Print("NIM : ")
	fmt.Scan(&nim)

	idx = binarySearchNIM(*data, n, nim)

	if idx != -1 {
		j = data