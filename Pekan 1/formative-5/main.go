package main

import "fmt"

func main() {
	//Soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)
	fmt.Printf("%s %d \n", "Luas Pesegi Panjang :", luas)
	fmt.Printf("%s %d \n", "Keliling Pesegi Panjang :", keliling)
	fmt.Printf("%s %d \n", "Volume Pesegi Panjang :", volume)

	fmt.Println()
	//Soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john)

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah)

	fmt.Println()
	//Soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	buahFavorit("John", buah...)

	fmt.Println()
	//Soal 4
	var dataFilm = []map[string]string{}
	var tambahDataFilm = func(film string, duration string, genre string, year string) (string, string, string, string) {
		dataFilm = []map[string]string{
			{
				"film":     film,
				"duration": duration,
				"genre":    genre,
				"year":     year,
			},
		}
		return film, duration, genre, year
	}
	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}

func luasPersegiPanjang(p int, l int) int {
	var luas int = p * l
	return luas
}

func kelilingPersegiPanjang(p int, l int) int {
	var keliling = 2 * (p + l)
	return keliling
}

func volumeBalok(p int, l int, t int) int {
	var volume = p * l * t
	return volume
}

func introduce(name string, gender string, job string, age string) (textIntroduce string) {
	textIntroduce = `Pak ` + name + ` adalah seorang ` + job + ` yang berusia ` + age
	return textIntroduce
}

func buahFavorit(name string, buah ...string) {
	fmt.Printf("%s %s %s", "Halo nama saya ", name, " dan buah favorit saya adalah ")
	for _, buah := range buah {
		fmt.Printf("%s", `"`+buah+`", `)
	}
}
