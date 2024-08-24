package tofu

import (
	"testing"
)

// Test for creating progress bar
func Test_ProgressBar(t *testing.T) {
	p, err := New(40, "limeWire", true)
	if err != nil {
		t.Errorf("Failed to create progress bar: %v", err)
	}
	total := 1000000
	for a := 0; a < total; a++ {
		percent := float32(a) / float32(total)
		p.ProgressBar(percent)
		p.PrintProgressBar(percent)
	}
}
