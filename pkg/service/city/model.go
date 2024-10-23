package city

type City struct {
	Id   string `gorm:"primaryKey"`
	Name string
}
