package game

var _ Game = &GameMock{}

type GameMock struct {
	outError  error
	outBlocks [][]BlockType
}

func (g *GameMock) Fill(_ int, _ int, _ BlockType) error {
	return g.outError
}

func (g *GameMock) Blocks() [][]BlockType {
	return g.outBlocks
}
