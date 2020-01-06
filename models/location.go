package models

type Location struct {
	City         map[string]string `json:"city"`
	Country      map[string]string `json:"country"`
	Subdivisions map[string]string `json:"Subdivisions"`
	Latitude     float64           `json:"latitude"`
	Longitude    float64           `json:"longitude"`
}

func (location *Location) GetCity(name string) string {
	return location.City[name]
}
