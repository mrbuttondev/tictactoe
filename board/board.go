package board

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var board [9]string
var sX string = "XXX"
var sO string = "OOO"

// MakeAMove is used to pick the place on the board
func MakeAMove(player string) bool {
	place := pickPlace()
	board[place] = player
	displayBoard()

	isMatch := validateBoard()
	return isMatch
}

// InitBoard is used to initialize empty board
func InitBoard() {
	board = [9]string{}
}

func pickPlace() int {
	isValid := false
	var place string
	for !isValid {
		fmt.Println("Pick place from 1 to 9: ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		place = input.Text()
		isValid = validatePlace(place)
	}
	placement, _ := strconv.Atoi(place)
	return placement - 1
}

func validateBoard() bool {

	cb := currentBoard()

	for _, val := range cb {
		if val == sX || val == sO {
			return true
		}
	}
	return false
}

func validatePlace(place string) bool {
	p, err := strconv.Atoi(place)
	if err != nil {
		fmt.Println("Not a number. Please try again.")
		return false
	}

	if p < 1 || p > 9 {
		fmt.Println("Number is out of range. Please try again.")
		return false
	}

	if board[p-1] != "" {
		fmt.Println("Place is taken. Please try again.")
		return false
	}
	return true
}

func displayBoard() {
	delimeterH := "-"
	delimeterV := "|"
	fmt.Printf("\n%1s%1s%1s%1s%1s", board[0], delimeterV, board[1], delimeterV, board[2])
	fmt.Printf("\n%1s%1s%1s%1s%1s", delimeterH, delimeterH, delimeterH, delimeterH, delimeterH)
	fmt.Printf("\n%1s%1s%1s%1s%1s", board[3], delimeterV, board[4], delimeterV, board[5])
	fmt.Printf("\n%1s%1s%1s%1s%1s", delimeterH, delimeterH, delimeterH, delimeterH, delimeterH)
	fmt.Printf("\n%1s%1s%1s%1s%1s", board[6], delimeterV, board[7], delimeterV, board[8])
	fmt.Printf("\n")
}

func currentBoard() [8]string {
	h1 := board[0] + board[1] + board[2]
	h2 := board[3] + board[4] + board[5]
	h3 := board[6] + board[7] + board[8]
	v1 := board[0] + board[3] + board[6]
	v2 := board[1] + board[4] + board[7]
	v3 := board[2] + board[5] + board[8]
	x1 := board[0] + board[4] + board[8]
	x2 := board[2] + board[4] + board[6]

	return [8]string{h1, h2, h3, v1, v2, v3, x1, x2}
}
