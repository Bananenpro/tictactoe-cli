package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("========== TicTacToe ==========")
}

func PrintBoard(board board) {
	fmt.Printf("\n %s | %s | %s \n---+---+---\n %s | %s | %s \n---+---+---\n %s | %s | %s \n\n", board[0], board[1], board[2], board[3], board[4], board[5], board[6], board[7], board[8])
}

func (g *Game) InputFieldIndex() int {
	scanner := bufio.NewScanner(os.Stdin)
	index := -1

	for index < 0 || index > 8 || g.board[index] != cellEmpty {
		fmt.Print("Which field do you want to mark? ")
		if scanner.Scan() {
			i, err := strconv.Atoi(scanner.Text())
			if err == nil {
				index = i
			}
		}
	}

	return index
}
