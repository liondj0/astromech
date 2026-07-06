package voice

import (
	"astromech/hardware/speaker"
	"strings"
	"time"
)

type Voice struct {
	speaker *speaker.Speaker
}

func (voice *Voice) Speak(sentance string) {
	upperCase := strings.ToUpper(sentance)
	characters := []rune(upperCase)

	for _, character := range characters {
		letterSound, ok := letterShapes[character]
		if !ok {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		letterSound.total = int(sampleRate * (letterSound.duration * 0.65))
		voice.speaker.PlaySound(&letterSound)
	}
}

func NewVoice(speaker *speaker.Speaker) *Voice {
	return &Voice{speaker: speaker}
}
