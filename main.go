package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"user-service/config"
	"user-service/controller"
)

func main() {
	dns := os.Getenv("MYSQL_STRING_CONNECTION")
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	appContext := config.NewAppContext(db)
	db = db.Debug()

	r := gin.Default()
	v1 := r.Group("/v1")
	v1.POST("/login", controller.Login(appContext))

	r.Run()
}
