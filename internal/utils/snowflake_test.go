package utils

import "testing"

func TestCreateId(t *testing.T) {
	id, err := CreateID()
	if err != nil {
		t.Error(err)
	}
	t.Logf("id: %d", id)
}

func BenchmarkCreateId(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		id, err := CreateID()
		if err != nil {
			b.Error(err)
		}
		b.Logf("id: %d", id)
	}
}
