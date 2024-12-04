package view

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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
