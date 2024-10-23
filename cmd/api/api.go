package api

import (
	"fmt"
	"net/http"

	"github.com/pgallardor/dondejugar/pkg/service/city"
	"github.com/pgallardor/dondejugar/pkg/service/event"
	"github.com/pgallardor/dondejugar/pkg/service/game"
	"github.com/pgallardor/dondejugar/pkg/service/store"
	"gorm.io/gorm"
)

type APIServer struct {
	addr string
	db   *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (api *APIServer) Run() {
	fmt.Println("Starting DondeJugar API")
	mux := http.NewServeMux()
	cityHandler := city.NewHandler(api.db)
	storeHandler := store.NewHandler(api.db)
	gameHandler := game.NewHandler(api.db)
	eventHandler := event.NewHandler(api.db)

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			fmt.Fprintf(w, "API OK!")
		} else {
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	})

	cityHandler.RegisterRoutes(mux)
	storeHandler.RegisterRoutes(mux)
	gameHandler.RegisterRoutes(mux)
	eventHandler.RegisterRoutes(mux)

	if err := http.ListenAndServe(api.addr, mux); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Listening in Port 8080")
	}
}
