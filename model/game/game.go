package game

import (
	"math/rand"
	"sync"
	"talpidae-backend/model/user"
	"time"
)

const (
	TreasureCnt      = 3
	ArrowCnt         = TreasureCnt * 3
	MaxNumberOfUsers = 4
	Height           = 150
	Width            = 80
)

type Game interface {
	Fill(string, BlockType, int, int) error
	Blocks() [][]BlockType
	Logs() []FillLog
	Users() []*user.User
	JoinUser(*user.User) error
	LeaveUser(*user.User) error
}

type FillLog struct {
	Height int
	Width  int
	Value  BlockType
	UserId string
}

type gameImpl struct {
	sync.Mutex
	blocks [][]BlockType
	logs   []FillLog
	users  []*user.User
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
	FakeArrowUp    BlockType = 8
	FakeArrowDown  BlockType = 9
	FakeArrowLeft  BlockType = 10
	FakeArrowRight BlockType = 11
	TrapArrow      BlockType = 12
	TrapTreasure   BlockType = 13
)

func validateBlockType(x BlockType) bool {
	all := []BlockType{FakeArrowUp, FakeArrowDown, FakeArrowLeft, FakeArrowRight, TrapArrow, TrapTreasure}
	for _, bt := range all {
		if bt == x {
			return true
		}
	}
	return false
}

func isFillable(x BlockType) bool {
	if x == Treasure || x == ArrowUp || x == ArrowDown || x == ArrowLeft || x == ArrowRight ||
		x == FakeArrowUp || x == FakeArrowDown || x == FakeArrowLeft || x == FakeArrowRight ||
		x == TrapArrow || x == TrapTreasure {
		return false
	}
	return true
}

func randomBlockType() BlockType {
	x := rand.Intn(10)
	if x < 5 {
		return SakuSaku
	} else if x < 8 {
		return KachiKachi
	} else {
		return GochiGochi
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func New(height int, width int) (Game, error) {
	blocks := [][]BlockType{}
	for i := 0; i < height; i++ {
		row := []BlockType{}
		for j := 0; j < width; j++ {
			row = append(row, randomBlockType())
		}
		blocks = append(blocks, row)
	}
	newGame := &gameImpl{blocks: blocks, logs: []FillLog{}}

	// treasures
	cnt := 0
	for {
		h := rand.Intn(height)
		w := rand.Intn(width)
		if !isFillable(newGame.blocks[h][w]) {
			continue
		}
		newGame.blocks[h][w] = Treasure
		newGame.surroundWith(GochiGochi, h, w, 1)
		newGame.surroundWith(KachiKachi, h, w, 2)
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
		if !isFillable(newGame.blocks[h][w]) {
			continue
		}
		newGame.blocks[h][w] = newGame.closestTreasureArrowDirection(h, w)
		newGame.surroundWith(KachiKachi, h, w, 1)
		cnt += 1
		if cnt >= ArrowCnt {
			break
		}
	}

	return newGame, nil
}

func (g *gameImpl) surroundWith(b BlockType, h, w, n int) {
	round_indexes := g.roundIndexes(h, w, n)
	for _, idx := range round_indexes {
		if !isFillable(g.blocks[idx[0]][idx[1]]) {
			continue
		}
		g.blocks[idx[0]][idx[1]] = b
	}
}

func (g *gameImpl) roundIndexes(h, w, n int) [][2]int {
	type index struct {
	}
	indexes := [][2]int{}

	// top & down line
	for i := w - n + 1; i < w+n; i++ {
		indexes = append(indexes, [2]int{h - n, i})
		indexes = append(indexes, [2]int{h + n, i})
	}
	// left & right line
	for i := h - n + 1; i < h+n; i++ {
		indexes = append(indexes, [2]int{i, w + n})
		indexes = append(indexes, [2]int{i, w - n})
	}
	// corners
	indexes = append(indexes, [2]int{h - n, w - n})
	indexes = append(indexes, [2]int{h + n, w - n})
	indexes = append(indexes, [2]int{h - n, w + n})
	indexes = append(indexes, [2]int{h + n, w + n})

	valid_indexes := [][2]int{}
	for _, v := range indexes {
		if v[0] >= 0 && v[0] < len(g.blocks) && v[1] >= 0 && v[1] < len(g.blocks[0]) {
			valid_indexes = append(valid_indexes, v)
		}
	}
	return valid_indexes
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

	if !isFillable(g.blocks[height][width]) {
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

func (g *gameImpl) Users() []*user.User {
	return g.users
}

func (g *gameImpl) JoinUser(u *user.User) error {
	g.Lock()
	defer g.Unlock()
	if u == nil {
		return InvalidArgumentErr
	}
	g.users = append(g.users, u)
	return nil
}

func (g *gameImpl) LeaveUser(u *user.User) error {
	g.Lock()
	defer g.Unlock()
	if u == nil {
		return InvalidArgumentErr
	}

	for i, user := range g.users {
		if u == user {
			g.users = append(g.users[:i], g.users[i+1:]...)
			return nil
		}
	}

	return InvalidArgumentErr
}
