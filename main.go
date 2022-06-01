package main

import (
	"crowdfunding-rest-api/user"
	"crowdfunding-rest-api/handler"
	"log"
	_ "net/http"

	_ "fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	_ "github.com/google/uuid"
)

func main() {
	
	dsn := "root:@tcp(127.0.0.1:3306)/db_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)
	
	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/user", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)

	router.Run()
}