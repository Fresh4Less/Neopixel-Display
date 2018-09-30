package main

import (
	"flag"
	"github.com/fresh4less/neopixel-display/neopixeldisplay"
	"time"
)

//tester is a main function because creating tests for animations takes more time than I'm willing to invest for this tiny library
func main() {
	var display neopixeldisplay.PixelDisplay
	mode := flag.String("m", "console", "Mode (neopixel or console)")
	flag.Parse()
	if *mode == "console" {
		display = neopixeldisplay.NewConsoleColorDisplay(64, [][]int{[]int{8, 8}})
	} else {
		display = neopixeldisplay.NewNeopixelDisplay(18, 64, 255)
	}
	screen := neopixeldisplay.NewScreenView(display, 0, 8, 8, 1.0, neopixeldisplay.Error)

	//frame := screen.GetFrame()
	//frame.SetRect(0,0,neopixeldisplay.MakeColorFrame(8,8, neopixeldisplay.MakeColor(255,0,0)), neopixeldisplay.Error)
	//frame.SetRect(1,1,neopixeldisplay.MakeColorFrame(7,7, neopixeldisplay.MakeColor(0,255,0)), neopixeldisplay.Error)
	//frame.Set(2,2, neopixeldisplay.MakeColor(255,255,255), neopixeldisplay.Error)
	//frame.Set(3,3, neopixeldisplay.MakeColor(0,0,0), neopixeldisplay.Error)
	//frame.Set(4,4, neopixeldisplay.MakeColor(0,0,1), neopixeldisplay.Error)
	//frame.Set(5,5, neopixeldisplay.MakeColor(0,0,254), neopixeldisplay.Error)
	//frame.Set(6,6, neopixeldisplay.MakeColor(0,0,255), neopixeldisplay.Error)
	//screen.Draw()

	transView := neopixeldisplay.NewTransitionView(screen.GetFrame())
	frame1 := transView.AddTransition(time.Second, neopixeldisplay.Slide).Frame
	frame1.SetRect(0, 0, neopixeldisplay.MakeColorFrame(8, 8, neopixeldisplay.MakeColor(255, 0, 0)), neopixeldisplay.Error)
	frame2 := transView.AddTransition(time.Second, neopixeldisplay.Slide).Frame
	frame2.SetRect(0, 0, neopixeldisplay.MakeColorFrame(8, 8, neopixeldisplay.MakeColor(0, 255, 0)), neopixeldisplay.Error)
	frame3 := transView.AddTransition(time.Second*3, neopixeldisplay.Slide).Frame
	frame3.SetRect(0, 0, neopixeldisplay.MakeColorFrame(8, 8, neopixeldisplay.MakeColor(0, 0, 255)), neopixeldisplay.Error)
	transView.TransitionTo(0, neopixeldisplay.None, true)
	//time.Sleep(time.Millisecond * 1750)
	//frame2.SetRect(0,0, neopixeldisplay.MakeColorNumberChar2x3(8, neopixeldisplay.MakeColor(255,255,255), neopixeldisplay.MakeColor(0,100,100)), neopixeldisplay.Error)
	//frame2.Draw()

	layerView := neopixeldisplay.NewLayerView(frame2)
	layer1 := layerView.AddLayer(neopixeldisplay.Add)
	layer1.Frame.SetRect(0, 0, neopixeldisplay.MakeColorFrame(7, 7, neopixeldisplay.MakeColor(255, 0, 0)), neopixeldisplay.Error)
	layer2 := layerView.AddLayer(neopixeldisplay.Add)
	layer2.Frame.SetRect(2, 2, neopixeldisplay.MakeColorFrame(5, 5, neopixeldisplay.MakeColor(0, 255, 0)), neopixeldisplay.Error)
	layer3 := layerView.AddLayer(neopixeldisplay.Overwrite)
	layer3.Frame.SetRect(4, 4, neopixeldisplay.MakeColorFrame(3, 3, neopixeldisplay.MakeColor(0, 0, 255)), neopixeldisplay.Error)
	layer4 := layerView.AddLayer(neopixeldisplay.SetWhite)
	layer4.Frame.SetRect(6, 6, neopixeldisplay.MakeColorFrame(2, 2, neopixeldisplay.MakeColor(100, 100, 100)), neopixeldisplay.Error)
	layerView.Draw()

	animationView := neopixeldisplay.NewAnimationView(frame3)
	animationView.PlayAnimation(
		[]neopixeldisplay.ColorFrame{
			neopixeldisplay.MakeColorFrame(8, 8, neopixeldisplay.MakeColorHue(0)),
			neopixeldisplay.MakeColorFrame(8, 8, neopixeldisplay.MakeColorHue(30)),
			neopixeldisplay.MakeColorFrame(8, 8, neopixeldisplay.MakeColorHue(60)),
			neopixeldisplay.MakeColorFrame(8, 8, neopixeldisplay.MakeColorHue(90)),
			neopixeldisplay.MakeColorFrame(8, 8, neopixeldisplay.MakeColorHue(120)),
			neopixeldisplay.MakeColorFrame(8, 8, neopixeldisplay.MakeColorHue(150)),
			neopixeldisplay.MakeColorFrame(8, 8, neopixeldisplay.MakeColorHue(180)),
			neopixeldisplay.MakeColorFrame(8, 8, neopixeldisplay.MakeColorHue(210)),
			neopixeldisplay.MakeColorFrame(8, 8, neopixeldisplay.MakeColorHue(240)),
		},
		10, true)
	time.Sleep(time.Millisecond * 1750)
	layerView.DeleteLayer(layer2)

	select {}
}
