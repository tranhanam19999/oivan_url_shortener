package urlshortener

import "errors"

var (
	ErrURLNotFound = errors.New("url not found")
	ErrInvalidURL  = errors.New("invalid url")
)
