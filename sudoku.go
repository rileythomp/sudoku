// Sudoku helper functions

// firstEmptyCell          - Returns if the given board is completed and the
//                           index of the first empty cell if it is not
// BoardIsValid            - Returns whether the given board is a valid sudoku board
// PrintSudokuTestTableRow - Prints a sudoku test row with info like the number
//                           of tests, time and difficulty
// TestSolver              - Tests a solver on a given number of boards of a
//                           given difficulty

package sudoku

import (
	"fmt"
	"time"
	"os"
	"bufio"
	"strings"
	"database/sql"
	"log"
	"sync"

	"github.com/google/uuid"
)

var primes = map[int]int{
	1: 2,
	2: 3,
	3: 5,
	4: 7,
	5: 11,
	6: 13,
	7: 17,
	8: 19,
	9: 23,
}

var primeIndices = map[int]string {
	2: "1",
	3: "2",
	5: "3",
	7: "4",
	11: "5",
	13: "6",
	17: "7",
	19: "8",
	23: "9",
}

var maxVal int = 2*3*5*7*11*13*17*19*23

// Returns if the given board is completed, and
// the index of the first empty cell if it is not
func firstEmptyCell(board string) (bool, int) {
	index := strings.Index(board, "0")
	return index < 0, index
}

// Returns whether the given board is a valid sudoku board
func BoardIsValid(board string) bool {
	valid := true	

	for i := 0; i < 9 && valid; i++ {
		rowSet := make(map[string]bool)
		colSet := make(map[string]bool)
		quadSet := make(map[string]bool)
		
		for j := 0; j < 9; j++ {
			rowIndex := 9*i+j
			colIndex := 9*j+i
			quadIndex := 18*(i/3) + 3*i + 9*(j/3) + (j%3)

			rowVal := string(board[rowIndex])
			colVal := string(board[colIndex])
			quadVal := string(board[quadIndex])

			if _, ok := rowSet[rowVal]; ok && rowVal != "0" {
				valid = false
				break
			}
			if _, ok := colSet[colVal]; ok && colVal != "0" {
				valid = false
				break
			}
			if _, ok := quadSet[quadVal]; ok && quadVal != "0" {
				valid = false
				break
			}

			rowSet[rowVal], colSet[colVal], quadSet[quadVal] = true, true, true
		}
	}

	return valid
}

// Prints a sudoku test row with info like
// the number of tests, time and difficulty
func PrintSudokuTestTableRow(
	name string,
	numBoards int,
	difficulty string,
	totalTime time.Duration,
	maxTime time.Duration,
	minTime time.Duration,
) {
	secondFloat := float64(time.Second)
	totalSeconds := float64(totalTime)/secondFloat

	fmt.Printf("| %-20s | %8s | %10s | %11s | %10s | %10s | %10s |\n", " ", " ", " ", " ", " ", " ", " ")
	stats := fmt.Sprintf(
		"| %-20s | %8d | %10s | %10fs | %9fs | %9fs | %9fs |",
		name,
		numBoards,
		difficulty,
		totalSeconds,
		totalSeconds/float64(numBoards),
		float64(maxTime)/secondFloat,
		float64(minTime)/secondFloat,
	)
	fmt.Println(stats)
	fmt.Println("|______________________|__________|____________|_____________|____________|____________|____________|")
}

// Tests a solver on a given number of boards of a given difficulty
func TestSolver(
	fileName string,
	solver (func(string) (string, bool)),
	solverName string,
	difficulty string,
	numBoards int,
) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var totalTime time.Duration = 0
	var maxTime time.Duration = 0
	var minTime time.Duration = 1000000000
	i := 0

	for scanner.Scan() && i < numBoards{
		board := scanner.Text()
		i++

		boardStart  := time.Now()
		_, solved := solver(board)
		solveTime := time.Since(boardStart)

		if !solved {
			fmt.Println("Error: Board not solved")
			fmt.Println(board)
			break
		}

		maxTime, minTime = MaxTime(maxTime, solveTime), MinTime(minTime, solveTime)
		totalTime += solveTime
	}

	PrintSudokuTestTableRow(solverName, i, difficulty, totalTime, maxTime, minTime)
}

func AddTests(
	wg *sync.WaitGroup,
	fileName string,
	solver (func(string) (string, bool)),
	solverName string,
	difficulty string,
	numBoards int,
) {
	defer wg.Done()

	log.Print(fmt.Sprintf("Initiating %d %s %s tests", numBoards, difficulty, solverName))

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer file.Close()

	DB_URL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", DB_URL)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	scanner := bufio.NewScanner(file)
	solves := 0

	insertTests := "INSERT INTO tests (id, solver, boardstr, numclues, difficulty, solvetime) VALUES "

	for scanner.Scan() && solves < numBoards {
		solves++
		board := scanner.Text()

		boardStart  := time.Now()
		_, solved := solver(board)
		solveTime := time.Since(boardStart)

		if !solved {
			fmt.Println("Error: Board not solved")
			fmt.Println(board)
			break
		}

		insertTests += fmt.Sprintf("('%v', '%v', '%v', %v, '%v', %v), ", uuid.New(), solverName, board, 81-strings.Count(board, "0"), difficulty, solveTime.Seconds())
	}

	insertTests = strings.TrimSuffix(insertTests, ", ") + ";"
	_, err = db.Exec(insertTests)
	if err != nil {
		panic(err)
	}

	log.Print(fmt.Sprintf("Wrote %d %s %s tests to DB", solves, difficulty, solverName))
}