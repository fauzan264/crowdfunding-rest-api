package main

import (
	_ "fmt"
	"log"
	"net/http"
	_ "net/http"
	"strings"

	"github.com/fauzan264/crowdfunding-rest-api/auth"
	"github.com/fauzan264/crowdfunding-rest-api/campaign"
	"github.com/fauzan264/crowdfunding-rest-api/db"
	"github.com/fauzan264/crowdfunding-rest-api/handler"
	"github.com/fauzan264/crowdfunding-rest-api/helper"
	"github.com/fauzan264/crowdfunding-rest-api/user"
	"github.com/golang-jwt/jwt/v4"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
)

func main() {

	err := db.Connection()
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db.Conn)
	campaignRepository := campaign.NewRepository(db.Conn)

	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/user", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/email_checker", userHandler.CheckEmailAvailability)
	api.POST("/avatar", authMiddleware(authService, userService), userHandler.UploadAvatar)

	// campaign
	api.GET("/campaign", campaignHandler.GetCampaigns)
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
