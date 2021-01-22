package game

import (
	"math/rand"
	"time"
)

const (
	OtakaraCnt   = 3
	YazirushiCnt = OtakaraCnt * 3
)

type Game interface {
	Fill(int, int, BlockType) error
	Blocks() [][]BlockType
}

type gameImpl struct {
	blocks [][]BlockType
}

type BlockType = string

var (
	// Sakusaku BlockType = "sakusaku"
	// Katikati BlockType = "katikati"
	// Gotigoti BlockType = "gotigoti"
	Otakara  BlockType = "otakara"
	Wanawana BlockType = "wanawana"
	Yazirusi BlockType = "yazirusi"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func New(height int, width int) (Game, error) {
	if height*width < OtakaraCnt+YazirushiCnt {
		return nil, InvalidArgumentErr
	}
	blocks := make([][]BlockType, height)
	for i := 0; i < height; i++ {
		blocks[i] = make([]BlockType, width)
	}

	// otakara
	cnt := 0
	for {
		if !random_fill(blocks, height, width, Otakara) {
			continue
		}
		cnt += 1
		if cnt >= OtakaraCnt {
			break
		}
	}
	// yazirusi
	for {
		if !random_fill(blocks, height, width, Yazirusi) {
			continue
		}
		cnt += 1
		if cnt >= YazirushiCnt {
			break
		}
	}

	return &gameImpl{blocks}, nil
}

func random_fill(blocks [][]BlockType, maxHeight int, maxWidth int, value BlockType) bool {
	rh := rand.Intn(maxHeight)
	rw := rand.Intn(maxWidth)
	if blocks[rh][rw] != "" {
		return false
	}
	blocks[rh][rw] = value
	return true

}

func (g *gameImpl) Fill(height int, width int, value BlockType) error {
	if height < 0 || height >= len(g.blocks) {
		return InvalidArgumentErr
	}

	if width < 0 || width >= len(g.blocks[height]) {
		return InvalidArgumentErr
	}

	if g.blocks[height][width] == Otakara || g.blocks[height][width] == Yazirusi {
		return InvalidArgumentErr
	}

	g.blocks[height][width] = value
	return nil
}

func (g *gameImpl) Blocks() [][]BlockType {
	return g.blocks
}
