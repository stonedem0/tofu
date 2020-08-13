package tofu

import (
	"fmt"
	"strings"

	"github.com/logrusorgru/aurora"
)

const (
	softPink   = 213
	purple     = 57
	hideCursor = "\033[?25l"
	showCursor = "\033[?25h"
)

//ProgressBar struct
type ProgressBar struct {
	Width          int
	Fg             string
	Bg             string
	Color          uint8
	ShowPercentage bool
}

// New returns a new ProgressBar
func New(width int) ProgressBar {
	p := ProgressBar{}
	return p
}

//ProgressBar returns created progess bar based on parameters
func (p *ProgressBar) ProgressBar(percent float32) string {
	if p.Width < 0 || p.Fg == "" || p.Bg == "" || p.Color > 0 {
		p.Width = 40
		p.Fg = "▇"
		p.Bg = "░"
		p.ShowPercentage = true
		p.Color = softPink
	}
	filled := int(float32(p.Width) * float32(percent))
	unfilled := p.Width - filled
	fgBar := aurora.Index(p.Color, strings.Repeat(p.Bg, unfilled)).String()
	bgBar := aurora.Index(p.Color, strings.Repeat(p.Fg, filled)).String()
	if p.ShowPercentage {
		return fmt.Sprintf("\r %s%s %d %s", bgBar, fgBar, aurora.Index(p.Color, int(percent*100)), aurora.Index(p.Color, "%"))
	}

	return fmt.Sprintf("\r %s%s", bgBar, fgBar)
}

//PrintProgressBar prints the result of ProgressBar function
func (p *ProgressBar) PrintProgressBar(percent float32) {

	//Hiding terminal cursor
	fmt.Printf(hideCursor)
	fmt.Printf("%s", p.ProgressBar(percent))
}

// CleanUp resets terminal default params
func CleanUp() {
	fmt.Printf(showCursor)
}

func main() {
	CleanUp()
}
