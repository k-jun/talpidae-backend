package storage

import (
	"talpidae-backend/model/game"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGameStorage(t *testing.T) {
	_ = NewGameStorage()
}

func TestGameStorageAdd(t *testing.T) {
	cases := []struct {
		name        string
		beforeGames map[string]game.Game
		inKey       string
		inGame      game.Game
		afterGames  map[string]game.Game
		outError    error
	}{
		{
			name:        "success",
			beforeGames: make(map[string]game.Game),
			inKey:       "test",
			inGame:      &game.GameMock{},
			afterGames:  map[string]game.Game{"test": &game.GameMock{}},
		},
		{
			name:        "failure",
			beforeGames: map[string]game.Game{"test": &game.GameMock{}},
			inKey:       "test",
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
		beforeGames map[string]game.Game
		inKey       string
		afterGames  map[string]game.Game
		outError    error
	}{
		{
			name:        "success",
			beforeGames: map[string]game.Game{"test": &game.GameMock{}},
			inKey:       "test",
			afterGames:  make(map[string]game.Game),
		},
		{
			name:        "failure",
			beforeGames: map[string]game.Game{},
			inKey:       "test",
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
		beforeGames map[string]game.Game
		inKey       string
		outGame     game.Game
		outError    error
	}{
		{
			name:        "success",
			beforeGames: map[string]game.Game{"test": &game.GameMock{}},
			inKey:       "test",
			outGame:     &game.GameMock{},
		},
		{
			name:        "failure",
			beforeGames: map[string]game.Game{},
			inKey:       "test",
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
