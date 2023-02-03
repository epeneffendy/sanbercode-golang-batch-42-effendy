package main

import (
	"fmt"
)

func main() {
	//Soal 1
	var nilaiSoal1 = 10
	for i := 1; i <= nilaiSoal1; i++ {
		if i%2 == 0 {
			fmt.Println("Berkualitas")
		} else {
			if i%3 == 0 {
				fmt.Println("I Love Coding ")
			} else {
				fmt.Println("Santai")
			}
		}
	}

	fmt.Println()
	//Soal 2
	var nilaiSoal2 = 10
	for x := 1; x <= nilaiSoal2; x++ {
		for y := 1; y <= x; y++ {
			fmt.Printf("%s", "#")
		}
		fmt.Println()
	}

	fmt.Println()
	//Soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	var newKalimat = kalimat[2:]
	fmt.Println(newKalimat)

	//Soal 4
	var sayuran = []string{}
	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkun", "Kubis", "Seledri", "Tauge", "Timun")

	for ind, z := range sayuran {
		fmt.Println(ind+1, z)
	}

	fmt.Println()
	//Soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	for a, b := range satuan {
		fmt.Println(a, "=", b)
	}

	var panjang int = satuan["panjang"]
	var lebar int = satuan["lebar"]
	var tinggi int = satuan["tinggi"]
	var volume int = panjang * lebar * tinggi
	fmt.Println("Volume Balok :", volume)
}
