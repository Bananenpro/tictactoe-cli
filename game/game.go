package game

import (
	"fmt"
	"strings"

	"github.com/Bananenpro/tictactoe-cli/network"
)

type cellState int

const (
	cellEmpty      cellState = 0
	cellCross      cellState = 1
	cellCircle     cellState = 2
	cellCrossWin   cellState = 3
	cellCircleWin  cellState = 4
	cellCrossLose  cellState = 5
	cellCircleLose cellState = 6
	cellCrossTie   cellState = 7
	cellCircleTie  cellState = 8
)

func (s cellState) String() string {
	switch s {
	case cellEmpty:
		return " "
	case cellCross:
		return "X"
	case cellCircle:
		return "O"
	case cellCrossWin:
		return "\033[0;32mX\033[0m"
	case cellCircleWin:
		return "\033[0;32mO\033[0m"
	case cellCrossLose:
		return "\033[0;31mX\033[0m"
	case cellCircleLose:
		return "\033[0;31mO\033[0m"
	case cellCrossTie:
		return "\033[0;34mX\033[0m"
	case cellCircleTie:
		return "\033[0;34mO\033[0m"
	default:
		return "?"
	}
}

type board [9]cellState

func BoardFromString(str string) board {
	board := board{}

	if len(str) != 9 {
		fmt.Println("Error: invalid board string:", str)
		return board
	}

	for i, r := range str {
		state := cellState(r - '0')
		if state < 0 || state > 2 {
			fmt.Printf("Error: invalid cell state for cell %d: %d\n", i, state)
		} else {
			board[i] = state
		}
	}

	return board
}

type Game struct {
	board            board
	serverConnection *network.ServerConnection
	running          bool
}

func New(serverConnection *network.ServerConnection) *Game {
	return &Game{
		serverConnection: serverConnection,
	}
}

func (g *Game) Start() {
	ClearTerminal()
	PrintBoard(g.board)
	g.running = true
}

func (g *Game) Stop() {
	g.running = false
	g.serverConnection.Close()
}

func (g *Game) IsRunning() bool {
	return g.running
}

func (g *Game) HandleCommand(cmd string) {
	if cmd == "your-turn" {
		g.serverConnection.ClickField(g.InputFieldIndex())
	} else if cmd == "their-turn" {
		fmt.Println("Waiting for opponent...")
	} else if strings.HasPrefix(cmd, "board:") {
		parts := strings.Split(cmd, ":")
		if len(parts) != 2 {
			fmt.Println("Invalid command:", cmd)
			return
		}
		g.board = BoardFromString(parts[1])
		ClearTerminal()
		PrintBoard(g.board)
	} else if cmd == "opponent-disconnected" {
		ClearTerminal()
		PrintBoard(g.board)
		fmt.Println("Your opponent disconnected!")
		g.Stop()
	} else if strings.HasPrefix(cmd, "winner:") || strings.HasPrefix(cmd, "loser:") {
		parts := strings.Split(cmd, ":")
		if len(parts) != 2 {
			fmt.Println("Invalid command:", cmd)
			return
		}

		winner := strings.HasPrefix(cmd, "winner:")

		ClearTerminal()
		for _, r := range parts[1] {
			index := int(r - '0')
			if index < 0 || index > 8 {
				fmt.Println("Error: invalid cell index:", index)
			} else {
				if winner {
					if g.board[index] == cellCross {
						g.board[index] = cellCrossWin
					} else {
						g.board[index] = cellCircleWin
					}
				} else {
					if g.board[index] == cellCross {
						g.board[index] = cellCrossLose
					} else {
						g.board[index] = cellCircleLose
					}
				}
			}
		}
		PrintBoard(g.board)
		if winner {
			fmt.Println("You win!")
		} else {
			fmt.Println("You lose!")
		}
		g.askAgain()
	} else if cmd == "tie" {
		ClearTerminal()
		for i, c := range g.board {
			if c == cellCross {
				g.board[i] = cellCrossTie
			} else {
				g.board[i] = cellCircleTie
			}
		}
		PrintBoard(g.board)
		fmt.Println("Tie!")
		g.askAgain()
	}
}

func (g *Game) askAgain() {
	fmt.Println()
	if AskYesNo("Would you like to play again?") {
		g.serverConnection.Send("again")
		fmt.Println("Waiting for decision of opponent...")
	} else {
		g.serverConnection.Close()
	}
}
