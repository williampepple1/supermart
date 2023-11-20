package routes

import (
	"supermart/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// "supermart/middleware"
)

func SetupProductRoutes(r *gin.Engine, db *gorm.DB) {
	authorized := r.Group("/branch")
	// authorized.Use(middleware.Authorize(db))
	{
		authorized.GET("/", controllers.ListProducts(db))
		authorized.GET("/:id", controllers.GetProduct(db))
		authorized.POST("/", controllers.CreateProduct(db))
		authorized.PUT("/:id", controllers.UpdateProduct(db))
		authorized.DELETE("/:id", controllers.DeleteProduct(db))
	}
}
