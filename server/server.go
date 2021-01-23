package server

import (
	"fmt"
	"net/http"
	"talpidae-backend/server/handler"
	"talpidae-backend/storage"

	"github.com/gorilla/mux"
)

func NewServer(port int) http.Server {
	router := mux.NewRouter()
	gameStorage := storage.NewGameStorage()
	attachHandlers(router, gameStorage)
	s := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	return s
}

func attachHandlers(mux *mux.Router, gameStorage storage.GameStorage) {
	mux.HandleFunc("/start", handler.GameStart(gameStorage)).Methods(http.MethodGet)
	mux.HandleFunc("/status", handler.GameStatus(gameStorage)).Methods(http.MethodGet)
	mux.HandleFunc("/fill", handler.GameFill(gameStorage)).Methods(http.MethodPost)
}
