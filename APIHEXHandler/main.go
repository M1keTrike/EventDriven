package main

import (
	"log"
	"os"

	"github.com/M1keTrike/EventDriven/database"
	"github.com/M1keTrike/EventDriven/offers/dependencies"

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

	offerDeps := dependencies.NewOfferDependencies(db)
	offerDeps.Execute(r)

	r.Run(":" + PORT)

}
