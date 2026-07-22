package main

import (
	hardware2 "astromech/hardware"
	"astromech/hardware/motor"
	"time"
)

func main() {
	hardware := hardware2.InitHardware()
	//voice := voice2.NewVoice(hardware.Speaker)
	//voice.Speak(`Hello there!`)
	//time.Sleep(1 * time.Second)
	//voice.Speak(`General Kenobi!`)
	//time.Sleep(1 * time.Second)
	println("Turning head right")
	hardware.HeadMotor.TurnDegrees(120, motor.Right)
	time.Sleep(1 * time.Second)
	println("Turning head left")
	hardware.HeadMotor.TurnDegrees(240, motor.Left)
	time.Sleep(1 * time.Second)
	hardware.HeadMotor.TurnDegrees(120, motor.Right)
}
