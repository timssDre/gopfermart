package api

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"mBoxMini/internal/users"
	"net/http"
)

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
	if userInfoData.New == true {
		c.JSON(http.StatusConflict, gin.H{"error": "fail login already taken"})
		return
	}

	user := users.User{
		ID:    userInfoData.ID,
		New:   userInfoData.New,
		Token: userInfoData.Token,
	}
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
	if userID != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Login already taken"})
		return
	}
	err = s.BoxService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.SetCookie("userID", user.Token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
