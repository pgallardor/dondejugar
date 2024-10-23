package store

import (
	"github.com/pgallardor/dondejugar/pkg/service/city"
)

type Store struct {
	Id      string `gorm:"primaryKey"`
	Name    string `gorm:"index:name_idx,unique"`
	City    city.City
	CityId  string
	Commune string
	Address string
	Socials string
	LogoRef string
}
