package models

import (
	"time"

	"gorm.io/gorm"
)

type Difficulty string

const (
	Easy   Difficulty = "easy"
	Medium Difficulty = "medium"
	Hard   Difficulty = "hard"
)

type Puzzle struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Difficulty   Difficulty     `json:"difficulty" gorm:"not null"`
	StartingGrid string         `json:"starting_grid" gorm:"not null"` // 81 characters representing the initial board
	Solution     string         `json:"solution" gorm:"not null"`      // 81 characters representing the complete solution
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}
