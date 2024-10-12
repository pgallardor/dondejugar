package store

import (
	"net/http"

	"github.com/pgallardor/dondejugar/pkg/utils"
	"gorm.io/gorm"
)

type Handler struct {
	storeService *StoreService
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		storeService: NewStoreService(db),
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /stores", h.GetStores)
}

func (h *Handler) GetStores(w http.ResponseWriter, r *http.Request) {
	stores := h.storeService.GetStores()
	utils.WriteJSON(w, 200, stores)
}
