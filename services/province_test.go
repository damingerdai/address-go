package service

import (
	"testing"
)

// go test -bench=. -benchtime="60s"
// go test -benchmem -run=^$ damingerdai/address/services -bench=. -benchtime="60s"
func BenchmarkListProvinces(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ListProvinces()
	}
}
