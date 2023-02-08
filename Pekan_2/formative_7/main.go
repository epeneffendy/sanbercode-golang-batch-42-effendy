package main

import (
	"fmt"
)

type buah struct {
	nama       string
	warna      string
	adaBijinya bool
	harga      int
}

type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (s segitiga) luasSegitiga() {
	var luasSegitiga int = int(float32(0.5 * float32(s.alas) * float32(s.tinggi)))
	fmt.Println("Luas Segitiga : ", luasSegitiga)
}

func (p persegi) luasPersegi() {
	var luasPersegi int = p.sisi * p.sisi
	fmt.Println("Luas Persegi : ", luasPersegi)
}

func (pp persegiPanjang) luasPersegiPanjang() {
	var luasPersegiPanjang int = pp.panjang * pp.lebar
	fmt.Println("Luas Persegi Panjang : ", luasPersegiPanjang)
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p phone) setWarna() {
	fmt.Println("Name : ", p.name)
	fmt.Println("Brand : ", p.brand)
	fmt.Println("Tahun : ", p.year)
	fmt.Println("Warna : ", p.colors)
}

type movie struct {
	film, genre    string
	duration, year int
}

func main() {
	//Soal 1

	var buahNanas = buah{"Nanas", "Kuning", false, 9000}
	var buahJeruk = buah{"Jeruk", "Oranye", true, 8000}
	var buahSemangka = buah{"Semangka", "Hijau & Merah", true, 10000}
	var buahPisang = buah{"Pisang", "Kuning", false, 5000}

	fmt.Println("Nama		Warna		Ada Bijinya		Harga")
	fmt.Println(buahNanas.nama, "		", buahNanas.warna, "	", buahNanas.adaBijinya, "			", buahNanas.harga)
	fmt.Println(buahJeruk.nama, "		", buahJeruk.warna, "	", buahJeruk.adaBijinya, "			", buahJeruk.harga)
	fmt.Println(buahSemangka.nama, "		", buahSemangka.warna, "	", buahSemangka.adaBijinya, "			", buahSemangka.harga)
	fmt.Println(buahPisang.nama, "		", buahPisang.warna, "	", buahPisang.adaBijinya, "			", buahPisang.harga)

	fmt.Println()
	//Soal 2
	var luasSegitiga = segitiga{10, 5}
	luasSegitiga.luasSegitiga()

	var luasPersegi = persegi{5}
	luasPersegi.luasPersegi()

	var luasPersegiPanjang = persegiPanjang{5, 3}
	luasPersegiPanjang.luasPersegiPanjang()

	fmt.Println()
	//Soal 3
	var arr_warna = []string{"Black"}
	var warna = phone{"Redmi 8", "Xiaomi", 2022, arr_warna}
	warna.setWarna()

	fmt.Println()
	//Soal 4

	var dataFilm = []map[string]string{}
	var tambahDataFilm = func(film, duration, genre, year string) {
		dataFilm = append(dataFilm, map[string]string{
			"film":     film,
			"duration": duration,
			"genre":    genre,
			"year":     year,
		})
	}
	
	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	fmt.Println()
	for ind, item := range dataFilm {
		fmt.Println(ind+1, ". title :", item["film"])
		fmt.Println("   ", "duration :", item["duration"])
		fmt.Println("   ", "genre :", item["genre"])
		fmt.Println("   ", "year :", item["year"])

	}
}
