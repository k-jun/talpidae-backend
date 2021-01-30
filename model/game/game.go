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
	Fill(string, BlockType, int, int) error
	Blocks() [][]BlockType
	Logs() []FillLog
}

type FillLog struct {
	Height int
	Width  int
	Value  BlockType
	UserId string
}

type gameImpl struct {
	blocks [][]BlockType
	logs   []FillLog
}

type BlockType int

const (
	SakuSaku       BlockType = iota
	KachiKachi     BlockType = 1
	GochiGochi     BlockType = 2
	Treasure       BlockType = 3
	ArrowUp        BlockType = 4
	ArrowDown      BlockType = 5
	ArrowLeft      BlockType = 6
	ArrowRight     BlockType = 7
	WanaArrowUp    BlockType = 8
	WanaArrowDown  BlockType = 9
	WanaArrowLeft  BlockType = 10
	WanaArrowRight BlockType = 11
)

func validateBlockType(x BlockType) bool {
	all := []BlockType{WanaArrowUp, WanaArrowDown, WanaArrowLeft, WanaArrowRight}
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
	newGame := &gameImpl{blocks: blocks, logs: []FillLog{}}

	// treasures
	cnt := 0
	for {
		h := rand.Intn(height)
		w := rand.Intn(width)
		if blocks[h][w] != SakuSaku {
			continue
		}
		blocks[h][w] = Treasure
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
		if blocks[h][w] != SakuSaku {
			continue
		}
		blocks[h][w] = newGame.closestTreasureArrowDirection(h, w)
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

func (g *gameImpl) Fill(userId string, value BlockType, height int, width int) error {
	if !validateBlockType(value) || userId == "" {
		return InvalidArgumentErr
	}

	if height < 0 || height >= len(g.blocks) {
		return InvalidArgumentErr
	}

	if width < 0 || width >= len(g.blocks[height]) {
		return InvalidArgumentErr
	}

	if g.blocks[height][width] != SakuSaku {
		return InvalidArgumentErr
	}

	g.blocks[height][width] = value
	g.logs = append(g.logs, FillLog{Height: height, Width: width, Value: value, UserId: userId})
	return nil
}

func (g *gameImpl) Blocks() [][]BlockType {
	return g.blocks
}

func (g *gameImpl) Logs() []FillLog {
	return g.logs
}
