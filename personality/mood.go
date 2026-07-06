package personality

import "math"
import "astromech/util"

type Mood int

const (
	Idle Mood = iota
	Happy
	Excited
	Scared
	Curious
	Annoyed
	Sad
	Angry
	Worried
)

var Moods = map[Mood]string{
	Idle:    "idle",
	Happy:   "happy",
	Excited: "excited",
	Scared:  "scared",
	Curious: "curious",
	Annoyed: "annoyed",
	Sad:     "sad",
	Angry:   "angry",
	Worried: "worried",
}

type PadCoords struct{ pleasure, arousal, dominance float64 }

var padClamp = util.ClampBuilder(-1, 1)

func newCoords(coords PadCoords) PadCoords {
	return PadCoords{
		pleasure:  padClamp(coords.pleasure),
		arousal:   padClamp(coords.arousal),
		dominance: padClamp(coords.dominance),
	}
}

var moodCoords = map[Mood]PadCoords{
	Idle:    {0.0, -0.7, 0.0},
	Happy:   {0.7, 0.2, 0.3},
	Excited: {0.8, 0.8, 0.4},
	Scared:  {-0.7, 0.7, -0.8},
	Curious: {0.2, 0.3, 0.1},
	Annoyed: {-0.4, 0.4, 0.6},
	Sad:     {-0.7, -0.7, -0.3},
	Angry:   {-0.8, 0.8, 0.8},
	Worried: {-0.5, 0.3, -0.6},
}

func calcMoodFromCoords(coords PadCoords) Mood {

	best := Idle
	bestDistance := math.MaxFloat64

	for mood, moodCoord := range moodCoords {
		pleasureDistance := moodCoord.pleasure - coords.pleasure
		arousalDistance := moodCoord.arousal - coords.arousal
		dominanceDistance := moodCoord.dominance - coords.dominance

		distance := math.Sqrt(pleasureDistance*pleasureDistance + arousalDistance*arousalDistance + dominanceDistance*dominanceDistance)
		if distance < bestDistance {
			best = mood
			bestDistance = distance
		}
	}

	return best
}
