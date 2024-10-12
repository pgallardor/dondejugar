package game

import "gorm.io/gorm"

type GameService struct {
	db *gorm.DB
}

func NewGameService(db *gorm.DB) *GameService {
	return &GameService{
		db: db,
	}
}

func (s *GameService) GetGames() []Game {
	games := []Game{}
	s.db.Find(&games)
	return games
}
