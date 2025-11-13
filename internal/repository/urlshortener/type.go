package repository

// Truthfully I do not know how to name this file haha so I just put type.go
// type.go is for defining the input gonna be used for the repository layer

type UpdateMappingInput struct {
	ID           int64
	OriginalURL  string
	ShortenedURL string
}

type FindOneInput struct {
	OriginalURL  string
	ShortenedURL string
}
