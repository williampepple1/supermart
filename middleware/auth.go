package middleware

import (
	"os"
	"strings"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"net/http"
	"supermart/models"
)

// JWT key used to create the signature.
var jwtKey = os.Getenv("JWT_SECRET_KEY")

func init() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func Authorize(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header not provided"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(401, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		tokenStr := parts[1]

		claims := &jwt.StandardClaims{}
		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		if err != nil || !tkn.Valid {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		// Check if the user is active
		var user models.User

		if err := db.Where("id = ?", claims.Id).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		if !user.Active {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not active"})
			c.Abort()
			return
		}

		// Set the username in the context for subsequent handlers to use
		c.Set("username", claims.Subject)
		c.Set("userId", claims.Id)
		c.Next() // proceed to next middleware or handler
	}
}
