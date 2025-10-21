package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"ukk-lab2/models"
)

func JurusanDetailHandler(w http.ResponseWriter, r *http.Request) {	
	idby := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idby)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	var jurus models.DataJurusan
	gagal := false
	for _, j := range models.IsiJurusan {
		if j.ID == id {
			jurus = j
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
		Jurusan models.DataJurusan
	}{
		Title: jurus.NamaJ,
		Jurusan: jurus,
	}

	tmpl, err := template.ParseFiles("templates/jurusan_detail.html", "templates/layout.html")
	if err != nil {
		http.Error(w, "Template error: " + err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, page)
	err != nil {
		http.Error(w, "Render error: " + err.Error(), http.StatusInternalServerError)
	}
}