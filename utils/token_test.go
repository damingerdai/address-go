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
	secretKey := []byte("damingerida")
	expiresAt := time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix()

	if token, err := CreateToken(&user, secretKey, expiresAt); err != nil {
		t.Errorf("fail to create token: %s", err)
	} else {
		t.Logf("token: %s", token)
	}

}
