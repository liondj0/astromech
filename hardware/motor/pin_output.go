package motor

import "astromech/hardware/pin"

type pinOutput struct {
	leftPin  *pin.Digital
	rightPin *pin.Digital
}

func (output *pinOutput) setLowLeft() {
	output.leftPin.SetLow()
}

func (output *pinOutput) setHighLeft() {
	output.leftPin.SetHigh()
}

func (output *pinOutput) setLowRight() {
	output.rightPin.SetLow()
}

func (output *pinOutput) setHighRight() {
	output.rightPin.SetHigh()
}

func newPinOutput(leftPin, rightPin *pin.Digital) *pinOutput {
	return &pinOutput{leftPin: leftPin, rightPin: rightPin}
}
