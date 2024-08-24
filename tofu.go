package tofu

import (
	"fmt"
	"strings"

	"github.com/logrusorgru/aurora"
)

const (
	hideCursor = "\033[?25l" /* ANSI escape code to hide the cursor */
	showCursor = "\033[?25h" /* ANSI escape code to show the cursor */
	clearLine  = "\033[K"    /* ANSI escape code to clear the line */
	defaultFg  = "█"         /* Default foreground character for the progress bar */
	defaultBg  = "░"         /* Default background character for the progress bar */
)

type ProgressBar struct {
	Width          int      /* Width of the progress bar */
	Fg             string   /* Foreground character for the progress bar */
	Bg             string   /* Background character for the progress bar */
	ShowPercentage bool     /* Flag to show percentage on the progress bar */
	Theme          string   /* Theme for colors */
	AddGradient    bool     /* Flag to add gradient effect */
	Colors         []string /* Slice of colors for gradient */
}

// New returns a new ProgressBar with default settings
func New(width int, theme string, addGradient bool) ProgressBar {
	var colors []string
	switch theme {
	case "purpleHaze":
		colors = PurpleHaze
	case "pastelCore":
		colors = PastelCore
	case "limeWire":
		colors = LimeWire
	default:
		theme = "heatWave" // Default theme
		colors = HeatWave
	}

	return ProgressBar{
		Width:          width,
		Fg:             defaultFg,
		Bg:             defaultBg,
		ShowPercentage: true,
		Theme:          theme,
		AddGradient:    addGradient,
		Colors:         colors,
	}
}

// createGradient generates a gradient if AddGradient is true
func (p *ProgressBar) createGradient(filled int) string {
	if len(p.Colors) == 0 {
		return strings.Repeat(p.Fg, filled) /* No colors defined; use default character */
	}

	if !p.AddGradient {
		/* No gradient needed; use the first color in the theme */
		return strings.Repeat(p.Colors[0]+p.Fg+p.Colors[len(p.Colors)-1], filled)
	}

	colorCount := len(p.Colors) - 1 /* Last color is used for resetting color */
	segmentSize := p.Width / colorCount
	fgBar := ""

	/* Build the progress bar with gradient colors */
	for i := 0; i < filled; i++ {
		colorIndex := i / segmentSize
		if colorIndex >= colorCount {
			colorIndex = colorCount - 1 /* Ensure index is within bounds */
		}
		fgBar += p.Colors[colorIndex] + p.Fg + p.Colors[len(p.Colors)-1] /* Append color and reset color */
	}

	return fgBar
}

// ProgressBar returns the created progress bar string based on the percentage
func (p *ProgressBar) ProgressBar(percent float32) string {
	filled := int(float32(p.Width) * percent)
	unfilled := p.Width - filled
	fgBar := p.createGradient(filled)
	bgBar := strings.Repeat(p.Bg, unfilled)

	if p.ShowPercentage {
		return fmt.Sprintf("\r %s%s %d %s", fgBar, bgBar, int(percent*100), aurora.BrightWhite("%"))
	}
	return fmt.Sprintf("\r %s%s", fgBar, bgBar)
}

// PrintProgressBar prints the progress bar to stdout
func (p *ProgressBar) PrintProgressBar(percent float32) {
	/* Hide the terminal cursor */
	fmt.Printf(hideCursor)
	fmt.Printf("%s", p.ProgressBar(percent))
}

// CleanUp resets terminal default parameters
func CleanUp() {
	fmt.Printf(showCursor)
}
