package database

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/your-org/backend-golang-player-service/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	// Use :memory: for in-memory database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate the schema
	err = db.AutoMigrate(&models.Player{})
	if err != nil {
		return nil, err
	}

	// Load data from CSV
	file, err := os.Open("Player.csv")
	if err != nil {
		return nil, fmt.Errorf("failed to open Player.csv: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Skip header row
	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV header: %v", err)
	}

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV records: %v", err)
	}

	for _, record := range records {
		// Convert string values to appropriate types
		homeRuns, _ := strconv.Atoi(record[4])
		rbi, _ := strconv.Atoi(record[5])
		batAvg, _ := strconv.ParseFloat(record[3], 64)

		player := models.Player{
			Name:     record[0],
			Team:     record[1],
			Position: record[2],
			BatAvg:   batAvg,
			HomeRuns: homeRuns,
			RBI:      rbi,
		}

		if err := db.Create(&player).Error; err != nil {
			return nil, fmt.Errorf("failed to create player record: %v", err)
		}
	}

	return db, nil
}