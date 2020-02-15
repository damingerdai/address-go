package utils

import (
	"damingerdai/address/internal/models"
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	user := models.User{
		Username: "daming",
		Password: "123456",
	}
	secretKey := []byte("damingeridai")
	expiresAt := time.Now().Add(time.Second * 5).Unix()

	if token, err := CreateToken(&user, secretKey, expiresAt); err != nil {
		t.Errorf("fail to create token: %s", err)
	} else {
		time.Sleep(time.Second * 10)
		if b := VerifyToken(token, secretKey); !b {
			t.Logf("token: %s", token)
		}
	}

}

func BenchmarkToken(b *testing.B) {
	secretKey := []byte("damingeridai")
	expiresAt := time.Now().Add(time.Second * 5).Unix()
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
