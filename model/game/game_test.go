package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	_, err := New(10, 10)
	assert.NoError(t, err)
}

func TestFill(t *testing.T) {
	cases := []struct {
		name         string
		beforeBlocks [][]BlockType
		inUserId     string
		inBlockType  BlockType
		inHeight     int
		inWidth      int
		afterBlocks  [][]BlockType
		afterLogs    []FillLog
		outError     error
	}{
		{
			name:         "success",
			beforeBlocks: [][]BlockType{{SakuSaku, SakuSaku}, {SakuSaku, SakuSaku}},
			inUserId:     "38f6e080-14df-3efa-abe2-9b01943eebd0",
			inBlockType:  FakeArrowUp,
			inHeight:     0,
			inWidth:      1,
			afterBlocks:  [][]BlockType{{SakuSaku, FakeArrowUp}, {SakuSaku, SakuSaku}},
			afterLogs:    []FillLog{{UserId: "38f6e080-14df-3efa-abe2-9b01943eebd0", Value: FakeArrowUp, Height: 0, Width: 1}},
			outError:     nil,
		},
		{
			name:         "failure: invalid position",
			beforeBlocks: [][]BlockType{{Treasure, SakuSaku}, {SakuSaku, SakuSaku}},
			inUserId:     "38f6e080-14df-3efa-abe2-9b01943eebd0",
			inBlockType:  FakeArrowUp,
			inHeight:     0,
			inWidth:      0,
			outError:     InvalidArgumentErr,
		},
		{
			name:         "failure: invalid UserId",
			beforeBlocks: [][]BlockType{{SakuSaku, SakuSaku}, {SakuSaku, SakuSaku}},
			inBlockType:  FakeArrowUp,
			inHeight:     0,
			inWidth:      0,
			outError:     InvalidArgumentErr,
		},
		{
			name:         "failure: invalid block type",
			beforeBlocks: [][]BlockType{{SakuSaku, SakuSaku}, {SakuSaku, SakuSaku}},
			inUserId:     "38f6e080-14df-3efa-abe2-9b01943eebd0",
			inBlockType:  Treasure,
			inHeight:     0,
			inWidth:      0,
			outError:     InvalidArgumentErr,
		},
		{
			name:         "failure: invalid height",
			beforeBlocks: [][]BlockType{{SakuSaku, SakuSaku}, {SakuSaku, SakuSaku}},
			inUserId:     "38f6e080-14df-3efa-abe2-9b01943eebd0",
			inBlockType:  FakeArrowUp,
			inHeight:     2,
			inWidth:      1,
			outError:     InvalidArgumentErr,
		},
		{
			name:         "failure: invalid width",
			beforeBlocks: [][]BlockType{{SakuSaku, SakuSaku}, {SakuSaku, SakuSaku}},
			inUserId:     "38f6e080-14df-3efa-abe2-9b01943eebd0",
			inHeight:     0,
			inWidth:      2,
			inBlockType:  FakeArrowUp,
			outError:     InvalidArgumentErr,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			g := gameImpl{blocks: c.beforeBlocks}
			err := g.Fill(c.inUserId, c.inBlockType, c.inHeight, c.inWidth)
			if err != nil {
				assert.Equal(t, c.outError, err)
				return
			}
			assert.Equal(t, c.afterBlocks, g.blocks)
			assert.Equal(t, c.afterLogs, g.logs)

		})
	}
}
