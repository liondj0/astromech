package speaker

import (
	"astromech/hardware/pin"
	"github.com/gopxl/beep/v2"
	"os"
)

const sampleRate = 44100
const cycleLength = 1024

type output interface {
	playSound(streamer beep.Streamer)
}

func newOutput(pinIndex int) output {
	if os.Getenv("USE_PINS") == "true" {
		return newPinOutput(pin.NewPWM(pinIndex, cycleLength, sampleRate*cycleLength))
	}
	return newNativeOutput()
}

type Speaker struct {
	out output
}

func (speaker *Speaker) PlaySound(streamer beep.Streamer) {
	speaker.out.playSound(streamer)
}

func NewSpeaker(pinIndex int) *Speaker {
	speaker := Speaker{
		out: newOutput(pinIndex),
	}

	return &speaker
}
