package motor

import (
	"astromech/hardware/pin"
	"os"
	"time"
)

type Direction int

const (
	Left  Direction = iota
	Right Direction = 1
)

type output interface {
	setLowLeft()
	setHighLeft()
	setLowRight()
	setHighRight()
}

type Motor struct {
	output output
	rpm    int
}

func (motor *Motor) TurnMilliseconds(milliseconds int, direction Direction) {
	if direction == Left {
		motor.output.setLowRight()
		motor.output.setHighLeft()
	} else {
		motor.output.setLowLeft()
		motor.output.setHighRight()
	}
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
	if direction == Left {
		motor.output.setLowLeft()
	} else {
		motor.output.setLowRight()
	}
}

func (motor *Motor) TurnDegrees(degrees int, direction Direction) {
	if degrees < 1 {
		println("Cannot turn less then 1 degree")
		return
	}
	fullCircleSeconds := 60 / motor.rpm
	degreesToSeconds := float64(fullCircleSeconds*degrees) / 360.0
	motor.TurnMilliseconds(int(degreesToSeconds*1000), direction)
}

func NewMotor(leftPin, rightPin, rpm int) *Motor {
	if os.Getenv("USE_PINS") != "true" {
		return &Motor{output: newNativeOutput(), rpm: rpm}
	}
	return &Motor{output: newPinOutput(pin.NewDigital(leftPin), pin.NewDigital(rightPin)), rpm: rpm}
}
