package services

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"solearn.ru/m/v2/models"
	"time"
)

// RegisterUser Регистрация
func RegisterUser(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if err := models.DB.Where("email=?", input.Email).First(&input).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"msg": "Электронная почта уже занята"})
		return
	}

	input.Password = MD5(input.Password)

	models.DB.Create(&input)

	emailCheck := models.EmailCheck{UserID: input.ID, UUID: uuid.New().String()}
	models.DB.Create(&emailCheck)

	SendEmailUUID(input.Email, emailCheck.UUID)


	c.JSON(http.StatusOK, gin.H{"msg": "Регистрация прошла успешно."})
}

// LoginUser Авторизация
func LoginUser(c *gin.Context) {
	var input models.UserLoginData

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	var user models.User

	if err := models.DB.Where("email=?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Пользователь не найден"})
		return
	}

	if user.Password != MD5(input.Password) {
		c.JSON(http.StatusForbidden, gin.H{"msg": "Неверный пароль"})
		return
	}


	//if user.Role == 1 {
	//	c.JSON(http.StatusForbidden, gin.H{"msg": "Подтвердите адрес электронной почты"})
	//	return
	//}

	tokens := models.Token{UserID: user.ID, Refresh: CreateTokenRefresh()}

	models.DB.Create(&tokens)
	c.SetCookie("refresh_token", tokens.Refresh, 60*60*24*30, "/", os.Getenv("cookie_http"), false, true) // if https: secure = true
	c.JSON(http.StatusOK, gin.H{"access": CreateToken(user)})

}

// CreateToken Создание JWT
func CreateToken(user models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": user.ID,
		"email":  user.Email,
		"exp":    time.Now().Add(time.Minute * 15).Unix(),
	})

	jwtToken, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return jwtToken
}

// CreateTokenRefresh Создание Refresh токена
func CreateTokenRefresh() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 7200).Unix(),
	})

	jwtToken, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return jwtToken
}

// Logout Выход из аккаунта
func Logout(c *gin.Context) {

	token := models.Token{}

	refreshToken, err := c.Cookie("refresh_token")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	if err := models.DB.Where("refresh=?", refreshToken).First(&token).Error; err == nil {
		models.DB.Where("user_id=?", token.UserID).Delete(&token)
	}

	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"msg": "Вы успешно вышли из системы"})
}

// CheckToken Проверка JWT
func CheckToken(token string, c *gin.Context) bool {
	type MyCustomClaims struct {
		ID    uint   `json:"userid"`
		Email string `json:"email"`
		jwt.StandardClaims
	}

	tokenParse, _ := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	var user models.User

	err := models.DB.Where("id=?", tokenParse.Claims.(*MyCustomClaims).ID).First(&user).Error

	if _, ok := tokenParse.Claims.(*MyCustomClaims); ok && tokenParse.Valid && err == nil {
		c.Header("Role", string(user.Role))
		return true
	}
	return false
}

// Refresh Создание JWT с помощью Refresh токена
func Refresh(c *gin.Context) {
	tokenRefresh, err := c.Cookie("refresh_token")

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"msg": "Рефреш токен не найден"})
		return
	}

	token := models.Token{}

	if err := models.DB.Where("refresh=?", tokenRefresh).First(&token).Error; err == nil && CheckToken(tokenRefresh, c) {
		user := models.User{}

		models.DB.Where("id=?", token.UserID).First(&user)

		newToken := CreateToken(user)

		c.JSON(http.StatusOK, gin.H{"access": newToken})
		return
	} else if !CheckToken(tokenRefresh, c) {
		models.DB.Delete(&token)
	}

	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusConflict, gin.H{"msg": "Рефреш токен не валидный"})
}

func Activate(c *gin.Context) {
	emailCheck := models.EmailCheck{}
	if err := models.DB.Where("uuid=?", c.Param("uuid")).First(&emailCheck).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Подтверждение не существует или недействительно!"})
		return
	}

	user := models.User{}

	if err := models.DB.Where("id=?", emailCheck.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Аккаунт не найден!"})
		models.DB.Delete(&emailCheck)
		return
	}

	user.Role = 2

	models.DB.Updates(user)
	models.DB.Where("user_id", emailCheck.UserID).Delete(emailCheck)

	c.JSON(http.StatusOK, gin.H{"msg": "Аккаунт успешно активирован!"})
}

func MD5(data string) string {
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}

func Verification(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Не найден токен."})
		return
	}
	if CheckToken(token, c) {
		c.JSON(http.StatusOK, gin.H{"msg": "Успешный успех"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Ваша сессия истекла. Пожалуйста, войдите заново."})
	}
}
