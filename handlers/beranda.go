package handlers

import (
	"html/template"
	"net/http"
	"ukk-lab2/models"
	"ukk-lab2/utils"
)

func BerandaHandler(w http.ResponseWriter, r *http.Request) {
	maxGaleri := 3
	galeriTampil := models.IsiGaleri
	if len(galeriTampil) > maxGaleri {
		galeriTampil = galeriTampil[:maxGaleri]
	}

	var jurusanDgnWarna []models.DataJurusan
	for _, j := range models.IsiJurusan {
		j.Warna = utils.AmabilWarna(j.NamaJ)
		jurusanDgnWarna = append(jurusanDgnWarna, j)
	}

	page := struct {
		Title string
		Jurusan []models.DataJurusan
		Berita []models.DataBerita
		Galeri []models.DataGaleri
		Profil models.DataProfil
		JumlahJ int
		JumlahS int
		JumlahG int
	}{
		Title: "Beranda",
		Jurusan: jurusanDgnWarna,
		Berita: models.IsiBerita,
		Profil: models.IsiProfil,
		Galeri: galeriTampil,
		JumlahJ: len(jurusanDgnWarna),
		JumlahS: len(models.IsiSiswa),
		JumlahG: len(models.IsiGuru),
	}

	tmpl, err := template.ParseFiles("templates/beranda.html","templates/layout.html")
	if err != nil {
		http.Error(w, "Template error: " + err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, page)
	err != nil {
		http.Error(w, "Render error: " + err.Error(), http.StatusInternalServerError)
	}
}