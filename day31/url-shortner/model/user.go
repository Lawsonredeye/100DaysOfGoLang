package model

import "gorm.io/gorm"

type ShortLink struct {
	gorm.Model
	Url      string
	ShortUrl string
}
