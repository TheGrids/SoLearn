package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"solearn.ru/m/v2/models"
	"solearn.ru/m/v2/services"
	"time"
)

func main() {
	r := gin.Default()

	models.ConnectionDataBase()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://solearn.ru", "http://127.0.0.1:5173"},
		AllowHeaders:     []string{"Role", "Origin", "Content-Type", "Authorization"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		MaxAge:           1 * time.Minute,
	}))

	r.POST("/register", services.RegisterUser)
	r.GET("/activate/:uuid", services.Activate)
	r.POST("/login", services.LoginUser)
	r.GET("/logout", services.Logout)
	r.POST("/refresh", services.Refresh)

	//r.Static("/image", "./image")
	//r.GET("/new_products", services.GetNewProducts)
	//r.GET("/popular_products", services.GetPopularProducts)
	//r.GET("/products", services.SearchProducts)
	//r.GET("/product/:id", services.Product)

	r.Run()
}
