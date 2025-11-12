package model

type URLShortener struct {
	Base
	OriginalURL string `gorm:"type:text;not null"`
	ShortURL    string `gorm:"type:varchar(255);uniqueIndex;not null"`
}
