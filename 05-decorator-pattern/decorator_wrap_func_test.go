package decorator

import "testing"

func Double(n int) int {
	return n * 2
}

func TestLogDecorate(t *testing.T) {
	f := LogDecorate(Double)
	n := f(5)
	if n != 10 {
		t.Error("decorate func error")
	}
}
