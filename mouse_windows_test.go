//go:build windows
// +build windows

package tea

import (
	"github.com/erikgeiser/coninput"
	"testing"
)

func TestMouseEventDoubleClick(t *testing.T) {
	// Previous state: no buttons pressed
	var prev coninput.ButtonState
	// Simulate a double click event at position (10,5) with left button
	e := coninput.MouseEventRecord{
		EventFlags:   coninput.DOUBLE_CLICK,
		ButtonState:  coninput.FROM_LEFT_1ST_BUTTON_PRESSED,
		MousePositon: coninput.Coord{X: 10, Y: 5},
	}

	got := mouseEvent(prev, e)
	want := MouseMsg{
		X:      10,
		Y:      5,
		Button: MouseButtonLeft,
		Action: MouseActionPress,
		Type:   MouseLeft,
	}

	if got != want {
		t.Fatalf("expected %#v, got %#v", want, got)
	}
}
