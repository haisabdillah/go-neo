package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type claims struct {
	jwt.RegisteredClaims
	Payload payload `json:"payload"`
}

type payload struct {
	ID          uint     `json:"id"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization") // Correct usage of Header
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header required"})
			c.Abort() // Stop the execution of the next handlers
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Authorization header format"})
			c.Abort() // Stop the execution of the next handlers
			return
		}
		tokenStr := parts[1]

		claims := &claims{}
		var jwtKey = []byte(os.Getenv("JWT_SECRET"))

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "unexpected signing method"})
				c.Abort() // Stop the execution of the next handlers
			}
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token signature"})
				c.Abort() // Stop the execution of the next handlers
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token signature"})
			c.Abort() // Stop the execution of the next handlers
			return

		}
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort() // Stop the execution of the next handlers
			return
		}

		c.Set("authID", claims.Payload.ID)
		c.Set("authRole", claims.Payload.Role)
		c.Set("authPermissions", claims.Payload.Permissions)
		c.Next()
	}

}
