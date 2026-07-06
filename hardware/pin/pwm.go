package pin

import (
	"github.com/stianeikeland/go-rpio/v4"
)

type PWM struct {
	pin         rpio.Pin
	cycleLength uint32
	frequency   int
}

func (pwm *PWM) SetDutyCycle(dutyCycle uint32) {
	if dutyCycle > pwm.cycleLength {
		dutyCycle = pwm.cycleLength
	}
	pwm.pin.DutyCycle(dutyCycle, pwm.cycleLength)
}

func (pwm *PWM) CycleLength() uint32 {
	return pwm.cycleLength
}

func NewPWM(pin int, cycleLength uint32, frequency int) *PWM {
	pwm := &PWM{rpio.Pin(pin), cycleLength, frequency}
	pwm.pin.Mode(rpio.Pwm)
	pwm.pin.Freq(frequency)
	return pwm
}
