package models

import (
	"strconv"
	"strings"
)

type City struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	CityId int    `json:"cityId"`
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
