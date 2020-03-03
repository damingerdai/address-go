package dao

import (
	"fmt"
	"testing"
)

func TestListProvinces(t *testing.T) {
	fmt.Printf("db: %v", db)
	err := db.Ping()
	if err != nil {
		t.Log("db is no ready")
		return
	}
	trx := db.MustBegin()
	provinces, err := ListProvinces(trx)
	if err != nil {
		trx.Rollback()
		t.Error(err)
	}
	for _, province := range provinces {
		t.Log(province)
	}
	trx.Commit()
}

func BenchmarkListProvinces(b *testing.B) {
	err := db.Ping()
	if err != nil {
		b.Log("db is no ready")
		return
	}
	trx := db.MustBegin()
	for i := 0; i < b.N; i++ {
		_, err := ListProvinces(trx)
		if err != nil {
			trx.Rollback()
			b.Error(err)
		}
	}
	trx.Commit()
}

func TestGetProvince(t *testing.T) {
	t.Log(db)
	err := db.Ping()
	if err != nil {
		t.Log("db is no ready")
		return
	}
	trx := db.MustBegin()
	for i := 1; i <= 32; i++ {
		province, err := GetProvince(trx, i)
		if err != nil {
			trx.Rollback()
			t.Errorf("err: %v; i: %d", err, i)
		}
		if province != nil {
			t.Log(province)
		}

	}
	trx.Commit()
}

func BenchmarkGetProvince(b *testing.B) {
	err := db.Ping()
	if err != nil {
		b.Log("db is no ready")
		return
	}
	trx := db.MustBegin()
	for i := 0; i < b.N; i++ {
		for j := 1; j <= 32; j++ {
			province, err := GetProvince(trx, j)
			if err != nil {
				trx.Rollback()
				b.Errorf("err: %v; i: %d", err, i)
			}
			if province != nil {
				b.Log(province)
			}
		}
	}
	trx.Commit()
}
