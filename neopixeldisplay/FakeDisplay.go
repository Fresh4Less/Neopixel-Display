package neopixeldisplay

type FakeDisplay struct {
	count int
}

func (fd *FakeDisplay) Set(index int, color Color) {
}

func (fd *FakeDisplay) Show() {
}

func (fd *FakeDisplay) Count() int {
	return fd.count
}
