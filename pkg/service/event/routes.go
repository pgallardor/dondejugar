package event

import (
	"net/http"
	"strconv"

	"github.com/pgallardor/dondejugar/pkg/utils"
	"gorm.io/gorm"
)

type Handler struct {
	eventService *EventService
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		eventService: NewEventService(db),
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /event/weekly/{id}", h.GetWeeklyEvent)
	mux.HandleFunc("GET /events/weekly", h.GetWeeklyEvents)
	mux.HandleFunc("GET /store/{id}/weekly", h.GetStoreEvents)
}

func (h *Handler) GetWeeklyEvents(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	if !q.Has("ciudad") {
		utils.WriteJSON(w, 400, "Bad Request - No city provided")
		return
	}

	city := q.Get("ciudad")
	game := q.Get("juego")

	if game == "" {
		events := h.eventService.GetWeeklyEventsByCity(city)
		utils.WriteJSON(w, 200, events)
	} else {
		events := h.eventService.GetWeeklyEventsByCityAndGame(city, game)
		utils.WriteJSON(w, 200, events)
	}
}

func (h *Handler) GetWeeklyEvent(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.WriteJSON(w, 400, "Bad Request - invalid ID")
		return
	}
	event := h.eventService.GetWeeklyEventDetails(id)
	utils.WriteJSON(w, 200, event)
}

func (h *Handler) GetStoreEvents(w http.ResponseWriter, r *http.Request) {
	storeId := r.PathValue("id")
	events := h.eventService.GetWeeklyEventsByStore(storeId)
	utils.WriteJSON(w, 200, events)
}
