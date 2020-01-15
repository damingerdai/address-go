package dao

import "testing"

func TestGetCConnection(t *testing.T) {
	conn := GetConnection()
	conn.Ping()
}
