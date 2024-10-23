package main

import (
	"fmt"

	"github.com/pgallardor/dondejugar/pkg/db"
	"github.com/pgallardor/dondejugar/pkg/service/city"
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

	cityList := []*city.City{
		{
			Id:   "ccp",
			Name: "Concepción",
		},
		{
			Id:   "scl",
			Name: "Santiago",
		},
		{
			Id:   "scfk",
			Name: "Chillán",
		},
	}

	dbConn.AutoMigrate(&city.City{})
	dbConn.Create(cityList)

	storeList := []*store.Store{
		{
			Id:      "gomerry",
			Name:    "GoMerry Cards",
			CityId:  "ccp",
			Commune: "Concepción",
			Address: "Av. Chacabuco 782",
			Socials: `{"instagram": "gomerrycards", "web":"https://gomerrycards.cl"}`,
		},
		{
			Id:      "cala",
			Name:    "El Calabozo del Dragón",
			CityId:  "ccp",
			Commune: "Concepción",
			Address: "O'Higgins 950B",
			Socials: `{"instagram": "calabozo_tienda", "web": "https://www.calabozotienda.cl"}`,
		},
		{
			Id:      "dreephy",
			Name:    "Dreephy",
			Commune: "Concepción",
			CityId:  "ccp",
			Address: "Freire 381",
			Socials: `{"instagram": "dreephy.cl"}`,
		},
		{
			Id:      "gom-01",
			Name:    "Game of Magic",
			Commune: "Concepción",
			CityId:  "ccp",
			Address: "O'Higgins 1397",
			Socials: `{"instagram": "gameofmagictienda", "web":"https://www.gameofmagictienda.cl"}`,
		},
		{
			Id:      "frigos",
			Name:    "Frigos Store",
			Commune: "Concepción",
			CityId:  "ccp",
			Address: "Lincoyan 598, 3er Piso",
			Socials: `{"instagram": "frigos_store"}`,
		},
	}

	dbConn.AutoMigrate(&store.Store{})
	result := dbConn.Create(storeList)
	fmt.Println("Added " + string(result.RowsAffected) + " new stores")

	gameList := []*game.Game{
		{Id: "yugioh-01", Name: "Yu-Gi-Oh!", Abbr: "YGO"},
		{Id: "pokemon-01", Name: "Pokémon TCG", Abbr: "PTCG"},
		{Id: "myl-01", Name: "Mitos y Leyendas - Primer Bloque", Abbr: "MyL PBX"},
		{Id: "myl-02", Name: "Mitos y Leyendas - Imperio", Abbr: "MyL Imperio"},
		{Id: "myl-03", Name: "Mitos y Leyendas - Furia Extendido", Abbr: "MyL FX"},
		//{Id: "myl-04", Name: "Mitos y Leyendas - Bloque Furia"},
		{Id: "op-01", Name: "One Piece Card Game", Abbr: "OPCG"},
		{Id: "digimon-01", Name: "Digimon Card Game", Abbr: "Digimon"},
		{Id: "db-01", Name: "Dragon Ball Card Game", Abbr: "DBCG"},
		{Id: "bss-01", Name: "Battle Spirits Saga", Abbr: "BSS"},
		{Id: "fab-01", Name: "Flesh and Blood", Abbr: "FaB"},
		{Id: "mtg-01", Name: "Magic the Gathering - Standard", Abbr: "MtG Standard"},
		{Id: "mtg-02", Name: "Magic the Gathering - Commander", Abbr: "MtG Commander"},
		{Id: "mtg-03", Name: "Magic the Gathering - Modern", Abbr: "MtG Modern"},
		{Id: "mtg-04", Name: "Magic the Gathering - Pioneer", Abbr: "MtG Pioneer"},
		{Id: "altered-01", Name: "Altered", Abbr: "Altered"},
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
			WeekDay:     "Sábado",
			Time:        "17:00",
			EventName:   "Torneo Yu-Gi-Oh!",
			Description: "Torneo local de Yu-Gi-Oh",
			StoreID:     "gomerry",
			GameID:      "yugioh-01",
		}, {
			WeekDay:     "Domingo",
			Time:        "16:00",
			EventName:   "Torneo Yu-Gi-Oh!",
			Description: "Torneo local de Yu-Gi-Oh",
			StoreID:     "gomerry",
			GameID:      "yugioh-01",
		},
		{
			WeekDay:     "Lunes",
			Time:        "18:30",
			EventName:   "Torneo One Piece TCG",
			Description: "Torneo oficial de One Piece TCG",
			StoreID:     "gomerry",
			GameID:      "op-01",
		}, {
			WeekDay:     "Sábado",
			Time:        "12:00",
			EventName:   "Torneo One Piece TCG",
			Description: "Torneo oficial de One Piece TCG",
			StoreID:     "gomerry",
			GameID:      "op-01",
		},
		{
			WeekDay:     "Martes",
			Time:        "18:30",
			EventName:   "Torneo MyL PBX",
			Description: "Torneo MyL",
			StoreID:     "gomerry",
			GameID:      "myl-01",
		},
		{
			WeekDay:     "Miércoles",
			Time:        "17:30",
			EventName:   "Torneo BSS",
			Description: "",
			StoreID:     "gomerry",
			GameID:      "bss-01",
		},
		{
			WeekDay:     "Miércoles",
			Time:        "18:30",
			EventName:   "Torneo Pokémon TCG",
			Description: "",
			StoreID:     "gomerry",
			GameID:      "pokemon-01",
		}, {
			WeekDay:     "Jueves",
			Time:        "17:30",
			EventName:   "Torneo Digimon",
			Description: "DTCG",
			StoreID:     "gomerry",
			GameID:      "digimon-01",
		}, {
			WeekDay:     "Sábado",
			Time:        "16:30",
			EventName:   "Torneo Digimon",
			Description: "DTCG",
			StoreID:     "gomerry",
			GameID:      "digimon-01",
		},
		{
			WeekDay:     "Lunes",
			Time:        "15:30",
			EventName:   "Torneo Casual Altered",
			Description: "",
			StoreID:     "cala",
			GameID:      "altered-01",
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
		}, {
			WeekDay:     "Martes",
			Time:        "18:15",
			EventName:   "MTG Showdown",
			Description: "Torneo Magic",
			StoreID:     "cala",
			GameID:      "mtg-01",
		}, {
			WeekDay:     "Miércoles",
			Time:        "15:30",
			EventName:   "Tarde Casual MTG Commander",
			Description: "Toda la tarde",
			StoreID:     "cala",
			GameID:      "mtg-02",
		}, {
			WeekDay:     "Jueves",
			Time:        "18:30",
			EventName:   "Torneo MyL Primer Bloque",
			Description: "Torneo MyL",
			StoreID:     "cala",
			GameID:      "myl-01",
		}, {
			WeekDay:     "Viernes",
			Time:        "18:45",
			EventName:   "MTG Pioneer",
			Description: "",
			StoreID:     "cala",
			GameID:      "mtg-04",
		}, {
			WeekDay:     "Sábado",
			Time:        "16:30",
			EventName:   "MTG Pioneer",
			Description: "Torneo MyL",
			StoreID:     "cala",
			GameID:      "mtg-04",
		},
	}

	dbConn.AutoMigrate(&event.WeeklyEvent{})
	result2 := dbConn.Create(&eventList)
	if result.Error != nil {
		fmt.Println(result2.Error)
	}
}
