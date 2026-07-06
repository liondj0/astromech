package speaker

import (
	"astromech/hardware/pin"
	"github.com/gopxl/beep/v2"
	"time"
)

type pinOutput struct {
	pwm *pin.PWM
}

func (speaker *pinOutput) normalizeCycleValue(buffer [2]float64, cycleLength uint32) uint32 {
	sample := (buffer[0] + buffer[1]) * 0.5
	return uint32((sample+1.0)*0.5*float64(cycleLength) + 0.5)
}

func (speaker *pinOutput) playSound(streamer beep.Streamer) {
	cycleLength := speaker.pwm.CycleLength()
	samplePeriod := time.Second / time.Duration(sampleRate)

	buffer := make([][2]float64, 256)

	for {
		n, ok := streamer.Stream(buffer)
		for i := 0; i < n; i++ {
			speaker.pwm.SetDutyCycle(speaker.normalizeCycleValue(buffer[i], cycleLength))
			time.Sleep(samplePeriod)
		}
		if !ok {
			break
		}
	}
	speaker.pwm.SetDutyCycle(cycleLength / 2)
}

func newPinOutput(pwm *pin.PWM) *pinOutput {
	return &pinOutput{pwm: pwm}
}
