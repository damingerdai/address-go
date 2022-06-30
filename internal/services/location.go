package service

import (
	"github.com/damingerdai/address-go/internal/dao"
	"github.com/damingerdai/address-go/internal/models"
	"log"
)

func GetLocationByIP(ip string) *models.Location {
	record, err := dao.GetCityByIp(ip)
	if err != nil {
		return nil
	}
	log.Printf("%v", record)
	location := new(models.Location)
	location.Ip = ip
	location.City = record.City.Names
	location.Country = record.Country.Names
	location.Latitude = record.Location.Latitude
	location.Longitude = record.Location.Longitude
	return location
}
