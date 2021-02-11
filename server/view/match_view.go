package view

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
)

type MatchUser struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type MatchGame struct {
	Id uuid.UUID `json:"game_id"`
}

type MatchGameUser struct {
	GameId uuid.UUID `json:"game_id"`
	UserId uuid.UUID `json:"user_id"`
}

func FromMatchUser(r *http.Request) (*MatchUser, error) {
	var body MatchUser
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &body)
	if err != nil {
		return nil, err
	}

	return &body, nil
}

func ToMatchGame(id uuid.UUID) MatchGame {
	return MatchGame{Id: id}
}

func FromMatchGameUser(r *http.Request) (*MatchGameUser, error) {
	var body MatchGameUser
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &body)
	if err != nil {
		return nil, err
	}

	return &body, nil
}
