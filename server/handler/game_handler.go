package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"talpidae-backend/model/game"
	"talpidae-backend/server/view"
	"talpidae-backend/storage"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var (
	GAME_ID            = "game_id"
	GAME_ID_DEFAULT, _ = uuid.Parse("1cb0490d-5434-37dd-85f6-7d344cfc7f50")
)

func GameStart(gs storage.GameStorage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		newGame, err := game.New(game.Height, game.Width)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		gid := retrieveIdFromPath(r)
		_ = gs.Remove(gid)
		err = gs.Add(gid, newGame)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		gv := view.ToGameField(newGame)

		bytes, err := json.Marshal(gv)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Write(bytes)
	}
}

func GameField(gs storage.GameStorage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		gid := retrieveIdFromPath(r)
		g, err := gs.Find(gid)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		gv := view.ToGameField(g)

		bytes, err := json.Marshal(gv)
		if err != nil {
			log.Println(err)
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
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		gid := retrieveIdFromPath(r)
		g, err := gs.Find(gid)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		err = g.Fill(body.UserId, body.Value, body.H, body.W)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}
}

func GameLogs(gs storage.GameStorage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		gid := retrieveIdFromPath(r)
		g, err := gs.Find(gid)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		gv := view.ToGamePositionArray(g)

		bytes, err := json.Marshal(gv)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Write(bytes)
	}
}

func retrieveIdFromPath(r *http.Request) uuid.UUID {
	vars := mux.Vars(r)
	gid := vars[GAME_ID]

	guuid, err := uuid.Parse(gid)
	if err != nil {
		log.Println(err)
		guuid = GAME_ID_DEFAULT
	}

	return guuid
}
