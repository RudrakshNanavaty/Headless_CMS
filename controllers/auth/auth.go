package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"headless-cms/initializers"
	"headless-cms/types"
	"net/http"
	"os"
	"time"
)

func SignUp(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":       "Invalid request",
			"error_message": err.Error(),
		})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error hashing password",
			"error_message": err.Error(),
		})
		return
	}
	user.Password = string(hashedPassword)
	// Save the user to the database
	newUser := types.User{
		Username: user.Username,
		Password: user.Password,
	}
	saved := initializers.DB.Create(&newUser)
	if saved.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error saving user",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "User saved successfully",
	})
}

func Login(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var err error
	err = c.BindJSON(&user)
	var dbUser types.User
	retrieved := initializers.DB.First(&dbUser, "username=?", user.Username)
	if retrieved.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Invalid Credentials",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Credentials",
		})
		return
	}

	// creating token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  dbUser.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET not found")
	}
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error signing token",
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 60*60*24, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}
