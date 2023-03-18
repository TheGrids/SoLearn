package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"solearn.ru/m/v2/models"
)

func GetCourses(c *gin.Context) {
	if _, role, err := CheckAccessToken(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	} else if role == 0 {
		c.JSON(http.StatusConflict, gin.H{"msg": "Токен недействителен или отсутствует."})
		return
	} else if role < 2 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Недостаточно прав."})
		return
	}

	var courses []models.Course

	models.DB.Order("id DESC").Find(&courses)

	c.JSON(http.StatusOK, gin.H{"data": courses})
}
