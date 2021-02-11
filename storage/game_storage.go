package storage

import (
	"talpidae-backend/model/game"
	"talpidae-backend/model/user"

	"github.com/google/uuid"
)

type GameStorage interface {
	Add(uuid.UUID, game.Game) error
	Remove(uuid.UUID) error
	Find(uuid.UUID) (game.Game, error)
	RandomMatch(*user.User) (uuid.UUID, error)
}

func NewGameStorage() GameStorage {
	return &gameStorageImpl{
		games: make(map[uuid.UUID]game.Game),
	}
}

type gameStorageImpl struct {
	games map[uuid.UUID]game.Game
}

func (gs *gameStorageImpl) Add(key uuid.UUID, g game.Game) error {
	if gs.games[key] != nil {
		return GameStorageInvalidArgumentErr
	}
	gs.games[key] = g
	return nil
}

func (gs *gameStorageImpl) Remove(key uuid.UUID) error {
	if gs.games[key] == nil {
		return GameStorageInvalidArgumentErr
	}

	delete(gs.games, key)
	return nil
}

func (gs *gameStorageImpl) Find(key uuid.UUID) (game.Game, error) {
	if gs.games[key] == nil {
		return nil, GameStorageInvalidArgumentErr
	}

	return gs.games[key], nil
}

func (gs *gameStorageImpl) RandomMatch(u *user.User) (uuid.UUID, error) {
	for key, g := range gs.games {
		if len(g.Users()) < game.MaxNumberOfUsers {
			err := g.JoinUser(u)
			return key, err
		}
	}

	ng, err := game.New(game.Height, game.Width)
	if err != nil {
		return uuid.Nil, err
	}
	err = ng.JoinUser(u)
	if err != nil {
		return uuid.Nil, err
	}
	nid := uuid.New()
	err = gs.Add(nid, ng)
	return nid, err
}
