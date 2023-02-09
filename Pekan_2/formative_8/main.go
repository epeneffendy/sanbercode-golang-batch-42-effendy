package main

import (
	"fmt"
	"math"
	"strconv"
)

type segitigaSamaSisi struct {
	alas, tinggi int
}

func (s segitigaSamaSisi) luas() int {
	return int(float32(0.5 * float32(s.alas) * float32(s.tinggi)))
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

type persegiPanjang struct {
	panjang, lebar int
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

type tabung struct {
	jariJari, tinggi float64
}

func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

type balok struct {
	panjang, lebar, tinggi float64
}

func (b balok) volume() float64 {
	return b.panjang * b.lebar * b.tinggi
}

func (b balok) luasPermukaan() float64 {
	return 2 * ((b.panjang * b.lebar) + (b.lebar * b.tinggi) + (b.panjang * b.tinggi))
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p phone) detailPhone() {
	fmt.Println("Name : ", p.name)
	fmt.Println("Brand : ", p.brand)
	fmt.Println("Year : ", p.year)

	var newColor string
	for _, a := range p.colors {
		newColor = a
	}
	fmt.Println("Colors : ", newColor)
}

type showPhone interface {
	detailPhone()
}

func luasPersegi(sisi int, type_sisi bool) (setence string) {
	if type_sisi {
		setence = "luas persegi dengan sisi 2 cm adalah 4 cm"
	} else if !type_sisi {
		setence = strconv.Itoa(sisi)
	} else if sisi == 0 && type_sisi {
		setence = "Maaf anda belum menginput sisi dari persegi"
	} else {
		setence = "nil"
	}

	return setence
}

func main() {
	//Soal 1
	var bangunDatar hitungBangunDatar

	fmt.Println("===== Hitung Bangun Datar")
	fmt.Println("===== Segitiga")
	bangunDatar = segitigaSamaSisi{6, 10}
	fmt.Println("Luas Segitiga Sama Sisi      :", bangunDatar.luas())
	fmt.Println("Keliling Segitiga Sama Sisi  :", bangunDatar.keliling())

	fmt.Println("===== Persegi Panjang")
	bangunDatar = persegiPanjang{6, 10}
	fmt.Println("Luas Persegi Panjang      :", bangunDatar.luas())
	fmt.Println("Keliling Persegi Panjang  :", bangunDatar.keliling())

	fmt.Println()
	var bangunRuang hitungBangunRuang
	fmt.Println("===== Hitung Bangun Ruang")
	fmt.Println("===== Tabung")
	bangunRuang = tabung{7, 10}
	fmt.Println("Luas Permukaan Tabung      :", bangunRuang.luasPermukaan())
	fmt.Println("Volume tabung  :", bangunRuang.volume())

	fmt.Println("===== Balok")
	bangunRuang = balok{panjang: 6, lebar: 3, tinggi: 3}
	fmt.Println("Luas Permukaan Balok      :", bangunRuang.luasPermukaan())
	fmt.Println("Volume Balok  :", bangunRuang.volume())

	fmt.Println()
	//Soal 2
	var newPhone showPhone
	var arr_warna = []string{"Mystic Bronze", "Mystic White", "Mystic Black"}
	newPhone = phone{"Samsung Galaxy Note 20", "Samsung Galaxy Note 20", 2020, arr_warna}
	newPhone.detailPhone()

	fmt.Println()
	//Soal 3
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

}
