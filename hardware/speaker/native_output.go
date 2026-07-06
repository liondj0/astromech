package speaker

import (
	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/speaker"
)

type nativeOutput struct{}

func (o *nativeOutput) playSound(streamer beep.Streamer) {
	done := make(chan struct{})
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		close(done)
	})))
	<-done
}

func newNativeOutput() *nativeOutput {
	speaker.Init(beep.SampleRate(sampleRate), sampleRate/10)
	return &nativeOutput{}
}
