package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/davidcl24/history_service/app/models"
)

type HistoryElementHandler struct {
}

func (h *HistoryElementHandler) ListUserHistoryElements(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(chi.URLParam(r, "user_id"))

	err := json.NewEncoder(w).Encode(models.GetAllUserHistoryElements(userId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h HistoryElementHandler) GetHistoryElement(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	historyElement := models.GetHistoryElementByID(id)
	if historyElement == nil {
		http.Error(w, "Element not found", http.StatusNotFound)
	}
	err := json.NewEncoder(w).Encode(historyElement)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (h HistoryElementHandler) CreateHistoryElement(w http.ResponseWriter, r *http.Request) {
	var historyElement models.HistoryElement
	err := json.NewDecoder(r.Body).Decode(&historyElement)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	models.AddHistoryElement(&historyElement)
	err = json.NewEncoder(w).Encode(historyElement)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (h HistoryElementHandler) UpdateHistoryElement(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var historyElement models.HistoryElement
	err := json.NewDecoder(r.Body).Decode(&historyElement)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedHistoryElement := models.UpdateHistoryElement(id, historyElement)
	if updatedHistoryElement == nil {
		http.Error(w, "Element not found", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(updatedHistoryElement)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (h HistoryElementHandler) DeleteHistoryElement(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	historyElement := models.DeleteHistoryElement(id)
	if historyElement == nil {
		http.Error(w, "Element not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h HistoryElementHandler) ClearUserHistory(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(chi.URLParam(r, "user_id"))
	historyElements := models.ClearUserHistoryElements(userId)
	if historyElements == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
