package storage

import "talpidae-backend/model/game"

type GameStorage interface {
	Add(string, game.Game) error
	Remove(string) error
	Find(string) (game.Game, error)
}

func NewGameStorage() GameStorage {
	return &gameStorageImpl{
		games: make(map[string]game.Game),
	}
}

type gameStorageImpl struct {
	games map[string]game.Game
}

func (gs *gameStorageImpl) Add(key string, g game.Game) error {
	if gs.games[key] != nil {
		return GameStorageInvalidArgumentErr
	}
	gs.games[key] = g
	return nil
}

func (gs *gameStorageImpl) Remove(key string) error {
	if gs.games[key] == nil {
		return GameStorageInvalidArgumentErr
	}

	delete(gs.games, key)
	return nil
}

func (gs *gameStorageImpl) Find(key string) (game.Game, error) {
	if gs.games[key] == nil {
		return nil, GameStorageInvalidArgumentErr
	}

	return gs.games[key], nil
}
