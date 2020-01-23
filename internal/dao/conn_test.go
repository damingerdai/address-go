package dao

import "testing"

func TestGetConnection(t *testing.T) {
	conn := GetConnection()
	conn.Ping()
	defer conn.Close()
	rows, err := conn.Query("SELECT _id FROM city")
	if err != nil {
		t.Error(err)
	}
	defer rows.Close()
}

func BenchmarkGetConnection(b *testing.B) {
	for i := 1; i < b.N; i++ {
		conn := GetConnection()
		conn.Ping()
		defer conn.Close()
	}
}
