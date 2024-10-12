package event

import (
	"github.com/pgallardor/dondejugar/pkg/service/game"
	"github.com/pgallardor/dondejugar/pkg/service/store"
)

type WeeklyEvent struct {
	Id          int `gorm:"primaryKey;autoincrement:true"`
	WeekDay     string
	Time        string
	EventName   string
	Description string
	StoreID     string
	Store       store.Store
	GameID      string
	Game        game.Game
}

/*
type UniqueEvent struct {
	Id             int
	EventTimestamp time.Time
	EventName      string
	Description    string
	Store          store.Store
	Game           game.Game
}
*/
