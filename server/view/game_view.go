package view

import "talpidae-backend/model/game"

type Position struct {
	H     int    `json:"h"`
	W     int    `json:"w"`
	Value string `json:"value"`
}

type GameView struct {
	Positions []Position `json:"positions"`
}

func GameStatus(g game.Game) GameView {

	blocks := g.Blocks()
	pos := []Position{}

	for i := 0; i < len(blocks); i++ {
		for j := 0; j < len(blocks[i]); j++ {
			if blocks[i][j] != "" {
				pos = append(pos, Position{H: i, W: j, Value: blocks[i][j]})
			}
		}
	}

	return GameView{Positions: pos}
}
