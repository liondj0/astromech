package hardware

import (
	"astromech/hardware/servo"
	"astromech/hardware/speaker"
)

const speakerPinIndex = 18
const servoPinIndex = 19

type Hardware struct {
	Speaker *speaker.Speaker
	Servo   *servo.Servo
}

func InitHardware() *Hardware {
	return &Hardware{Speaker: speaker.NewSpeaker(speakerPinIndex), Servo: servo.NewServo(servoPinIndex)}
}
