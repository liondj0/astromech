package servo

import (
	"astromech/hardware/pin"
)

type pinOutput struct {
	pwm         *pin.PWM
	frequency   int
	cycleLength uint32
}

func (output *pinOutput) rotate(angle float64) {
	pulseWidthMs := 1.0 + (angleClamp(angle) / 180.0)
	periodMs := 1000.0 / float64(output.frequency)
	dutyCycle := uint32((pulseWidthMs / periodMs) * float64(output.cycleLength))
	output.pwm.SetDutyCycle(dutyCycle)
}

func newPinOutput(pwm *pin.PWM, frequency int, cycleLength uint32) *pinOutput {
	return &pinOutput{pwm, frequency, cycleLength}
}
