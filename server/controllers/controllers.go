package controllers

import (
	"log"
	"net/http"

	"github.com/DEMYSTIF/gin-postgres-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOne(c *gin.Context, db *gorm.DB) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
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
	var user models.User
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Query database for user with given username
	result := db.Where("username = ?", username).First(&user)
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

// func ReadAll(c *gin.Context, db *gorm.DB) {
// 	var certificates []models.Certificate
// 	result := db.Find(&certificates)
// 	if result.Error != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, certificates)
// }

// func ReadOne(c *gin.Context, db *gorm.DB) {
// 	var oldUser models.User
// 	param := c.Param("id")
// 	id, err := strconv.Atoi(param)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
// 		return
// 	}

// 	result := db.First(&oldUser, "id = ?", id)
// 	if result.Error != nil {
// 		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, oldUser)
// }

// func UpdateOne(c *gin.Context, db *gorm.DB) {
// 	var update models.Certificate
// 	param := c.Param("id")
// 	id, err := strconv.Atoi(param)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
// 		return
// 	}

// 	result := db.First(&update, "id = ?", id)
// 	if result.Error != nil {
// 		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
// 		return
// 	}

// 	if err := c.ShouldBindJSON(&update); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
// 		return
// 	}

// 	result = db.Save(&update)
// 	if result.Error != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, update)
// }

// func DeleteOne(c *gin.Context, db *gorm.DB) {
// 	var trash models.Certificate
// 	param := c.Param("id")
// 	id, err := strconv.Atoi(param)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
// 		return
// 	}

// 	result := db.First(&trash, "id = ?", id)
// 	if result.Error != nil {
// 		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
// 		return
// 	}

// 	result = db.Delete(&models.Certificate{}, id)
// 	if result.Error != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, trash)
// }
