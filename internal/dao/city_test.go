package dao

import "testing"

func TestListCities(t *testing.T) {
	cities, err := ListCities()
	if err != nil {
		t.Error(err)
	}
	if len(cities) != 343 {
		t.Error("no enough cities")
	}
	for i, city := range cities {
		t.Log(i)
		t.Log(city)
		if city.Id != i+1 {
			t.Errorf("fail to get city whose id is %d", i)
		}
	}
}

func BenchmarkListCities(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cities, err := ListCities()
		if err != nil {
			b.Error(err)
		}
		if len(cities) != 343 {
			b.Errorf("no enough cities: %d", len(cities))
		}
	}
}

func TestGetCity(t *testing.T) {
	for i := 1; i <= 350; i++ {
		city, err := GetCity(i)
		if err != nil {
			t.Error(err)
		}
		if city.Id != i {
			t.Errorf("fail to get city whose id is %d", i)
		}
	}
}

func BenchmarkGetCity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 1; i < 600; j++ {
			city, err := GetCity(j)
			if err != nil {
				b.Error(err)
			}
			if city.Id != j {
				b.Errorf("fail to get city whose id is %d", j)
			}
		}
	}
}
