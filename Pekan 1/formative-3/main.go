package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	var panjang, _ = strconv.Atoi(panjangPersegiPanjang)
	var lebar, _ = strconv.Atoi(lebarPersegiPanjang)

	var luasPersegiPanjang int = panjang * lebar
	var kelilingPersegiPanjang int = 2 * (panjang * lebar)

	var alas, _ = strconv.ParseFloat(alasSegitiga, 32)
	var tinggi, _ = strconv.ParseFloat(tinggiSegitiga, 32)
	var luasSegitiga int = int(float32(0.5 * alas * tinggi))

	fmt.Printf("Luas Persegi Panjang : %d\n", luasPersegiPanjang)
	fmt.Printf("Keliling Persegi Panjang : %d\n", kelilingPersegiPanjang)
	fmt.Printf("Keliling Persegi Panjang : %d\n", luasSegitiga)

	//Soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	switch {
	case nilaiJohn >= 80:
		fmt.Println("Index Nilai John : A")
	case (nilaiJohn >= 70) && (nilaiJohn < 80):
		fmt.Println("Index Nilai John : B")
	case (nilaiJohn >= 60) && (nilaiJohn < 70):
		fmt.Println("Index Nilai John : C")
	case (nilaiJohn >= 50) && (nilaiJohn < 60):
		fmt.Println("Index Nilai John : D")
	case nilaiJohn < 50:
		fmt.Println("Index Nilai John : E")
	}

	switch {
	case nilaiDoe >= 80:
		fmt.Println("Index Nilai Doe : A")
	case (nilaiDoe >= 70) && (nilaiDoe < 80):
		fmt.Println("Index Nilai Doe : B")
	case (nilaiDoe >= 60) && (nilaiDoe < 70):
		fmt.Println("Index Nilai Doe : C")
	case (nilaiDoe >= 50) && (nilaiDoe < 60):
		fmt.Println("Index Nilai Doe : D")
	case nilaiDoe < 50:
		fmt.Println("Index nilaiDoe Doe : E")
	}

	//Soal 3
	var tanggal = 31
	var bulan = 5
	var tahun = 1994

	switch bulan {
	case 1:
		fmt.Println("Tanggal Lahir Saya : ", strconv.Itoa(tanggal), "Januari", strconv.Itoa(tahun))
	case 2:
		fmt.Println("Tanggal Lahir Saya : ", strconv.Itoa(tanggal), "Februari", strconv.Itoa(tahun))
	case 3:
		fmt.Println("Tanggal Lahir Saya : ", strconv.Itoa(tanggal), "Maret", strconv.Itoa(tahun))
	case 4:
		fmt.Println("Tanggal Lahir Saya : ", strconv.Itoa(tanggal), "April", strconv.Itoa(tahun))
	case 5:
		fmt.Println("Tanggal Lahir Saya : ", strconv.Itoa(tanggal), "Mei", strconv.Itoa(tahun))
	case 6:
		fmt.Println("Tanggal Lahir Saya : ", strconv.Itoa(tanggal), "Juni", strconv.Itoa(tahun))
	case 7:
		fmt.Println("Tanggal Lahir Saya : ", strconv.Itoa(tanggal), "Juli", strconv.Itoa(tahun))
	case 8:
		fmt.Println("Tanggal Lahir Saya : ", strconv.Itoa(tanggal), "Agustus", strconv.Itoa(tahun))
	case 9:
		fmt.Println("Tanggal Lahir Saya : ", strconv.Itoa(tanggal), "September", strconv.Itoa(tahun))
	case 10:
		fmt.Println("Tanggal Lahir Saya : ", strconv.Itoa(tanggal), "Oktober", strconv.Itoa(tahun))
	case 11:
		fmt.Println("Tanggal Lahir Saya : ", strconv.Itoa(tanggal), "November", strconv.Itoa(tahun))
	case 12:
		fmt.Println("Tanggal Lahir Saya : ", strconv.Itoa(tanggal), "Desember", strconv.Itoa(tahun))
	}

	//Soal 4
	var tahunLahir = 1994

	if tahunLahir >= 1944 && tahunLahir <= 1964 {
		fmt.Println("Kelahiran tahun ", strconv.Itoa(tahunLahir), "adalah Baby Boomer")
	} else if tahunLahir >= 1965 && tahunLahir <= 1979 {
		fmt.Println("Kelahiran tahun ", strconv.Itoa(tahunLahir), "adalah Generasi X")
	} else if tahunLahir >= 1980 && tahunLahir <= 1994 {
		fmt.Println("Kelahiran tahun ", strconv.Itoa(tahunLahir), "adalah Generasi Y (Milenials)")
	} else if tahunLahir >= 1995 && tahunLahir <= 2015 {
		fmt.Println("Kelahiran tahun ", strconv.Itoa(tahunLahir), "adalah Generasi Z")
	}

}
