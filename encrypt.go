package main

import (
	"crypto/md5"
	"fmt"
	"io"
	// "strconv"
)

func md5Encrypt(s string) string {
	h := md5.New()
	io.WriteString(h, s)

	md5Str := fmt.Sprintf("%x", h.Sum(nil))
	return md5Str
}

func sha1Encrypt(s string) string {
	h := sha1.New()
	io.WriteString(h, s)

	sha1Str := fmt.Sprintf("%x", h.Sum(nil))
	return sha1Str
}

func sha256Encrypt(s string) string {
	h := sha1.New()
	io.WriteString(h, s)

	sha256Str := fmt.Sprintf("%x", h.Sum(nil))
	return sha256
}
func main() {
	pwd := md5Encrypt("admin123")
	fmt.Println(pwd)
}
