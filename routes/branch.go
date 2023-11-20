package routes

import (
	"supermart/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// "supermart/middleware"
)

func SetupBranchRoutes(r *gin.Engine, db *gorm.DB) {
	authorized := r.Group("/branch")
	// authorized.Use(middleware.Authorize(db))
	{
		authorized.GET("/", controllers.ListBranches(db))
		authorized.GET("/:id", controllers.GetBranch(db))
		authorized.POST("/", controllers.CreateBranch(db))
		authorized.PUT("/:id", controllers.UpdateBranch(db))
		authorized.DELETE("/:id", controllers.DeleteBranch(db))
	}
}
