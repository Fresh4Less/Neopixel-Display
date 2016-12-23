package neopixeldisplay

import(
	"fmt"
)

type ScreenView struct {
	Display PixelDisplay
	Offset, Width, Height int
	Brightness float32
	overflowMode ColorOverflowMode
	frame ColorFrame
}

//ScreenView maps a 2d grid of colors onto the 1d led array
func NewScreenView(display PixelDisplay, offset, width, height int, brightness float32, overflowMode ColorOverflowMode) *ScreenView {
	if offset < 0 || width < 0 || height < 0 || offset+width*height > display.Count() {
		panic(fmt.Sprintf("NewScreenView: invalid pixel dimensions (%v,%v,%v)", offset, width, height))
	}

	sv := ScreenView{display, offset, width, height, brightness, overflowMode, MakeColorFrame(width, height, MakeColor(0,0,0))}
	sv.frame.parent = &sv

	return &sv
}

func (sv *ScreenView) GetFrame() *ColorFrame {
	return &sv.frame
}

func (sv *ScreenView) Draw() {
	for i := 0; i < sv.Height; i++ {
		for j := 0; j < sv.Width; j++ {
			color := sv.frame.Get(j,i, sv.overflowMode)
			sv.Display.Set(sv.Offset+sv.Height*i+j, MakeColor(
				uint32(float32(color.GetRed())*sv.Brightness),
				uint32(float32(color.GetGreen())*sv.Brightness),
				uint32(float32(color.GetBlue())*sv.Brightness)))
		}
	}
	sv.Display.Show()
}
