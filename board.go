package arena

import (
	"errors"
	"fmt"
)

const NumRows = 6
const NumColumns = 7
const NumConsecutive = 4

// Checks fourup board state
// Author: Kevin Burke <kev@inburke.com>

// row varies, column does not.
func checkVerticalWin(column int, board [NumRows][NumColumns]int) bool {
	checkRowInColumn := func(column int, row int, board [NumRows][NumColumns]int) bool {
		initColor := board[row][column]
		for k := 0; k < NumConsecutive; k++ {
			if row+k >= NumRows {
				return false
			}
			value := board[row+k][column]
			if value == Empty || value != initColor {
				return false
			}
		}
		// if we get here and haven't broken, seen 4 in a row of the same color
		return true
	}

	for row := 0; row <= (NumRows - NumConsecutive); row++ {
		initColor := board[row][column]
		if initColor == Empty {
			continue
		}
		if checkRowInColumn(column, row, board) {
			return true
		}
	}
	return false
}

func checkHorizontalWin(row int, board [NumRows][NumColumns]int) bool {
	checkColumnInRow := func(row int, column int, board [NumRows][NumColumns]int) bool {
		initColor := board[row][column]
		for k := 0; k < NumConsecutive; k++ {
			if column+k >= NumColumns {
				return false
			}
			if board[row][column+k] != initColor {
				return false
			}
		}
		// if we get here and haven't broken, seen 4 in a row of the same color
		return true
	}
	for column := 0; column < NumConsecutive; column++ {
		initColor := board[row][column]
		if initColor == Empty {
			continue
		}
		if checkColumnInRow(row, column, board) {
			return true
		}
	}
	return false
}

// check squares down and to the right for a match
func checkSoutheastDiagonalWin(row int, column int, board [NumRows][NumColumns]int) bool {
	initColor := board[row][column]
	if initColor == Empty {
		return false
	}
	for i := 0; i < NumConsecutive; i++ {
		if board[row+i][column+i] != initColor {
			return false
		}
	}
	return true
}

func checkSouthwestDiagonalWin(row int, column int, board [NumRows][NumColumns]int) bool {
	initColor := board[row][column]
	if initColor == Empty {
		return false
	}
	for i := 0; i < NumConsecutive; i++ {
		if board[row+i][column-i] != initColor {
			return false
		}
	}
	return true
}

// Checks if a connect four exists
// I'm sure there's some more efficient way to conduct these checks, but at
// modern computer speeds, it really doesn't matter
func GameOver(board [NumRows][NumColumns]int) bool {
	for column := 0; column < NumColumns; column++ {
		if checkVerticalWin(column, board) {
			return true
		}
	}

	for row := 0; row < NumRows; row++ {
		if checkHorizontalWin(row, board) {
			return true
		}
	}
	for row := 0; row <= (NumRows - NumConsecutive); row++ {
		for column := 0; column <= (NumColumns - NumConsecutive); column++ {
			if checkSoutheastDiagonalWin(row, column, board) {
				return true
			}
		}
	}
	for column := (NumColumns - NumConsecutive); column < NumColumns; column++ {
		for row := 0; row <= (NumRows - NumConsecutive); row++ {
			if checkSouthwestDiagonalWin(row, column, board) {
				return true
			}
		}
	}
	return false
}

func GetStringBoard(board *[NumRows][NumColumns]int) [NumRows][NumColumns]string {
	var stringBoard [NumRows][NumColumns]string
	for row := 0; row < NumRows; row++ {
		for column := 0; column < NumColumns; column++ {
			if board[row][column] == Empty {
				stringBoard[row][column] = ""
			} else if board[row][column] == Red {
				stringBoard[row][column] = "R"
			} else if board[row][column] == Black {
				stringBoard[row][column] = "B"
			} else {
				panic(fmt.Sprintf("invalid value", board[row][column], "for a board"))
			}
		}
	}
	return stringBoard
}
func IsBoardFull(board [NumRows][NumColumns]int) bool {
	// will check the top row, which is always the last to fill up.
	for column := 0; column < NumColumns; column++ {
		if board[0][column] == Empty {
			return false
		}
	}
	return true
}

// Returns error if the move is invalid
func ApplyMoveToBoard(move int, playerId int, bp *[NumRows][NumColumns]int) (*[NumRows][NumColumns]int, error) {
	if move >= NumColumns || move < 0 {
		return bp, errors.New(fmt.Sprintf("Move %d is invalid", move))
	}
	for i := NumRows - 1; i >= 0; i-- {
		if bp[i][move] == 0 {
			bp[i][move] = playerId
			return bp, nil
		}
	}
	return bp, errors.New(fmt.Sprintf("No room in column %d for a move", move))
}

func InitializeBoard() *[NumRows][NumColumns]int {
	// Board is initialized to be filled with zeros.
	var board [NumRows][NumColumns]int
	return &board
}
