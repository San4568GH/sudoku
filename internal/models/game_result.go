package models

import (
	"time"

	"gorm.io/gorm"
)

type GameMode string

const (
	PlayMode  GameMode = "play"
	LearnMode GameMode = "learn"
)

type GameResult struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	UserID        uint           `json:"user_id" gorm:"not null"`
	User          User           `json:"user" gorm:"foreignKey:UserID"`
	PuzzleID      uint           `json:"puzzle_id" gorm:"not null"`
	Puzzle        Puzzle         `json:"puzzle" gorm:"foreignKey:PuzzleID"`
	Mode          GameMode       `json:"mode" gorm:"not null"`
	Score         int            `json:"score" gorm:"default:0"`
	TimeSeconds   int            `json:"time_seconds" gorm:"default:0"`
	Completed     bool           `json:"completed" gorm:"default:false"`
	UsedHints     bool           `json:"used_hints" gorm:"default:false"`
	UsedAutoSolve bool           `json:"used_auto_solve" gorm:"default:false"`
	Disqualified  bool           `json:"disqualified" gorm:"default:false"`
	FinalGrid     string         `json:"final_grid" gorm:"not null"` // 81 characters representing the final board state
	StartedAt     time.Time      `json:"started_at"`
	CompletedAt   *time.Time     `json:"completed_at"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}
