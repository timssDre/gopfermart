package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mBoxMini/internal/users"
	"net/http"
)

var ErrNoRows = errors.New("sql: no rows in result set")

func (s *RestAPI) Registration(c *gin.Context) {
	userInfo, exists := c.Get("userInfo")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	userInfoData, ok := userInfo.(*users.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	user := users.User{
		ID:  userInfoData.ID,
		New: userInfoData.New,
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	userID, err := s.BoxService.GetUser(user.Login)
	if err != nil {
		if errors.Is(err, ErrNoRows) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "the user was not found."})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userID != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Login already taken"})
		return
	}
	err = s.BoxService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
