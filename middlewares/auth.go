package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"headless-cms/config/roles"
	"net/http"
	"os"
)

func validateToken(tokenString string, c *gin.Context) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return nil, fmt.Errorf("JWT_SECRET not set")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if token == nil {
		return nil, fmt.Errorf("token is nil")
	}
	return token, nil
}

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(401, gin.H{
			"message":       "Unauthorized",
			"error_message": "No token provided",
		})
		c.Abort()
		return
	}
	// Validate the token
	token, err := validateToken(tokenString, c)
	if err != nil {
		c.SetCookie("Authorization", "", -1, "", "", false, true)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		c.Set("ID", claims["ID"])
		c.Next()
	} else {
		c.SetCookie("Authorization", "", -1, "", "", false, true)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}
}

func RequireSuperAdmin(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(401, gin.H{
			"message":       "Unauthorized",
			"error_message": "No token provided",
		})
		c.Abort()
		return
	}
	// Validate the token
	token, err := validateToken(tokenString, c)
	if err != nil {
		c.SetCookie("Authorization", "", -1, "", "", false, true)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		c.Set("ID", claims["ID"])
		role := claims["role"]
		if role != roles.SuperAdmin {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	} else {
		c.SetCookie("Authorization", "", -1, "", "", false, true)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}
}

func RequireAdmin(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(401, gin.H{
			"message":       "Unauthorized",
			"error_message": "No token provided",
		})
		c.Abort()
		return
	}
	// Validate the token
	token, err := validateToken(tokenString, c)
	if err != nil {
		c.SetCookie("Authorization", "", -1, "", "", false, true)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		c.Set("ID", claims["ID"])
		role := claims["role"]
		if role != roles.SuperAdmin && role != roles.Admin {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	} else {
		c.SetCookie("Authorization", "", -1, "", "", false, true)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}
}
