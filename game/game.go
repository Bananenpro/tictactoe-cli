package game

import (
	"fmt"

	"github.com/Bananenpro/tictactoe-cli/network"
)

type cellState int

const (
	cellEmpty  cellState = 0
	cellCross  cellState = 1
	cellCircle cellState = 2
)

type board [9]cellState

func BoardFromString(str string) board {
	board := board{}

	if len(str) != 9 {
		fmt.Println("Error: invalid board string:", str)
		return board
	}

	for i := 0; i < 9; i++ {
		state := cellState(str[i])
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
}

func New(serverConnection *network.ServerConnection) *Game {
	return &Game{
		serverConnection: serverConnection,
	}
}

func (g *Game) Start() {
	ClearTerminal()
	PrintBoard(g.board)
}
