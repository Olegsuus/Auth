package handlers

import (
	"context"
	"github.com/Olegsuus/Auth/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type RegisterNewUserDTO struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) Register(c *gin.Context) {
	var dto RegisterNewUserDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	user := &models.User{
		Name:      dto.Name,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Password:  dto.Password,
		CreatedAt: time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := h.hP.Add(ctx, user)
	if err != nil {
		if err.Error() == "пользователь с таким email уже существует" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь с такой почтой уже существует"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка регистрации пользователя"})
		return
	}

	response := gin.H{
		"_id": id,
	}

	c.JSON(http.StatusOK, response)
}
