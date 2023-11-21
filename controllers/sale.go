package controllers

import (
	"net/http"
	"supermart/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ListSales(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var sales []models.Sale
		if err := db.Find(&sales).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve sales"})
			return
		}
		c.JSON(http.StatusOK, sales)
	}
}

func CreateSale(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newSale models.Sale
		if err := c.ShouldBindJSON(&newSale); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&newSale).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating sale"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Sale created successfully", "sale": newSale})

	}
}

func GetSale(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var sale models.Sale
		if err := db.Where("id = ?", id).First(&sale).Error; err != nil {
			c.JSON(404, gin.H{"error": "Sale not found"})
			return
		}

		c.JSON(http.StatusOK, sale)
	}
}

func UpdateSale(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var sale models.Sale
		saleId := c.Param("id")

		if err := db.First(&sale, saleId).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Sale not found"})
			return
		}

		if err := c.ShouldBindJSON(&sale); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Save(&sale)

		c.JSON(http.StatusOK, gin.H{"message": "Sale updated successfully", "sale": sale})
	}
}

func DeleteSale(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var sale models.Sale
		saleId := c.Param("id")

		if err := db.First(&sale, saleId).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Sale not found"})
			return
		}

		db.Delete(&sale)

		c.JSON(http.StatusOK, gin.H{"message": "Sale deleted successfully"})
	}
}
