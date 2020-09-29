package board

import "testing"

func TestInitBoard(t *testing.T) {
	InitBoard()

	for i, val := range board {
		if val != "" {
			t.Errorf("Field %d is not empty", i)
		}
	}
}
