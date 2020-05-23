package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/logrusorgru/aurora"
)

const (
	softPink = 213
	purple   = 57
)

type ProgressBar struct {
	width int
	fg    string
	bg    string
	color uint8
}

func (p *ProgressBar) PrintProgressBarSeconds(seconds time.Duration) {
	total := int(seconds) * 4
	for a := 0; a <= total; a++ {
		percent := float32(a) / float32(total)
		filled := int(float32(p.width) * float32(percent))
		unfilled := p.width - filled
		fgBar := aurora.Index(p.color, strings.Repeat(p.bg, unfilled)).String()
		bgBar := aurora.Index(p.color, strings.Repeat(p.fg, filled)).String()
		fmt.Printf("\r %s%s %d %s", bgBar, fgBar, aurora.Index(p.color, int(percent*100)), aurora.Index(p.color, "%"))
		time.Sleep(time.Second / seconds)
	}
}
func (p *ProgressBar) PrintProgressBarArray(seconds time.Duration, ar []string) {
	// total := int(seconds) * 4
	for _, m := range ar {
		// fmt.Printf("\r %s", m)
		fmt.Printf("%s  ", m)
		// 	filled := int(float32(p.width) * float32(percent))
		// 	unfilled := p.width - filled
		// 	fgBar := aurora.Index(p.color, strings.Repeat(p.bg, unfilled)).String()
		// 	bgBar := aurora.Index(p.color, strings.Repeat(p.fg, filled)).String()
		// 	fmt.Printf("\r %s%s %d %s", bgBar, fgBar, aurora.Index(p.color, int(percent*100)), aurora.Index(p.color, "%"))
		time.Sleep(time.Second / seconds)
		// }
	}
}
func main() {
	// p := ProgressBar{40, "▇", "░", softPink}
	moon := []string{"🌑", "🌒", "🌓", "🌔", "🌕", "🌖", "🌗", "🌘", "🌑"}
	p := ProgressBar{10, "🌝 ", "🌚 ", softPink}
	// p.PrintProgressBarSeconds(10)
	p.PrintProgressBarArray(5, moon)

}
