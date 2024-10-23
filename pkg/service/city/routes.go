package city

import (
	"net/http"

	"github.com/pgallardor/dondejugar/pkg/utils"
	"gorm.io/gorm"
)

type Handler struct {
	cityService *CityService
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		cityService: NewCityService(db),
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /cities", h.GetCities)
}

func (h *Handler) GetCities(w http.ResponseWriter, r *http.Request) {
	cities := h.cityService.GetAllCities()
	utils.WriteJSON(w, 200, cities)
}
