package bowling

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreEmpty(t *testing.T) {
	assert := assert.New(t)
	emptyGame := Game{}

	assert.Equal(0, emptyGame.Score())
}

func TestScoreAfterFrame(t *testing.T) {
	assert := assert.New(t)
	game := Game{}

	game.Roll(3)
	game.Roll(8)

	assert.Equal(11, game.Score())
	assert.Len(game.frames, 1)
}

func TestScoreAfterManyFrames(t *testing.T) {
	assert := assert.New(t)
	game := Game{}

	game.Roll(3)
	game.Roll(3)
	game.Roll(1)
	game.Roll(2)
	game.Roll(3)
	game.Roll(4)

	assert.Equal(16, game.Score())
	assert.Len(game.frames, 3)
}

func TestScoreDuringFrame(t *testing.T) {
	assert := assert.New(t)
	game := Game{}

	game.Roll(3)

	assert.Equal(3, game.Score())
	assert.Len(game.frames, 1)
}

func TestScoreStrike(t *testing.T) {
	assert := assert.New(t)
	game := Game{}

	game.Roll(10)
	assert.Equal(10, game.Score())

	game.Roll(3)
	assert.Equal(16, game.Score())
	assert.Equal(2, len(game.frames))

	game.Roll(3)
	assert.Equal(22, game.Score())
	assert.Equal(2, len(game.frames))
}

func TestScoreSpare(t *testing.T) {
	assert := assert.New(t)
	game := Game{}

	game.Roll(5)
	assert.Equal(5, game.Score())

	game.Roll(5)
	assert.Equal(10, game.Score())

	game.Roll(3)
	assert.Equal(16, game.Score())
	assert.Equal(2, len(game.frames))

	game.Roll(4)
	assert.Equal(20, game.Score())
	assert.Equal(2, len(game.frames))
}

func TestPerfectGame(t *testing.T) {
	assert := assert.New(t)
	game := Game{}

	for i := 0; i < 11; i++ {
		game.Roll(10)
	}

	assert.Equal(210, game.Score())
}

func TestMaxFramesWithoutBonus(t *testing.T) {
	assert := assert.New(t)
	game := Game{}

	for i := 0; i < 20; i++ {
		game.Roll(1)
	}

	assert.Error(game.Roll(1))
}

func TestMaxFramesWithStrikeBonus(t *testing.T) {
	assert := assert.New(t)
	game := Game{}

	for i := 0; i < 18; i++ {
		game.Roll(1)
	}
	game.Roll(10)

	assert.Nil(game.Roll(1))
}

func TestMaxFramesWithSpareBonus(t *testing.T) {
	assert := assert.New(t)
	game := Game{}

	for i := 0; i < 18; i++ {
		game.Roll(1)
	}
	game.Roll(5)
	game.Roll(5)

	assert.Nil(game.Roll(1))
}
