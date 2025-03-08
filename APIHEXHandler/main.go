package main

import (
	"log"
	"os"
	"time"

	"github.com/M1keTrike/EventDriven/database"
	"github.com/M1keTrike/EventDriven/offers/dependencies"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}
	defer db.Close()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	offerDeps := dependencies.NewOfferDependencies(db)
	offerDeps.Execute(r)

	r.Run(":" + PORT)
}
