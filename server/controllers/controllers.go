package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/DEMYSTIF/gin-postgres-api/models"
	"github.com/DEMYSTIF/gin-postgres-api/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateOne(c *gin.Context, db *gorm.DB) {
	player := new(models.Player)
	if err := c.ShouldBindJSON(player); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}
	// check if entries have empty values
	if strings.TrimSpace(player.UserName) == "" {
		c.AbortWithStatusJSON(430, gin.H{"message": "Username is not allowed to be empty"})
		return
	}
	if strings.TrimSpace(player.Password) == "" {
		c.AbortWithStatusJSON(430, gin.H{"message": "Password is not allowed to be empty"})
		return
	}
	if strings.TrimSpace(player.Network) == "" {
		c.AbortWithStatusJSON(430, gin.H{"message": "Network is not allowed to be empty"})
		return
	}
	if strings.TrimSpace(player.Address) == "" {
		c.AbortWithStatusJSON(430, gin.H{"message": "Address is not allowed to be empty"})
		return
	}
	// Check if the network value is valid
	if player.Network != "ethereum" && player.Network != "stellar" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Network must be either 'ethereum' or 'stellar'"})
		return
	}

	// Check address length for network
	if player.Network == "ethereum" {
		if len(player.Address) != 42 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Length of 'ethereum' address should be 42"})
			return
		}
	} else {
		if len(player.Address) != 56 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Length of 'stellar' address should be 56"})
			return
		}
	}
	// AutoMigrate will create the table if it doesn't exist
	if err := db.AutoMigrate(&models.Player{}); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(player.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	player.Password = string(passwordHash)

	result := db.Create(player)
	if result.Error != nil {
		log.Println("Error occurred:", result.Error)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	if player.Network == "ethereum" {
		if err := services.IssueNFT(player); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Unable to mint NFT"})
			return
		}
	}

	c.IndentedJSON(http.StatusCreated, player)
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
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(401, gin.H{"error": "Invalid password"})
		return
	}

	// Login successful
	c.JSON(200, gin.H{"message": "Login successful"})
}
