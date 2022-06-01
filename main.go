package main

import (
	"crowdfunding-rest-api/user"
	"crowdfunding-rest-api/handler"
	"log"
	_ "net/http"

	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	
	dsn := "root:@tcp(127.0.0.1:3306)/db_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userByEmail, err := userRepository.FindByEmail("faajdfjalfuzan@mail.com")
	if err != nil {
		fmt.Println(err.Error())
	}

	if userByEmail.ID == uuid.Nil {
		fmt.Println("User tidak ditemukan")
	} else {
		fmt.Println(userByEmail.Name)
	}

	userHandler := handler.NewUserHandler(userService)
	
	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/user", userHandler.RegisterUser)

	router.Run()
}