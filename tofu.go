package tofu

import (
	"errors"
	"fmt"
	"strings"

	"github.com/logrusorgru/aurora"
)

const (
	hideCursor = "\033[?25l" /* ansi escape code to hide the cursor */
	showCursor = "\033[?25h" /* ansi escape code to show the cursor */
	clearLine  = "\033[K"    /* ansi escape code to clear the line */
	defaultFg  = "█"         /* default foreground character for the progress bar */
	defaultBg  = "░"         /* default background character for the progress bar */
)

var (
	// define color themes
	PurpleHaze = []string{"\033[38;5;57m", "\033[38;5;93m", "\033[38;5;99m", "\033[0m"}
	PastelCore = []string{"\033[38;5;153m", "\033[38;5;159m", "\033[38;5;165m", "\033[0m"}
	LimeWire   = []string{"\033[38;5;118m", "\033[38;5;154m", "\033[38;5;190m", "\033[0m"}
	HeatWave   = []string{"\033[38;5;196m", "\033[38;5;202m", "\033[38;5;208m", "\033[0m"}
)

type ProgressBar struct {
	Width          int      /* width of the progress bar */
	Fg             string   /* foreground character for the progress bar */
	Bg             string   /* background character for the progress bar */
	ShowPercentage bool     /* flag to show percentage on the progress bar */
	Theme          string   /* theme for colors */
	AddGradient    bool     /* flag to add gradient effect */
	Colors         []string /* slice of colors for gradient */
}

// New creates a new progress bar with default settings or returns an error if the input is invalid
func New(width int, theme string, addGradient bool) (ProgressBar, error) {
	if width <= 0 {
		return ProgressBar{}, errors.New("width must be a positive integer")
	}

	var colors []string
	switch theme {
	case "purpleHaze":
		colors = PurpleHaze
	case "pastelCore":
		colors = PastelCore
	case "limeWire":
		colors = LimeWire
	default:
		if theme != "heatWave" {
			return ProgressBar{}, errors.New("invalid theme provided")
		}
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
	}, nil
}

// createGradient generates a gradient if AddGradient is true
func (p *ProgressBar) createGradient(filled int) string {
	if len(p.Colors) == 0 {
		return strings.Repeat(p.Fg, filled) /* no colors defined; use default character */
	}

	if !p.AddGradient {
		/* no gradient needed; use the first color in the theme */
		return strings.Repeat(p.Colors[0]+p.Fg+p.Colors[len(p.Colors)-1], filled)
	}

	colorCount := len(p.Colors) - 1 /* last color is used for resetting color */
	segmentSize := p.Width / colorCount
	fgBar := ""

	/* build the progress bar with gradient colors */
	for i := 0; i < filled; i++ {
		colorIndex := i / segmentSize
		if colorIndex >= colorCount {
			colorIndex = colorCount - 1 /* ensure index is within bounds */
		}
		fgBar += p.Colors[colorIndex] + p.Fg + p.Colors[len(p.Colors)-1] /* append color and reset color */
	}

	return fgBar
}

// ProgressBar returns the created progress bar string based on the percentage
func (p *ProgressBar) ProgressBar(percent float32) string {
	if percent < 0 || percent > 1 {
		return "" /* return empty string for invalid percentage values */
	}

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
	if percent < 0 || percent > 1 {
		return /* do nothing for invalid percentage values */
	}

	/* hide the terminal cursor */
	fmt.Printf(hideCursor)
	fmt.Printf("%s", p.ProgressBar(percent))
}

// CleanUp resets terminal default parameters
func CleanUp() {
	fmt.Printf(showCursor)
}
