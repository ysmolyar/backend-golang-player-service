package main

import (
	"log"

	"github.com/your-org/backend-golang-player-service/internal/handlers"
	"github.com/your-org/backend-golang-player-service/internal/database"
	"github.com/your-org/backend-golang-player-service/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize DB
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize services
	playerService := services.NewPlayerService(db)
	chatService, err := services.NewChatService()
	if err != nil {
		log.Fatalf("Failed to initialize chat service: %v", err)
	}

	// Initialize handlers
	playerHandler := handlers.NewPlayerHandler(playerService)
	chatHandler := handlers.NewChatHandler(chatService)

	// Initialize router
	r := gin.Default()

	// Routes
	v1 := r.Group("/v1")
	{
		// Player routes
		v1.GET("/players", playerHandler.ListPlayers)
		v1.GET("/players/:id", playerHandler.GetPlayer)

		// Chat routes
		v1.GET("/chat/list-models", chatHandler.ListModels)
		v1.POST("/chat/generate", chatHandler.Generate)
	}

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 