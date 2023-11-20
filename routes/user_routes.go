package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"supermart/controllers"
)

func SetupUserRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/register", controllers.RegisterUser(db))
	r.POST("/login", controllers.LoginUser(db))
}
