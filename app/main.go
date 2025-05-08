package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/davidcl24/history_service/app/handlers"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Mount("api/history", HistoryRouters())
	http.ListenAndServe(":4000", router)
}

func HistoryRouters() chi.Router {
	router := chi.NewRouter()
	historyElementHandler := handlers.HistoryElementHandler{}
	router.Use(middleware.Logger)
	router.Get("/user/{user_id}", historyElementHandler.ListUserHistoryElements)
	router.Get("/{id}", historyElementHandler.GetHistoryElement)
	router.Post("/", historyElementHandler.CreateHistoryElement)
	router.Patch("/{id}", historyElementHandler.UpdateHistoryElement)
	router.Delete("/{id}", historyElementHandler.DeleteHistoryElement)
	router.Delete("/user/{user_id}", historyElementHandler.ClearUserHistory)
	return router
}
