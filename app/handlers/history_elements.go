// It handles HTTP requests and makes a call to the desired CRUD operation
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/davidcl24/history_service/app/models"
)

type HistoryElementHandler struct {
	DB *models.DB
}

// Handles the specific GET HTTP request and returns a list of all history elements from the desired user in a JSON list.
func (h *HistoryElementHandler) ListUserHistoryElements(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(chi.URLParam(r, "user_id"))

	elements, err := h.DB.GetAllUserHistoryElements(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(elements)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// Handles the specific GET HTTP request and returns a history element with a movie from the desired user in a JSON list.
func (h *HistoryElementHandler) GetUserMovieHistoryElement(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(chi.URLParam(r, "user_id"))
	movieId, _ := strconv.Atoi(chi.URLParam(r, "movie_id"))

	historyElement, err := h.DB.GetMovieHistoryElementFromUser(userId, movieId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if historyElement == nil {
		http.Error(w, "Element not found", http.StatusNotFound)
	} else {
		err = json.NewEncoder(w).Encode(historyElement)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

// Handles the specific GET HTTP request and returns a history element with an episode from the desired user in a JSON list.
func (h *HistoryElementHandler) GetUserEpisodeHistoryElement(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(chi.URLParam(r, "user_id"))
	episodeId, _ := strconv.Atoi(chi.URLParam(r, "episode_id"))

	historyElement, err := h.DB.GetEpisodeHistoryElementFromUser(userId, episodeId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if historyElement == nil {
		http.Error(w, "Element not found", http.StatusNotFound)
	} else {
		err = json.NewEncoder(w).Encode(historyElement)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

// Handles the specific GET HTTP request and returns a specific history element as a JSON.
func (h *HistoryElementHandler) GetHistoryElement(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	historyElement, err := h.DB.GetHistoryElementByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if historyElement == nil {
		http.Error(w, "Element not found", http.StatusNotFound)
	} else {
		err = json.NewEncoder(w).Encode(historyElement)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

// Handles the specific POST HTTP request to create a new history element and returns it as a JSON if successful.
func (h *HistoryElementHandler) CreateHistoryElement(w http.ResponseWriter, r *http.Request) {
	var historyElement models.HistoryElement

	err := json.NewDecoder(r.Body).Decode(&historyElement)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	createdElement, err := h.DB.AddHistoryElement(&historyElement)
	if err != nil {
		http.Error(w, "Failed to insert element: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(createdElement)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// Handles the specific PATCH HTTP request to update an existing history element and returns it as a JSON if successful.
func (h *HistoryElementHandler) UpdateHistoryElement(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var historyElement models.HistoryElement
	err := json.NewDecoder(r.Body).Decode(&historyElement)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedHistoryElement, err := h.DB.UpdateHistoryElement(id, historyElement)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
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

// Handles the specific DELETE HTTP request to remove a specific history element and returns the desired HTTP status code
// if successful.
func (h *HistoryElementHandler) DeleteHistoryElement(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	historyElement, err := h.DB.DeleteHistoryElement(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if historyElement == nil {
		http.Error(w, "Element not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Handles the specific DELETE HTTP request to remove all histroy from a use and returns the desired HTTP status code
// if successful.
func (h *HistoryElementHandler) ClearUserHistory(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(chi.URLParam(r, "user_id"))
	historyElements, err := h.DB.ClearUserHistoryElements(userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if historyElements == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
