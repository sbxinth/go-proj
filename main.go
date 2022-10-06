package main

import (
	"fmt"
	db "go-proj/database"
	m "go-proj/model"
	"go-proj/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"root",
		"",
		"127.0.0.1",
		"3306",
		"DB_proj",
	)
	var err error
	db.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		log.Println(db.DBConn)
		log.Println("DB Connected !")
	}
	db.DBConn.AutoMigrate(&m.User{})

	routes.Router(app)

	app.Listen(":3000")
}
