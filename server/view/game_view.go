package view

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"talpidae-backend/model/game"
)

type Log struct {
	H      int            `json:"h"`
	W      int            `json:"w"`
	Value  game.BlockType `json:"value"`
	UserId string         `json:"user_id"`
}

type GameField struct {
	Field                [][]game.BlockType `json:"field"`
	CurrentNumberOfUsers int                `json:"current_number_of_users"`
}

type GameLogs struct {
	Logs []Log `json:"logs"`
}

type GameCell struct {
	H      int            `json:"h"`
	W      int            `json:"w"`
	Value  game.BlockType `json:"value"`
	UserId string         `json:"user_id"`
}

func ToGameField(g game.Game) GameField {
	return GameField{Field: g.Blocks(), CurrentNumberOfUsers: len(g.Users())}
}

func ToGamePositionArray(g game.Game) GameLogs {
	logs := []Log{}
	for _, v := range g.Logs() {
		logs = append(logs, Log{UserId: v.UserId, Value: v.Value, H: v.Height, W: v.Width})
	}
	return GameLogs{Logs: logs}
}

func FromGameCell(r *http.Request) (*GameCell, error) {
	var body GameCell
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
