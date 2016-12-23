package neopixeldisplay

import(
	"time"
	"github.com/fresh4less/neopixel-display/ws2811"
)

type PixelDisplay interface {
	Set(index int, color Color)
	Show()
	Count() int
}

type NeopixelDisplay struct {
	count int
	cooloffTimer *time.Timer
	timerRunning bool
	renderAfterTimer bool
}

const DisplayCooloffTime = 5*time.Millisecond

//interface to the neopixel API
//The neopixels stop working if we try to call Render too many times, so we debounce signals to at most one per 5ms
func NewNeopixelDisplay(gpioPin, ledCount, brightness int) *NeopixelDisplay {
	err := ws2811.Init(gpioPin, ledCount, brightness)
	if err != nil {
		panic(err)
	}
	display := NeopixelDisplay{ledCount, time.NewTimer(DisplayCooloffTime), false, false}
	return &display
}

func (nd *NeopixelDisplay) Count() int {
	return nd.count
}

func (nd *NeopixelDisplay) Set(index int, color Color) {
	ws2811.SetLed(index, uint32(color))
}

func (nd *NeopixelDisplay) Show() {
	if nd.timerRunning {
		if !nd.renderAfterTimer {
			nd.renderAfterTimer = true
			go func() {
				<-nd.cooloffTimer.C
				nd.render()
			}()
		}
	} else {
	nd.render()
	}
}

func (nd *NeopixelDisplay) render() {
	err := ws2811.Render()
	if err != nil {
		panic(err)
	}

	//set timer
	nd.cooloffTimer.Reset(DisplayCooloffTime)
	nd.timerRunning = true
	nd.renderAfterTimer = false
}
