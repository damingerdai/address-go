package config

import "os"

var hmacSampleSecret []byte

func GetHmacSampleSecret() []byte {
	if len(hmacSampleSecret) == 0 {
		fetchHmacSampleSecret()
	}
	return hmacSampleSecret
}

func fetchHmacSampleSecret() {
	secret := os.Getenv("Secret")

	if len(secret) == 0 {
		panic("env Secret is required")
	}

	hmacSampleSecret = []byte(secret)

}
