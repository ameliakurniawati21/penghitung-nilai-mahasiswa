package main

import "fmt"

// bacaAngka membaca input sebagai teks, lalu mengubahnya menjadi
// bilangan bulat secara manual. Mengembalikan -1 jika input bukan angka.
func bacaAngka(pesan string) int {
	var teks string
	fmt.Print(pesan)
	fmt.Scan(&teks)

	if len(teks) == 0 || len(teks) > 9 {
		return -1
	}
	hasil := 0
	for i := 0; i < len(teks); i++ {
		if teks[i] < '0' || teks[i] > '9' {
			return -1
		}
		hasil = hasil*10 + int(teks[i]-'0')
	}
	return hasil
}

// hurufMutu mengubah nilai angka menjadi huruf mutu.
func hurufMutu(nilai float64) string {
	if nilai >= 85 {
		return "A"
	} else if nilai >= 70 {
		return "B"
	} else if nilai >= 55 {
		return "C"
	} else if nilai >= 40 {
		return "D"
	}
	return "E"
}

func main() {
	var nama string
	var nilai [10]int
	lagi := "y"

	for lagi == "y" {
		fmt.Println("\n===== PENGHITUNG NILAI AKHIR MAHASISWA =====")
		fmt.Print("Nama mahasiswa (tanpa spasi): ")
		fmt.Scan(&nama)

		// Validasi jumlah mata kuliah (1-10)
		jumlah := bacaAngka("Jumlah mata kuliah (1-10): ")
		for jumlah < 1 || jumlah > 10 {
			fmt.Println("Error: jumlah harus bilangan bulat 1 sampai 10.")
			jumlah = bacaAngka("Jumlah mata kuliah (1-10): ")
		}

		// Input nilai tiap mata kuliah (0-100)
		total := 0
		tertinggi := 0
		terendah := 100
		for i := 0; i < jumlah; i++ {
			n := bacaAngka(fmt.Sprint("Nilai mata kuliah ke-", i+1, " (0-100): "))
			for n < 0 || n > 100 {
				fmt.Println("Error: nilai harus bilangan bulat 0 sampai 100.")
				n = bacaAngka(fmt.Sprint("Nilai mata kuliah ke-", i+1, " (0-100): "))
			}
			nilai[i] = n
			total = total + n
			if n > tertinggi {
				tertinggi = n
			}
			if n < terendah {
				terendah = n
			}
		}

		// Hitung rata-rata, huruf mutu, dan status kelulusan
		rata := float64(total) / float64(jumlah)
		mutu := hurufMutu(rata)
		lulus := rata >= 55

		// Cetak hasil
		fmt.Println("\n----- HASIL -----")
		fmt.Println("Nama            :", nama)
		for i := 0; i < jumlah; i++ {
			fmt.Printf("Mata kuliah %2d  : %3d (%s)\n", i+1, nilai[i], hurufMutu(float64(nilai[i])))
		}
		fmt.Println("Nilai tertinggi :", tertinggi)
		fmt.Println("Nilai terendah  :", terendah)
		fmt.Printf("Rata-rata       : %.2f\n", rata)
		fmt.Println("Huruf mutu      :", mutu)
		if lulus {
			fmt.Println("Status          : LULUS")
		} else {
			fmt.Println("Status          : TIDAK LULUS")
		}

		// Tanya apakah ingin menghitung lagi (validasi y/n)
		fmt.Print("\nHitung mahasiswa lain? (y/n): ")
		fmt.Scan(&lagi)
		for lagi != "y" && lagi != "n" {
			fmt.Println("Error: ketik y atau n.")
			fmt.Print("Hitung mahasiswa lain? (y/n): ")
			fmt.Scan(&lagi)
		}
	}
	fmt.Println("Selesai. Terima kasih!")
}
