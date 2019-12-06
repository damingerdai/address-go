package token

import (
	"damingerdai/address/models"
	"testing"
	"time"
)

func TestCreateToken(t *testing.T) {
	user := models.User{
		Username: "daming",
		Password: "123456",
	}
	secretKey := []byte("damingeridai")
	expiresAt := time.Date(2020, 10, 10, 12, 0, 0, 0, time.UTC).Unix()

	if token, err := CreateToken(&user, secretKey, expiresAt); err != nil {
		t.Errorf("fail to create token: %s", err)
	} else {
		t.Logf("token: %s", token)
	}

}

func BenchmarkToken(b *testing.B) {
	secretKey := []byte("damingeridai")
	expiresAt := time.Date(2020, 10, 10, 12, 0, 0, 0, time.UTC).Unix()
	for i := 0; i < b.N; i++ {
		user := models.User{
			Username: "daming",
			Password: "123456",
		}
		if token, err := CreateToken(&user, secretKey, expiresAt); err != nil {
			b.Errorf("fail to create token: %s", err)
		} else if ok := VerifyToken(token, secretKey); ok != true {
			b.Error("fail to verify token")
		}
	}
}
