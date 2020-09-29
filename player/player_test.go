package player

import "testing"

// func TestChoosePlayers(t *testing.T) {
// 	ChoosePlayers()
// }

func TestValidatePlayer(t *testing.T) {
	isValid := validatePlayer("x")

	if !isValid {
		t.Errorf("X should be a valid sign")
	}

	isValid = validatePlayer("o")
	if !isValid {
		t.Errorf("O should be a valid sign")
	}

	isValid = validatePlayer("12")
	if isValid {
		t.Errorf("Anything besides X or O is not valid")
	}
}

func TestAssignPlayers(t *testing.T) {
	pX := Player{
		First:  "X",
		Second: "O",
	}

	pXR := assignPlayers("X")
	if pX.First != pXR.First && pX.Second != pXR.Second {
		t.Errorf("Players X, O wrongly assigned")
	}

	pO := Player{
		First:  "O",
		Second: "X",
	}
	pOR := assignPlayers("O")

	if pO.First != pOR.First && pO.Second != pOR.Second {
		t.Errorf("Players O, X wrongly assigned")
	}
}
