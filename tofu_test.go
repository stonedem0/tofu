package tofu

import (
	"testing"
	"time"
)

func Test_ProgressBar(t *testing.T) {
	p, err := New(40, BubbleGum, true, true)
	if err != nil {
		t.Fatalf("failed to create progress bar: %v", err)
	}
	total := 200
	for a := 0; a <= total; a++ {
		percent := float32(a) / float32(total)
		p.ProgressBar(percent)
		p.PrintProgressBar(percent)
		time.Sleep(10 * time.Millisecond)
	}
	p.Finalize()
}
