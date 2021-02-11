package handler

import (
	"encoding/json"
	"net/http"
	"talpidae-backend/model/user"
	"talpidae-backend/server/view"
	"talpidae-backend/storage"
)

func MatchJoin(gs storage.GameStorage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := view.FromMatchUser(r)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		u := user.New(body.Id, body.Name)
		gid, err := gs.RandomMatch(u)

		mv := view.ToMatchGame(gid)
		bytes, err := json.Marshal(mv)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Write(bytes)
	}
}

func MatchLeave(gs storage.GameStorage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := view.FromMatchGameUser(r)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		u := user.New(body.UserId, "")

		g, err := gs.Find(body.UserId)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		err = g.LeaveUser(u)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}
}
