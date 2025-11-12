package usrepository

// Truthfully I do not know how to name this file haha so I just put type.go
// type.go is for defining the input gonna be used for the repository layer

type SaveURLInput struct {
	OriginalURL string
	ShortURL    string
}

type FindOriginalURLInput struct {
	ShortURL string
}
