package game

import (
	"fmt"
	"strings"

	"github.com/Bananenpro/tictactoe-cli/network"
)

type cellState int

const (
	cellEmpty  cellState = 0
	cellCross  cellState = 1
	cellCircle cellState = 2
)

func (s cellState) String() string {
	switch s {
	case cellEmpty:
		return " "
	case cellCross:
		return "X"
	case cellCircle:
		return "O"
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
		// TODO: stop game
	} else if cmd == "winner" {
		// TODO: display winner message and ask if player wants to repeat
	} else if cmd == "loser" {
		// TODO: display winner message and ask if player wants to repeat
	}
}
