package sudoku

import (
	"errors"
	"math/rand"
	"time"

	"gorm.io/gorm"

	"sudoku/internal/models"
)

type Service struct {
	db *gorm.DB
}

type Board [9][9]int

type Move struct {
	Row    int    `json:"row"`
	Col    int    `json:"col"`
	Value  int    `json:"value"`
	Reason string `json:"reason"`
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

// Convert string representation to Board
func StringToBoard(s string) Board {
	var board Board
	for i := 0; i < 81; i++ {
		row := i / 9
		col := i % 9
		board[row][col] = int(s[i] - '0')
	}
	return board
}

// Convert Board to string representation
func BoardToString(board Board) string {
	var s string
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			s += string(rune(board[i][j] + '0'))
		}
	}
	return s
}

// Validate if a move is valid
func (s *Service) IsValidMove(board Board, row, col, value int) bool {
	// Check row
	for j := 0; j < 9; j++ {
		if j != col && board[row][j] == value {
			return false
		}
	}

	// Check column
	for i := 0; i < 9; i++ {
		if i != row && board[i][col] == value {
			return false
		}
	}

	// Check 3x3 box
	boxRow := (row / 3) * 3
	boxCol := (col / 3) * 3
	for i := boxRow; i < boxRow+3; i++ {
		for j := boxCol; j < boxCol+3; j++ {
			if (i != row || j != col) && board[i][j] == value {
				return false
			}
		}
	}

	return true
}

// Get valid candidates for a cell
func (s *Service) GetCandidates(board Board, row, col int) []int {
	if board[row][col] != 0 {
		return []int{}
	}

	var candidates []int
	for value := 1; value <= 9; value++ {
		if s.IsValidMove(board, row, col, value) {
			candidates = append(candidates, value)
		}
	}
	return candidates
}

// Find naked singles (cells with only one candidate)
func (s *Service) FindNakedSingles(board Board) []Move {
	var moves []Move
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				candidates := s.GetCandidates(board, i, j)
				if len(candidates) == 1 {
					moves = append(moves, Move{
						Row:    i,
						Col:    j,
						Value:  candidates[0],
						Reason: "Naked Single",
					})
				}
			}
		}
	}
	return moves
}

// Get hint for a specific cell
func (s *Service) GetHint(board Board, row, col int) (*Move, error) {
	if board[row][col] != 0 {
		return nil, errors.New("cell is already filled")
	}

	// Try to find a naked single
	nakedSingles := s.FindNakedSingles(board)
	for _, move := range nakedSingles {
		if move.Row == row && move.Col == col {
			return &move, nil
		}
	}

	// If no naked single, return a random valid candidate
	candidates := s.GetCandidates(board, row, col)
	if len(candidates) == 0 {
		return nil, errors.New("no valid candidates for this cell")
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(candidates))

	return &Move{
		Row:    row,
		Col:    col,
		Value:  candidates[randomIndex],
		Reason: "Hint",
	}, nil
}

// Solve puzzle using backtracking
func (s *Service) SolvePuzzle(board Board) (Board, bool) {
	var solved Board
	copy(solved[:], board[:])

	if s.solve(&solved) {
		return solved, true
	}
	return board, false
}

func (s *Service) solve(board *Board) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for value := 1; value <= 9; value++ {
					if s.IsValidMove(*board, i, j, value) {
						board[i][j] = value
						if s.solve(board) {
							return true
						}
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

// Validate if a board is complete and correct
func (s *Service) ValidateSolution(board Board) bool {
	// Check if all cells are filled
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return false
			}
		}
	}

	// Check rows
	for i := 0; i < 9; i++ {
		seen := make(map[int]bool)
		for j := 0; j < 9; j++ {
			if seen[board[i][j]] {
				return false
			}
			seen[board[i][j]] = true
		}
	}

	// Check columns
	for j := 0; j < 9; j++ {
		seen := make(map[int]bool)
		for i := 0; i < 9; i++ {
			if seen[board[i][j]] {
				return false
			}
			seen[board[i][j]] = true
		}
	}

	// Check 3x3 boxes
	for boxRow := 0; boxRow < 9; boxRow += 3 {
		for boxCol := 0; boxCol < 9; boxCol += 3 {
			seen := make(map[int]bool)
			for i := boxRow; i < boxRow+3; i++ {
				for j := boxCol; j < boxCol+3; j++ {
					if seen[board[i][j]] {
						return false
					}
					seen[board[i][j]] = true
				}
			}
		}
	}

	return true
}

// Calculate score based on correct moves
func (s *Service) CalculateScore(initialBoard, finalBoard Board) int {
	score := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if initialBoard[i][j] == 0 && finalBoard[i][j] != 0 {
				// Check if the move is correct
				if s.IsValidMove(finalBoard, i, j, finalBoard[i][j]) {
					score += 10
				}
			}
		}
	}
	return score
}

// Get random puzzle by difficulty
func (s *Service) GetRandomPuzzle(difficulty models.Difficulty) (*models.Puzzle, error) {
	var puzzle models.Puzzle
	err := s.db.Where("difficulty = ?", difficulty).Order("RANDOM()").First(&puzzle).Error
	if err != nil {
		return nil, err
	}
	return &puzzle, nil
}
