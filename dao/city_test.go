package dao

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("./../.env")
	fmt.Println("load env file")
}

// go test -timeout 30s damingerdai/address/dao -v
func TestGetCity(t *testing.T) {
	for i := 1; i < 100; i++ {
		city, err := GetCity(i)
		if err != nil {
			t.Error(err)
		}
		if city.Id != i {
			t.Errorf("fail to get city whose id is %d", i)
		}
		//fmt.Println(city.String())
	}
}
