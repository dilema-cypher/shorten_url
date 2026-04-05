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
		utils.SetErrorOnResponse(w, "invalid request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}


	shortURL, err := h.service.Shorten(r.Context(), req.URL)
	if err != nil {
		utils.SetErrorOnResponse(w, err.Error())
		w.WriteHeader(utils.StatusCodeByError(err))
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"short_url": shortURL,
	})
}

func (h *Handler) RedirectByUrl(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		utils.SetErrorOnResponse(w, "id is required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := h.service.Get(r.Context(), id)
	if err != nil {
		utils.SetErrorOnResponse(w, err.Error())
		w.WriteHeader(utils.StatusCodeByError(err))
		return
	}

	http.Redirect(w,r,url,http.StatusMovedPermanently)
}