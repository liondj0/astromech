package voice

import "math"

const sampleRate = 44100

type tone struct {
	freqStart  float64
	freqEnd    float64
	amplitude  float64
	wobbleRate float64
	pos        int
	total      int
	duration   float64
}

func (tone *tone) Stream(samples [][2]float64) (n int, ok bool) {
	for i := range samples {
		if tone.pos >= tone.total {
			return i, false
		}
		progress := float64(tone.pos) / float64(tone.total)
		freq := tone.freqStart + (tone.freqEnd-tone.freqStart)*progress

		wobble := 1.0 + 0.3*math.Sin(2*math.Pi*tone.wobbleRate*progress)
		sample := tone.amplitude * wobble * math.Sin(2*math.Pi*freq*float64(tone.pos)/sampleRate)

		envelope := 1.0
		fadeLen := tone.total / 10
		if tone.pos < fadeLen {
			envelope = float64(tone.pos) / float64(fadeLen)
		} else if tone.pos > tone.total-fadeLen {
			envelope = float64(tone.total-tone.pos) / float64(fadeLen)
		}

		sample *= envelope
		samples[i][0] = sample
		samples[i][1] = sample
		tone.pos++
	}
	return len(samples), true
}

func (tone *tone) Err() error { return nil }
