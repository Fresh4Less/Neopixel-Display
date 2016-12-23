package neopixeldisplay

import(
	"fmt"
	"errors"
	"github.com/gilliek/go-xterm256/xterm256"
)

type ConsoleColorDisplay struct {
	count int
	rects [][]int
	colors []Color
}

//this is a debugging tool meant for drawing without physical pixels.
//takes an array of (width,height) pairs which specify the dimensions of ScreenViews that will be used
func NewConsoleColorDisplay(count int, rects [][]int) *ConsoleColorDisplay {
	return &ConsoleColorDisplay{count, rects, make([]Color, count)}
}

func (ccd *ConsoleColorDisplay) Set(index int, color Color) {
	ccd.colors[index] = color
}

func (ccd *ConsoleColorDisplay) Show() {
	offset := 0
	fmt.Print("---\n")
	for _, rect := range ccd.rects {
		for y := 0; y < rect[1]; y++ {
			for x := 0; x < rect[0]; x++ {
				color := ccd.colors[offset + y*rect[1] + x]
				consoleColor, _ := xterm256.NewColor(getConsoleColor(color.GetRed()), getConsoleColor(color.GetGreen()),getConsoleColor(color.GetBlue()))
				_ , err := xterm256.Print(consoleColor, "\u2022")
				if err != nil {
					panic(errors.New(fmt.Sprintf("neopixeldisplay: ConsoleColorDisplay: Show: Printf error: %v", err)))
				}
			}
			fmt.Print("\n")
		}
		fmt.Print("\n")

		offset += rect[0]*rect[1]
	}
	fmt.Print("---\n")
}

func (ccd *ConsoleColorDisplay) Count() int {
	return ccd.count
}

//maps 0-255 to 0-5. 0->0, 255->5, otherwise, maps evenly from 1-4
func getConsoleColor(color uint32) int {
	if color == 0 {
		return 0
	}
	if color == 255 {
		return 5
	}
	return (int(color)*4)/255 + 1
}
