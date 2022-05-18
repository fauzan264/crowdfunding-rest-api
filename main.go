package main

import(
	"crowdfunding-rest-api/user"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/db_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connection to database success")

	// create variable for struct
	var users []user.User

	// find db
	db.Find(&users)

	for _, user := range users {
		fmt.Println(user.Name)
		fmt.Println(user.Email)
		fmt.Println("==================")
	}	

}