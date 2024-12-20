package day4

import (
	"strconv"
	"strings"
)

type Number struct {
	value   string
	checked bool
}

type Board struct {
	boardId int
	numbers [][]Number
}

func ComputeFinalScore(input string) int {
	drawnNumbers, boards := parseInput(input)
	winningBoard, lastDrawnNumber := playGame(drawnNumbers, boards)
	return winningBoard.sumUnmarkedNumbers() * lastDrawnNumber
}

func ComputeFinalScoreLetSquidWin(input string) int {
	drawnNumbers, boards := parseInput(input)
	winningBoard, lastDrawnNumber := playGameLetSquidWin(drawnNumbers, boards)
	return winningBoard.sumUnmarkedNumbers() * lastDrawnNumber
}

func playGame(drawnNumbers []string, boards []Board) (*Board, int) {
	for _, drawnNumber := range drawnNumbers {
		for _, board := range boards {
			isUpdated := board.updateBoard(drawnNumber)
			if isUpdated {
				if board.isWinning() {
					drawnNumberValue, _ := strconv.Atoi(drawnNumber)
					return &board, drawnNumberValue
				}
			}
		}
	}
	return nil, -1
}

func playGameLetSquidWin(drawnNumbers []string, boards []Board) (*Board, int) {
	winningBoards := make(map[int]bool)
	for _, drawnNumber := range drawnNumbers {
		for _, board := range boards {
			if _, ok := winningBoards[board.boardId]; ok {
				continue
			}
			isUpdated := board.updateBoard(drawnNumber)
			if isUpdated {
				if board.isWinning() {
					winningBoards[board.boardId] = true
					if len(winningBoards) == len(boards) {
						drawnNumberValue, _ := strconv.Atoi(drawnNumber)
						return &board, drawnNumberValue
					}
				}
			}
		}
	}
	return nil, -1
}

func (board Board) updateBoard(drawnNumber string) bool {
	for y, row := range board.numbers {
		for x, number := range row {
			if number.value == drawnNumber {
				board.numbers[y][x].checked = true
				return true
			}
		}
	}
	return false
}

func (board Board) sumUnmarkedNumbers() int {
	sum := 0
	for _, rows := range board.numbers {
		for _, number := range rows {
			if !number.checked {
				numberValue, _ := strconv.Atoi(number.value)
				sum += numberValue
			}
		}
	}
	return sum
}

func (board Board) isWinning() bool {
	winningRow := board.hasWinningRow()
	winningColumn := board.hasWinningColumn()
	return winningRow || winningColumn
}

func (board Board) hasWinningRow() bool {
	for y := 0; y < len(board.numbers); y++ {
		isWinningRow := true
		for x := 0; x < len(board.numbers[y]); x++ {
			if !board.numbers[y][x].checked {
				isWinningRow = false
				break
			}
		}
		if isWinningRow {
			return true
		}
	}
	return false
}

func (board Board) hasWinningColumn() bool {
	for x := 0; x < len(board.numbers[0]); x++ {
		isWinningColumn := true
		for y := 0; y < len(board.numbers); y++ {
			if !board.numbers[y][x].checked {
				isWinningColumn = false
				break
			}
		}
		if isWinningColumn {
			return true
		}
	}
	return false
}

func parseInput(input string) ([]string, []Board) {
	lines := strings.Split(input, "\n")
	drawnNumers := strings.Split(lines[0], ",")
	var boards []Board
	boardId := 0
	for i := 2; i < len(lines); i += 6 {
		board := parseBoard(lines[i:i+5], boardId)
		boards = append(boards, board)
		boardId += 1
	}
	return drawnNumers, boards
}

func parseBoard(boardNumbersLines []string, boardId int) Board {
	boardNumbers := make([][]Number, len(boardNumbersLines))
	for idx, line := range boardNumbersLines {
		numbers := strings.Split(line, " ")
		var boardNumberRow []Number
		for _, number := range numbers {
			if number == "" {
				continue
			}
			boardNumberRow = append(boardNumberRow, Number{number, false})
		}
		boardNumbers[idx] = boardNumberRow
	}
	return Board{boardId, boardNumbers}
}
