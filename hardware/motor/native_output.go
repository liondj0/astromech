package motor

type nativeOutput struct{}

func (output *nativeOutput) setLowLeft() {
	println("set low left")
}

func (output *nativeOutput) setHighLeft() {
	println("set high left")
}

func (output *nativeOutput) setLowRight() {
	println("set low right")
}

func (output *nativeOutput) setHighRight() {
	println("set high right")
}

func newNativeOutput() *nativeOutput {
	return &nativeOutput{}
}
