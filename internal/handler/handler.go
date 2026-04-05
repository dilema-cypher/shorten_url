package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dilema-cypher/shorten_url/internal/models"
	"github.com/dilema-cypher/shorten_url/internal/service"
	"github.com/dilema-cypher/shorten_url/internal/utils"
)

type Handler struct {
	service service.ShortenerService
}

func NewHandler(svc service.ShortenerService) *Handler {
	return &Handler{service: svc}
}

func (h *Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	var req models.URL

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	shortURL, err := h.service.Shorten(r.Context(), req.URL)
	if err != nil {
		w.WriteHeader(utils.StatusCodeByError(err))
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"short_url": shortURL,
	})
}

func (h *Handler) RedirectByUrl(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "id is required"})
		return
	}

	url, err := h.service.Get(r.Context(), id)
	if err != nil {
		w.WriteHeader(utils.StatusCodeByError(err))
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	http.Redirect(w,r,url,http.StatusMovedPermanently)
}