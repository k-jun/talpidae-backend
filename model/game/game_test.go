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
		inHeight     int
		inWidth      int
		inBlockType  BlockType
		afterBlocks  [][]BlockType
		outError     error
	}{
		{
			name:         "success",
			beforeBlocks: [][]BlockType{{SakuSaku, SakuSaku}, {SakuSaku, SakuSaku}},
			inHeight:     0,
			inWidth:      1,
			inBlockType:  Treasure,
			afterBlocks:  [][]BlockType{{SakuSaku, Treasure}, {SakuSaku, SakuSaku}},
			outError:     nil,
		},
		{
			name:         "failure: invalid position",
			beforeBlocks: [][]BlockType{{Treasure, SakuSaku}, {SakuSaku, SakuSaku}},
			inHeight:     0,
			inWidth:      0,
			inBlockType:  Treasure,
			outError:     InvalidArgumentErr,
		},
		{
			name:         "failure: invalid block type",
			beforeBlocks: [][]BlockType{{SakuSaku, SakuSaku}, {SakuSaku, SakuSaku}},
			inHeight:     0,
			inWidth:      0,
			inBlockType:  -1,
			outError:     InvalidArgumentErr,
		},
		{
			name:         "failure: invalid height",
			beforeBlocks: [][]BlockType{{SakuSaku, SakuSaku}, {SakuSaku, SakuSaku}},
			inHeight:     2,
			inWidth:      1,
			inBlockType:  Treasure,
			outError:     InvalidArgumentErr,
		},
		{
			name:         "failure: invalid width",
			beforeBlocks: [][]BlockType{{SakuSaku, SakuSaku}, {SakuSaku, SakuSaku}},
			inHeight:     0,
			inWidth:      2,
			inBlockType:  Treasure,
			outError:     InvalidArgumentErr,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			g := gameImpl{blocks: c.beforeBlocks}
			err := g.Fill(c.inHeight, c.inWidth, c.inBlockType)
			if err != nil {
				assert.Equal(t, c.outError, err)
				return
			}
			assert.Equal(t, c.afterBlocks, g.blocks)

		})
	}
}
