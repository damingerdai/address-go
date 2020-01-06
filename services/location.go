package service

import "damingerdai/address/dao"
import "damingerdai/address/models"

func GetLocationByIP(ip string) *models.Location {
	record, err := dao.GetCityByIp(ip)
	if err != nil {
		return nil
	}
	location := new(models.Location)
	location.City = record.City.Names
	location.Country = record.Country.Names
	location.Latitude = record.Location.Latitude
	location.Longitude = record.Location.Longitude
	return location
}
