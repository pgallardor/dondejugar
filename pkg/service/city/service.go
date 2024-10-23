package city

import "gorm.io/gorm"

type CityService struct {
	db *gorm.DB
}

func NewCityService(db *gorm.DB) *CityService {
	return &CityService{
		db: db,
	}
}

func (s *CityService) GetAllCities() []City {
	var cities []City
	s.db.Find(&cities)
	return cities
}
