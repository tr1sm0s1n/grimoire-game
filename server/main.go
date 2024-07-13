package main

import (
	"log"

	"github.com/DEMYSTIF/gin-postgres-api/controllers"
	"github.com/DEMYSTIF/gin-postgres-api/db"
	"github.com/DEMYSTIF/gin-postgres-api/middlewares"
	"github.com/DEMYSTIF/gin-postgres-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		log.Fatal("Failed to connect the database")
	}

	db.AutoMigrate(&models.Player{})

	router := gin.Default()
	router.Static("/", "./dist")
	router.Use(middlewares.Authority())

	router.POST("/register", func(ctx *gin.Context) {
		controllers.CreateOne(ctx, db)
	})
	router.POST("/login", func(ctx *gin.Context) {
		controllers.VerifyLogin(ctx, db)
	})

	router.Run("0.0.0.0:8080")
}
