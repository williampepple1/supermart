package controllers

import (
	"net/http"
	"supermart/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ListBranches(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var branchs []models.Branch
		if err := db.Find(&branchs).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve branchs"})
			return
		}
		c.JSON(http.StatusOK, branchs)
	}
}

func CreateBranch(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newBranch models.Branch
		if err := c.ShouldBindJSON(&newBranch); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&newBranch).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating branch"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Branch created successfully", "branch": newBranch})

	}
}

func GetBranch(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var branch models.Branch
		if err := db.Where("id = ?", id).First(&branch).Error; err != nil {
			c.JSON(404, gin.H{"error": "Branch not found"})
			return
		}

		c.JSON(http.StatusOK, branch)
	}
}

func UpdateBranch(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var branch models.Branch
		branchId := c.Param("id")

		if err := db.First(&branch, branchId).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Branch not found"})
			return
		}

		if err := c.ShouldBindJSON(&branch); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Save(&branch)

		c.JSON(http.StatusOK, gin.H{"message": "Branch updated successfully", "branch": branch})
	}
}

func DeleteBranch(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var branch models.Branch
		branchId := c.Param("id")

		if err := db.First(&branch, branchId).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Branch not found"})
			return
		}

		db.Delete(&branch)

		c.JSON(http.StatusOK, gin.H{"message": "Branch deleted successfully"})
	}
}
