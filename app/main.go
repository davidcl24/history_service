package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/davidcl24/history_service/app/config"
	"github.com/davidcl24/history_service/app/handlers"
	"github.com/davidcl24/history_service/app/models"

	_ "github.com/lib/pq"
)

var router *chi.Mux
var db *sql.DB

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	dbConfig := config.NewEnvDBConfig()

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Database)

	db, _ = sql.Open("postgres", connectionString)
	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
}

func main() {
	router.Mount("/api/history", historyRouters())
	http.ListenAndServe(":7500", router)
}

func historyRouters() chi.Router {
	router := chi.NewRouter()
	dbWrapper := models.NewDB(db)
	historyElementHandler := handlers.HistoryElementHandler{DB: dbWrapper}

	router.Use(middleware.Logger)

	router.Get("/user/{user_id}", historyElementHandler.ListUserHistoryElements)
	router.Get("/user/{user_id}/movie/{movie_id}", historyElementHandler.GetUserMovieHistoryElement)
	router.Get("/user/{user_id}/episode/{episode_id}", historyElementHandler.GetUserEpisodeHistoryElement)
	router.Get("/{id}", historyElementHandler.GetHistoryElement)
	router.Post("/", historyElementHandler.CreateHistoryElement)
	router.Patch("/{id}", historyElementHandler.UpdateHistoryElement)
	router.Delete("/{id}", historyElementHandler.DeleteHistoryElement)
	router.Delete("/user/{user_id}", historyElementHandler.ClearUserHistory)

	return router
}
