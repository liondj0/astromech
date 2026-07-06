package servo

import (
	"astromech/hardware/pin"
	"astromech/util"
	"os"
)

type Servo struct {
	output output
}

var angleClamp = util.ClampBuilder(0, 180)

type output interface {
	rotate(angle float64)
}

func (servo *Servo) SetAngle(angle float64) {
	servo.output.rotate(angle)
}

func NewServo(pinIndex int) *Servo {
	const frequency = 50
	const cycleLength = 1024
	if os.Getenv("USE_PINS") == "true" {
		return &Servo{
			output: newPinOutput(pin.NewPWM(pinIndex, cycleLength, frequency), frequency, cycleLength),
		}
	}
	return &Servo{
		output: newNativeOutput(frequency, cycleLength),
	}
}
