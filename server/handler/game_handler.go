package handler

import (
	"encoding/json"
	"io/ioutil"
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
		gv := view.GameStatus(newGame)

		bytes, err := json.Marshal(gv)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Write(bytes)
	}
}

func GameStatus(gs storage.GameStorage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		g, err := gs.Find(DEBUG_KEY)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		gv := view.GameStatus(g)

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
		var body view.Position
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(bytes, &body)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		g, err := gs.Find(DEBUG_KEY)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		err = g.Fill(body.H, body.W, body.Value)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}
}
