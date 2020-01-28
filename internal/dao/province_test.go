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
