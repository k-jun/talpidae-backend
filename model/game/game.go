package game

import (
	"math/rand"
	"time"
)

const (
	TreasureCnt = 3
	ArrowCnt    = TreasureCnt * 3
)

type Game interface {
	Fill(int, int, BlockType) error
	Blocks() [][]BlockType
}

type gameImpl struct {
	blocks [][]BlockType
}

type BlockType string

var (
	Treasure   BlockType = "treasure"
	ArrowRight BlockType = "arrow-right"
	ArrowLeft  BlockType = "arrow-left"
	ArrowUp    BlockType = "arrow-up"
	ArrowDown  BlockType = "arrow-down"
)

func validateBlockType(x BlockType) bool {
	all := []BlockType{Treasure, ArrowUp, ArrowDown, ArrowLeft, ArrowRight}
	for _, bt := range all {
		if bt == x {
			return true
		}
	}
	return false
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func New(height int, width int) (Game, error) {
	if height*width < TreasureCnt+ArrowCnt {
		return nil, InvalidArgumentErr
	}
	blocks := make([][]BlockType, height)
	for i := 0; i < height; i++ {
		blocks[i] = make([]BlockType, width)
	}
	newGame := &gameImpl{blocks}

	// treasures
	cnt := 0
	for {
		h := rand.Intn(height)
		w := rand.Intn(width)
		if err := newGame.Fill(h, w, Treasure); err != nil {
			continue
		}
		cnt += 1
		if cnt >= TreasureCnt {
			break
		}
	}
	// arrows
	cnt = 0
	for {
		h := rand.Intn(height)
		w := rand.Intn(width)
		if err := newGame.Fill(h, w, newGame.closestTreasureArrowDirection(h, w)); err != nil {
			continue
		}
		cnt += 1
		if cnt >= ArrowCnt {
			break
		}
	}

	return newGame, nil
}

func (g *gameImpl) closestTreasureArrowDirection(h int, w int) BlockType {
	abs := func(x int) int {
		if x < 0 {
			x = -x
		}
		return x
	}
	pos := [2]int{len(g.blocks) * 3, len(g.blocks[0]) * 3}
	for i := 0; i < len(g.blocks); i++ {
		for j := 0; j < len(g.blocks[i]); j++ {
			if g.blocks[i][j] == Treasure {
				if abs(h-i)+abs(w-j) < abs(h-pos[0])+abs(w-pos[1]) {
					pos = [2]int{i, j}
				}
			}
		}
	}

	y := pos[0] - h
	x := pos[1] - w
	if abs(y) >= abs(x) {
		if y > 0 {
			return ArrowDown
		} else {
			return ArrowUp
		}
	} else {
		if x > 0 {
			return ArrowRight
		} else {
			return ArrowLeft
		}

	}
}

func (g *gameImpl) Fill(height int, width int, value BlockType) error {
	if !validateBlockType(value) {
		return InvalidArgumentErr
	}

	if height < 0 || height >= len(g.blocks) {
		return InvalidArgumentErr
	}

	if width < 0 || width >= len(g.blocks[height]) {
		return InvalidArgumentErr
	}

	if g.blocks[height][width] != "" {
		return InvalidArgumentErr
	}

	g.blocks[height][width] = value
	return nil
}

func (g *gameImpl) Blocks() [][]BlockType {
	return g.blocks
}
