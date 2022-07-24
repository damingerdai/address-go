package api

import (
	"os"
)

var hmacSampleSecret []byte

func init() {
	secret := os.Getenv("Secret")

	// if len(secret) == 0 {
	// 	panic("env Secret is required")
	// }

	hmacSampleSecret = []byte(secret)
}

func GetHmacSampleSecret() []byte {
	return hmacSampleSecret
}
