package main

import (
	"log"

	"github.com/DEMYSTIF/gin-postgres-api/src/controllers"
	"github.com/DEMYSTIF/gin-postgres-api/src/db"
	"github.com/DEMYSTIF/gin-postgres-api/src/middlewares"
	"github.com/DEMYSTIF/gin-postgres-api/src/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		log.Fatal("Failed to connect the database")
	}

	db.AutoMigrate(&models.Certificate{})

	router := gin.Default()
	router.Use(middlewares.Authority())
	router.POST("/create", func(ctx *gin.Context) {
		controllers.CreateOne(ctx, db)
	})
	router.GET("/read", func(ctx *gin.Context) {
		controllers.ReadAll(ctx, db)
	})
	router.GET("/read/:id", func(ctx *gin.Context) {
		controllers.ReadOne(ctx, db)
	})
	router.PUT("/update/:id", func(ctx *gin.Context) {
		controllers.UpdateOne(ctx, db)
	})
	router.DELETE("/delete/:id", func(ctx *gin.Context) {
		controllers.DeleteOne(ctx, db)
	})
	router.Run("localhost:8080")
}
