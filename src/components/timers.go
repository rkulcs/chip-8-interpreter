package components

type Timer struct {
	Value byte
}

func (timer *Timer) Decrement() {
	if timer.Value == 0 {
		return
	} else {
		timer.Value--
	}
}
