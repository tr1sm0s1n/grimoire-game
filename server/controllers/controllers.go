package controllers

import (
	"log"
	"net/http"
	"strconv"

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
		}
	} else {
		if len(user.Address) != 56 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Length of 'stellar' address should be 56"})
		}
	}

	result := db.Create(&user)
	if result.Error != nil {
		log.Println("Error occurred:", result.Error)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	c.IndentedJSON(http.StatusCreated, user)
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

func ReadOne(c *gin.Context, db *gorm.DB) {
	var oldUser models.User
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	result := db.First(&oldUser, "id = ?", id)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}

	c.IndentedJSON(http.StatusOK, oldUser)
}

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
