package bowling

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrike(t *testing.T) {
	assert := assert.New(t)
	frame := Frame{firstRoll: 10}

	assert.True(frame.Strike())
	assert.False(frame.Spare())
}

func TestSpare(t *testing.T) {
	assert := assert.New(t)
	frame := Frame{firstRoll: 5, secondRoll: 5}

	assert.False(frame.Strike())
	assert.True(frame.Spare())
}
