# Astromech

A 50cm R2-D2-style autonomous home companion robot, built from scratch — 3D-printed body, custom electronics, and a hand-written Go brain. Every layer (mechanical, electrical, software) is designed and built end-to-end by hand.

## Layout

- `hardware/` — GPIO, PWM, servo, and speaker drivers (Raspberry Pi)
- `voice/` — speech/tone synthesis
- `personality/` — behavior and mood engine
- `util/` — shared helpers

## Requirements

- Go 1.23+
- Raspberry Pi (or compatible SBC) for GPIO access via [go-rpio](https://github.com/stianeikeland/go-rpio)

## Run

```
go run main.go
```
