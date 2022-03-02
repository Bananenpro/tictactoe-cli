package game

import (
	"errors"
)

// returns 1 if cross wins, 0 if tie, -1 if circle wins and an error if the game is not over
func (b *board) getScore() (int, error) {
	for i := 0; i < 3; i++ {
		// top to bottom
		if b[0+i] != cellEmpty && b[0+i] == b[3+i] && b[0+i] == b[6+i] {
			if b[0+i] == cellCross {
				return 1, nil
			} else {
				return -1, nil
			}
		}

		// left to right
		if b[i*3+0] != cellEmpty && b[i*3+0] == b[i*3+1] && b[i*3+0] == b[i*3+2] {
			if b[i*3+0] == cellCross {
				return 1, nil
			} else {
				return -1, nil
			}
		}
	}

	// top left to bottom right
	if b[0] != cellEmpty && b[0] == b[4] && b[0] == b[8] {
		if b[0] == cellCross {
			return 1, nil
		} else {
			return -1, nil
		}
	}

	// top right to bottom left
	if b[2] != cellEmpty && b[2] == b[4] && b[2] == b[6] {
		if b[2] == cellCross {
			return 1, nil
		} else {
			return -1, nil
		}
	}

	// tie
	tie := true
	for _, cell := range b {
		if cell == cellEmpty {
			tie = false
			break
		}
	}
	if tie {
		return 0, nil
	}

	return 0, errors.New("Game is not over")
}

type move struct {
	index int
	score int
}

func minimax(board board, currentSign cellState) move {
	score, err := board.getScore()
	if err == nil {
		return move{
			index: -1,
			score: score,
		}
	}

	moves := make([]move, 0, len(board))

	for i := range board {
		if board[i] == cellEmpty {
			board[i] = currentSign

			if currentSign == cellCross {
				score = minimax(board, cellCircle).score
			} else {
				score = minimax(board, cellCross).score
			}

			moves = append(moves, move{
				index: i,
				score: score,
			})

			board[i] = cellEmpty
		}
	}

	var bestMove move

	if currentSign == cellCross {
		bestMove.score = -2
		for _, m := range moves {
			if m.score > bestMove.score {
				bestMove = m
			}
		}
	} else {
		bestMove.score = 2
		for _, m := range moves {
			if m.score < bestMove.score {
				bestMove = m
			}
		}
	}

	return bestMove
}

func (g *Game) AIGetNextFieldIndex() int {
	return minimax(g.board, g.sign).index
}
