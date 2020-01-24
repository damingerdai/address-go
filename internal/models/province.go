package models

import (
	"strconv"
	"strings"
)

type Province struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	ProvinceId int    `json:"provinceId"`
}

func (province *Province) String() string {
	var b strings.Builder
	b.WriteString("Province[id=")
	b.WriteString(strconv.Itoa(province.Id))
	b.WriteString(",name=")
	b.WriteString(province.Name)
	b.WriteString(",ProvinceId=")
	b.WriteString(strconv.Itoa(province.ProvinceId))
	b.WriteString("]")
	return b.String()
}
