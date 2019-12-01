package dao

import "testing"

import "github.com/joho/godotenv"

import "fmt"

func init() {
	godotenv.Load("./../.env")
	fmt.Println("load env file")
}

// go test -timeout 30s damingerdai/address/dao -v
func TestListProvinces(t *testing.T) {
	fmt.Println("load env file")
	provinces := ListProvinces()

	for i, v := range provinces {
		fmt.Printf("i:%d \t", i)
		fmt.Printf("province: %s\n", v)
	}
}
