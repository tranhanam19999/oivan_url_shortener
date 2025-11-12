package httpurlshortener

type EncodeURLInput struct {
	URL string `json:"url" binding:"required,url"`
}

type DecodeURLInput struct {
	URL string `json:"url" binding:"required,url"`
}
