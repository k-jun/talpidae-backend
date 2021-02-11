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
	mux.HandleFunc("/", handler.GameField(gameStorage)).Methods(http.MethodGet)
	mux.HandleFunc("/game/start", handler.GameStart(gameStorage)).Methods(http.MethodGet)
	mux.HandleFunc("/game/field", handler.GameField(gameStorage)).Methods(http.MethodGet)
	mux.HandleFunc("/game/logs", handler.GameLogs(gameStorage)).Methods(http.MethodGet)
	mux.HandleFunc("/game/fill", handler.GameFill(gameStorage)).Methods(http.MethodPost)
	mux.HandleFunc("/match/join", handler.MatchJoin(gameStorage)).Methods(http.MethodPost)
	mux.HandleFunc("/match/leave", handler.MatchLeave(gameStorage)).Methods(http.MethodPost)
}
