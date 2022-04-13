// Sudoku solvers

package sudoku

import (
	"strconv"
)

// Solves with only recursive backtracking
func SolveBoardv0(board string) (string, bool) {
	return recursiveBacktrackv1(board, false)
}

// Makes logical deductions, then backtracks
func SolveBoardv1(board string) (string, bool) {
	partialBoard, _, _, _, flip := LogicalSolvev1(board);
	return recursiveBacktrackv1(partialBoard, flip)
}

// Makes logical deductions, then backtracks with flipping
func SolveBoardv2(board string) (string, bool) {
	partialBoard, _, _, _, flip := logicalSolvev2(board, 40);
	return recursiveBacktrackv1(partialBoard, flip)
}

// Makes logical deductions, then backtracks efficiently
func SolveBoardv3(board string) (string, bool) {
	return recursiveBacktrackv2(LogicalSolvev1(board))
}

// Makes logical deductions, then backtracks efficiently with flipping
func SolveBoardv4(board string) (string, bool) {
	return recursiveBacktrackv2(logicalSolvev2(board, 40))
}

// Standard recursive backtracking
func recursiveBacktrackv1(board string, flipBoard bool) (string, bool) {
	filled, index := firstEmptyCell(board)

	if filled && BoardIsValid(board) {
		return board, true
	}

	for val := 1; val <= 9; val++ {
		newboard := board[:index] + strconv.Itoa(val) + board[index+1:]
		if (BoardIsValid(newboard)) {
			board = newboard
			if solvedBoard, isSolved := recursiveBacktrackv1(board, false); isSolved {
				if flipBoard {
					solvedBoard = reverse(solvedBoard)
				}
				return solvedBoard, isSolved
			}
		}		
	}

	return board, false
}

// Recursive backtracking but only with possible values
func recursiveBacktrackv2(
	board string,
	rowVals [9]int,
	colVals [9]int,
	quadVals [9]int,
	flipBoard bool,
) (string, bool) {
	filled, index := firstEmptyCell(board)

	if filled && BoardIsValid(board) {
		return board, true
	}

	r, c := index/9, index%9
	q := (r/3)*3 + (c/3)
	primeVal := maxVal/LCM(rowVals[r], colVals[c], quadVals[q])
	primeFactors := primeFactors(primeVal)

	for i := 0; i < len(primeFactors); i++ {
		primeFactor := primeFactors[i]
		if val, ok := primeIndices[primeFactor]; ok {
			newboard := board[:index] + val + board[index+1:]
			if (BoardIsValid(newboard)) {
				board = newboard

				rowVals[r] *= primeFactor
				colVals[c] *= primeFactor
				quadVals[q] *= primeFactor

				if solvedBoard, isSolved := recursiveBacktrackv2(board, rowVals, colVals, quadVals, false); isSolved {
					if flipBoard {
						solvedBoard = reverse(solvedBoard)
					}
					return solvedBoard, isSolved
				}

				rowVals[r] /= primeFactor
				colVals[c] /= primeFactor
				quadVals[q] /= primeFactor;
			}
		}
	}

	return board, false
}

// "Only possible value" deductions
func LogicalSolvev1(board string) (string, [9]int, [9]int, [9]int, bool) {
	rowVals := [9]int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	colVals := [9]int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	quadVals := [9]int{1, 1, 1, 1, 1, 1, 1, 1, 1}

	for i := 0; i < len(board); i++ {
		r, c := i/9, i%9
		q := (r/3)*3 + (c/3)
		val := string(board[i])
		
		if val != "0" {
			num, _ := strconv.Atoi(val)
			prime := primes[num]
			rowVals[r] *= prime
			colVals[c] *= prime
			quadVals[q] *= prime
		}
	}

	solving := true
	for solving {
		solving = false
		for i := 0; i < len(board); i++ {
			if string(board[i]) != "0" {
				continue
			}
	
			r, c := i/9, i%9
			q := (r/3)*3 + (c/3)
	
			primeVal := maxVal/LCM(rowVals[r], colVals[c], quadVals[q])
			if val, ok := primeIndices[primeVal]; ok {
				solving = true
				board = board[:i] + val + board[i+1:]
				rowVals[r] *= primeVal
				colVals[c] *= primeVal
				quadVals[q] *= primeVal
			}
		}
	}

	return board, rowVals, colVals, quadVals, false
}

// "Only possible value" deductions and flippingCheck
// board for more efficient recursive backtracking
func logicalSolvev2(board string, flipLimit int) (string, [9]int, [9]int, [9]int, bool) {
	rowVals := [9]int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	colVals := [9]int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	quadVals := [9]int{1, 1, 1, 1, 1, 1, 1, 1, 1}

	for i := 0; i < len(board); i++ {
		val := string(board[i])
		if val != "0" {
			num, _ := strconv.Atoi(val)
			prime := primes[num]
			r, c := i/9, i%9
			q := (r/3)*3 + (c/3)
			rowVals[r] *= prime
			colVals[c] *= prime
			quadVals[q] *= prime
		}
	}

	solving := true
	topClues, bottomClues := 0, 0
	for solving {
		solving = false
		topClues, bottomClues = 0, 0
		for i := 0; i < len(board); i++ {
			if string(board[i]) != "0" {
				if i < flipLimit {
					topClues++
				} else {
					bottomClues++
				}
				continue
			}
	
			r, c := i/9,  i%9
			q := (r/3)*3 + (c/3)
			primeVal := maxVal/LCM(rowVals[r], colVals[c], quadVals[q])
			if val, ok := primeIndices[primeVal]; ok {
				solving = true
				board = board[:i] + val + board[i+1:]
				rowVals[r] *= primeVal
				colVals[c] *= primeVal
				quadVals[q] *= primeVal
			}
		}
	}

	if bottomClues > topClues {
		// reverse board for more efficient backtracking
		board = reverse(board)
		for i, j := 0, len(rowVals)-1; i < j; i, j = i+1, j-1 {
			rowVals[i], rowVals[j] = rowVals[j], rowVals[i]
			colVals[i], colVals[j] = colVals[j], colVals[i]
			quadVals[i], quadVals[j] = quadVals[j], quadVals[i]
		}
	}

	return board, rowVals, colVals, quadVals, bottomClues > topClues
}
