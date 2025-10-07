package handlers

import (
	"html/template"
	"net/http"
	"path"
	"strconv"
	"ukk-lab2/models"
)

func BeritaDetailHandler(w http.ResponseWriter, r *http.Request) {
	idby := path.Base(r.URL.Path)
	id, err := strconv.Atoi(idby)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	var berita models.DataBerita
	gagal := false
	for _, b := range models.IsiBerita {
		if  b.ID == id {
			berita = b
			gagal = true
			break
		}
	}

	if !gagal {
		http.NotFound(w, r)
		return
	}

	page := struct {
		Title string
		Berita models.DataBerita
	}{
		Title: berita.Judul,
		Berita: berita,
	}

	tmpl, err := template.ParseFiles("templates/berita_detail.html", "templates/layout.html")
	if err != nil {
		http.Error(w, "Template error: " + err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, page)
	err != nil {
		http.Error(w, "Render error: " + err.Error(), http.StatusInternalServerError)
	}
}