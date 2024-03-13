package api

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"mBoxMini/internal/users"
	"net/http"
)

func (s *RestAPI) Registration(c *gin.Context) {
	authorization, exists := c.Get("authorization")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if authorization == true {
		c.JSON(http.StatusConflict, gin.H{"error": "fail login already user"})
		return
	}

	user := users.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	userID, err := s.BoxService.GetUser(user.Login)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	}
	if userID.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Login already taken"})
		return
	}

	err = s.BoxService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.SetCookie("login", user.Login, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (s *RestAPI) Authentication(c *gin.Context) {
	authorization, exists := c.Get("authorization")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if authorization == true {
		c.JSON(http.StatusConflict, gin.H{"error": "fail login already user"})
		return
	}

	user := users.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	userID, err := s.BoxService.GetUser(user.Login)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashPassword, err := user.PasswordHashToString(userID.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "password could not be verified"})
		return
	}

	if err = user.ComparedPass(hashPassword, user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect password"})
		return
	}
	c.SetCookie("login", user.Login, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
