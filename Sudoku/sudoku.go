package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) < 9 || len(arg) > 9 {
		fail()
		z01.PrintRune('\n')
		return
	}
	for _, i := range arg {
		if len(i) != 9 {
			fail()
			z01.PrintRune('\n')
			return
		}
	}
	board := [][]int{
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	for idx, i := range arg {
		for idxj, j := range i {
			if checkInt(j) {
				if isValidPlace(board, int(j-'0'), idx, idxj) {
					board[idx][idxj] = int(j - '0')
				} else {
					fail()
					z01.PrintRune('\n')
					return
				}
			} else if checkDot(j) {
				board[idx][idxj] = 0
			} else {
				fail()
				z01.PrintRune('\n')
				return
			}
		}
	}
	if solveSudoku(board) {
		finBoard := board
		printBoard(finBoard)
	} else {
		fail()
		z01.PrintRune('\n')
	}
}

// 1) Все что связанно с sudoku
// Здесь мы проверяем доску на свободные клетки и ставим цифру (Вообщем главная функция которая решает задачу)
func solveSudoku(board [][]int) bool {
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if board[row][column] == 0 {
				for try := 1; try <= 9; try++ {
					if isValidPlace(board, try, row, column) {
						board[row][column] = try
						if solveSudoku(board) {
							return true
						} else {
							board[row][column] = 0
						}
					}
				}
				return false
			}
		}
	}
	return true
}

// Проверка всех 3 условий в 1 функций
func isValidPlace(board [][]int, num, row, column int) bool {
	return !checkRow(board, num, row) &&
		!checkColumn(board, num, column) &&
		!checkBox(board, num, row, column)
}

// Проверка по строке
func checkRow(board [][]int, num, row int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return true
		}
	}
	return false
}

// Проверка по колонне
func checkColumn(board [][]int, num, column int) bool {
	for i := 0; i < 9; i++ {
		if board[i][column] == num {
			return true
		}
	}
	return false
}

// Проверка внутри квадрата 3х3
func checkBox(board [][]int, num, row, column int) bool {
	localRow := row - row%3
	localColumn := column - column%3
	for i := localRow; i < localRow+3; i++ {
		for j := localColumn; j < localColumn+3; j++ {
			if board[i][j] == num {
				return true
			}
		}
	}
	return false
}

// 2) Все что связанно с выводом и проверкой
// Проверка на цифру
func checkInt(r rune) bool {
	if r >= 48 && r <= 57 {
		return true
	}
	return false
}
func fail() {
	r := []rune("Error")
	for _, i := range r {
		z01.PrintRune(i)
	}
}
func printBoard(b [][]int) {
	for _, i := range b {
		for idxj, j := range i {
			z01.PrintRune(rune(j + '0'))
			if idxj != 8 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
func checkDot(r rune) bool {
	if r == 46 {
		return true
	}
	return false
}
