package storage

import (
	"talpidae-backend/model/game"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewGameStorage(t *testing.T) {
	_ = NewGameStorage()
}

var (
	testKey, _ = uuid.FromBytes([]byte("bbc142ff-4686-38ba-b4dc-76b82c8da544"))
)

func TestGameStorageAdd(t *testing.T) {
	cases := []struct {
		name        string
		beforeGames map[uuid.UUID]game.Game
		inKey       uuid.UUID
		inGame      game.Game
		afterGames  map[uuid.UUID]game.Game
		outError    error
	}{
		{
			name:        "success",
			beforeGames: make(map[uuid.UUID]game.Game),
			inKey:       testKey,
			inGame:      &game.GameMock{},
			afterGames:  map[uuid.UUID]game.Game{testKey: &game.GameMock{}},
		},
		{
			name:        "failure",
			beforeGames: map[uuid.UUID]game.Game{testKey: &game.GameMock{}},
			inKey:       testKey,
			inGame:      &game.GameMock{},
			outError:    GameStorageInvalidArgumentErr,
		},
	}

	for _, c := range cases {

		t.Run(c.name, func(t *testing.T) {
			gs := gameStorageImpl{games: c.beforeGames}
			err := gs.Add(c.inKey, c.inGame)
			if err != nil {
				assert.Equal(t, c.outError, err)
				return
			}
			assert.Equal(t, c.afterGames, gs.games)
		})
	}
}

func TestGameStorageRemove(t *testing.T) {
	cases := []struct {
		name        string
		beforeGames map[uuid.UUID]game.Game
		inKey       uuid.UUID
		afterGames  map[uuid.UUID]game.Game
		outError    error
	}{
		{
			name:        "success",
			beforeGames: map[uuid.UUID]game.Game{testKey: &game.GameMock{}},
			inKey:       testKey,
			afterGames:  make(map[uuid.UUID]game.Game),
		},
		{
			name:        "failure",
			beforeGames: map[uuid.UUID]game.Game{},
			inKey:       testKey,
			outError:    GameStorageInvalidArgumentErr,
		},
	}

	for _, c := range cases {

		t.Run(c.name, func(t *testing.T) {
			gs := gameStorageImpl{games: c.beforeGames}
			err := gs.Remove(c.inKey)
			if err != nil {
				assert.Equal(t, c.outError, err)
				return
			}
			assert.Equal(t, c.afterGames, gs.games)
		})
	}
}

func TestGameStorageFind(t *testing.T) {
	cases := []struct {
		name        string
		beforeGames map[uuid.UUID]game.Game
		inKey       uuid.UUID
		outGame     game.Game
		outError    error
	}{
		{
			name:        "success",
			beforeGames: map[uuid.UUID]game.Game{testKey: &game.GameMock{}},
			inKey:       testKey,
			outGame:     &game.GameMock{},
		},
		{
			name:        "failure",
			beforeGames: map[uuid.UUID]game.Game{},
			inKey:       testKey,
			outError:    GameStorageInvalidArgumentErr,
		},
	}

	for _, c := range cases {

		t.Run(c.name, func(t *testing.T) {
			gs := gameStorageImpl{games: c.beforeGames}
			g, err := gs.Find(c.inKey)
			if err != nil {
				assert.Equal(t, c.outError, err)
				return
			}
			assert.Equal(t, c.outGame, g)
		})
	}
}
