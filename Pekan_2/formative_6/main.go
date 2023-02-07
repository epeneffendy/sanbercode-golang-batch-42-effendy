package main

import "fmt"

func main() {
	//Soal 1
	var luasLingkaran float64
	var kelilingLingakaran float64
	var jari_jari float64 = 7
	const pi float64 = 3.14

	luasLingkaran = luas(&jari_jari, pi)
	kelilingLingakaran = keliling(&jari_jari, pi)

	fmt.Println("Luas Lingkaran : ", luasLingkaran)
	fmt.Println("Keliling Lingkaran : ", kelilingLingakaran)

	//Soal 2
	var sentence string
	John := introduce(&sentence, "John", "laki-laki", "penulis", "30")
	sarah := introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(John)
	fmt.Println(sarah)

	//Soal 3
	var buah = []string{}
	arr_buah := [7]string{"jeruk", "semangka", "mangga", "strawberry", "durian", "manggis", "alpukat"}

	var tambahBuah = func(name_buah *[7]string) {
		for _, item := range *name_buah {
			buah = append(buah, item)
		}
	}
	tambahBuah(&arr_buah)

	fmt.Println()
	for _, item := range buah {
		fmt.Println(item)
	}

	////Soal 4
	var dataFilm = []map[string]string{}
	var tambahDataFilm = func(film, duration, genre, year string) {
		dataFilm = append(dataFilm, map[string]string{
			"film":     film,
			"duration": duration,
			"genre":    genre,
			"year":     year,
		})
	}
	fmt.Println()
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

func luas(jari_jari *float64, pi float64) (luas float64) {
	*jari_jari = 14
	luas = pi * *jari_jari * *jari_jari
	return luas
}

func keliling(jari_jari *float64, pi float64) (keliling float64) {
	*jari_jari = 14
	keliling = pi * (2 * *jari_jari)
	return keliling
}

func introduce(kalimat *string, name, gender, jobs, age string) (textIntroduce string) {
	var calling string
	if gender == "Perempuan" {
		calling = "Bu"
	} else {
		calling = "Pak"
	}
	textIntroduce = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", calling, name, jobs, age)
	return textIntroduce
}
