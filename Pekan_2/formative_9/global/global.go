package global

import (
	"fmt"
	"math"
	"strconv"
)

type SegitigaSamaSisi struct {
	alas, tinggi int
}

func (s SegitigaSamaSisi) Luas() int {
	return int(float32(0.5 * float32(s.alas) * float32(s.tinggi)))
}

func (s SegitigaSamaSisi) Keliling() int {
	return s.alas * 3
}

type PersegiPanjang struct {
	panjang, lebar int
}

func (p PersegiPanjang) Luas() int {
	return p.panjang * p.lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.panjang + p.lebar)
}

type Tabung struct {
	jariJari, tinggi float64
}

func (t Tabung) Volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

type Balok struct {
	panjang, lebar, tinggi float64
}

func (b Balok) Volume() float64 {
	return b.panjang * b.lebar * b.tinggi
}

func (b Balok) LuasPermukaan() float64 {
	return 2 * ((b.panjang * b.lebar) + (b.lebar * b.tinggi) + (b.panjang * b.tinggi))
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

type Phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p Phone) DetailPhone() {
	fmt.Println("Name : ", p.name)
	fmt.Println("Brand : ", p.brand)
	fmt.Println("Year : ", p.year)

	var newColor string
	for _, a := range p.colors {
		newColor = a
	}
	fmt.Println("Colors : ", newColor)
}

type ShowPhone interface {
	detailPhone()
}

func LuasPersegi(sisi int, type_sisi bool) (setence string) {
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
	var bangunDatar HitungBangunDatar

	fmt.Println("===== Hitung Bangun Datar")
	fmt.Println("===== Segitiga")
	bangunDatar = SegitigaSamaSisi{6, 10}
	fmt.Println("Luas Segitiga Sama Sisi      :", bangunDatar.Luas())
	fmt.Println("Keliling Segitiga Sama Sisi  :", bangunDatar.Keliling())

	fmt.Println("===== Persegi Panjang")
	bangunDatar = PersegiPanjang{6, 10}
	fmt.Println("Luas Persegi Panjang      :", bangunDatar.Luas())
	fmt.Println("Keliling Persegi Panjang  :", bangunDatar.Keliling())

	fmt.Println()
	var bangunRuang HitungBangunRuang
	fmt.Println("===== Hitung Bangun Ruang")
	fmt.Println("===== Tabung")
	bangunRuang = Tabung{7, 10}
	fmt.Println("Luas Permukaan Tabung      :", bangunRuang.LuasPermukaan())
	fmt.Println("Volume tabung  :", bangunRuang.Volume())

	fmt.Println("===== Balok")
	bangunRuang = Balok{panjang: 6, lebar: 3, tinggi: 3}
	fmt.Println("Luas Permukaan Balok      :", bangunRuang.LuasPermukaan())
	fmt.Println("Volume Balok  :", bangunRuang.Volume())

	fmt.Println()
	//Soal 2
	var newPhone ShowPhone
	var arr_warna = []string{"Mystic Bronze", "Mystic White", "Mystic Black"}
	newPhone = phone{"Samsung Galaxy Note 20", "Samsung Galaxy Note 20", 2020, arr_warna}
	newPhone.detailPhone()

	fmt.Println()
	//Soal 3
	fmt.Println(LuasPersegi(4, true))
	fmt.Println(LuasPersegi(8, false))
	fmt.Println(LuasPersegi(0, true))
	fmt.Println(LuasPersegi(0, false))

}
