package utils

import (
	"crypto/sha256"
	"encoding/binary"
	"strings"
)

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func StringToBase62(s string) string {
	hash := sha256.Sum256([]byte(s))
	n := binary.BigEndian.Uint64(hash[:8])

	if n == 0 {
		n = uint64(len(s))
	}

	return IntToBase62(n)
}

func IntToBase62(n uint64) string {
	if n == 0 {
		return string(base62Chars[0])
	}

	var result strings.Builder
	for n > 0 {
		result.WriteByte(base62Chars[n%62])
		n /= 62
	}

	return reverseString(result.String())
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}