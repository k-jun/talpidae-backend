package game

var _ Game = &GameMock{}

type GameMock struct {
	OutError  error
	OutBlocks [][]BlockType
}

func (g *GameMock) Fill(_ int, _ int, _ BlockType) error {
	return g.OutError
}

func (g *GameMock) Blocks() [][]BlockType {
	return g.OutBlocks
}
