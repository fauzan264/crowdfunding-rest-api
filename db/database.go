package db

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Conn       *gorm.DB
	DBUser     = os.Getenv("DB_USER_CROWDFUNDING")
	DBPassword = os.Getenv("DB_PASSWORD_CROWDFUNDING")
	DBHost     = os.Getenv("DB_HOST_CROWDFUNDING")
	DBPort     = os.Getenv("DB_PORT_CROWDFUNDING")
	DBName     = os.Getenv("DB_NAME_CROWDFUNDING")
)

func Connection() (err error) {
	port, err := strconv.Atoi(DBPort)
	if err != nil {
		return err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, port, DBName)
	Conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	return nil
}
