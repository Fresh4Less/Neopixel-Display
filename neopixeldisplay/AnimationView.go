package neopixeldisplay

import(
	"time"
)

type AnimationView struct {
	target *ColorFrame
	animationIndex int
}

func NewAnimationView(target *ColorFrame) *AnimationView {
	return &AnimationView{target, 0}
}

//cancels previous animation
func (av *AnimationView) PlayAnimation(frames []ColorFrame, fps float32, loop bool) chan int {
	av.animationIndex++
	animationIndex := av.animationIndex
	doneChan := make(chan int)
	go func() {
		loopCount := 0
		for loop || loopCount == 0 {
			for _, frame := range frames {
				av.target.SetRect(0, 0, frame, Error)
				av.target.Draw()
				time.Sleep(time.Duration(1000.0/fps)*time.Millisecond)
				if av.animationIndex != animationIndex {
					//someone has started a new animation and cancelled this one
					doneChan <- animationIndex
					return
				}
			}
		}
		doneChan <- animationIndex
	}()
	return doneChan
}
