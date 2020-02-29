package dao

import (
	"net"

	"github.com/oschwald/geoip2-golang"
)

func GetCityByIp(ipStr string) (*geoip2.City, error) {
	geoReader, err := geoip2.Open("/usr/bin/data/GeoLite2-City.mmdb")
	if err != nil {
		return nil, err
	}
	defer geoReader.Close()

	ip := net.ParseIP(ipStr)
	record, err := geoReader.City(ip)

	if err != nil {
		return nil, err
	}

	return record, nil

}
