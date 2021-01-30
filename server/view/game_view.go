package view

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"talpidae-backend/model/game"
)

type Position struct {
	H      int            `json:"h"`
	W      int            `json:"w"`
	Value  game.BlockType `json:"value"`
	UserId string         `json:"user_id"`
}

type GameView struct {
	Field [][]game.BlockType `json:"field"`
}

func ToGameField(g game.Game) GameView {
	return GameView{Field: g.Blocks()}
}

func FromGamePosition(r *http.Request) (*Position, error) {
	var body Position
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
