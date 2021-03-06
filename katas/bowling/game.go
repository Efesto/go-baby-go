package bowling

import (
	"fmt"
)

var maxPins int = 10
var maxFrames int = 10

// Game represents a bowling game
type Game struct {
	frames []Frame
}

// Roll takes number of pins knocked down
func (g *Game) Roll(pins int) error {
	if len(g.frames) > 0 {
		lastFrame := &g.frames[len(g.frames)-1]

		if lastFrame.finished {
			if len(g.frames) < maxFrames || lastFrame.Strike() || lastFrame.Spare() {
				g.appendFrame(pins)
			} else {
				return fmt.Errorf("Cannot roll more")
			}
		} else {
			lastFrame.secondRoll = pins
			lastFrame.finished = true
		}
	} else {
		g.appendFrame(pins)
	}

	return nil
}

func (g *Game) appendFrame(pins int) {
	newFrame := Frame{firstRoll: pins}
	if newFrame.Strike() {
		newFrame.finished = true
	}

	g.frames = append(g.frames, newFrame)
}

// Score calculates game's score
func (g *Game) Score() int {
	total := 0
	addStrikeBonus := false
	addFirstRollBonus := false
	addSpareBonus := false
	for index, frame := range g.frames {
		if index < 10 {
			total += frame.firstRoll + frame.secondRoll
		}

		if addFirstRollBonus {
			total += frame.firstRoll
			addFirstRollBonus = false
		}

		if addStrikeBonus && index < 11 {
			if frame.Strike() {
				total += frame.firstRoll
				addFirstRollBonus = true
			} else {
				total += frame.firstRoll + frame.secondRoll
			}
		}

		if addSpareBonus {
			total += frame.firstRoll
		}

		addStrikeBonus = frame.Strike()
		addSpareBonus = frame.Spare()
	}

	return total
}
