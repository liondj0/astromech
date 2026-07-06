package personality

import "time"

type Config struct {
	decaySpeed  time.Duration
	decayFactor float64
}

type Engine struct {
	coords       PadCoords
	mood         Mood
	config       Config
	OnMoodChange func()
}

func (engine *Engine) UpdateMood(coords PadCoords) {
	engine.coords = newCoords(coords)
	engine.mood = calcMoodFromCoords(engine.coords)
	engine.OnMoodChange()
}

func (engine *Engine) Nudge(coords PadCoords) {
	pleasure := engine.coords.pleasure + coords.pleasure
	arousal := engine.coords.arousal + coords.arousal
	dominance := engine.coords.dominance + coords.dominance
	engine.UpdateMood(PadCoords{pleasure: pleasure, arousal: arousal, dominance: dominance})
}

func (engine *Engine) Decay() {
	idleCoords := moodCoords[Idle]
	coords := engine.coords
	decayedCoords := PadCoords{}
	decayedCoords.pleasure = coords.pleasure + ((idleCoords.pleasure - coords.pleasure) * engine.config.decayFactor)
	decayedCoords.arousal = coords.arousal + ((idleCoords.arousal - coords.arousal) * engine.config.decayFactor)
	decayedCoords.dominance = coords.dominance + ((idleCoords.dominance - coords.dominance) * engine.config.decayFactor)
	engine.UpdateMood(decayedCoords)
}

func (engine *Engine) Mood() string {
	return Moods[engine.mood]
}

func NewEngine() *Engine {
	return &Engine{
		mood:         Excited,
		coords:       moodCoords[Excited],
		OnMoodChange: func() {},
		config: Config{
			decaySpeed:  2 * time.Minute,
			decayFactor: 0.05,
		},
	}
}
