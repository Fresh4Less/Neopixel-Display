package neopixeldisplay

import(
	"time"
)

type FrameTransition int
const(
	None = iota
	Slide
)

type TransitionView struct {
	target *ColorFrame
	currentFrame int
	transitioning bool
	transitionIndex int
	views []*TransitionData
}

type TransitionData struct {
	Frame *ColorFrame
	Duration time.Duration
	Transition FrameTransition
	x, y int
}

func NewTransitionView(target *ColorFrame) *TransitionView {
	return &TransitionView{target, 0, false, 0, nil}
}

func (tv *TransitionView) AddTransition(duration time.Duration, transition FrameTransition) *TransitionData {
	frame := MakeColorFrame(tv.target.Width, tv.target.Height, MakeColor(0,0,0))
	frame.parent = tv
	data := &TransitionData{&frame, duration, transition, 0, 0}
	tv.views = append(tv.views, data)

	if len(tv.views) == 1 {
		tv.TransitionTo(0, None, true)
	}
	return data
}

func (tv *TransitionView) GetTransition(index int) *TransitionData {
	return tv.views[index]
}

//if cycle=true, continues cycling from this frame
func (tv *TransitionView) TransitionTo(frameIndex int, transition FrameTransition, cycle bool) {
	tv.transitionIndex++
	transitionIndex := tv.transitionIndex

	beginIndex := (((frameIndex-1)%len(tv.views))+ len(tv.views)) % len(tv.views) //double modulus prevents negative index

	tv.transitioning = true
	// async transition
	go func() {
		switch transition {
		case None:
			tv.views[frameIndex].x = 0
			tv.views[frameIndex].y = 0
			tv.Draw()
		case Slide:
			//for now just transition left to right
			tv.views[frameIndex].x = tv.target.Width
			for i := 0; i < tv.target.Width; i++ {
				tv.views[beginIndex].x--
				tv.views[frameIndex].x--
				tv.Draw()

				//TODO: don't hardcode transition duration
				time.Sleep(time.Duration(500/8)*time.Millisecond)
				if tv.transitionIndex != transitionIndex {
					//someone else started a transition while we were sleeping--cancel the rest of this transition
					return
				}
			}
		}
		//transition done, sleep until next cycle
		tv.transitioning = false
		tv.currentFrame = frameIndex
		if cycle {
			time.Sleep(tv.views[frameIndex].Duration)
			if tv.transitionIndex == transitionIndex {
				nextFrameIndex := (frameIndex+1)%len(tv.views)
				tv.TransitionTo(nextFrameIndex, tv.views[nextFrameIndex].Transition, cycle)
			}
		}
	}()
}

func (tv *TransitionView) Draw() {
	tv.target.SetAll(MakeColor(0,0,0))

	frame := tv.views[tv.currentFrame]
	tv.target.SetRect(frame.x, frame.y, *frame.Frame, Clip)
	if tv.transitioning {
		//draw next frame
		nextFrame := tv.views[(tv.currentFrame + 1) % len(tv.views)]
		tv.target.SetRect(nextFrame.x, nextFrame.y, *nextFrame.Frame, Clip)
	}

	tv.target.Draw()
}
