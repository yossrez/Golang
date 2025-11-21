package math

// Angka armstrong ini adalah angka bilangan
// bulat yang jumlah pangkat tiga digitnya sama
// dengan bilangan itu sendiri
// sebagai contoh
// 3 ** 3 + 7 ** 3 + 1 ** 3 = 371
// maka 371 adalah angka armstrong

import (
	"math"
	"strconv"
)

// fungsi AngkaArmstrong memeriksa apakah suatu bilangan murupakan bilangan armstrong
// sebagai contoh
// 1 ** 3 + 5 ** 3 + 3 ** 3 = 1 + 125 + 27 = 153
// maka 153 adalah angka armstrong
func AngkaArmstrong(angka int) bool {
	var digitTerakhir int
	var jumlah int = 0
	var sementara int = angka

	// menghitung jumlah digit pada bilangan tersebut
	// konversi ke string kemudian hitung panjangnya
	jumlahDigit := float64(len(strconv.Itoa(angka)))

	// ekstrak digit satu per satu dari kanan
	// hingga semua digit telah diproses
	for sementara > 0 {
		// ambil digit paling kanan
		digitTerakhir = sementara % 10

		// tambahkan digit ^ jumlahDigit
		jumlah += int(math.Pow(float64(digitTerakhir), jumlahDigit))

		// kemudian hapus digit paling kanan dengan pembagian integer
		sementara /= 10
	}

	// jika hasil sama maka adalah angka armstrong, jika tidak maka return false
	return angka == jumlah
}
