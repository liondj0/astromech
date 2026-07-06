package main

import (
	hardware2 "astromech/hardware"
	voice2 "astromech/voice"
	"time"
)

func main() {
	hardware := hardware2.InitHardware()
	voice := voice2.NewVoice(hardware.Speaker)
	voice.Speak(`Hello there!`)
	time.Sleep(1 * time.Second)
	voice.Speak(`General Kenobi!`)
}
