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

type ProgressBar struct {
	width   int
	fg      string
	bg      string
	color   uint8
	percent float32
}

//ProgressBar returns created progess bar based on parameters
func (p *ProgressBar) ProgressBar(percent float32, width int, color uint8, fg string, bg string) string {

	//Creating progress bar based on arguments:
	p.width = width
	p.fg = fg
	p.bg = bg
	p.color = color
	p.percent = percent

	//Calculting filled/unfilled space
	filled := int(float32(p.width) * float32(percent))
	unfilled := p.width - filled
	fgBar := aurora.Index(p.color, strings.Repeat(p.bg, unfilled)).String()
	bgBar := aurora.Index(p.color, strings.Repeat(p.fg, filled)).String()
	return fmt.Sprintf("\r %s%s %d %s", bgBar, fgBar, aurora.Index(p.color, int(percent*100)), aurora.Index(p.color, "%"))
}

//PrintProgressBar prints the result of ProgressBar function
func (p *ProgressBar) PrintProgressBar() {

	//Hiding terminal cursor
	fmt.Printf(hideCursor)

	fmt.Printf("%s", p.ProgressBar(p.percent, p.width, p.color, p.fg, p.bg))
	//Showing terminal cursor
}

// CleanUp resets terminal default params
func (p *ProgressBar) CleanUp() {
	fmt.Printf(showCursor)
}

//TODO PrintLoader, Loader

//PrintLoader ...
func (p *ProgressBar) PrintLoader(percent float32, ar []string, total int) {
	for i := 0; i <= total; i++ {
		for _, m := range ar {
			fmt.Printf("\r %s", m)
			// time.Sleep(time.Second / percent)
		}
	}
}
func main() {
	p := ProgressBar{}
	// moon := []string{"🌑", "🌒", "🌓", "🌔", "🌕", "🌖", "🌗", "🌘", "🌑"}
	// // clock := []string{"🕛", "🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚", "🕛"}
	// p.PrintLoader(1, moon, 5)
	// total := 100
	// for i := 0; i <= total; i++ {
	// 	// p.PrintLoader(float32(i)/float32(total), moon, total)
	// 	p.ProgressBar(float32(i)/float32(total), 40, softPink, "▇", "░")
	// 	p.PrintProgressBar()
	// }
	p.CleanUp()
}
