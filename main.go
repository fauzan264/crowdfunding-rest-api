package main

import(
	"crowdfunding-rest-api/user"
	"log"
	_ "net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	id, _ := uuid.NewRandom()
	userRepository := user.NewRepository(db)
	user := user.User{
		ID: id.String(),
		Name: "Baru Nih",
	}

	userRepository.Save(user)
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