package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("========== TicTacToe ==========")
}

func PrintBoard(board board) {
	fmt.Printf("\n  1   2   3 \nA %s | %s | %s \n ---+---+---\nB %s | %s | %s \n ---+---+---\nC %s | %s | %s \n\n", board[0], board[1], board[2], board[3], board[4], board[5], board[6], board[7], board[8])
}

func (g *Game) InputFieldIndex() int {
	scanner := bufio.NewScanner(os.Stdin)
	index := -1

	for index < 0 || index > 8 || g.board[index] != cellEmpty {
		fmt.Print("Which field do you want to mark? ")
		if scanner.Scan() {
			parts := strings.Split(strings.ToLower(scanner.Text()), "")
			if len(parts) == 2 {
				row := -1
				column := -1
				for i := 0; i < 2; i++ {
					if strings.ContainsAny(parts[i], "123") {
						column, _ = strconv.Atoi(parts[i])
						column--
					} else if strings.ContainsAny(parts[i], "abc") {
						switch parts[i] {
						case "a":
							row = 0
						case "b":
							row = 1
						case "c":
							row = 2
						}
					}
				}
				if row != -1 && column != -1 {
					index = row*3 + column
				}
			}
		}
	}

	return index
}

func AskYesNo(question string) bool {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(question, " [y/n]: ")
		if scanner.Scan() {
			text := strings.ToLower(scanner.Text())
			if text == "y" || text == "yes" {
				return true
			}
			if text == "n" || text == "no" {
				return false
			}
		} else {
			break
		}
	}

	return false
}
