package menu

import (
	"testing"
)

func TestValidatePlayGameAnswer(t *testing.T) {
	validAnswer, play := validatePlayGameAnswer("Y")

	if !validAnswer {
		t.Errorf("Anwer 'Y' should be accepted")
	}

	if !play {
		t.Errorf("Anwer 'Y' should return true")
	}

	validAnswer, play = validatePlayGameAnswer("N")
	if !validAnswer {
		t.Errorf("Anwer 'N' should be accepted")
	}
	if play {
		t.Errorf("Anwer 'N' should return false")
	}

	validAnswer, play = validatePlayGameAnswer("12")
	if validAnswer {
		t.Errorf("Anything besides 'Y' or 'N' should not be accepted")
	}

	if play {
		t.Errorf("Anything should return false")
	}

}
