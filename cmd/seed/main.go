package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"sudoku/internal/models"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Get database URL from environment
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate models
	if err := db.AutoMigrate(&models.User{}, &models.Puzzle{}, &models.GameResult{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Sample puzzles
	puzzles := []models.Puzzle{
		{
			Difficulty:   models.Easy,
			StartingGrid: "530070000600195000098000060800060003400803001700020006060000280000419005000080079",
			Solution:     "534678912672195348198342567859761423426853791713924856961537284287419635345286179",
		},
		{
			Difficulty:   models.Easy,
			StartingGrid: "009000000000000000000000000000000000000000000000000000000000000000000000000000000",
			Solution:     "123456789456789123789123456234567891567891234891234567345678912678912345912345678",
		},
		{
			Difficulty:   models.Medium,
			StartingGrid: "800000000003600000070090200050007000000045700000100030001000068008500010090000400",
			Solution:     "812753649943682175675491283154237896369845721287169534521974368438526917796318452",
		},
		{
			Difficulty:   models.Medium,
			StartingGrid: "000000000000003085001020000000507000004000100090000000500000073002010000000040009",
			Solution:     "987654321246173985351928746128537694634892157795461832519286473472319568863745219",
		},
		{
			Difficulty:   models.Hard,
			StartingGrid: "000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			Solution:     "123456789456789123789123456234567891567891234891234567345678912678912345912345678",
		},
		{
			Difficulty:   models.Hard,
			StartingGrid: "000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			Solution:     "123456789456789123789123456234567891567891234891234567345678912678912345912345678",
		},
	}

	// Insert puzzles
	for _, puzzle := range puzzles {
		if err := db.Create(&puzzle).Error; err != nil {
			log.Printf("Failed to create puzzle: %v", err)
		} else {
			log.Printf("Created puzzle with ID: %d, Difficulty: %s", puzzle.ID, puzzle.Difficulty)
		}
	}

	log.Println("Database seeding completed!")
}
