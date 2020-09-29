package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PlayGameQuestion is used to ask a user about play willingness
func PlayGameQuestion(txt string) bool {
	var answer string = ""
	// isPlayGame := false
	input := bufio.NewScanner(os.Stdin)

	for counter := 0; counter < 5; counter++ {
		fmt.Print(txt, " (Y/N) ")
		input.Scan()
		answer = input.Text()

		validAnswer, play := validatePlayGameAnswer(answer)

		if validAnswer {
			return play
		}
	}
	fmt.Println("Max number of attemps reached!")
	return true
}

func validatePlayGameAnswer(answer string) (bool, bool) {
	validAnswer := strings.ToUpper(answer) == "Y" || strings.ToUpper(answer) == "N"
	play := false
	if !validAnswer {
		fmt.Println("Answer is incorrect. Try once again")
	}

	if strings.ToUpper(answer) == "Y" {
		play = true
	}
	return validAnswer, play
}
