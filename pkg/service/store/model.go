package store

type Store struct {
	Id      string `gorm:"primaryKey"`
	Name    string `gorm:"index:name_idx,unique"`
	City    string `gorm:"index:city_idx"`
	Address string
	LogoRef string
}
