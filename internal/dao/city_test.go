package dao

import (
	jdbc "database/sql"
	"testing"
)

func TestListCities(t *testing.T) {
	err := db.Ping()
	if err != nil {
		t.Log("db is no ready")
		return
	}
	trx := db.MustBegin()
	cityDao := CityDao{Trx: trx}
	cities, err := cityDao.ListCities()
	if err != nil {
		trx.Rollback()
		t.Error(err)
		return
	}
	if len(*cities) != 343 {
		t.Error("no enough cities")
	}
	for i, city := range *cities {
		t.Log(i)
		t.Log(city)
		if city.Id != i+1 {
			t.Errorf("fail to get city whose id is %d", i)
		}
	}

	trx.Commit()
}

func BenchmarkListCities(b *testing.B) {
	err := db.Ping()
	if err != nil {
		b.Log("db is no ready")
		return
	}
	trx := db.MustBegin()
	cityDao := CityDao{Trx: trx}
	for i := 0; i < b.N; i++ {
		cities, err := cityDao.ListCities()
		if err != nil {
			trx.Rollback()
			b.Error(err)
			return
		}
		if len(*cities) != 343 {
			b.Errorf("no enough cities: %d", len(*cities))
		}
	}
	trx.Commit()
}

func TestGetCity(t *testing.T) {
	err := db.Ping()
	if err != nil {
		t.Log("db is no ready")
		return
	}
	trx := db.MustBegin()
	cityDao := CityDao{Trx: trx}
	for i := 0; i <= 350; i++ {
		city, err := cityDao.GetCity(i)
		if err != nil {
			if err == jdbc.ErrNoRows {
				t.Log(err)
			} else {
				trx.Rollback()
				t.Error(err)
				return
			}
		}
		if city.Id != i {
			t.Errorf("fail to get city whose id is %d", i)
		}
	}
	trx.Commit()
}

func BenchmarkGetCity(b *testing.B) {
	err := db.Ping()
	if err != nil {
		b.Log("db is no ready")
		return
	}
	trx := db.MustBegin()
	for i := 0; i < b.N; i++ {
		for j := 1; i < 600; j++ {
			city, err := GetCity(trx, j)
			if err != nil {
				if err == jdbc.ErrNoRows {
					b.Log(err)
				} else {
					trx.Rollback()
					b.Error(err)
					return
				}
			}
			if city.Id != j {
				b.Errorf("fail to get city whose id is %d", j)
			}
		}
	}

	trx.Commit()
}
