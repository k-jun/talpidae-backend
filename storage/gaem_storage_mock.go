package storage

import (
	"talpidae-backend/model/game"
)

var _ GameStorage = &GameStorageMock{}

type GameStorageMock struct {
	GameMock  game.Game
	ErrorMock error
}

func (g *GameStorageMock) Add(_ string, _ game.Game) error {
	return g.ErrorMock
}

func (g *GameStorageMock) Remove(_ string) error {
	return g.ErrorMock
}

func (g *GameStorageMock) Find(_ string) (game.Game, error) {
	return g.GameMock, g.ErrorMock
}
