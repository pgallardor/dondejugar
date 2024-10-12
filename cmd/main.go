package main

import (
	"github.com/pgallardor/dondejugar/cmd/api"
	"github.com/pgallardor/dondejugar/config"
	"github.com/pgallardor/dondejugar/pkg/db"
	"gorm.io/driver/postgres"
)

func main() {
	dbConf := config.InitConfig()
	dbConn, err := db.NewDbStorage(postgres.Config{
		DSN: dbConf.GetSSN(),
	})

	if err != nil {
		panic(err.Error())
	}

	api := api.NewAPIServer(":8080", dbConn)
	api.Run()
}
