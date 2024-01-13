package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mBoxMini/internal/users"
	"net/http"
)

var ErrNoRows = errors.New("sql: no rows in result set")

func (s *RestAPI) Registration(c *gin.Context) {
	user := users.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	userID, err := s.BoxService.GetUser(user.Login)
	if err != nil && !errors.As(err, &ErrNoRows) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if userID != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Login already taken"})
		return
	}
	err = s.BoxService.CreateUser(&user)
	c.JSON(http.StatusOK, "test")
}
