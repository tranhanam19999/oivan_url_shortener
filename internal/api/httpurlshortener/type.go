package httpurlshortener

// [....]Input is used for bindinbg the HTTP request's payload
// Could be body, query params, path params, etc
// @name EncodeURLInput
type EncodeURLInput struct {
	URL string `json:"url" binding:"required,url"`
}

type DecodeURLInput struct {
	URL string `json:"url" binding:"required,url"`
}
