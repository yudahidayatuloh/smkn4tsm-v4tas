package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"ukk-lab2/models"
)

func EskulDetailHandler(w http.ResponseWriter, r *http.Request) {
	idby := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idby)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	var eskul models.DataEskul
	gagal := false
	for _, e := range models.IsiEskul {
		if e.ID == id {
			eskul = e
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
		Eskul models.DataEskul
	}{
		Title: eskul.NamaE,
		Eskul: eskul,
	}

	tmpl, err := template.ParseFiles("templates/eskul_detail.html", "templates/layout.html")
	if err != nil {
		http.Error(w, "Template error: " + err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, page)
	err != nil {
		http.Error(w, "Render error: " + err.Error(), http.StatusInternalServerError)
	}
}