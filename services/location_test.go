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

// go test -benchmem -run=^$ damingerdai/address/services -bench BenchmarkGetLocationByIP
func BenchmarkGetLocationByIP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ip := "81.2.69.142"
		location := GetLocationByIP(ip)
		b.Log(location)
	}

}
