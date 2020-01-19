package dao

import "testing"

func TestGetConnection(t *testing.T) {
	conn := GetConnection()
	conn.Ping()
	defer conn.Close()
}

func BenchmarkGetConnection(b *testing.B) {
	for i := 1; i < b.N; i++ {
		conn := GetConnection()
		conn.Ping()
		defer conn.Close()
	}
}
