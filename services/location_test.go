package service

import (
	"testing"
)

// go test -timeout 30s damingerdai/address/services -run TestGetLocationByIP -v
func TestGetLocationByIP(t *testing.T) {
	ip := "81.2.69.142"
	location := GetLocationByIP(ip)
	t.Log(location)
}
