package dao

import "testing"

func TestListProvinces(t *testing.T) {
	provinces, err := ListProvinces()
	if err != nil {
		t.Error(err)
	}
	for _, province := range provinces {
		t.Log(province)
	}
}

func BenchmarkListProvinces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ListProvinces()
		if err != nil {
			b.Error(err)
		}
	}
}

func TestGetProvince(t *testing.T) {
	for i := 1; i <= 32; i++ {
		province, err := GetProvince(i)
		if err != nil {
			t.Errorf("err: %v; i: %d", err, i)
		}
		if province != nil {
			t.Log(province)
		}

	}
}
