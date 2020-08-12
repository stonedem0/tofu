package tofu

import (
	"testing"
)

//Test for creating progress bar
func Test_ProgressBar(t *testing.T) {
	p := ProgressBar{}
	total := 100
	for a := 0; a < total; a++ {
		p.ProgressBar(float32(a)/float32(total), 40, softPink, "▇", "░")

	}
}
