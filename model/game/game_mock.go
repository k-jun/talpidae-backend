package game

import "talpidae-backend/model/user"

var _ Game = &GameMock{}

type GameMock struct {
	ErrorMock  error
	BlocksMock [][]BlockType
	LogsMock   []FillLog
	UsersMock  []*user.User
}

func (g *GameMock) Fill(_ string, _ BlockType, _ int, _ int) error {
	return g.ErrorMock
}

func (g *GameMock) Blocks() [][]BlockType {
	return g.BlocksMock
}

func (g *GameMock) Logs() []FillLog {
	return g.LogsMock
}

func (g *GameMock) Users() []*user.User {
	return g.UsersMock
}

func (g *GameMock) JoinUser(_ *user.User) error {
	return g.ErrorMock
}

func (g *GameMock) LeaveUser(_ *user.User) error {
	return g.ErrorMock
}
