package utils

import (
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
