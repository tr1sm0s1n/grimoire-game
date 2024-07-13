package controllers

import (
	"net/http"
	"strconv"

	"github.com/DEMYSTIF/gin-postgres-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOne(c *gin.Context, db *gorm.DB) {
	var newCertificate models.Certificate
	if err := c.ShouldBindJSON(&newCertificate); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	result := db.Create(&newCertificate)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Duplicate Key"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newCertificate)
}

func ReadAll(c *gin.Context, db *gorm.DB) {
	var certificates []models.Certificate
	result := db.Find(&certificates)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.IndentedJSON(http.StatusOK, certificates)
}

func ReadOne(c *gin.Context, db *gorm.DB) {
	var oldCertificate models.Certificate
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	result := db.First(&oldCertificate, "id = ?", id)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}

	c.IndentedJSON(http.StatusOK, oldCertificate)
}

func UpdateOne(c *gin.Context, db *gorm.DB) {
	var update models.Certificate
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	result := db.First(&update, "id = ?", id)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}

	if err := c.ShouldBindJSON(&update); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	result = db.Save(&update)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.IndentedJSON(http.StatusOK, update)
}

func DeleteOne(c *gin.Context, db *gorm.DB) {
	var trash models.Certificate
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	result := db.First(&trash, "id = ?", id)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}

	result = db.Delete(&models.Certificate{}, id)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.IndentedJSON(http.StatusOK, trash)
}
