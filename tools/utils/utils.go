package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func EncodeBase62(num int64) string {
	if num == 0 {
		return string(base62Chars[0])
	}

	result := ""
	for num > 0 {
		remainder := num % 62
		result = string(base62Chars[remainder]) + result
		num = num / 62
	}
	return result
}

func DecodeBase62(str string) int64 {
	var result int64
	for _, c := range str {
		result = result*62 + int64(strings.IndexRune(base62Chars, c))
	}
	return result
}

func IsValidUrl(u string) bool {
	parsed, err := url.ParseRequestURI(u)
	if err != nil {
		return false
	}

	// Only allow http or https
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return false
	}

	if parsed.Host == "" {
		return false
	}

	return true
}

func BuildUrlFromRequest(r *http.Request) string {
	fullURL := r.URL
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	reqUrl := scheme + "://" + r.Host + fullURL.RequestURI()

	return reqUrl
}

func BuildShortenUrlFromConfig(in BuildUrlFromConfigInput) string {
	if in.Stage == "local" {
		return fmt.Sprintf("%s:%s/r", in.Host, in.Port)
	}

	return fmt.Sprintf("%s/r", in.Host)
}
