package pin

import "github.com/stianeikeland/go-rpio/v4"

type Digital struct {
	pin rpio.Pin
}

func (d Digital) SetHigh() {
	d.pin.High()
}

func (d Digital) SetLow() {
	d.pin.Low()
}

func NewDigital(pin int) *Digital {
	d := &Digital{
		pin: rpio.Pin(pin),
	}
	d.pin.Output()
	return d
}
