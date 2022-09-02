package main

import (
	"fmt"
	_ "fmt"
	"log"
	"net/http"
	_ "net/http"
	"strings"

	"github.com/fauzan264/crowdfunding-rest-api/auth"
	"github.com/fauzan264/crowdfunding-rest-api/campaign"
	"github.com/fauzan264/crowdfunding-rest-api/handler"
	"github.com/fauzan264/crowdfunding-rest-api/helper"
	"github.com/fauzan264/crowdfunding-rest-api/user"
	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	campaignRepository := campaign.NewRepository(db)

	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	authService := auth.NewService()

	// userId := uuid.MustParse("adfjakdf")

	// if len(userId) < 16 {
	// 	log.Fatal("Wrong UUID")
	// }

	campaigns, _ := campaignService.FindCampaigns("")
	fmt.Println(len(campaigns))

	// id := uuid.MustParse("d0837349-eb4f-4864-acc6-e06f6c4676b6")
	// fmt.Println(authService.GenerateToken(id))

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/user", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/email_checker", userHandler.CheckEmailAvailability)
	api.POST("/avatar", authMiddleware(authService, userService), userHandler.UploadAvatar)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := uuid.MustParse(claim["user_id"].(string))

		user, err := userService.GetUserById(userId)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
