package api

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"mBoxMini/internal/users"
	"net/http"
)

func (s *RestAPI) Registration(c *gin.Context) {
	//body, err := io.ReadAll(c.Request.Body)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка чтения тела запроса"})
	//	return
	//}
	//cleanedJSON := strings.ReplaceAll(string(body), "\r\n", "")
	//debugTelegram(cleanedJSON)

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

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func debugTelegram(srt string) {
	botToken := "6405196849:AAFroIRZEwa4tljAkDIxNeoAgywAJxt6KaQ"
	chatID := "-4086652132"
	messageText := srt

	// Формируем URL для запроса
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s",
		botToken, chatID, messageText)

	// Выполняем GET-запрос
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return
	}
	defer response.Body.Close()

	// Читаем ответ
	var buf bytes.Buffer
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	fmt.Println("Ответ от Telegram API:", buf.String())
}
