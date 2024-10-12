package game

import (
	"net/http"

	"github.com/pgallardor/dondejugar/pkg/utils"
	"gorm.io/gorm"
)

type Handler struct {
	gameService *GameService
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		gameService: NewGameService(db),
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /games", h.GetGames)
}

func (h *Handler) GetGames(w http.ResponseWriter, r *http.Request) {
	games := h.gameService.GetGames()
	utils.WriteJSON(w, 200, games)
}
