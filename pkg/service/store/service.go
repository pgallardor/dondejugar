package store

import (
	"gorm.io/gorm"
)

type StoreService struct {
	db *gorm.DB
}

func NewStoreService(db *gorm.DB) *StoreService {
	return &StoreService{
		db: db,
	}
}

func (s *StoreService) GetStores() []Store {
	store := []Store{}
	s.db.Find(&store)
	return store
}
