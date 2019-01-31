package util

import (
	"encoding/base64"
	"log"
)

// Base64 return base64 string
func Base64(d string) string {
	en := base64.URLEncoding.EncodeToString([]byte(d))

	dec, err := base64.URLEncoding.DecodeString(en)
	if err != nil {
		log.Println("decode error:", err)
	}
	log.Printf("%s base64(%s, %s)", d, en, dec)
	return en
}
