package main

import (
	"fmt"

	"github.com/pgallardor/dondejugar/pkg/db"
	"github.com/pgallardor/dondejugar/pkg/service/event"
	"github.com/pgallardor/dondejugar/pkg/service/game"
	"github.com/pgallardor/dondejugar/pkg/service/store"
	"gorm.io/driver/postgres"
)

func main() {
	dbConn, err := db.NewDbStorage(postgres.Config{
		DSN: "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable options='--client_encoding=UTF8'",
	})

	if err != nil {
		panic(err.Error())
	}

	storeList := []*store.Store{
		{
			Id:      "gomerry",
			Name:    "GoMerry Cards",
			City:    "Concepción",
			Address: "Av. Chacabuco 782",
		},
		{
			Id:      "cala",
			Name:    "El Calabozo del Dragón",
			City:    "Concepción",
			Address: "O'Higgins 950B",
		},
		{
			Id:      "dreephy",
			Name:    "Dreephy",
			City:    "Concepción",
			Address: "Freire 381",
		},
	}

	dbConn.AutoMigrate(&store.Store{})
	result := dbConn.Create(storeList)
	fmt.Println("Added " + string(result.RowsAffected) + " new stores")

	gameList := []*game.Game{
		{Id: "yugioh-01", Name: "Yu-Gi-Oh!"},
		{Id: "pokemon-01", Name: "Pokémon TCG"},
		{Id: "myl-01", Name: "Mitos y Leyendas - Primer Bloque"},
		{Id: "myl-02", Name: "Mitos y Leyendas - Imperio"},
		{Id: "myl-03", Name: "Mitos y Leyendas - Furia Extendido"},
		//{Id: "myl-04", Name: "Mitos y Leyendas - Bloque Furia"},
		{Id: "op-01", Name: "One Piece Card Game"},
		{Id: "digimon-01", Name: "Digimon Card Game"},
		{Id: "db-01", Name: "Dragon Ball Card Game"},
		{Id: "bss-01", Name: "Battle Spirits Saga"},
		{Id: "fab-01", Name: "Flesh and Blood"},
		{Id: "mtg-01", Name: "Magic the Gathering - Standard"},
		{Id: "mtg-02", Name: "Magic the Gathering - Commander"},
		{Id: "mtg-03", Name: "Magic the Gathering - Modern"},
		{Id: "mtg-04", Name: "Magic the Gathering - Pioneer"},
		{Id: "altered-01", Name: "Altered"},
	}

	dbConn.AutoMigrate(&game.Game{})
	result1 := dbConn.Create(gameList)
	fmt.Printf("Added %v new games\n", result1.RowsAffected)

	eventList := []*event.WeeklyEvent{
		{
			WeekDay:     "Miércoles",
			Time:        "17:30",
			EventName:   "Torneo Yu-Gi-Oh!",
			Description: "Torneo local de Yu-Gi-Oh",
			StoreID:     "gomerry",
			GameID:      "yugioh-01",
		},
		{
			WeekDay:     "Jueves",
			Time:        "17:30",
			EventName:   "Torneo Yu-Gi-Oh!",
			Description: "Torneo local de Yu-Gi-Oh",
			StoreID:     "gomerry",
			GameID:      "yugioh-01",
		}, {
			WeekDay:     "Viernes",
			Time:        "17:30",
			EventName:   "Torneo Yu-Gi-Oh!",
			Description: "Torneo local de Yu-Gi-Oh",
			StoreID:     "gomerry",
			GameID:      "yugioh-01",
		}, {
			WeekDay:     "Lunes",
			Time:        "18:30",
			EventName:   "Torneo One Piece TCG",
			Description: "Torneo oficial de One Piece TCG",
			StoreID:     "gomerry",
			GameID:      "op-01",
		}, {
			WeekDay:     "Martes",
			Time:        "18:30",
			EventName:   "Torneo One Piece TCG",
			Description: "Torneo OPTCG (sin app)",
			StoreID:     "gomerry",
			GameID:      "op-01",
		}, {
			WeekDay:     "Martes",
			Time:        "18:30",
			EventName:   "Torneo MyL PBX",
			Description: "Torneo MyL",
			StoreID:     "gomerry",
			GameID:      "myl-01",
		}, {
			WeekDay:     "Lunes",
			Time:        "17:30",
			EventName:   "Yu-Gi-Oh! Avanzando",
			Description: "",
			StoreID:     "cala",
			GameID:      "yugioh-01",
		}, {
			WeekDay:     "Martes",
			Time:        "18:30",
			EventName:   "Torneo MyL FX",
			Description: "Torneo MyL",
			StoreID:     "cala",
			GameID:      "myl-03",
		},
	}

	dbConn.AutoMigrate(&event.WeeklyEvent{})
	result2 := dbConn.Create(&eventList)
	if result.Error != nil {
		fmt.Println(result2.Error)
	}
}
