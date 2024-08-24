package tofu

import (
	"testing"
)

// Test for creating progress bar
func Test_ProgressBar(t *testing.T) {
	p := New(40, "limeWire", true)
	total := 1000000
	for a := 0; a < total; a++ {
		percent := float32(a) / float32(total)
		p.ProgressBar(percent)
		p.PrintProgressBar(percent)
	}
}
