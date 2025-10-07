package main

import (
	"fmt"
	"log"
	"net/http"
	"ukk-lab2/handlers"
)

func main () {
	http.HandleFunc("/", handlers.BerandaHandler)
	http.HandleFunc("/berita/", handlers.BeritaDetailHandler)
	http.HandleFunc("/profil", handlers.ProfilHandler)
	http.HandleFunc("/ekstrakulikuler", handlers.EskulHandler)
	http.HandleFunc("/galeri", handlers.GaleriHandler)

	fmt.Println("Server Ready Derr: http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}