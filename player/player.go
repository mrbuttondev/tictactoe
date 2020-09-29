package player

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Player struct
type Player struct {
	First  string
	Second string
}

// ChoosePlayers is used to assign player to X or O
func ChoosePlayers() (bool, Player) {
	var players Player
	var sign string
	input := bufio.NewScanner(os.Stdin)

	for counter := 0; counter < 5; counter++ {
		fmt.Print("Chose sign for player 1: X or O: ")
		input.Scan()
		sign = input.Text()
		if validatePlayer(sign) {
			players = assignPlayers(sign)
			return true, players
		}
	}
	fmt.Println("Max number of attemps reached!")
	return false, Player{}
}

func validatePlayer(player string) bool {
	isCorrect := strings.ToUpper(player) == "X" || strings.ToUpper(player) == "O"
	if !isCorrect {
		fmt.Println("Incorrect sign. Try once again")
	}
	return isCorrect
}

func assignPlayers(player string) Player {
	if strings.ToUpper(player) == "X" {
		return Player{"X", "O"}
	}

	return Player{"O", "X"}
}
