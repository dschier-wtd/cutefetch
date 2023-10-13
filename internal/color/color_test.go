package color

import "testing"

func TestColorVars(t *testing.T) {
	colors := []string{Reset, Red, Green, Yellow, Blue, Purple, Cyan, Gray, White}
	for _, color := range colors {
		if color == "" {
			t.Errorf("Expected color, but got empty string")
		}
	}
}
