package repository

// Truthfully I do not know how to name this file haha so I just put type.go
// type.go is for defining the input gonna be used for the repository layer

type CreateInput struct {
	OriginalURL  string
	ShortenedURL string
}

type FindOneInput struct {
	URL string
}
