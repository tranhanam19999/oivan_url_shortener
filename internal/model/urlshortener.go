package model

type URLShortener struct {
	Base
	OriginalURL string `gorm:"type:text;"`
	ShortURL    string `gorm:"type:varchar(255);uniqueIndex;"`
}
