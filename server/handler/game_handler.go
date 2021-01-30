package handler

import (
	"encoding/json"
	"net/http"
	"talpidae-backend/model/game"
	"talpidae-backend/server/view"
	"talpidae-backend/storage"
)

const (
	DEBUG_KEY = "debug-key"
	HEIGHT    = 150
	WIDTH     = 80
)

func GameStart(gs storage.GameStorage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		newGame, err := game.New(HEIGHT, WIDTH)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		_ = gs.Remove(DEBUG_KEY)
		err = gs.Add(DEBUG_KEY, newGame)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		gv := view.ToGameField(newGame)

		bytes, err := json.Marshal(gv)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Write(bytes)
	}
}

func GameField(gs storage.GameStorage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		g, err := gs.Find(DEBUG_KEY)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		gv := view.ToGameField(g)

		bytes, err := json.Marshal(gv)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Write(bytes)
	}
}

func GameFill(gs storage.GameStorage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := view.FromGameCell(r)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		g, err := gs.Find(DEBUG_KEY)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		err = g.Fill(body.UserId, body.Value, body.H, body.W)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}
}

func GameLogs(gs storage.GameStorage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		g, err := gs.Find(DEBUG_KEY)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		gv := view.ToGamePositionArray(g)

		bytes, err := json.Marshal(gv)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Write(bytes)
	}
}
