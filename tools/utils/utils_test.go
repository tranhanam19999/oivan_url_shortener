package utils

import (
	"testing"
)

func TestEncodeBase62(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{0, "0"},
		{1, "1"},
		{61, "Z"},
		{62, "10"},
		{12345, "3d7"},
		{136745, "zzz"},
	}

	for _, tt := range tests {
		got := EncodeBase62(tt.input)
		if got != tt.expected {
			t.Errorf("EncodeBase62(%d) = %s; want %s", tt.input, got, tt.expected)
		}
	}
}

func TestDecodeBase62(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"0", 0},
		{"1", 1},
		{"Z", 61},
		{"10", 62},
		{"3d7", 12345},
		{"zzz", 136745},
	}

	for _, tt := range tests {
		got := DecodeBase62(tt.input)
		if got != tt.expected {
			t.Errorf("DecodeBase62(%s) = %d; want %d", tt.input, got, tt.expected)
		}
	}
}

func TestEncodeDecodeBase62(t *testing.T) {
	numbers := []int64{0, 1, 61, 62, 12345, 238327, 987654321}

	for _, n := range numbers {
		encoded := EncodeBase62(n)
		decoded := DecodeBase62(encoded)
		if decoded != n {
			t.Errorf("EncodeDecodeBase62 failed: original %d -> encoded %s -> decoded %d", n, encoded, decoded)
		}
	}
}

func TestIsValidUrl(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"https://google.com", true},
		{"http://example.com/path?query=1", true},
		{"ftp://fileserver.com/file.txt", true},
		{"google.com", false},    // no scheme
		{"htp://bad.com", false}, // invalid scheme
		{"", false},
		{"://invalid.com", false},
	}

	for _, tt := range tests {
		got := IsValidUrl(tt.input)
		if got != tt.expected {
			t.Errorf("IsValidUrl(%q) = %v; want %v", tt.input, got, tt.expected)
		}
	}
}
