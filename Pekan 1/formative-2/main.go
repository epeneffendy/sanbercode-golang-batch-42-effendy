package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//Soal 1
	var textPertama = "Bootcamp"
	var textKedua = "Digital"
	var textKetiga = "Skill"
	var textKeempat = "Sanbercode"
	var textKelima = "Golang"
	var sentenceSoal1 = textPertama + ` ` + textKedua + ` ` + textKetiga + ` ` + textKeempat + ` ` + textKelima
	fmt.Println(sentenceSoal1)

	//Soal 2
	halo := "Halo Dunia"
	newHelo := strings.ReplaceAll(halo, "Dunia", "Golang")
	fmt.Println(newHelo)

	//Soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	var newKataKedua = strings.Replace(kataKedua, "s", "S", 1)
	var newKataKetiga = strings.Replace(kataKetiga, "r", "R", 1)
	var newKataKeempat = strings.ToUpper(kataKeempat)
	var sentenceSoal3 = kataPertama + ` ` + newKataKedua + ` ` + newKataKetiga + ` ` + newKataKeempat
	fmt.Println(sentenceSoal3)

	//Soal 4
	var angkaPertama = 8
	var angkaKedua = 5
	var angkaKetiga = 6
	var angkaKeempat = 7

	var isEqual = angkaPertama + angkaKedua + angkaKetiga + angkaKeempat
	fmt.Printf("Hasil Penjumlahan : %d\n", isEqual)

	//Soal 5
	kalimat := "halo halo bandung"
	angka := 2021

	kalimat = `"Hi Hi bandung"`
	intAngka := strconv.Itoa(angka)
	fmt.Println(kalimat, "-", intAngka)
}
