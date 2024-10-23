package event

import (
	"github.com/pgallardor/dondejugar/pkg/service/game"
	"github.com/pgallardor/dondejugar/pkg/service/store"
	"gorm.io/gorm"
)

type EventService struct {
	db *gorm.DB
}

func NewEventService(db *gorm.DB) *EventService {
	return &EventService{
		db: db,
	}
}

func (s *EventService) GetWeeklyEventsByCity(cityId string) []WeeklyEvent {
	var event []WeeklyEvent
	s.db.InnerJoins("Store", s.db.Statement.Where(&store.Store{CityId: cityId})).Find(&event)
	return event
}

func (s *EventService) GetWeeklyEventsByStore(storeId string) []WeeklyEvent {
	var event []WeeklyEvent
	s.db.InnerJoins("Store", s.db.Statement.Where(&store.Store{Id: storeId})).
		InnerJoins("Game").Find(&event)
	return event
}

func (s *EventService) GetWeeklyEventsByCityAndGame(cityId string, gameId string) []WeeklyEvent {
	var event []WeeklyEvent
	s.db.InnerJoins("Store", s.db.Statement.Where(&store.Store{CityId: cityId})).
		InnerJoins("Game", s.db.Statement.Where(&game.Game{Id: gameId})).Find(&event)
	return event
}

func (s *EventService) GetWeeklyEventDetails(eventId int) WeeklyEvent {
	var event WeeklyEvent
	s.db.Joins("Game").Joins("Store").Find(&event, eventId)
	return event
}

// func (s *EventService) GetUniqueEventsByGame(storeId string, start time.Time, end time.Time) [] {}
