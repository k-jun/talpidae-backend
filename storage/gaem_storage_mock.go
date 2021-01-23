package storage

import (
	"talpidae-backend/model/game"
)

var _ GameStorage = &GameStorageMock{}

type GameStorageMock struct {
	OutGame  game.Game
	OutError error
}

func (g *GameStorageMock) Add(_ string, _ game.Game) error {
	return g.OutError
}

func (g *GameStorageMock) Remove(_ string) error {
	return g.OutError
}

func (g *GameStorageMock) Find(_ string) (game.Game, error) {
	return g.OutGame, g.OutError
}
