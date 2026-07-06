package servo

import "os"

type nativeOutput struct {
	frequency   int
	cycleLength uint32
}

func (output *nativeOutput) rotate(angle float64) {
	pulseWidthMs := 1.0 + (angleClamp(angle) / 180.0)
	periodMs := 1000.0 / float64(output.frequency)
	dutyCycle := uint32((pulseWidthMs / periodMs) * float64(output.cycleLength))
	if os.Getenv("USE_PINS") == "true" {
		println(`Servo rotated to angle: `, angle)
		println(`Servo rotated with duty cycle: `, dutyCycle)
		return
	}
}

func newNativeOutput(frequency int, cycleLength uint32) *nativeOutput {
	return &nativeOutput{frequency: frequency, cycleLength: cycleLength}
}
