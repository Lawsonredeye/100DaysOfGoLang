package view

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"url-shortener/controller"
	"url-shortener/db"
	"url-shortener/model"
)

func CreateLink(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	body, _ := io.ReadAll(r.Body)
	var link model.ShortLink

	err = json.Unmarshal(body, &link)

	result := db.First(&link, "url = ?", link.Url)
	if result.RowsAffected == 0 {
		s := controller.Shortener(link.Url)
		shortlink := model.ShortLink{
			Url:      link.Url,
			ShortUrl: s,
		}
		db.Create(&shortlink)
		w.Write([]byte(s))
		return
	}
	w.Write([]byte(link.ShortUrl))
}

func RedirectLink(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// check the path url and then check if the shortlink exists in the database
	// if it does then call the redirect function from the database
	// else write to the console in correct url path.

	s := r.URL.Path
	s = strings.TrimPrefix(s, "/")

	link := model.ShortLink{ShortUrl: s}

	db, _ := db.ConnectDB()
	result := db.First(&link, "short_url = ?", s)

	if result.RowsAffected < 1 {
		http.Error(w, "Bad request, check url", 400)
		return
	}
	log.Printf("User successfully redirected to: %v", link.Url)
	http.Redirect(w, r, link.Url, http.StatusFound)
	// w.Write([]byte("Hello User"))
}
