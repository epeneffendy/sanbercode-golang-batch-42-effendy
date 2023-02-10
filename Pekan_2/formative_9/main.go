package main

import (
	"fmt"
	"formative_9/global"
)

func main() {
	//Soal 1
	var bangunDatar global.HitungBangunDatar

	fmt.Println("===== Hitung Bangun Datar")
	fmt.Println("===== Segitiga")
	bangunDatar = global.SegitigaSamaSisi{6, 10}
	fmt.Println("Luas Segitiga Sama Sisi      :", bangunDatar.Luas())
	fmt.Println("Keliling Segitiga Sama Sisi  :", bangunDatar.Keliling())

	fmt.Println("===== Persegi Panjang")
	bangunDatar = global.PersegiPanjang{6, 10}
	fmt.Println("Luas Persegi Panjang      :", bangunDatar.Luas())
	fmt.Println("Keliling Persegi Panjang  :", bangunDatar.Keliling())

	fmt.Println()
	var bangunRuang global.HitungBangunRuang
	fmt.Println("===== Hitung Bangun Ruang")
	fmt.Println("===== Tabung")
	bangunRuang = global.Tabung{7, 10}
	fmt.Println("Luas Permukaan Tabung      :", bangunRuang.LuasPermukaan())
	fmt.Println("Volume tabung  :", bangunRuang.Volume())

	fmt.Println("===== Balok")
	bangunRuang = global.Balok{panjang: 6, lebar: 3, tinggi: 3}
	fmt.Println("Luas Permukaan Balok      :", bangunRuang.LuasPermukaan())
	fmt.Println("Volume Balok  :", bangunRuang.Volume())

	fmt.Println()
	//Soal 2
	var newPhone global.ShowPhone
	var arr_warna = []string{"Mystic Bronze", "Mystic White", "Mystic Black"}
	newPhone = global.Phone{"Samsung Galaxy Note 20", "Samsung Galaxy Note 20", 2020, arr_warna}
	newPhone.global.detailPhone()

	fmt.Println()
	//Soal 3
	fmt.Println(global.LuasPersegi(4, true))
	fmt.Println(global.LuasPersegi(8, false))
	fmt.Println(global.LuasPersegi(0, true))
	fmt.Println(global.LuasPersegi(0, false))

}
