package game

var _ Game = &GameMock{}

type GameMock struct {
	ErrorMock  error
	BlocksMock [][]BlockType
	LogsMock   []FillLog
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
