package services

import (
	"math/rand"
	"time"

	"github.com/your-org/backend-golang-player-service/internal/models"
	"gorm.io/gorm"
)

type PlayerService struct {
	db *gorm.DB
}

func NewPlayerService(db *gorm.DB) *PlayerService {
	rand.Seed()
	return &PlayerService{db: db}
}

func (s *PlayerService) GetAllPlayers() ([]models.Player, error) {
	var players []models.Player
	if err := s.db.Find(&players).Error; err != nil {
		return nil, err
	}
	return players, nil
}

func (s *PlayerService) GetPlayerByID(id string) (*models.Player, error) {
	// Simulate random latency between 0 and 2 seconds
	latency := rand.Intn(2000) // 0-2000ms
	time.Sleep(time.Duration(latency) * time.Millisecond)

	var player models.Player
	if err := s.db.First(&player, id).Error; err != nil {
		return nil, err
	}
	return &player, nil
} 