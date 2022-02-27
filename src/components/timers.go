package components

type DelayTimer struct {
	Value byte
}

func (timer *DelayTimer) Decrement() {
	if timer.Value == 0 {
		return
	} else {
		timer.Value--
	}
}
