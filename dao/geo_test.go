package dao

import (
	"fmt"
	"testing"
)

// go test -timeout 30s damingerdai/address/dao -run TestGetCityByIp -v
func TestGetCityByIp(t *testing.T) {
	ip := "81.2.69.142"
	record, err := GetCityByIp(ip)
	if err != nil {
		t.Error(err)
	}
	t.Log(record.Country)
	fmt.Printf("Portuguese (BR) city name: %v\n", record.City.Names["pt-BR"])
	if len(record.Subdivisions) > 0 {
		fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
	}
	fmt.Printf("Russian country name: %v\n", record.Country.Names["ru"])
	fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
	fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
	// Output:
	// Portuguese (BR) city name: Londres
	// English subdivision name: England
	// Russian country name: Великобритания
	// ISO country code: GB
	// Time zone: Europe/London
	// Coordinates: 51.5142, -0.0931
}

// damingerdai/address/dao -bench BenchmarkGetCityByIp
func BenchmarkGetCityByIp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ip := "81.2.69.142"
		record, err := GetCityByIp(ip)
		if err != nil {
			b.Error(err)
		}
		b.Log(record.Country)
	}
}
