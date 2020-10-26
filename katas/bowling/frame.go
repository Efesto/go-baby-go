package bowling

// Frame represents a game's frame
type Frame struct {
	firstRoll  int
	secondRoll int
	finished   bool
}

// Strike returns true if frame contains a strike
func (f *Frame) Strike() bool {
	return f.firstRoll == maxPins
}

// Spare returns true if a frame contains a spare
func (f *Frame) Spare() bool {
	return f.firstRoll+f.secondRoll == maxPins && !f.Strike()
}
