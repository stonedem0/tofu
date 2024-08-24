package tofu

import (
	"errors"
	"fmt"
	"strings"
)

const (
	hideCursor = "\033[?25l" /* ansi escape code to hide the cursor */
	showCursor = "\033[?25h" /* ansi escape code to show the cursor */
	clearLine  = "\033[K"    /* ansi escape code to clear the line */
	defaultFg  = "█"         /* default foreground character for the progress bar */
	defaultBg  = "░"         /* default background character for the progress bar */
)

const (
	PurpleHaze = "purpleHaze"
	PastelCore = "pastelCore"
	LimeWire   = "limeWire"
	HeatWave   = "heatWave"
)

var (
	// define color themes
	purpleHaze = []string{"\033[38;5;57m", "\033[38;5;93m", "\033[38;5;99m", "\033[0m"}
	pastelCore = []string{"\033[38;5;225m", "\033[38;5;189m", "\033[38;5;153m", "\033[38;5;117m", "\033[38;5;81m", "\033[38;5;45m", "\033[0m"}
	limeWire   = []string{"\033[38;5;226m", "\033[38;5;190m", "\033[38;5;154m", "\033[38;5;118m", "\033[38;5;82m", "\033[38;5;46m", "\033[0m"}
	heatWave   = []string{"\033[38;5;196m", "\033[38;5;202m", "\033[38;5;208m", "\033[38;5;214m", "\033[38;5;220m", "\033[38;5;226m", "\033[0m"}
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
	case PurpleHaze:
		colors = purpleHaze
	case PastelCore:
		colors = pastelCore
	case LimeWire:
		colors = limeWire
	default:
		if theme != HeatWave {
			return ProgressBar{}, errors.New("invalid theme provided")
		}
		colors = heatWave
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

	// ANSI escape code for bright white text
	percentText := fmt.Sprintf("\033[97m%d%%\033[0m", int(percent*100))

	if p.ShowPercentage {
		return fmt.Sprintf("\r %s%s %s", fgBar, bgBar, percentText)
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
