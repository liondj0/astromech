package hardware

import (
	"astromech/hardware/motor"
	"astromech/hardware/servo"
	"astromech/hardware/speaker"
	"github.com/stianeikeland/go-rpio/v4"
)

const speakerPinIndex = 18
const servoPinIndex = 19
const headMotorLeftRotationPin = 17
const headMotorRightRotationPin = 27
const headMotorRPM = 30

type Hardware struct {
	Speaker   *speaker.Speaker
	Servo     *servo.Servo
	HeadMotor *motor.Motor
}

func InitHardware() *Hardware {
	if err := rpio.Open(); err != nil {
		println("failed to open speaker")
	}
	return &Hardware{
		Speaker:   speaker.NewSpeaker(speakerPinIndex),
		Servo:     servo.NewServo(servoPinIndex),
		HeadMotor: motor.NewMotor(headMotorLeftRotationPin, headMotorRightRotationPin, headMotorRPM),
	}
}
