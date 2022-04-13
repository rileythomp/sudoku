package sudoku

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
)

// /board?board= - returns the board sent
func GetBoard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetBoard")

	keys, ok := r.URL.Query()["board"]

	if !ok || len(keys[0]) < 1 {
		fmt.Println("Query Error")
		fmt.Println("/GetBoard\n")
		return
	}

	board := keys[0]

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(board)

	fmt.Println("/GetBoard\n")
}

// /validate?board= - returns whether or not the board sent is a valid 
// 							 sudoku board
func ValidateBoard(w http.ResponseWriter, r  *http.Request) {
	fmt.Println("ValidateBoard")

	keys, ok := r.URL.Query()["board"]

	if !ok || len(keys[0]) < 1 {
		fmt.Println("Query Error")
		fmt.Println("/ValidateBoard\n")
		return
	}

	board := keys[0]

	valid := BoardIsValid(board)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(valid)

	fmt.Println("/ValidateBoard\n")
}

// /partial?board= - returns the board sent with only logical deductions made
func PartialSolve(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PartialSolve")

	keys, ok := r.URL.Query()["board"]

	if !ok || len(keys[0]) < 1 {
		fmt.Println("Query Error")
		fmt.Println("/PartialSolve\n")
		return
	}

	board := keys[0]

	partialBoard, _, _, _, _ := LogicalSolvev1(board)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(partialBoard)

	fmt.Println("/PartialSolve\n")
}

// /solve?board= - returns the board sent completely solved
func SolveBoard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SolveBoard")

	keys, ok := r.URL.Query()["board"]

	if !ok || len(keys[0]) < 1 {
		fmt.Println("Query Error")
		fmt.Println("/ValidateBoard\n")
		return
	}

	board := keys[0]

	solvedBoard, _ := SolveBoardv4(board)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(solvedBoard)

	fmt.Println("/SolveBoard\n")
}

// /test - runs a test on many boards of different difficulty
func TestSolvers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TestingSolvers\n")

	fmt.Println(" ___________________________________________________________________________________________________")
	fmt.Printf("| %-20s | %8s | %10s | %11s | %10s | %10s | %10s |\n", " ", " ", " ", " ", " ", " ", " ")
	header := fmt.Sprintf(
		"| %-20s | %8s | %8s | %11s | %10s | %10s | %10s |\n",
		"Solver Name",
		"# Boards",
		"Difficulty",
		"Total time",
		"Avg. time",
		"Max. time",
		"Min. time",
	)
	fmt.Print(header)
	fmt.Println("|______________________|__________|____________|_____________|____________|____________|____________|")

	// TestSolver("tests/easy1000.txt", SolveBoardv0, "v0", "easy", 20)
	// TestSolver("tests/easy1000.txt", SolveBoardv1, "v1", "easy", 200)
	// TestSolver("tests/easy1000.txt", SolveBoardv2, "v2", "easy", 1000)
	// TestSolver("tests/easy1000.txt", SolveBoardv3, "v3", "easy", 100)
	// TestSolver("tests/easy1000.txt", SolveBoardv4, "v4", "easy", 100)


	// TestSolver("tests/med1000.txt", SolveBoardv0, "v0", "medium", 5)
	// TestSolver("tests/med1000.txt", SolveBoardv1, "v1", "medium", 50)
	// TestSolver("tests/med1000.txt", SolveBoardv2, "v2", "medium", 500)
	// TestSolver("tests/med1000.txt", SolveBoardv3, "v3", "medium", 50)
	// TestSolver("tests/med1000.txt", SolveBoardv4, "v4", "medium", 50)

	// TestSolver("tests/hard1000.txt", SolveBoardv0, "v0", "hard", 1)
	// TestSolver("tests/hard1000.txt", SolveBoardv1, "v1", "hard", 10)
	// TestSolver("tests/hard1000.txt", SolveBoardv2, "v2", "hard", 100)
	// TestSolver("tests/hard1000.txt", SolveBoardv3, "v3", "hard", 10)
	// TestSolver("tests/hard1000.txt", SolveBoardv4, "v4", "hard", 10)


	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(0)

	fmt.Println("/TestingSolvers\n")
}

func AddTest(w http.ResponseWriter, r *http.Request) {
	log.Print("AddTest\n")

	// var wg sync.WaitGroup

	// wg.Add(15)

	// go AddTests(&wg, "tests/easy1000.txt", SolveBoardv0, "v0", "easy", 1000);
	// go AddTests(&wg, "tests/easy1000.txt", SolveBoardv1, "v1", "easy", 1000);
	// go AddTests(&wg, "tests/easy1000.txt", SolveBoardv2, "v2", "easy", 1000);
	// go AddTests(&wg, "tests/easy1000.txt", SolveBoardv3, "v3", "easy", 1000);
	// go AddTests(&wg, "tests/easy1000.txt", SolveBoardv4, "v4", "easy", 1000);

	// go AddTests(&wg, "tests/med1000.txt", SolveBoardv0, "v0", "medium", 1000);
	// go AddTests(&wg, "tests/med1000.txt", SolveBoardv1, "v1", "medium", 1000);
	// go AddTests(&wg, "tests/med1000.txt", SolveBoardv2, "v2", "medium", 1000);
	// go AddTests(&wg, "tests/med1000.txt", SolveBoardv3, "v3", "medium", 1000);
	// go AddTests(&wg, "tests/med1000.txt", SolveBoardv4, "v4", "medium", 1000);

	// go AddTests(&wg, "tests/hard1000.txt", SolveBoardv0, "v0", "hard", 1000);
	// go AddTests(&wg, "tests/hard1000.txt", SolveBoardv1, "v1", "hard", 1000);
	// go AddTests(&wg, "tests/hard1000.txt", SolveBoardv2, "v2", "hard", 1000);
	// go AddTests(&wg, "tests/hard1000.txt", SolveBoardv3, "v3", "hard", 1000);
	// go AddTests(&wg, "tests/hard1000.txt", SolveBoardv4, "v4", "hard", 1000);

	// wg.Wait()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(0)

	log.Print("/AddTest\n")
}