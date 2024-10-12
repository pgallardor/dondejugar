package game

type Game struct {
	Id       string `gorm:"primaryKey"`
	Name     string `gorm:"index:name_idx,unique"`
	ImageRef string
}
