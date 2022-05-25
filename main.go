package main

import(
	"crowdfunding-rest-api/user"
	"log"
	_ "net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

func main() {
	// router := gin.Default()
	// router.GET("/user", handlerUser)
	// router.Run()
	
	dsn := "root:@tcp(127.0.0.1:3306)/db_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	id, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.ID = id
	userInput.Name = "Ahmad"
	userInput.Email = "ahmad@mail.com"
	userInput.Occupation = "penulis kode"
	userInput.Password = "password"
	userInput.CreatedAt = time.Now()

	userService.RegisterUser(userInput)
	
}

// func handlerUser(c *gin.Context) {
// 	dsn := "root:@tcp(127.0.0.1:3306)/db_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	// create variable for struct
// 	var users []user.User

// 	// find db
// 	db.Find(&users)

// 	// convert to JSON
// 	c.JSON(http.StatusOK, users)
// }