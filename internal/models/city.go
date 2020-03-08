package models

import (
	"strconv"
	"strings"
)

type City struct {
	Id     int    `json:"id" db:"_id"`
	Name   string `json:"name" db:"name"`
	CityId int    `json:"cityId" db:"city_id"`
}

func (city *City) String() string {
	var b strings.Builder
	b.WriteString("City[id=")
	b.WriteString(strconv.Itoa(city.Id))
	b.WriteString(",name=")
	b.WriteString(city.Name)
	b.WriteString(",CityId=")
	b.WriteString(strconv.Itoa(city.CityId))
	b.WriteString("]")
	return b.String()
}
