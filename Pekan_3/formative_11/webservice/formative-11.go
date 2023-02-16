package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	var jarijari float64 = 7
	var tinggi float64 = 10

	hitungVolume := volume(jarijari, tinggi)
	hitungLuasAlas := luasAlas(jarijari)
	hitungKelilingAlas := kelilingAlas(jarijari)

	_, err4 := fmt.Fprintf(w, "jari-jari : %.f, tinngi : %.f, luas alas : %.f, keliling alas : %.f", jarijari, tinggi, hitungVolume, hitungLuasAlas, hitungKelilingAlas)
	if err4 != nil {
		fmt.Println(err4.Error())
		return
	}
}

func luasAlas(jari float64) float64 {
	return math.Pi * math.Pow(jari, 2)
}

func volume(jari, tinggi float64) float64 {
	return luasAlas(jari) * tinggi
}

func kelilingAlas(jari float64) float64 {
	return 2 * math.Pi * jari
}

func main() {
	fmt.Println("=================Soal 4====================")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome")
	})
	http.HandleFunc("/index", index)
	fmt.Println("Server is running at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
