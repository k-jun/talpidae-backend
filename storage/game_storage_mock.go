package storage

import (
	"talpidae-backend/model/game"
	"talpidae-backend/model/user"

	"github.com/google/uuid"
)

var _ GameStorage = &GameStorageMock{}

type GameStorageMock struct {
	GameMock  game.Game
	ErrorMock error
	UUIDMock  uuid.UUID
}

func (g *GameStorageMock) Add(_ uuid.UUID, _ game.Game) error {
	return g.ErrorMock
}

func (g *GameStorageMock) Remove(_ uuid.UUID) error {
	return g.ErrorMock
}

func (g *GameStorageMock) Find(_ uuid.UUID) (game.Game, error) {
	return g.GameMock, g.ErrorMock
}

func (g *GameStorageMock) RandomMatch(_ *user.User) (uuid.UUID, error) {
	return g.UUIDMock, g.ErrorMock
}
