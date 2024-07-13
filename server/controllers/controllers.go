package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/DEMYSTIF/gin-postgres-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOne(c *gin.Context, db *gorm.DB) {
	var user models.Player
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}
	log.Println(user)
	// check if entries have empty values
	if strings.TrimSpace(user.UserName) == "" {
		c.AbortWithStatusJSON(430, gin.H{"message": "Username is not allowed to be empty"})
		return
	}
	if strings.TrimSpace(user.Password) == "" {
		c.AbortWithStatusJSON(430, gin.H{"message": "Password is not allowed to be empty"})
		return
	}
	if strings.TrimSpace(user.Network) == "" {
		c.AbortWithStatusJSON(430, gin.H{"message": "Network is not allowed to be empty"})
		return
	}
	if strings.TrimSpace(user.Address) == "" {
		c.AbortWithStatusJSON(430, gin.H{"message": "Address is not allowed to be empty"})
		return
	}
	// Check if the network value is valid
	if user.Network != "ethereum" && user.Network != "stellar" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Network must be either 'ethereum' or 'stellar'"})
		return
	}

	// Check address length for network
	if user.Network == "ethereum" {
		if len(user.Address) != 42 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Length of 'ethereum' address should be 42"})
			return
		}
	} else {
		if len(user.Address) != 56 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Length of 'stellar' address should be 56"})
			return
		}
	}
	// AutoMigrate will create the table if it doesn't exist
	if err := db.AutoMigrate(&models.Player{}); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	result := db.Create(&user)
	if result.Error != nil {
		log.Println("Error occurred:", result.Error)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	// Mint NFT for User

	c.IndentedJSON(http.StatusCreated, user)
}
func VerifyLogin(c *gin.Context, db *gorm.DB) {
	type LogInData struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
	}
	var user models.Player
	var loginData LogInData
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}
	username := loginData.UserName
	password := loginData.Password

	// Query database for user with given username
	result := db.Where("user_name = ?", username).First(&user)
	if result.Error != nil {
		c.JSON(401, gin.H{"error": "User not found"})
		return
	}

	// Verify password
	if user.Password != password {
		c.JSON(401, gin.H{"error": "Invalid password"})
		return
	}

	// Login successful
	c.JSON(200, gin.H{"message": "Login successful"})
}
