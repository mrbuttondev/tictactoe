package main

import (
	"fmt"
	"game/board"
	"game/menu"
	"game/player"
	"os"
)

func main() {
	isPlay := menu.PlayGameQuestion("Would you like to play?")

	for isPlay {
		isChoosen, players := player.ChoosePlayers()
		if !isChoosen {
			fmt.Println("Exit")
			os.Exit(0)
		}
		board.InitBoard()
		for i := 0; i < 9; i++ {
			var player string
			if i%2 == 0 {
				player = players.First
			} else {
				player = players.Second
			}
			isMatch := board.MakeAMove(player)

			if i >= 4 && isMatch == true {
				fmt.Println("Player ", player, " win!")
				break
			}

			if i == 8 {
				fmt.Println("No one won. Draw!")
			}
		}

		isPlay = menu.PlayGameQuestion("Would you like to play again?")
		if !isPlay {
			break
		}
	}
	fmt.Println("Bye bye!")
}
