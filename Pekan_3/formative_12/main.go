package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type NilaiMahasiswa struct {
	ID          uint   `json:"id"`
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mataKuliah"`
	IndeksNilai string `json:"indeksNilai"`
	Nilai       uint   `json:"nilai"`
}

var listOfMahasiswa = []NilaiMahasiswa{
	{1, "Effendy Anwar", "Alogirma Dasar", "A", 80},
}

func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//encode
		dataMahasiswa, err := json.Marshal(listOfMahasiswa)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataMahasiswa)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func PostMahasiswa(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	//var newMahasiswa = NilaiMahasiswa{ID: int(len(listOfMahasiswa))}
	var newMahasiswa NilaiMahasiswa
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			//Parse from json
			decodeJson := json.NewDecoder(r.Body)
			if err := decodeJson.Decode(&newMahasiswa); err != nil {
				log.Fatal(err)
			}
		} else {
			//Parse from form data
			nama := r.PostFormValue("nama")
			mataKuliah := r.PostFormValue("mataKuliah")
			getNilai := r.PostFormValue("nilai")
			nilai, _ := strconv.Atoi(getNilai)

			var indexNilai string
			if nilai >= 80 {
				indexNilai = "A"
			} else if nilai >= 70 && nilai < 80 {
				indexNilai = "B"
			} else if nilai >= 60 && nilai < 70 {
				indexNilai = "C"
			} else if nilai >= 50 && nilai < 60 {
				indexNilai = "D"
			} else {
				indexNilai = "E"
			}
			newMahasiswa.Nama = nama
			newMahasiswa.MataKuliah = mataKuliah
			newMahasiswa.IndeksNilai = indexNilai
			newMahasiswa.Nilai = uint(nilai)
		}

		listOfMahasiswa = append(listOfMahasiswa, newMahasiswa)
		dataPerson, _ := json.Marshal(newMahasiswa)
		rw.Write(dataPerson)
		return
	}
	http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			rw.Write([]byte("Username dan Password tidak boleh kosong!"))
			return
		}

		if username == "admin" && password == "admin" {
			next.ServeHTTP(rw, r)
			return
		}

		rw.Write([]byte("Username atau Password tidak sesuai"))
	})
}

func main() {
	http.HandleFunc("/mahasiswa", GetMahasiswa)
	http.Handle("/post-mahasiswa", Auth(http.HandlerFunc(PostMahasiswa)))

	fmt.Println("Server is running at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
