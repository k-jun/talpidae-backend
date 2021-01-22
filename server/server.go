package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewServer(port int) http.Server {
	router := mux.NewRouter()
	attachHandlers(router)
	s := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	return s
}

func attachHandlers(mux *mux.Router) {
	mux.HandleFunc("/gamestart", handlers.GameStart()).Methods(http.MethodPost)
}
