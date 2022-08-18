package main

import (
	"fmt"
	_ "fmt"
	"log"
	_ "net/http"

	"github.com/fauzan264/crowdfunding-rest-api/auth"
	"github.com/fauzan264/crowdfunding-rest-api/handler"
	"github.com/fauzan264/crowdfunding-rest-api/user"

	"github.com/gin-gonic/gin"
	_ "github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/db_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	// id := uuid.MustParse("d0837349-eb4f-4864-acc6-e06f6c4676b6")
	// fmt.Println(authService.GenerateToken(id))

	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMDBiMWEyMjYtZDI1My00NDE1LTk0OGUtYWE5YzlmMzZjYzdjIn0.7PfwXxts0rmXHCBqyYhBE5TpSV81Mc4-7vaN0N9qBXo")
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println("ERROR")
		fmt.Println("ERROR")
	}

	if token.Valid {
		fmt.Println("VALID")
		fmt.Println("VALID")
		fmt.Println("VALID")
	} else {
		fmt.Println("INVALID")
		fmt.Println("INVALID")
		fmt.Println("INVALID")
	}

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/user", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/email_checker", userHandler.CheckEmailAvailability)
	api.POST("/avatar", userHandler.UploadAvatar)

	router.Run()
}
